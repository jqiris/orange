package mahjong

import (
	"github.com/jqiris/kungfu/v2/base"
	"github.com/jqiris/kungfu/v2/launch"
	"github.com/jqiris/kungfu/v2/rpc"
	"github.com/jqiris/kungfu/v2/treaty"
	"github.com/jqiris/orange/constant"
	"github.com/jqiris/orange/protos"
)

type MahjongServer struct {
	*base.ServerBase
}

func MahjongServerCreator(s *treaty.Server) (rpc.ServerEntity, error) {
	server := &MahjongServer{
		ServerBase: base.NewServerBase(s),
	}
	//reg inner handler
	server.Register(int32(protos.MsgId_MsgMjRoomCreate), server.MjRoomCreate)
	return server, nil
}

func init() {
	launch.RegisterCreator(constant.MahjongServer, MahjongServerCreator)
}
