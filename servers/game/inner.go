package game

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jqiris/kungfu/v2/utils"
	"github.com/jqiris/orange/constant"
	"github.com/jqiris/orange/model"
	"github.com/jqiris/orange/protos"
)

// //router
// func (s *GameServer) HttpRouter(app *gin.Engine) {
// 	app.GET("/get_server_info", s.GetServerInfo)
// 	app.GET("/create_room", s.CreateRoom)
// 	app.GET("/enter_room", s.EnterRoom)
// 	app.GET("/is_room_runing", s.IsRoomRuning)
// }

//向大厅服定时心跳
// func (s *GameServer) update() {
// 	ticker := time.NewTicker(time.Duration(constant.HTTP_TICK_TIME) * time.Millisecond)
// 	for {
// 		select {
// 		case <-ticker.C:
// 			s.gameServerInfo["load"] = []string{utils.IntToString(roomMgr.getTotallRooms())}
// 			url := fmt.Sprintf("http://%v:%v/register_gs?%v", constant.HALL_IP, constant.HALL_PORT, s.gameServerInfo.Encode())
// 			resp, err := http.Get(url)
// 			if err != nil {
// 				logger.Error(err)
// 				continue
// 			}
// 			logger.Infof("game http heart:%v", resp.StatusCode)
// 		}
// 	}
// }

func (s *GameServer) Error(c *gin.Context, msg string, args ...int) {
	errcode := 1
	if len(args) > 0 {
		errcode = args[0]
	}
	c.JSON(200, map[string]any{"errcode": errcode, "errmsg": msg})
}

func (s *GameServer) Success(c *gin.Context, data map[string]any) {
	data["errcode"] = 0
	data["errmsg"] = "ok"
	c.JSON(200, data)
}

// //获取服务信息
// func (s *GameServer) GetServerInfo(c *gin.Context) {
// 	serverId, sign := c.Query("serverid"), c.Query("sign")
// 	if serverId != constant.SERVER_ID || len(sign) < 1 {
// 		s.Error(c, "invalid parameters")
// 		return
// 	}
// 	md5 := utils.Md5(serverId + constant.ROOM_PRI_KEY)
// 	if md5 != sign {
// 		s.Error(c, "sign check failed")
// 		return
// 	}
// 	locations := roomMgr.getUserLocations()
// 	var arr []int
// 	for userId, location := range locations {
// 		arr = append(arr, int(userId), int(location.RoomId))
// 	}
// 	s.Success(c, map[string]any{"userroominfo": arr})
// }

//创建房间
func (s *GameServer) CreateRoom(req *protos.InnerCreateRoomReq) *protos.InnerMsgResp {
	resp := &protos.InnerMsgResp{Errcode: 1, Errmsg: "服务器忙"}
	userId, gems, sign, conf := int64(req.Userid), int32(req.Gems), req.Sign, req.Conf
	if userId < 1 || len(sign) < 1 || len(conf) < 1 {
		resp.Errmsg = "Invalid parameters"
		return resp
	}
	md5 := utils.Md5(fmt.Sprintf("%v%v%v%v", userId, conf, gems, constant.ROOM_PRI_KEY))
	if md5 != sign {
		resp.Errmsg = "sign check failed"
		return resp
	}
	var cfg model.GameConf
	err := json.Unmarshal([]byte(conf), &cfg)
	if err != nil {
		resp.Errmsg = err.Error()
		return resp
	}
	serverId, ip, port := s.Server.ServerId, s.Server.ServerIp, int32(s.Server.ClientPort)
	errcode, roomId := roomMgr.createRoom(userId, cfg, gems, serverId, ip, port)
	if errcode != 0 || roomId < 1 {
		resp.Errmsg = "create failed"
		return resp
	}
	resp.Errcode = 0
	resp.Errmsg = "ok"
	resp.Any = &protos.InnerMsgResp_CreateRoom{
		CreateRoom: &protos.InnerCreateRoomResp{
			Roomid:   int32(roomId),
			ServerId: s.Server.ServerId,
		},
	}
	return resp
}

//进入房间
func (s *GameServer) EnterRoom(req *protos.InnerEnterRoomReq) *protos.InnerMsgResp {
	resp := &protos.InnerMsgResp{Errcode: 1, Errmsg: "服务器忙"}
	userId, roomId, name, sign := int64(req.Userid), int32(req.Roomid), req.Name, req.Sign
	if userId < 1 || roomId < 1 || len(sign) < 1 {
		resp.Errmsg = "Invalid parameters"
		return resp
	}
	md5 := utils.Md5(fmt.Sprintf("%v%v%v%v", userId, name, roomId, constant.ROOM_PRI_KEY))
	if md5 != sign {
		resp.Errmsg = "sign check failed"
		return resp
	}
	ret := roomMgr.enterRoom(roomId, userId, name)
	if ret != 0 {
		switch ret {
		case 1:
			resp.Errmsg = "room is full"
			return resp
		case 2:
			resp.Errmsg = "can't find room"
			return resp
		case 3:
			resp.Errmsg = "build room from db failed"
			return resp
		}
	}
	token := tokenMgr.createToken(userId, 5000)
	resp.Errcode = 0
	resp.Errmsg = "ok"
	resp.Any = &protos.InnerMsgResp_EnterRoom{
		EnterRoom: &protos.InnerEnterRoomResp{
			Token: token,
		},
	}
	return resp
}

//房间是否运行
func (s *GameServer) IsRoomRuning(c *gin.Context) {
	roomId, sign := utils.StringToInt(c.Query("roomid")), c.Query("sign")
	if roomId < 1 || len(sign) < 1 {
		s.Error(c, "invalid parameters")
		return
	}
	md5 := utils.Md5(fmt.Sprintf("%v%v", roomId, constant.ROOM_PRI_KEY))
	if md5 != sign {
		s.Error(c, "sign check failed", 2)
		return
	}
	s.Success(c, map[string]any{"runing": true})
}
