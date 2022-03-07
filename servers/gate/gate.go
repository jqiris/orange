package gate

import (
	"github.com/jqiris/kungfu/v2/launch"
	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/packet/nano"
	"github.com/jqiris/kungfu/v2/plugin"
	"github.com/jqiris/kungfu/v2/rpc"
	"github.com/jqiris/kungfu/v2/tcpface"
	"github.com/jqiris/kungfu/v2/treaty"
	"github.com/jqiris/orange/constant"
)

//前端负载器
type GateServer struct {
	*rpc.ServerBase
}

func GateServerCreator(s *treaty.Server) (rpc.ServerEntity, error) {
	server := &GateServer{
		ServerBase: rpc.NewServerBase(s),
	}
	//reg plugin
	plug := plugin.NewServerConnector()
	plug.RouteHandler = func(s tcpface.IServer) {
		rs := s.GetMsgHandler()
		router := rs.(*nano.MsgHandle)
		err := router.Register(server)
		if err != nil {
			logger.Fatal(err)
		}
	}
	server.AddPlugin(plug)
	return server, nil
}

func init() {
	launch.RegisterCreator(constant.GateServer, GateServerCreator)
}
