package mahjong

import (
	"github.com/jqiris/kungfu/v2/base"
	"github.com/jqiris/kungfu/v2/launch"
	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/rpc"
	"github.com/jqiris/kungfu/v2/treaty"
	"github.com/jqiris/orange/constant"
	"github.com/jqiris/orange/protos"
)

type MahjongServer struct {
	*base.ServerBase
}

func (m *MahjongServer) HandleSelfEvent(req *rpc.MsgRpc) []byte {
	resp, err := m.DealMsg(rpc.CodeTypeProto, m.Rpc, req)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return resp
}

func (m *MahjongServer) HandleBroadcastEvent(req *rpc.MsgRpc) []byte {
	return nil
}

func MahjongServerCreator(s *treaty.Server) (rpc.ServerEntity, error) {
	server := &MahjongServer{
		ServerBase: base.NewServerBase(s),
	}
	server.SelfEventHandler = server.HandleSelfEvent
	server.BroadcastEventHandler = server.HandleBroadcastEvent
	//msg handler register
	server.Register(int32(protos.MsgId_MsgMjRoomCreate), server.MjRoomCreate)
	return server, nil
}

func init() {
	launch.RegisterCreator(constant.MahjongServer, MahjongServerCreator)
}
