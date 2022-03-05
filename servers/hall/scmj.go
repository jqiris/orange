package hall

import (
	"fmt"
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
	"github.com/jqiris/orange/servers/game"
	"github.com/jqiris/orange/tools"
)

func (h *HallServer) Error(c *gin.Context, msg string, args ...int) {
	errcode := 1
	if len(args) > 0 {
		errcode = args[0]
	}
	c.JSON(200, map[string]any{"errcode": errcode, "errmsg": msg})
}

func (h *HallServer) Success(c *gin.Context, data map[string]any) {
	if data == nil {
		data = make(map[string]any)
	}
	data["errcode"] = 0
	data["errmsg"] = "ok"
	c.JSON(200, data)
}
func (h *HallServer) checkAccount(account, sign string) bool {
	return len(account) > 0 && len(sign) > 0
}

func (h *HallServer) Login(c *gin.Context) {
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
		"userid":  data.Userid,
		"name":    data.Name,
		"lv":      data.Lv,
		"exp":     data.Exp,
		"coins":   data.Coins,
		"gems":    data.Gems,
		"ip":      ip,
		"sex":     data.Sex,
	}
	if data.Roomid > 0 {
		if _, err := database.GetRoomById(data.Roomid); err == nil {
			ret["roomid"] = data.Roomid
		} else if database.ErrRecordNotFound(err) {
			database.UpdateUser(data.Userid, map[string]interface{}{"roomid": 0})
		} else {
			logger.Error(err)
		}
	}
	h.Success(c, ret)
}

func (h *HallServer) CreateUser(c *gin.Context) {
	account, sign := c.Query("account"), c.Query("sign")
	if !h.checkAccount(account, sign) {
		h.Error(c, "unknown error")
		return
	}
	name := c.Query("name")
	coins, gems := 1000, 21

	if _, err := database.GetUserByAccount(account); err == nil {
		h.Error(c, "account have already exist.")
	}
	err := database.CreateUser(&model.TUser{
		Account: account,
		Name:    tools.Stringify(name),
		Sex:     0,
		Headimg: "null",
		Lv:      0,
		Exp:     0,
		Coins:   coins,
		Gems:    gems,
		Roomid:  0,
		History: "",
	})
	if err != nil {
		logger.Error(err)
		h.Error(c, "system error", 2)
		return
	}
	h.Success(c, nil)
}

func (h *HallServer) CreatePrivateRoom(c *gin.Context) {
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
	if data.Roomid > 0 {
		h.Error(c, "user is playing in room now.", -1)
		return
	}
	userId, name := data.Userid, data.Name
	reqCreateSign := utils.Md5(fmt.Sprintf("%v%v%v%v", userId, conf, data.Gems, constant.ROOM_PRI_KEY))
	reqCreate := &protos.InnerCreateRoomReq{
		Userid: int64(userId),
		Gems:   int64(data.Gems),
		Conf:   conf,
		Sign:   reqCreateSign,
	}
	respCreate := game.CreatePrivateRoom(h.Rpc, reqCreate)
	if respCreate.Errcode > constant.StatusOk {
		h.Error(c, respCreate.Errmsg, int(respCreate.Errcode))
		return
	}
	roomId := respCreate.GetCreateRoom().Roomid
	serverId := respCreate.GetCreateRoom().ServerId
	server := discover.GetServerById(serverId)
	if server == nil {
		h.Error(c, constant.ErrNotFoundGameServer)
		return
	}
	reqEnterSign := utils.Md5(fmt.Sprintf("%v%v%v%v", userId, name, roomId, constant.ROOM_PRI_KEY))
	reqEnter := &protos.InnerEnterRoomReq{
		Userid: int64(userId),
		Name:   name,
		Roomid: roomId,
		Sign:   reqEnterSign,
	}
	respEnter := game.EnterPrivateRoom(h.Rpc, reqEnter, server)
	if respEnter.Errcode > constant.StatusOk {
		h.Error(c, respEnter.Errmsg, int(respEnter.Errcode))
		return
	}
	err = database.UpdateUser(userId, map[string]any{"roomid": roomId})
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
	resp["sign"] = utils.Md5(fmt.Sprintf("%v%v%v%v", roomId, token, nowTime, constant.ROOM_PRI_KEY))
	h.Success(c, resp)
}

