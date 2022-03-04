package game

import (
	"net/http"
	"net/url"
	"time"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"github.com/jqiris/kungfu/v2/base"
	"github.com/jqiris/kungfu/v2/launch"
	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/rpc"
	"github.com/jqiris/kungfu/v2/treaty"
	"github.com/jqiris/orange/constant"
	"github.com/jqiris/orange/protos"
)

var (
	roomMgr  *RoomMgr
	tokenMgr *TokenMgr
	userMgr  *UserMgr
)

type GameServer struct {
	*base.ServerBase
	gameServerInfo url.Values
}

func (s *GameServer) AfterInit() {
	s.ServerBase.AfterInit()
	//socket handler
	allowOrigin := func(r *http.Request) bool {
		return true
	}
	sc := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				Client: &http.Client{
					Timeout: time.Minute,
				},
				CheckOrigin: allowOrigin,
			},
			&websocket.Transport{
				CheckOrigin: allowOrigin,
			},
		},
	},
	)
	//sockt router
	go s.SocketRouter(sc)

	//服务信息

	// s.gameServerInfo = url.Values{
	// 	"id":         {constant.SERVER_ID},
	// 	"clientip":   {constant.CLIENT_IP},
	// 	"clientport": {utils.IntToString(constant.CLIENT_PORT)},
	// 	"httpPort":   {utils.IntToString(constant.HTTP_PORT)},
	// 	"load":       {utils.IntToString(roomMgr.getTotallRooms())},
	// }
	// go s.update()
}

func (s *GameServer) HandleSelfEvent(req *rpc.MsgRpc) []byte {
	resp, err := s.DealMsg(rpc.CodeTypeProto, s.Rpc, req)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return resp
}

func (s *GameServer) HandleBroadcastEvent(req *rpc.MsgRpc) []byte {
	logger.Infof("game server HandleBroadcastEvent:%+v", req)
	return nil
}
func GameServerCreator(s *treaty.Server) (rpc.ServerEntity, error) {
	logger.Infof("world server creat:%+v", s)
	// //http handler
	// app := gin.Default()
	//server entity
	server := &GameServer{
		ServerBase: base.NewServerBase(s),
	}
	server.SelfEventHandler = server.HandleSelfEvent
	server.BroadcastEventHandler = server.HandleBroadcastEvent
	//reg init handler
	server.Register(int32(protos.InnerMsgId_InnerMsgCreateRoom), server.CreateRoom)
	server.Register(int32(protos.InnerMsgId_InnerMsgEnterRoom), server.EnterRoom)
	//cors
	// app.Use(cors.Default())
	//http router
	// server.HttpRouter(app)
	return server, nil
}

func init() {
	launch.RegisterCreator(constant.GameServer, GameServerCreator)

	//inti mgr
	roomMgr = NewRoomMgr()
	tokenMgr = NewTokenMgr()
	userMgr = NewUserMgr()
}
