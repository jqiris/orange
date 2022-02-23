package mahjong

import (
	"github.com/jqiris/orange/constant"
	"github.com/jqiris/orange/protos"
)

//MjRoomCreate 创建麻将牌桌
func (m *MahjongServer) MjRoomCreate(req *protos.MsgRequest) *protos.MsgResponse {
	resp := &protos.MsgResponse{Code: constant.StatusOk}

	return resp
}
