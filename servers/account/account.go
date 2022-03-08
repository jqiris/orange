package account

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jqiris/kungfu/v2/launch"
	"github.com/jqiris/kungfu/v2/plugin"
	"github.com/jqiris/kungfu/v2/rpc"
	"github.com/jqiris/kungfu/v2/treaty"
	"github.com/jqiris/orange/constant"
)

type Msg struct {
	Errcode int    `json:"errrcode"`
	ErrMsg  string `json:"errmsg"`
}

type ServerAccount struct {
	*rpc.ServerBase
}

func ServerAccountCreator(s *treaty.Server) (rpc.ServerEntity, error) {
	//server entity
	server := &ServerAccount{
		ServerBase: rpc.NewServerBase(s),
	}
	//http handler
	app := gin.Default()
	app.Use(cors.Default())
	server.Router(app)
	//reg plugin
	plug := plugin.NewServerHttp(app)
	server.AddPlugin(plug)
	return server, nil
}

func init() {
	launch.RegisterCreator(constant.AccountServer, ServerAccountCreator)
}

// Router router
func (s *ServerAccount) Router(app *gin.Engine) {
	app.GET("/register", s.Register)
	app.GET("/get_version", s.GetVersion)
	app.GET("/get_serverinfo", s.GetServerinfo)
	app.GET("/guest", s.Guest)
	app.GET("/auth", s.Auth)
	app.GET("/base_info", s.BaseInfo)
}
