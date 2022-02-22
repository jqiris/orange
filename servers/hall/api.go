package hall

import (
	"fmt"
	"time"

	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/rpc"
	"github.com/jqiris/kungfu/v2/utils"
	"github.com/jqiris/orange/constant"
	"github.com/jqiris/orange/memdb"
	"github.com/jqiris/orange/model"
	"github.com/jqiris/orange/protos"
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
	if dbServer := h.Rpc.Find(constant.DbServer, uid); dbServer != nil {
		reqBuilder := rpc.NewReqBuilder(dbServer).SetCodeType(rpc.CodeTypeJson).SetSuffix(rpc.JsonSuffix).SetReq(&model.DbUpdateMember{
			UserId: uid,
			Updates: map[string]interface{}{
				"login_time": nowTime,
			},
		}).Build()
		if err := h.Rpc.Publish(reqBuilder); err != nil {
			logger.Error(err)
		}
	} else {
		logger.Errorf("hall login 找不到数据库服务器 uid:%v", uid)
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
	return msg
}
