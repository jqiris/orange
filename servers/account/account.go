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

type AccountServer struct {
	*rpc.ServerBase
}

func AccountServerCreator(s *treaty.Server) (rpc.ServerEntity, error) {
	//server entity
	server := &AccountServer{
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
	launch.RegisterCreator(constant.AccountServer, AccountServerCreator)
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
