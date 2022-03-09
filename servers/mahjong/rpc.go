package mahjong

import (
	"encoding/json"
	"fmt"
	"github.com/jqiris/kungfu/v2/discover"

	"github.com/gin-gonic/gin"
	"github.com/jqiris/kungfu/v2/rpc"
	"github.com/jqiris/kungfu/v2/treaty"
	"github.com/jqiris/kungfu/v2/utils"
	"github.com/jqiris/orange/constant"
	"github.com/jqiris/orange/model"
	"github.com/jqiris/orange/protos"
)

var (
	builder = rpc.DefaultReqBuilder().SetCodeType(rpc.CodeTypeProto).SetSuffix(rpc.DefaultSuffix)
)

// CreatePrivateRoom 外部调用
func CreatePrivateRoom(s rpc.ServerRpc, req *protos.InnerCreateRoomReq) *protos.InnerMsgResp {
	resp := &protos.InnerMsgResp{Errcode: constant.StatusOk}
	server := discover.GetServerByTypeLoad(constant.MahjongServer) //根据最小负载选择服务器
	if server == nil {
		resp.Errcode = constant.StatusError
		resp.Errmsg = constant.ErrNotFoundMahjongServer
		return resp
	}
	b := builder.Build()
	if err := s.Request(b.SetMsgId(int32(protos.InnerMsgId_InnerMsgCreateRoom)).SetServer(server).SetReq(req).SetResp(resp).Build()); err != nil {
		resp.Errcode = constant.StatusError
		resp.Errmsg = err.Error()
		return resp
	}
	return resp
}

func EnterPrivateRoom(s rpc.ServerRpc, req *protos.InnerEnterRoomReq, server *treaty.Server) *protos.InnerMsgResp {
	resp := &protos.InnerMsgResp{Errcode: constant.StatusOk}
	b := builder.Build()
	if err := s.Request(b.SetMsgId(int32(protos.InnerMsgId_InnerMsgEnterRoom)).SetServer(server).SetReq(req).SetResp(resp).Build()); err != nil {
		resp.Errcode = constant.StatusError
		resp.Errmsg = err.Error()
		return resp
	}
	return resp
}

//内部处理
func (s *ServerMahjong) Error(c *gin.Context, msg string, args ...int) {
	errcode := 1
	if len(args) > 0 {
		errcode = args[0]
	}
	c.JSON(200, map[string]any{"errcode": errcode, "errmsg": msg})
}

func (s *ServerMahjong) Success(c *gin.Context, data map[string]any) {
	data["errcode"] = 0
	data["errmsg"] = "ok"
	c.JSON(200, data)
}

// CreateRoom 创建房间
func (s *ServerMahjong) CreateRoom(req *protos.InnerCreateRoomReq) *protos.InnerMsgResp {
	resp := &protos.InnerMsgResp{Errcode: 1, Errmsg: "服务器忙"}
	userId, gems, sign, conf := int64(req.UserId), int32(req.Gems), req.Sign, req.Conf
	if userId < 1 || len(sign) < 1 || len(conf) < 1 {
		resp.Errmsg = "Invalid parameters"
		return resp
	}
	md5 := utils.Md5(fmt.Sprintf("%v%v%v%v", userId, conf, gems, constant.RoomPriKey))
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
	if errcode != 0 || len(roomId) < 1 {
		resp.Errmsg = "create failed"
		return resp
	}
	resp.Errcode = 0
	resp.Errmsg = "ok"
	resp.Any = &protos.InnerMsgResp_CreateRoom{
		CreateRoom: &protos.InnerCreateRoomResp{
			RoomId:   roomId,
			ServerId: s.Server.ServerId,
		},
	}
	return resp
}

// EnterRoom 进入房间
func (s *ServerMahjong) EnterRoom(req *protos.InnerEnterRoomReq) *protos.InnerMsgResp {
	resp := &protos.InnerMsgResp{Errcode: 1, Errmsg: "服务器忙"}
	userId, roomId, name, sign := int64(req.UserId), req.RoomId, req.Name, req.Sign
	if userId < 1 || len(roomId) < 1 || len(sign) < 1 {
		resp.Errmsg = "Invalid parameters"
		return resp
	}
	md5 := utils.Md5(fmt.Sprintf("%v%v%v%v", userId, name, roomId, constant.RoomPriKey))
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