func (h *HallServer) EnterPrivateRoom(c *gin.Context) {
	account, sign := c.Query("account"), c.Query("sign")
	if !h.checkAccount(account, sign) {
		h.Error(c, "unknown error")
		return
	}
	roomId := utils.StringToInt(c.Query("roomid"))
	if roomId < 1 {
		h.Error(c, "parameters don't match api requirements", -1)
		return
	}
	data, err := database.GetUserByAccount(account)
	if err != nil {
		logger.Error(err)
		h.Error(c, "system error", -1)
		return
	}
	userId, name := data.Userid, data.Name

	room, err := database.GetRoomById(roomId)
	if err != nil {
		logger.Error(err)
		h.Error(c, "enter room failed.", -2)
		return
	}
	serverId := room.ServerId
	server := discover.GetServerById(serverId)
	if server == nil {
		h.Error(c, constant.ErrNotFoundGameServer)
		return
	}
	reqEnterSign := utils.Md5(fmt.Sprintf("%v%v%v%v", userId, name, roomId, constant.ROOM_PRI_KEY))
	reqEnter := &protos.InnerEnterRoomReq{
		Userid: int64(userId),
		Name:   name,
		Roomid: int32(roomId),
		Sign:   reqEnterSign,
	}
	respEnter := game.EnterPrivateRoom(h.Rpc, reqEnter, server)
	if respEnter.Errcode > constant.StatusOk {
		h.Error(c, respEnter.Errmsg, int(respEnter.Errcode))
		return
	}
	err = database.UpdateUser(userId, map[string]any{"roomid": roomId})
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
	resp["sign"] = utils.Md5(fmt.Sprintf("%v%v%v%v", roomId, token, nowTime, constant.ROOM_PRI_KEY))
	h.Success(c, resp)
}

func (h *HallServer) GetHistoryList(c *gin.Context) {
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

func (h *HallServer) GetGamesOfRoom(c *gin.Context) {
	account, sign := c.Query("account"), c.Query("sign")
	if !h.checkAccount(account, sign) {
		h.Error(c, "unknown error")
		return
	}
	uuid := c.Query("uuid")
	if len(uuid) == 0 {
		h.Error(c, "parameters don't match api requirements", -1)
		return
	}
	data, err := database.GetGameArchive(uuid)
	if err != nil {
		logger.Error(err)
		h.Error(c, "system error", -1)
		return
	}
	h.Success(c, map[string]any{"data": data})
}

func (h *HallServer) GetDetailOfGame(c *gin.Context) {
	account, sign := c.Query("account"), c.Query("sign")
	if !h.checkAccount(account, sign) {
		h.Error(c, "unknown error")
		return
	}
	uuid, index := c.Query("uuid"), utils.StringToInt(c.Query("index"))
	if len(uuid) == 0 || index < 0 {
		h.Error(c, "parameters don't match api requirements", -1)
		return
	}
	data, err := database.GetGameArchiveDetail(uuid, index)
	if err != nil {
		logger.Error(err)
		h.Error(c, "system error", -1)
		return
	}
	h.Success(c, map[string]any{"data": data})
}

func (h *HallServer) GetUserStatus(c *gin.Context) {
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

func (h *HallServer) GetMessage(c *gin.Context) {
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
	data, err := database.GetMessage(typ, version)
	if err != nil {
		logger.Error(err)
		h.Error(c, "get message failed")
		return
	}
	h.Success(c, map[string]any{"msg": data.Msg, "version": data.Version})
}
