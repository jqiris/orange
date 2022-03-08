package db

import (
	"github.com/jqiris/kungfu/v2/launch"
	"github.com/jqiris/kungfu/v2/rpc"
	"github.com/jqiris/kungfu/v2/treaty"
	"github.com/jqiris/orange/constant"
)

// ServerDb 异步数据库操作服务,同步服务请直接用database库
type ServerDb struct {
	*rpc.ServerBase
}

func ServerDbCreator(s *treaty.Server) (rpc.ServerEntity, error) {
	server := &ServerDb{
		ServerBase: rpc.NewServerBase(s),
	}
	//reg inner handler
	server.Register(constant.DbMsgIdUpdateMember, server.UpdateMember)
	return server, nil
}

func init() {
	launch.RegisterCreator(constant.DbServer, ServerDbCreator)
}
