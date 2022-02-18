package hall

import (
	"fmt"

	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/utils"
	"github.com/jqiris/orange/common"
	"github.com/jqiris/orange/protos"
	"github.com/spf13/viper"
)

func (h *HallServer) ChanResp(req *protos.MsgRequest) *protos.MsgResponse {
	logger.Infof("hall chanResp received: %v", req)
	msg := &protos.MsgResponse{Code: common.StatusOk}
	switch req.MsgId {
	case protos.MsgId_MsgLogin:
		msg = h.Login(req.GetLoginRequest())
	case protos.MsgId_MsgLogout:
		msg = h.Logout(req.GetLogoutRequest())
	case protos.MsgId_MsgOffline:
		msg = h.Offline(req.GetOfflineRequest())
	}
	return msg
}

func (h *HallServer) Login(req *protos.LoginRequest) *protos.MsgResponse {
	msg := &protos.MsgResponse{Code: common.StatusOk}
	tokenKey := viper.GetString("frame.connector.token_key")
	token := utils.Md5(fmt.Sprintf("%d|%s|%s", req.Uid, req.Nickname, tokenKey))
	if req.Token != token {
		msg.Code = common.StatusError
		msg.Msg = "token不正确"
		return msg
	}
	return msg
}

func (h *HallServer) Logout(req *protos.LogoutRequest) *protos.MsgResponse {
	msg := &protos.MsgResponse{Code: common.StatusOk}
	return msg
}

func (h *HallServer) Offline(req *protos.OfflineRequest) *protos.MsgResponse {
	msg := &protos.MsgResponse{Code: common.StatusOk}
	return msg
}
