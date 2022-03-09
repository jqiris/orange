package hall

import (
	"errors"
	"fmt"
	"github.com/jqiris/kungfu/v2/rpc"
	"github.com/jqiris/orange/servers/mahjong"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jqiris/kungfu/v2/discover"
	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/utils"
	"github.com/jqiris/orange/constant"
	"github.com/jqiris/orange/database"
	"github.com/jqiris/orange/model"
	"github.com/jqiris/orange/protos"
)

func (h *ServerHall) Error(c *gin.Context, msg string, args ...int) {
	errcode := 1
	if len(args) > 0 {
		errcode = args[0]
	}
	c.JSON(200, map[string]any{"errcode": errcode, "errmsg": msg})
}

func (h *ServerHall) Success(c *gin.Context, data map[string]any) {
	if data == nil {
		data = make(map[string]any)
	}
	data["errcode"] = 0
	data["errmsg"] = "ok"
	c.JSON(200, data)
}
func (h *ServerHall) checkAccount(account, sign string) bool {
	return len(account) > 0 && len(sign) > 0
}

func (h *ServerHall) Login(c *gin.Context) {
	account, sign := c.Query("account"), c.Query("sign")
	if !h.checkAccount(account, sign) {
		h.Error(c, "unknown error")
		return
	}
	ip := c.ClientIP()
	if strings.Contains(ip, "::ffff:") {
		ip = ip[7:]
	}
	data, err := database.GetUserByAccount(account)
	if err != nil {
		logger.Error(err)
		h.Success(c, nil)
		return
	}
	ret := map[string]any{
		"account": data.Account,
		"userid":  data.UserID,
		"name":    data.Name,
		"lv":      data.Lv,
		"exp":     data.Exp,
		"coins":   data.Coins,
		"gems":    data.Gems,
		"ip":      ip,
		"sex":     data.Sex,
	}
	if len(data.RoomID) > 0 {
		if _, err := database.GetMjRoomById(data.RoomID); err == nil {
			ret["roomid"] = data.RoomID
		} else if database.ErrRecordNotFound(err) {
			err = database.UpdateUser(data.UserID, map[string]any{"room_id": ""})
			if err != nil {
				logger.Error(err)
			}
		} else {
			logger.Error(err)
		}
	}
	//更新用户登录时间
	err = database.UpdateUser(data.UserID, map[string]any{"login_time": time.Now()})
	if err != nil {
		logger.Error(err)
	}
	h.Success(c, ret)
}

func (h *ServerHall) CreateUser(c *gin.Context) {
	account, sign := c.Query("account"), c.Query("sign")
	if !h.checkAccount(account, sign) {
		h.Error(c, "unknown error")
		return
	}
	name := c.Query("name")
	coins, gems := 1000, 21
	nowTime := time.Now()
	if _, err := database.GetUserByAccount(account); err == nil {
		h.Error(c, "account have already exist.")
		return
	}
	err := database.CreateUser(&model.UserMember{
		Account:   account,
		Name:      name,
		Sex:       0,
		Headimg:   "null",
		Lv:        0,
		Exp:       0,
		Coins:     coins,
		Gems:      gems,
		RegTime:   nowTime,
		LoginTime: nowTime,
	})
	if err != nil {
		logger.Error(err)
		h.Error(c, "system error", 2)
		return
	}
	h.Success(c, nil)
}

