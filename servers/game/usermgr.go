package game

import (
	"sync"

	socketio "github.com/googollee/go-socket.io"
)

type UserLocation struct {
	RoomId    int `json:"roomId"`
	SeatIndex int `json:"seatIndex"`
}

type UserMgr struct {
	userList   sync.Map
	userOnline int
}

func NewUserMgr() *UserMgr {
	return &UserMgr{
		userList:   sync.Map{},
		userOnline: 0,
	}
}

func (m *UserMgr) bind(userId int, c socketio.Conn) {
	m.userList.Store(userId, c)
	m.userOnline++
}

func (m *UserMgr) isOnline(userId int) bool {
	_, exist := m.userList.Load(userId)
	return exist
}

func (m *UserMgr) broadcastInRoom(event string, data interface{}, sender int, args ...bool) {
	includingSender := false
	if len(args) > 0 {
		includingSender = args[0]
	}
	roomId := roomMgr.getUserRoom(sender)
	if roomId < 1 {
		return
	}
	roomInfo := roomMgr.getRoom(roomId)
	if roomInfo == nil {
		return
	}
	for i := 0; i < len(roomInfo.Seats); i++ {
		rs := roomInfo.Seats[i]
		if rs.UserId == sender && !includingSender {
			continue
		}
		if socket := m.get(rs.UserId); socket != nil {
			if data != nil {
				socket.Emit(event, data)
			} else {
				socket.Emit(event)
			}
		}
	}
}

func (m *UserMgr) sendMsg(userId int, event string, msgdata ...interface{}) {
	socket := m.get(userId)
	if socket == nil {
		return
	}
	socket.Emit(event, msgdata...)
}

func (m *UserMgr) get(userId int) socketio.Conn {
	if v, ok := m.userList.Load(userId); ok {
		if socket, okv := v.(socketio.Conn); okv {
			return socket
		}
	}
	return nil
}

func (m *UserMgr) kickAllInRoom(roomId int) {
	if roomId < 1 {
		return
	}
	roomInfo := roomMgr.getRoom(roomId)
	if roomInfo == nil {
		return
	}

	for _, rs := range roomInfo.Seats {
		if rs.UserId > 0 {
			socket := m.get(rs.UserId)
			if socket != nil {
				m.del(rs.UserId)
				socket.Close()
			}
		}
	}
}
func (m *UserMgr) del(userId int) {
	m.userList.Delete(userId)
	m.userOnline--
}
