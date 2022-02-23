package db

import (
	"github.com/jqiris/kungfu/v2/base"
	"github.com/jqiris/kungfu/v2/launch"
	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/rpc"
	"github.com/jqiris/kungfu/v2/treaty"
	"github.com/jqiris/orange/constant"
)

//异步数据库操作服务,同步服务请直接用database库
type DbServer struct {
	*base.ServerBase
}

func (d *DbServer) HandleSelfEvent(req *rpc.MsgRpc) []byte {
	resp, err := d.DealMsg(rpc.CodeTypeJson, d.Rpc, req)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return resp
}

func (d *DbServer) HandleBroadcastEvent(req *rpc.MsgRpc) []byte {
	return nil
}

func DbServerCreator(s *treaty.Server) (rpc.ServerEntity, error) {
	server := &DbServer{
		ServerBase: base.NewServerBase(s),
	}
	server.SelfEventHandler = server.HandleSelfEvent
	server.BroadcastEventHandler = server.HandleBroadcastEvent
	//msg handler register
	server.Register(constant.DbMsgIdUpdateMember, server.UpdateMember)
	return server, nil
}

func init() {
	launch.RegisterCreator(constant.DbServer, DbServerCreator)
}
