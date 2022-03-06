package game

import (
	"net/http"
	"sync"
	"time"

	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"github.com/jqiris/kungfu/v2/base"
	"github.com/jqiris/kungfu/v2/base/plugin"
	"github.com/jqiris/kungfu/v2/launch"
	"github.com/jqiris/kungfu/v2/rpc"
	"github.com/jqiris/kungfu/v2/treaty"
	"github.com/jqiris/orange/constant"
	"github.com/jqiris/orange/protos"
)

var (
	roomMgr  *RoomMgr
	tokenMgr *TokenMgr
	userMgr  *UserMgr
	gameMgrs sync.Map
)

func RegGameMgr(typ string, gameMgr GameMgr) {
	gameMgrs.Store(typ, gameMgr)
}

func GetGameMgr(typ string) GameMgr {
	if v, ok := gameMgrs.Load(typ); ok {
		if res, okv := v.(GameMgr); okv {
			return res
		}
	}
	return nil
}

type GameServer struct {
	*base.ServerBase
}

func socketCreator() *plugin.ServerSocket {
	allowOrigin := func(r *http.Request) bool {
		return true
	}
	opts := &engineio.Options{
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
	}
	sc := plugin.NewServerSocket("/", opts)
	return sc
}

func GameServerCreator(s *treaty.Server) (rpc.ServerEntity, error) {
	server := &GameServer{
		ServerBase: base.NewServerBase(s),
	}
	//socket plugin
	plug := socketCreator()
	server.SocketRouter(plug)
	//reg plugin
	server.AddPlugin(plug)
	//reg inner handler
	server.Register(int32(protos.InnerMsgId_InnerMsgCreateRoom), server.CreateRoom)
	server.Register(int32(protos.InnerMsgId_InnerMsgEnterRoom), server.EnterRoom)
	return server, nil
}

func init() {
	launch.RegisterCreator(constant.GameServer, GameServerCreator)

	//inti mgr
	roomMgr = NewRoomMgr()
	tokenMgr = NewTokenMgr()
	userMgr = NewUserMgr()
}
