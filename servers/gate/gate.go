package gate

import (
	"github.com/jqiris/kungfu/v2/base"
	"github.com/jqiris/kungfu/v2/launch"
	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/packet/nano"
	"github.com/jqiris/kungfu/v2/rpc"
	"github.com/jqiris/kungfu/v2/tcpface"
	"github.com/jqiris/kungfu/v2/treaty"
	"github.com/jqiris/orange/constant"
)

//前端负载器
type GateServer struct {
	*base.ServerConnector
}

func (g *GateServer) HandleSelfEvent(req *rpc.MsgRpc) []byte {
	logger.Infof("gate server handleSelfEvent:%+v", req)
	return nil
}

func (g *GateServer) HandleBroadcastEvent(req *rpc.MsgRpc) []byte {
	logger.Infof("gate server HandleBroadcastEvent:%+v", req)
	return nil
}

func GateServerCreator(s *treaty.Server) (rpc.ServerEntity, error) {
	server := &GateServer{
		ServerConnector: base.NewServerConnector(s),
	}
	server.SelfEventHandler = server.HandleSelfEvent
	server.BroadcastEventHandler = server.HandleBroadcastEvent
	server.RouteHandler = func(s tcpface.IServer) {
		rs := s.GetMsgHandler()
		router := rs.(*nano.MsgHandle)
		err := router.Register(server)
		if err != nil {
			logger.Fatal(err)
		}
	}
	return server, nil
}

func init() {
	launch.RegisterCreator(constant.GateServer, GateServerCreator)
}