func (h *ServerHall) CreatePrivateRoom(c *gin.Context) {
	account, sign := c.Query("account"), c.Query("sign")
	if !h.checkAccount(account, sign) {
		h.Error(c, "unknown error")
		return
	}
	conf := c.Query("conf")
	data, err := database.GetUserByAccount(account)
	if err != nil {
		logger.Error(err)
		h.Error(c, "system error")
		return
	}
	if len(data.RoomID) > 0 {
		room, err := database.GetMjRoomById(data.RoomID)
		if err != nil && database.ErrRecordNotFound(err) {
			if err = database.UpdateUser(data.UserID, map[string]any{"room_id": ""}); err != nil {
				logger.Error(err)
			}
		} else {
			logger.Error(err, room)
			h.Error(c, "user is playing in room now.", -1)
			return
		}

	}
	userId, name := data.UserID, data.Name
	reqCreateSign := utils.Md5(fmt.Sprintf("%v%v%v%v", userId, conf, data.Gems, constant.RoomPriKey))
	reqCreate := &protos.InnerCreateRoomReq{
		UserId: userId,
		Gems:   int64(data.Gems),
		Conf:   conf,
		Sign:   reqCreateSign,
	}
	respCreate := mahjong.CreatePrivateRoom(h.Rpc, reqCreate)
	if respCreate.Errcode > constant.StatusOk {
		h.Error(c, respCreate.Errmsg, int(respCreate.Errcode))
		return
	}
	roomId := respCreate.GetCreateRoom().RoomId
	serverId := respCreate.GetCreateRoom().ServerId
	server := discover.GetServerById(serverId)
	if server == nil {
		h.Error(c, constant.ErrNotFoundMahjongServer)
		return
	}
	reqEnterSign := utils.Md5(fmt.Sprintf("%v%v%v%v", userId, name, roomId, constant.RoomPriKey))
	reqEnter := &protos.InnerEnterRoomReq{
		UserId: userId,
		Name:   name,
		RoomId: roomId,
		Sign:   reqEnterSign,
	}
	respEnter := mahjong.EnterPrivateRoom(h.Rpc, reqEnter, server)
	if respEnter.Errcode > constant.StatusOk {
		h.Error(c, respEnter.Errmsg, int(respEnter.Errcode))
		return
	}
	err = database.UpdateUser(userId, map[string]any{"room_id": roomId})
	if err != nil {
		logger.Error(err)
	}
	token := respEnter.GetEnterRoom().Token
	nowTime := time.Now().UnixMilli()
	resp := map[string]any{
		"roomid": roomId,
		"ip":     server.ServerIp,
		"port":   server.ClientPort,
		"token":  token,
		"time":   nowTime,
	}
	resp["sign"] = utils.Md5(fmt.Sprintf("%v%v%v%v", roomId, token, nowTime, constant.RoomPriKey))
	h.Success(c, resp)
}

func (h *ServerHall) EnterPrivateRoom(c *gin.Context) {
	account, sign := c.Query("account"), c.Query("sign")
	if !h.checkAccount(account, sign) {
		h.Error(c, "unknown error")
		return
	}
	roomId := c.Query("roomid")
	if len(roomId) < 1 {
		h.Error(c, "parameters don't match api requirements", -1)
		return
	}
	data, err := database.GetUserByAccount(account)
	if err != nil {
		logger.Error(err)
		h.Error(c, "system error", -1)
		return
	}
	userId, name := data.UserID, data.Name

	room, err := database.GetMjRoomById(roomId)
	if err != nil {
		logger.Error(err)
		h.Error(c, "enter room failed.", -2)
		return
	}
	serverId := room.ServerID
	server := discover.GetServerById(serverId)
	if server == nil {
		h.Error(c, constant.ErrNotFoundMahjongServer)
		return
	}
	reqEnterSign := utils.Md5(fmt.Sprintf("%v%v%v%v", userId, name, roomId, constant.RoomPriKey))
	reqEnter := &protos.InnerEnterRoomReq{
		UserId: userId,
		Name:   name,
		RoomId: roomId,
		Sign:   reqEnterSign,
	}
	respEnter := mahjong.EnterPrivateRoom(h.Rpc, reqEnter, server)
	if respEnter.Errcode > constant.StatusOk {
		h.Error(c, respEnter.Errmsg, int(respEnter.Errcode))
		return
	}
	err = database.UpdateUser(userId, map[string]any{"room_id": roomId})
	if err != nil {
		logger.Error(err)
	}
	token := respEnter.GetEnterRoom().Token
	nowTime := time.Now().UnixMilli()
	resp := map[string]any{
		"roomid": roomId,
		"ip":     server.ServerIp,
		"port":   server.ClientPort,
		"token":  token,
		"time":   nowTime,
	}
	resp["sign"] = utils.Md5(fmt.Sprintf("%v%v%v%v", roomId, token, nowTime, constant.RoomPriKey))
	h.Success(c, resp)
}

