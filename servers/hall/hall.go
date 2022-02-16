package hall

import (
	"github.com/jqiris/kungfu/v2/base"
	"github.com/jqiris/kungfu/v2/launch"
	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/rpc"
	"github.com/jqiris/kungfu/v2/treaty"
	"github.com/jqiris/orange/protos"
)

type HallServer struct {
	*base.ServerBase
	handler *rpc.Handler
}

func (h *HallServer) HandleSelfEvent(req *rpc.MsgRpc) []byte {
	resp, err := h.handler.DealMsg(rpc.CodeTypeProto, h.Rpc, req)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return resp
}
func (h *HallServer) HandleBroadcastEvent(req *rpc.MsgRpc) []byte {
	return nil
}

func HallServerCreator(s *treaty.Server) (rpc.ServerEntity, error) {
	server := &HallServer{
		ServerBase: base.NewServerBase(s),
	}
	server.SelfEventHandler = server.HandleSelfEvent
	server.BroadcastEventHandler = server.HandleBroadcastEvent
	//handle instance
	handler := rpc.NewHandler()
	handler.Register(int32(protos.MsgId_MsgChan), server.ChanResp)
	server.handler = handler
	return server, nil
}

func init() {
	launch.RegisterCreator("hall", HallServerCreator)
}
