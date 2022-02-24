package mahjong

import (
	"github.com/jqiris/kungfu/v2/rpc"
	"github.com/jqiris/orange/constant"
	"github.com/jqiris/orange/protos"
)

var (
	builder = rpc.DefaultReqBuilder().SetCodeType(rpc.CodeTypeProto).SetSuffix(rpc.DefaultSuffix)
)

func MjRoomCreate(s rpc.ServerRpc, req *protos.MjRoomCreateRequest) *protos.MsgResponse {
	resp := &protos.MsgResponse{Code: constant.StatusOk}
	uid := int(req.Uid)
	server := s.Find(constant.MahjongServer, uid)
	if server == nil {
		resp.Code = constant.StatusError
		resp.Msg = constant.ErrNotFoundMahjongServer
		return resp
	}
	b := builder.Build()
	if err := s.Request(b.SetMsgId(int32(protos.MsgId_MsgMjRoomCreate)).SetServer(server).SetReq(req).SetResp(resp).Build()); err != nil {
		resp.Code = constant.StatusError
		resp.Msg = err.Error()
		return resp
	}
	return resp
}
