package mahjong

import (
	"github.com/jqiris/kungfu/v2/logger"
	"sync"

	socketio "github.com/googollee/go-socket.io"
)

type UserMgr struct {
	userList   sync.Map
	userOnline int32
}

func NewUserMgr() *UserMgr {
	return &UserMgr{
		userList:   sync.Map{},
		userOnline: 0,
	}
}

func (m *UserMgr) bind(userId int64, c socketio.Conn) {
	m.userList.Store(userId, c)
	m.userOnline++
}

func (m *UserMgr) isOnline(userId int64) bool {
	_, exist := m.userList.Load(userId)
	return exist
}

func (m *UserMgr) broadcastInRoom(event string, data any, sender int64, args ...bool) {
	includingSender := false
	if len(args) > 0 {
		includingSender = args[0]
	}
	roomId := roomMgr.getUserRoom(sender)
	if len(roomId) < 1 {
		return
	}
	roomInfo := roomMgr.getRoom(roomId)
	if roomInfo == nil {
		return
	}
	for i := 0; i < len(roomInfo.Seats); i++ {
		rs := roomInfo.Seats[i]
		if rs.Userid == sender && !includingSender {
			continue
		}
		if socket := m.get(rs.Userid); socket != nil {
			if data != nil {
				socket.Emit(event, data)
			} else {
				socket.Emit(event)
			}
		}
	}
}

func (m *UserMgr) sendMsg(userId int64, event string, msgdata ...any) {
	socket := m.get(userId)
	if socket == nil {
		return
	}
	socket.Emit(event, msgdata...)
}

func (m *UserMgr) get(userId int64) socketio.Conn {
	if v, ok := m.userList.Load(userId); ok {
		if socket, okv := v.(socketio.Conn); okv {
			return socket
		}
	}
	return nil
}

func (m *UserMgr) kickAllInRoom(roomId string) {
	if len(roomId) < 1 {
		return
	}
	roomInfo := roomMgr.getRoom(roomId)
	if roomInfo == nil {
		return
	}

	for _, rs := range roomInfo.Seats {
		if rs.Userid > 0 {
			socket := m.get(rs.Userid)
			if socket != nil {
				m.del(rs.Userid)
				err := socket.Close()
				if err != nil {
					logger.Error(err)
				}
			}
		}
	}
}
func (m *UserMgr) del(userId int64) {
	m.userList.Delete(userId)
	m.userOnline--
}
