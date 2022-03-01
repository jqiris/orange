package game

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/jqiris/kungfu/v2/base"
	"github.com/jqiris/kungfu/v2/launch"
	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/rpc"
	"github.com/jqiris/kungfu/v2/treaty"
	"github.com/jqiris/orange/constant"
)

var (
	roomMgr  *RoomMgr
	tokenMgr *TokenMgr
	userMgr  *UserMgr
)

type GameServer struct {
	*base.ServerHttp
}

func (s *GameServer) HandleSelfEvent(req *rpc.MsgRpc) []byte {
	logger.Infof("game server handleSelfEvent:%+v", req)
	return nil
}

func (s *GameServer) HandleBroadcastEvent(req *rpc.MsgRpc) []byte {
	logger.Infof("game server HandleBroadcastEvent:%+v", req)
	return nil
}
func GameServerCreator(s *treaty.Server) (rpc.ServerEntity, error) {
	logger.Infof("world server creat:%+v", s)
	//http handler
	app := gin.Default()
	//server entity
	server := &GameServer{
		ServerHttp: base.NewServerHttp(s, app),
	}
	server.SelfEventHandler = server.HandleSelfEvent
	server.BroadcastEventHandler = server.HandleBroadcastEvent
	//cors
	app.Use(cors.Default())
	//http router
	server.HttpRouter(app)
	//socket handler
	sc := socketio.NewServer(nil)
	//sockt router
	server.SocketRouter(sc)

	return server, nil
}

func init() {
	launch.RegisterCreator(constant.GameServer, GameServerCreator)
}
