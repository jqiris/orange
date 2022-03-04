package game

import (
	"github.com/jqiris/kungfu/v2/rpc"
	"github.com/jqiris/kungfu/v2/treaty"
	"github.com/jqiris/orange/constant"
	"github.com/jqiris/orange/protos"
)

var (
	builder = rpc.DefaultReqBuilder().SetCodeType(rpc.CodeTypeProto).SetSuffix(rpc.DefaultSuffix)
)

func CreatePrivateRoom(s rpc.ServerRpc, req *protos.InnerCreateRoomReq) *protos.InnerMsgResp {
	resp := &protos.InnerMsgResp{Errcode: constant.StatusOk}
	uid := int(req.Userid)
	server := s.Find(constant.GameServer, uid)
	if server == nil {
		resp.Errcode = constant.StatusError
		resp.Errmsg = constant.ErrNotFoundGameServer
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
