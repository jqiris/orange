package world

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jqiris/kungfu/v2/base"
	"github.com/jqiris/kungfu/v2/launch"
	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/rpc"
	"github.com/jqiris/kungfu/v2/treaty"
)

type WorldServer struct {
	*base.ServerHttp
}

func (s *WorldServer) HandleSelfEvent(req *rpc.MsgRpc) []byte {
	logger.Infof("world server handleSelfEvent:%+v", req)
	return nil
}

func (s *WorldServer) HandleBroadcastEvent(req *rpc.MsgRpc) []byte {
	logger.Infof("world server HandleBroadcastEvent:%+v", req)
	return nil
}

func WorldServerCreator(s *treaty.Server) (rpc.ServerEntity, error) {
	logger.Infof("world server creat:%+v", s)
	//http handler
	app := gin.Default()
	//server entity
	server := &WorldServer{
		ServerHttp: base.NewServerHttp(s, app),
	}
	server.SelfEventHandler = server.HandleSelfEvent
	server.BroadcastEventHandler = server.HandleBroadcastEvent
	//router
	server.Router(app)
	return server, nil
}

func init() {
	launch.RegisterCreator("world", WorldServerCreator)
}

//router
func (s *WorldServer) Router(app *gin.Engine) {
	app.GET("/login", func(c *gin.Context) {
		msg := s.Login(c)
		c.JSON(http.StatusOK, msg)
	})
}
