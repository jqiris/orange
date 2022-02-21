package gate

import (
	"github.com/jqiris/kungfu/v2/discover"
	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/rpc"
	"github.com/jqiris/kungfu/v2/session"
	"github.com/jqiris/kungfu/v2/treaty"
	"github.com/jqiris/orange/constant"
	"github.com/jqiris/orange/protos"
)

func (g *GateServer) AfterInit() {
	g.ServerConnector.AfterInit()
	session.Lifetime.OnClosed(func(s *session.Session) {
		uid, addr := s.UID(), s.RemoteAddr()
		logger.Infof("gate offline,uid:%v,ip:%v", uid, addr)
		if uid > 0 {
			//通知大厅离线
			hall := s.Value("hall").(*treaty.Server)
			req := &protos.MsgRequest{
				MsgId: protos.MsgId_MsgOffline,
				Any: &protos.MsgRequest_OfflineRequest{
					OfflineRequest: &protos.OfflineRequest{
						Uid: uid,
					},
				},
			}
			resp := &protos.MsgResponse{Code: constant.StatusOk}
			reqForward := rpc.NewReqBuilder(hall).SetReq(req).SetResp(resp).SetMsgId(int32(protos.MsgId_MsgChan)).Build()
			if err := g.Rpc.Request(reqForward); err != nil {
				logger.Infof("gate offline uid:%v,err:%v", uid, err)
			} else {
				logger.Infof("gate offline uid:%v,resp:%+v", uid, resp)
			}
		}
	})
}

func (g *GateServer) ChanReq(s *session.Session, req *protos.MsgRequest) error {
	logger.Infof("gate chanReq uid:%v, received: %v", s.UID(), req)
	resp := &protos.MsgResponse{Code: constant.StatusOk}
	hall := discover.GetServerByType("hall", s.RemoteAddr().String())
	if hall == nil {
		resp.Code = constant.StatusError
		resp.Msg = "找不到大厅服务器"
		return s.Response(resp)
	}
	reqForward := rpc.NewReqBuilder(hall).SetReq(req).SetResp(resp).SetMsgId(int32(protos.MsgId_MsgChan)).Build()
	if err := g.Rpc.Request(reqForward); err != nil {
		resp.Code = constant.StatusError
		resp.Msg = err.Error()
		return s.Response(resp)
	}
	//成功处理绑定关系
	if resp.Code == constant.StatusOk {
		switch req.MsgId {
		case protos.MsgId_MsgLogin:
			s.Bind(req.GetLoginRequest().Uid)
			s.Set("hall", hall) //设置用户的大厅信息
		case protos.MsgId_MsgLogout:
			s.Close()
		}
	}
	return s.Response(resp)
}
