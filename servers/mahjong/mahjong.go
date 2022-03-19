package mahjong

import (
	"net/http"
	"sync"
	"time"

	"go.uber.org/atomic"

	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"github.com/jqiris/kungfu/v2/launch"
	"github.com/jqiris/kungfu/v2/plugin"
	"github.com/jqiris/kungfu/v2/rpc"
	"github.com/jqiris/kungfu/v2/treaty"
	"github.com/jqiris/orange/constant"
	"github.com/jqiris/orange/model"
	"github.com/jqiris/orange/protos"
)

var (
	roomMgr  *RoomMgr
	tokenMgr *TokenMgr
	userMgr  *UserMgr
	gameMgrs sync.Map
)

func init() {
	//inti mgr
	roomMgr = NewRoomMgr()
	tokenMgr = NewTokenMgr()
	userMgr = NewUserMgr()
}

type GameMahjong interface {
	SetReady(userId int64)
	HasBegan(roomId string) bool
	DissolveRequest(roomId string, userId int64) *model.MjRoom
	DissolveAgree(roomId string, userId int64, agree bool) *model.MjRoom
	DoDissolve(roomId string)
	HuanSanZhang(userId int64, p1, p2, p3 int32)
	DingQue(userId int64, que int32)
	ChuPai(userId int64, pai int32)
	Peng(userId int64)
	Gang(userId int64, pai int32)
	Hu(userId int64)
	Guo(userId int64)
}

func RegGameMgr(typ string, gameMgr GameMahjong) {
	gameMgrs.Store(typ, gameMgr)
}

func GetGameMgr(typ string) GameMahjong {
	if v, ok := gameMgrs.Load(typ); ok {
		if res, okv := v.(GameMahjong); okv {
			return res
		}
	}
	return nil
}

type ServerMahjong struct {
	*rpc.ServerBase
	totalConn *atomic.Int64
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

func ServerMahjongCreator(s *treaty.Server) (rpc.ServerEntity, error) {
	server := &ServerMahjong{
		ServerBase: rpc.NewServerBase(s),
		totalConn:  atomic.NewInt64(0),
	}
	//socket plugin
	plug := socketCreator()
	server.SocketRouter(plug)
	//reg plugin
	server.AddPlugin(plug)
	//reg inner handler
	server.Register(int32(protos.InnerMsgId_InnerMsgCreateRoom), server.CreateRoom)
	server.Register(int32(protos.InnerMsgId_InnerMsgEnterRoom), server.EnterRoom)
	server.Register(int32(protos.InnerMsgId_InnerMsgMaintain), server.ServerMaintain)
	return server, nil
}

func init() {
	launch.RegisterCreator(constant.MahjongServer, ServerMahjongCreator)
}
