package loader

import (
	"github.com/jqiris/kungfu/v2/base"
	"github.com/jqiris/kungfu/v2/launch"
	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/packet/nano"
	"github.com/jqiris/kungfu/v2/rpc"
	"github.com/jqiris/kungfu/v2/tcpface"
	"github.com/jqiris/kungfu/v2/treaty"
)

//前端负载器
type LoaderServer struct {
	*base.ServerConnector
}

func (s *LoaderServer) HandleSelfEvent(req *rpc.MsgRpc) []byte {
	logger.Infof("world server handleSelfEvent:%+v", req)
	return nil
}

func (s *LoaderServer) HandleBroadcastEvent(req *rpc.MsgRpc) []byte {
	logger.Infof("world server HandleBroadcastEvent:%+v", req)
	return nil
}

func LoaderServerCreator(s *treaty.Server) (rpc.ServerEntity, error) {
	server := &LoaderServer{
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
	launch.RegisterCreator("loader", LoaderServerCreator)
}
