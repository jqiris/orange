package db

import (
	"github.com/jqiris/kungfu/v2/base"
	"github.com/jqiris/kungfu/v2/launch"
	"github.com/jqiris/kungfu/v2/rpc"
	"github.com/jqiris/kungfu/v2/treaty"
	"github.com/jqiris/orange/constant"
)

//异步数据库操作服务,同步服务请直接用database库
type DbServer struct {
	*base.ServerBase
}

func DbServerCreator(s *treaty.Server) (rpc.ServerEntity, error) {
	server := &DbServer{
		ServerBase: base.NewServerBase(s),
	}
	//reg inner handler
	server.Register(constant.DbMsgIdUpdateMember, server.UpdateMember)
	return server, nil
}

func init() {
	launch.RegisterCreator(constant.DbServer, DbServerCreator)
}
