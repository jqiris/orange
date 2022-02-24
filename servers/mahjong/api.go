package mahjong

import (
	"errors"
	"time"

	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/orange/constant"
	"github.com/jqiris/orange/database"
	"github.com/jqiris/orange/model"
	"github.com/jqiris/orange/protos"
	"github.com/jqiris/orange/tools"
)

//MjRoomCreate 创建麻将牌桌
func (m *MahjongServer) MjRoomCreate(req *protos.MjRoomCreateRequest) *protos.MsgResponse {
	resp := &protos.MsgResponse{Code: constant.StatusOk}
	//生成房号
	roomCode, err := m.createRoomCode(0)
	if err != nil {
		logger.Error(err)
		resp.Code = constant.StatusError
		resp.Msg = constant.ErrServerBusy
		return resp
	}
	//创建房间
	sid := m.Rpc.GetServer().ServerId
	mahjongRoom := &model.GameMahjongRooms{
		Code:       roomCode,
		Sid:        sid,
		Gid:        int(req.Gid),
		Tid:        0,
		Landlord:   int(req.Uid),
		State:      0,
		Round:      0,
		CreateTime: time.Now(),
	}
	if err := database.CreateMahjongRoom(mahjongRoom); err != nil {
		logger.Error(err)
		resp.Code = constant.StatusError
		resp.Msg = constant.ErrServerBusy
		return resp
	}
	resp.Any = &protos.MsgResponse_MjRoomCreateResponse{
		MjRoomCreateResponse: &protos.MjRoomCreateResponse{
			Code: int64(roomCode),
			Sid:  sid,
			Tid:  0,
		},
	}
	return resp
}

func (m *MahjongServer) createRoomCode(max int) (int, error) {
	roomCode := tools.CreateRoomCode(constant.MahjongRoomCode, constant.MahjongRoomCodeLow, constant.MahjongRoomCodeHigh)
	//同步检查房号是否已经存在
	if _, err := database.GetMahjongRoom(roomCode); err == nil {
		if max < 5 {
			time.Sleep(time.Second)
			return m.createRoomCode(max + 1)
		}
		return roomCode, errors.New("mjRoomCreate createRoomCode repeated than max times")
	}
	return roomCode, nil
}
