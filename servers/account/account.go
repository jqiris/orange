package account

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jqiris/kungfu/v2/base"
	"github.com/jqiris/kungfu/v2/launch"
	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/rpc"
	"github.com/jqiris/kungfu/v2/treaty"
	"github.com/jqiris/orange/constant"
)

type Msg struct {
	Errcode int    `json:"errrcode"`
	ErrMsg  string `json:"errmsg"`
}

type AccountServer struct {
	*base.ServerHttp
}

func (s *AccountServer) HandleSelfEvent(req *rpc.MsgRpc) []byte {
	logger.Infof("world server handleSelfEvent:%+v", req)
	return nil
}

func (s *AccountServer) HandleBroadcastEvent(req *rpc.MsgRpc) []byte {
	logger.Infof("world server HandleBroadcastEvent:%+v", req)
	return nil
}

func WorldServerCreator(s *treaty.Server) (rpc.ServerEntity, error) {
	logger.Infof("world server creat:%+v", s)
	//http handler
	app := gin.Default()
	//server entity
	server := &AccountServer{
		ServerHttp: base.NewServerHttp(s, app),
	}
	server.SelfEventHandler = server.HandleSelfEvent
	server.BroadcastEventHandler = server.HandleBroadcastEvent
	//cors
	app.Use(cors.Default())
	//router
	server.Router(app)
	return server, nil
}

func init() {
	launch.RegisterCreator(constant.AccountServer, WorldServerCreator)
}

//router
func (s *AccountServer) Router(app *gin.Engine) {
	app.GET("/register", s.Register)
	app.GET("/get_version", s.GetVersion)
	app.GET("/get_serverinfo", s.GetServerinfo)
	app.GET("/guest", s.Guest)
	app.GET("/auth", s.Auth)
	app.GET("/base_info", s.BaseInfo)
}
