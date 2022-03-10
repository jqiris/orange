package hall

import (
	"github.com/jqiris/kungfu/v2/discover"
	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/orange/protos"
)

func (s *ServerHall) ServerMaintain(req *protos.InnerMaintainReq) {
	logger.Warnf("ServerMaintain notice:%+v", req)
	serverId, reqState := req.ServerId, req.ReqState
	server := s.Server
	if serverId != server.ServerId {
		logger.Errorf("maintain server not current server,current:%v,req:%v", server.ServerId, serverId)
		return
	}
	if reqState == 1 && !server.Maintained { //进入维护
		//删除服务负载量
		server.Maintained = true
		err := discover.Register(server)
		if err != nil {
			logger.Error(err)
		}
		logger.Warnf("服务器:%v,进入维护状态", serverId)
	} else if reqState == 2 && server.Maintained { //解除维护
		//设置服务负载量
		server.Maintained = false
		err := discover.Register(server)
		if err != nil {
			logger.Error(err)
		}
		logger.Warnf("服务器:%v,解除维护状态", serverId)
	}
}
