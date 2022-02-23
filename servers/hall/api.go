package hall

import (
	"fmt"
	"time"

	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/utils"
	"github.com/jqiris/orange/constant"
	"github.com/jqiris/orange/memdb"
	"github.com/jqiris/orange/protos"
	"github.com/jqiris/orange/servers/db"
	"github.com/spf13/viper"
)

func (h *HallServer) ChanResp(req *protos.MsgRequest) *protos.MsgResponse {
	logger.Infof("hall chanResp received: %v", req)
	msg := &protos.MsgResponse{Code: constant.StatusOk}
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
	msg := &protos.MsgResponse{Code: constant.StatusOk}
	tokenKey := viper.GetString("frame.connector.token_key")
	token := utils.Md5(fmt.Sprintf("%d|%s|%s", req.Uid, req.Nickname, tokenKey))
	if req.Token != token {
		msg.Code = constant.StatusError
		msg.Msg = "token不正确"
		return msg
	}
	uid := int(req.Uid)
	//查询用户数据库信息
	user, err := memdb.GetUserMember(uid)
	if err != nil {
		logger.Error(err)
		msg.Code = constant.StatusError
		msg.Msg = "无法找到用户信息"
		return msg
	}
	//异步更新用户登录时间
	nowTime := time.Now()
	if err := db.UpdateMember(h.Rpc, uid, map[string]interface{}{
		"login_time":    nowTime,
		"offline_state": 0,
	}); err != nil {
		logger.Error(err)
	}
	//返回前端数据
	msg.Any = &protos.MsgResponse_LoginResponse{
		LoginResponse: &protos.LoginResponse{
			Uid:       int64(user.UserID),
			Nickname:  user.Nickname,
			Avatar:    user.Avatar,
			Gender:    int32(user.Gender),
			RegTime:   user.RegTime.Unix(),
			LoginTime: nowTime.Unix(),
			IsNew:     user.IsNew,
		},
	}
	return msg
}

func (h *HallServer) Logout(req *protos.LogoutRequest) *protos.MsgResponse {
	msg := &protos.MsgResponse{Code: constant.StatusOk}
	return msg
}

func (h *HallServer) Offline(req *protos.OfflineRequest) *protos.MsgResponse {
	msg := &protos.MsgResponse{Code: constant.StatusOk}
	uid := int(req.Uid)
	//更新用户离线状态
	if err := db.UpdateMember(h.Rpc, uid, map[string]interface{}{
		"offline_state": 1,
		"offline_time":  time.Now(),
	}); err != nil {
		logger.Error(err)
	}
	//清除用户服务记录缓存
	h.Rpc.RemoveFindCache(uid)
	return msg
}
