package game

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	socketio "github.com/googollee/go-socket.io"
	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/utils"
	"github.com/jqiris/orange/constant"
)

type Msg struct {
	Errcode int                    `json:"errcode"`
	Errmsg  string                 `json:"errmsg"`
	Data    map[string]interface{} `json:"data"`
}

type SocketCtx struct {
	UserId  int
	GameMgr GameMgr
}

type LoginData struct {
	Token  string `json:"token"`
	Roomid int    `json:"roomid"`
	Time   int64  `json:"time"`
	Sign   string `json:"sign"`
}

func (s *GameServer) SocketRouter(sc *socketio.Server) {
	ns := "/"
	sc.OnConnect(ns, s.OnConnect)                         //连接
	sc.OnDisconnect(ns, s.OnDisconnect)                   //断开连接
	sc.OnEvent(ns, "login", s.Login)                      //登录
	sc.OnEvent(ns, "ready", s.Ready)                      //准备
	sc.OnEvent(ns, "huanpai", s.Huanpai)                  //换牌
	sc.OnEvent(ns, "dingque", s.Dingque)                  //定缺
	sc.OnEvent(ns, "chupai", s.Chupai)                    //出牌
	sc.OnEvent(ns, "peng", s.Peng)                        //碰
	sc.OnEvent(ns, "gang", s.Gang)                        //杠
	sc.OnEvent(ns, "hu", s.Hu)                            //胡
	sc.OnEvent(ns, "guo", s.Guo)                          //过
	sc.OnEvent(ns, "game_ping", s.GamePing)               //心跳
	sc.OnEvent(ns, "chat", s.Chat)                        //聊天
	sc.OnEvent(ns, "quick_chat", s.QuickChat)             //快速聊天
	sc.OnEvent(ns, "voice_msg", s.VoiceMsg)               //语音聊天
	sc.OnEvent(ns, "emoji", s.Emoji)                      //表情
	sc.OnEvent(ns, "exit", s.Exit)                        //退出房间
	sc.OnEvent(ns, "dispress", s.Dispress)                //解散房间
	sc.OnEvent(ns, "dissolve_request", s.DissolveRequest) //解散房间请求
	sc.OnEvent(ns, "dissolve_agree", s.DissolveAgree)     //解散房间同意
	sc.OnEvent(ns, "dissolve_reject", s.DissolveReject)   //解散房间不同意

	go sc.Serve()
	defer sc.Close()

	http.Handle("/socket.io/", sc)
	http.Handle("/hi", sc)
	logger.Infof("socket server start at:%v", constant.CLIENT_PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", constant.CLIENT_PORT), nil)
}

func (s *GameServer) OnConnect(c socketio.Conn) error {
	logger.Infof("Connect:%v", c.Context())
	return nil
}

func (s *GameServer) OnDisconnect(c socketio.Conn, reason string) {
	logger.Warnf("Disconnect:%v,reason:%v", c.Context(), reason)
	ctx := s.GetSocketCtx(c)
	if ctx == nil {
		return
	}
	userId := ctx.UserId
	//如果是旧链接断开，则不需要处理。
	if userMgr.get(userId) != c {
		return
	}

	var data = map[string]interface{}{
		"userid": userId,
		"online": false,
	}

	//通知房间内其它玩家
	userMgr.broadcastInRoom("user_state_push", data, userId)

	//清除玩家的在线信息
	userMgr.del(userId)
	ctx.UserId = 0
}
func (s *GameServer) DissolveReject(c socketio.Conn) {
	ctx := s.GetSocketCtx(c)
	if ctx == nil {
		return
	}
	userId := ctx.UserId
	roomId := roomMgr.getUserRoom(userId)
	if roomId < 1 {
		return
	}
	ret := ctx.GameMgr.dissolveAgree(roomId, userId, false)
	if ret != nil {
		userMgr.broadcastInRoom("dissolve_cancel_push", nil, userId, true)
	}
}

func (s *GameServer) DissolveAgree(c socketio.Conn) {
	ctx := s.GetSocketCtx(c)
	if ctx == nil {
		return
	}
	userId := ctx.UserId
	roomId := roomMgr.getUserRoom(userId)
	if roomId < 1 {
		return
	}
	ret := ctx.GameMgr.dissolveAgree(roomId, userId, true)
	if ret != nil {
		dr := ret.Dr
		remainingTime := (dr.EndTime - time.Now().UnixMilli()) / 1000
		data := map[string]interface{}{
			"time":   remainingTime,
			"states": dr.States,
		}
		userMgr.broadcastInRoom("dissolve_notice_push", data, userId, true)

		doAllAgree := true
		for i := 0; i < len(dr.States); i++ {
			if dr.States[i] == false {
				doAllAgree = false
				break
			}
		}

		if doAllAgree {
			ctx.GameMgr.doDissolve(roomId)
		}
	}
}

func (s *GameServer) DissolveRequest(c socketio.Conn) {
	ctx := s.GetSocketCtx(c)
	if ctx == nil {
		return
	}
	userId := ctx.UserId
	roomId := roomMgr.getUserRoom(userId)
	if roomId < 1 {
		return
	}
	if ctx.GameMgr.hasBegan(roomId) == false {
		return
	}
	ret := ctx.GameMgr.dissolveRequest(roomId, userId)
	if ret != nil {
		dr := ret.Dr
		remainingTime := (dr.EndTime - time.Now().UnixMilli()) / 1000
		data := map[string]interface{}{
			"time":   remainingTime,
			"states": dr.States,
		}
		userMgr.broadcastInRoom("dissolve_notice_push", data, userId, true)
	}
}

func (s *GameServer) Dispress(c socketio.Conn) {
	ctx := s.GetSocketCtx(c)
	if ctx == nil {
		return
	}
	userId := ctx.UserId
	roomId := roomMgr.getUserRoom(userId)
	if roomId < 1 {
		return
	}
	//如果游戏已经开始，则不可以
	if ctx.GameMgr.hasBegan(roomId) {
		return
	}

	//如果不是房主，则不能解散房间
	if roomMgr.isCreator(roomId, userId) == false {
		return
	}
	userMgr.broadcastInRoom("dispress_push", nil, userId, true)
	userMgr.kickAllInRoom(roomId)
	roomMgr.destroy(roomId)
	c.Close()
}

func (s *GameServer) Exit(c socketio.Conn) {
	ctx := s.GetSocketCtx(c)
	if ctx == nil {
		return
	}
	userId := ctx.UserId
	roomId := roomMgr.getUserRoom(userId)
	if roomId < 1 {
		return
	}
	if ctx.GameMgr.hasBegan(roomId) {
		return
	}
	//如果是房主，则只能走解散房间
	if roomMgr.isCreator(roomId, userId) {
		return
	}
	//通知其它玩家，有人退出了房间
	userMgr.broadcastInRoom("exit_notify_push", userId, userId, false)
	roomMgr.exitRoom(userId)
	userMgr.del(userId)
	c.Emit("exit_result")
	c.Close()
}

func (s *GameServer) Emoji(c socketio.Conn, msg string) {
	ctx := s.GetSocketCtx(c)
	if ctx == nil {
		return
	}
	userMgr.broadcastInRoom("emoji_push", map[string]interface{}{
		"sender":  ctx.UserId,
		"content": msg,
	}, ctx.UserId, true)
}

func (s *GameServer) VoiceMsg(c socketio.Conn, msg string) {
	ctx := s.GetSocketCtx(c)
	if ctx == nil {
		return
	}
	userMgr.broadcastInRoom("voice_msg_push", map[string]interface{}{
		"sender":  ctx.UserId,
		"content": msg,
	}, ctx.UserId, true)
}

func (s *GameServer) QuickChat(c socketio.Conn, msg string) {
	ctx := s.GetSocketCtx(c)
	if ctx == nil {
		return
	}
	userMgr.broadcastInRoom("quick_chat_push", map[string]interface{}{
		"sender":  ctx.UserId,
		"content": msg,
	}, ctx.UserId, true)
}
func (s *GameServer) Chat(c socketio.Conn, msg string) {
	ctx := s.GetSocketCtx(c)
	if ctx == nil {
		return
	}
	userMgr.broadcastInRoom("chat_push", map[string]interface{}{
		"sender":  ctx.UserId,
		"content": msg,
	}, ctx.UserId, true)
}

func (s *GameServer) IsLogined(c socketio.Conn) bool {
	ctx := s.GetSocketCtx(c)
	if ctx == nil {
		return false
	}
	return ctx.UserId > 0
}

func (s *GameServer) GetSocketCtx(c socketio.Conn) *SocketCtx {
	if v, ok := c.Context().(*SocketCtx); ok {
		return v
	}
	return nil
}

func (s *GameServer) Login(c socketio.Conn, msg string) {
	logger.Infof("Login:%v", msg)
	var data LoginData
	err := json.Unmarshal([]byte(msg), &data)
	if err != nil {
		logger.Error(err)
		c.Emit("login_result", Msg{Errcode: 1, Errmsg: "invalid parameters"})
		return
	}
	if s.IsLogined(c) {
		return //已经登录过
	}
	token := data.Token
	//roomId := utils.StringToInt(data.Roomid)
	roomId := data.Roomid
	rtime := data.Time
	sign := data.Sign
	md5 := utils.Md5(fmt.Sprintf("%v%v%v%v", roomId, token, rtime, constant.ROOM_PRI_KEY))
	if md5 != sign {
		c.Emit("login_result", Msg{Errcode: 2, Errmsg: "login failed invalid sign"})
		return
	}
	if !tokenMgr.isTokenVaild(token) {
		c.Emit("login_result", Msg{Errcode: 3, Errmsg: "token out of time"})
	}

	userId := tokenMgr.getUserId(token)
	roomId = roomMgr.getUserRoom(userId)
	userMgr.bind(userId, c)
	ctx := &SocketCtx{UserId: userId}
	c.SetContext(ctx)

	//返回房间信息
	roomInfo := roomMgr.getRoom(roomId)
	if roomInfo == nil {
		c.Emit("login_result", Msg{
			Errcode: 4,
			Errmsg:  "can't find the room",
		})
		return
	}
	seatIndex := roomMgr.getUserSeat(userId)
	roomInfo.Seats[seatIndex].Ip = c.RemoteAddr().String()
	var userData *Seat
	var seats []*Seat
	for i := 0; i < len(roomInfo.Seats); i++ {
		rs := roomInfo.Seats[i]
		online := false
		if rs.UserId > 0 {
			online = userMgr.isOnline(rs.UserId)
		}
		seat := &Seat{
			UserId:    rs.UserId,
			Score:     rs.Score,
			Name:      rs.Name,
			Ready:     rs.Ready,
			SeatIndex: i,
			Ip:        rs.Ip,
			Online:    online,
		}
		seats = append(seats, seat)
		if userId == rs.UserId {
			userData = seat
		}
	}
	ret := Msg{
		Errcode: 0,
		Errmsg:  "ok",
		Data: map[string]interface{}{
			"roomid":     roomInfo.Id,
			"conf":       roomInfo.Conf,
			"numofgames": roomInfo.NumOfGames,
			"seats":      seats,
		},
	}
	c.Emit("login_result", ret)

	//通知其他客户端
	userMgr.broadcastInRoom("new_user_comes_push", userData, userId)
	ctx.GameMgr = roomInfo.GameMgr
	//玩家上限，强制设置为true
	ctx.GameMgr.setReady(userId)
	c.Emit("login_finished")

	if roomInfo.Dr != nil {
		dr := roomInfo.Dr
		remainingTime := dr.EndTime - time.Now().Unix()
		data := map[string]interface{}{
			"time":   remainingTime,
			"states": dr.States,
		}
		userMgr.sendMsg(userId, "dissolve_notice_push", data)
	}
}

func (s *GameServer) Ready(c socketio.Conn) {
	if v := s.GetSocketCtx(c); v != nil {
		v.GameMgr.setReady(v.UserId)
		userMgr.broadcastInRoom("user_ready_push", map[string]interface{}{"userid": v.UserId, "ready": true}, v.UserId, true)
	}
}

//换牌
func (s *GameServer) Huanpai(c socketio.Conn, msg string) {
	ctx := s.GetSocketCtx(c)
	if ctx == nil {
		return
	}
	data := make(map[string]int)
	err := json.Unmarshal([]byte(msg), &data)
	if err != nil {
		logger.Error(err)
		return
	}
	p1, p2, p3 := data["p1"], data["p2"], data["p3"]
	if p1 < 0 || p2 < 0 || p3 < 0 {
		logger.Errorf("data err:%+v", data)
		return
	}
	ctx.GameMgr.huanSanZhang(ctx.UserId, p1, p2, p3)
}

//定缺
func (s *GameServer) Dingque(c socketio.Conn, que int) {
	ctx := s.GetSocketCtx(c)
	if ctx == nil {
		return
	}
	ctx.GameMgr.dingQue(ctx.UserId, que)
}

//出牌
func (s *GameServer) Chupai(c socketio.Conn, pai int) {
	ctx := s.GetSocketCtx(c)
	if ctx == nil {
		return
	}
	ctx.GameMgr.chuPai(ctx.UserId, pai)
}

//碰
func (s *GameServer) Peng(c socketio.Conn) {
	ctx := s.GetSocketCtx(c)
	if ctx == nil {
		return
	}
	ctx.GameMgr.peng(ctx.UserId)
}

//杠
func (s *GameServer) Gang(c socketio.Conn, pai int) {
	ctx := s.GetSocketCtx(c)
	if ctx == nil {
		return
	}
	if pai < 0 {
		return
	}
	ctx.GameMgr.gang(ctx.UserId, pai)
}

//胡
func (s *GameServer) Hu(c socketio.Conn) {
	ctx := s.GetSocketCtx(c)
	if ctx == nil {
		return
	}
	ctx.GameMgr.hu(ctx.UserId)
}

//过
func (s *GameServer) Guo(c socketio.Conn) {
	ctx := s.GetSocketCtx(c)
	if ctx == nil {
		return
	}
	ctx.GameMgr.guo(ctx.UserId)
}

//心跳
func (s *GameServer) GamePing(c socketio.Conn) {
	ctx := s.GetSocketCtx(c)
	if ctx == nil {
		return
	}
	logger.Infof("GamePing:%v", ctx.UserId)
	c.Emit("game_pong")
}