func (h *ServerHall) GetHistoryList(c *gin.Context) {
	account, sign := c.Query("account"), c.Query("sign")
	if !h.checkAccount(account, sign) {
		h.Error(c, "unknown error")
		return
	}
	data, err := database.GetUserByAccount(account)
	if err != nil {
		logger.Error(err)
		h.Error(c, "system error", -1)
		return
	}
	resp := map[string]any{
		"history": data.History,
	}
	h.Success(c, resp)
}

func (h *ServerHall) GetGamesOfRoom(c *gin.Context) {
	account, sign := c.Query("account"), c.Query("sign")
	if !h.checkAccount(account, sign) {
		h.Error(c, "unknown error")
		return
	}
	gameType := c.Query("game_type")
	uuid := c.Query("uuid")
	if len(uuid) == 0 || len(gameType) < 1 {
		h.Error(c, "parameters don't match api requirements", -1)
		return
	}
	data, err := database.GetMjActionArchive(gameType, uuid)
	if err != nil {
		logger.Error(err)
		h.Error(c, "system error", -1)
		return
	}
	h.Success(c, map[string]any{"data": data})
}

func (h *ServerHall) GetDetailOfGame(c *gin.Context) {
	account, sign := c.Query("account"), c.Query("sign")
	if !h.checkAccount(account, sign) {
		h.Error(c, "unknown error")
		return
	}
	gameType, uuid, index := c.Query("game_type"), c.Query("uuid"), utils.StringToInt32(c.Query("index"))
	if len(gameType) == 0 || len(uuid) == 0 || index < 0 {
		h.Error(c, "parameters don't match api requirements", -1)
		return
	}
	data, err := database.GetMjActionArchiveDetail(gameType, uuid, index)
	if err != nil {
		logger.Error(err)
		h.Error(c, "system error", -1)
		return
	}
	h.Success(c, map[string]any{"data": data})
}

func (h *ServerHall) GetUserStatus(c *gin.Context) {
	account, sign := c.Query("account"), c.Query("sign")
	if !h.checkAccount(account, sign) {
		h.Error(c, "unknown error")
		return
	}
	data, err := database.GetUserByAccount(account)
	if err != nil {
		logger.Error(err)
		h.Error(c, "get gems failed")
		return
	}
	h.Success(c, map[string]any{"gems": data.Gems})
}

func (h *ServerHall) GetMessage(c *gin.Context) {
	account, sign := c.Query("account"), c.Query("sign")
	if !h.checkAccount(account, sign) {
		h.Error(c, "unknown error")
		return
	}
	typ, version := c.Query("type"), c.Query("version")
	if len(typ) == 0 {
		h.Error(c, "parameters don't match api requirements", -1)
		return
	}
	data, err := database.GetGameNotice(typ, version)
	if err != nil {
		logger.Error(err)
		h.Error(c, "get message failed")
		return
	}
	h.Success(c, map[string]any{"msg": data.Msg, "version": data.Version})
}

func (h *ServerHall) Maintain(c *gin.Context) {
	serverId, reqState, sign := c.PostForm("serverId"), utils.StringToInt32(c.PostForm("reqState")), c.PostForm("sign")
	md5 := utils.Md5(fmt.Sprintf("%v-%v-%v", serverId, reqState, constant.MaintainKey))
	if md5 != sign {
		logger.Errorf("maintain sign is wrong, ip:%v", c.ClientIP())
		h.Success(c, nil)
		return
	}
	if reqState != 1 && reqState != 2 {
		h.Error(c, "invalid req state")
		return
	}
	if err := h.serverMaintain(serverId, reqState); err != nil {
		h.Error(c, err.Error())
		return
	}
	h.Success(c, nil)
}

func (h *ServerHall) serverMaintain(serverId string, reqState int32) error {
	//广播通知服务器进行维护
	server := discover.GetServerById(serverId, false) //不过滤维护状态
	if server == nil {
		return errors.New("server not found")
	}
	req := rpc.NewReqBuilder(server)
	err := h.Rpc.Publish(req.SetMsgId(int32(protos.InnerMsgId_InnerMsgMaintain)).SetReq(&protos.InnerMaintainReq{
		ServerId: serverId,
		ReqState: reqState,
	}).Build())
	if err != nil {
		return err
	}
	return nil
}
