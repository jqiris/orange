package gate

import (
	"github.com/jqiris/kungfu/v2/discover"
	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/rpc"
	"github.com/jqiris/kungfu/v2/session"
	"github.com/jqiris/orange/common"
	"github.com/jqiris/orange/protos"
)

func (g *GateServer) ChanReq(s *session.Session, req *protos.MsgRequest) error {
	logger.Infof("gate chanReq uid:%v, received: %v", s.UID(), req)
	resp := &protos.MsgResponse{Code: common.StatusOk}
	hall := discover.GetServerByType("hall", s.RemoteAddr().String())
	if hall == nil {
		resp.Code = common.StatusError
		resp.Msg = "找不到大厅服务器"
		return s.Response(resp)
	}
	reqForward := rpc.NewReqBuilder(hall).SetReq(req).SetResp(resp).SetMsgId(int32(protos.MsgId_MsgChan)).Build()
	if err := g.Rpc.Request(reqForward); err != nil {
		resp.Code = common.StatusError
		resp.Msg = err.Error()
		return s.Response(resp)
	}
	//成功处理绑定关系
	if resp.Code == common.StatusOk {
		switch req.MsgId {
		case protos.MsgId_MsgLogin:
			s.Bind(req.GetLoginRequest().Uid)
		case protos.MsgId_MsgLogout:
			s.Close()
		}
	}
	return s.Response(resp)
}
