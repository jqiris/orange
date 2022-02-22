package db

import (
	"github.com/jqiris/kungfu/v2/base"
	"github.com/jqiris/kungfu/v2/launch"
	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/rpc"
	"github.com/jqiris/kungfu/v2/treaty"
	"github.com/jqiris/orange/constant"
)

//异步数据库操作服务
type DbServer struct {
	*base.ServerBase
	handler *rpc.Handler
}

func (d *DbServer) HandleSelfEvent(req *rpc.MsgRpc) []byte {
	resp, err := d.handler.DealMsg(rpc.CodeTypeProto, d.Rpc, req)
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
	//handle instance
	handler := rpc.NewHandler()
	handler.Register(constant.DbMsgIdUpdateMember, server.UpdateMember)
	server.handler = handler
	return server, nil
}

func init() {
	launch.RegisterCreator(constant.DbServer, DbServerCreator)
}
