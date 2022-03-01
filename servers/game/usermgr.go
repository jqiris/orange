package game

import socketio "github.com/googollee/go-socket.io"

type UserLocation struct {
	RoomId    int `json:"roomId"`
	SeatIndex int `json:"seatIndex"`
}

type UserMgr struct {
	userList   map[int]socketio.Conn
	userOnline int
}

func NewUserMgr() *UserMgr {
	return &UserMgr{
		userList:   make(map[int]socketio.Conn),
		userOnline: 0,
	}
}

func (m *UserMgr) bind(userId int, c socketio.Conn) {
	m.userList[userId] = c
	m.userOnline++
}

func (m *UserMgr) isOnline(userId int) bool {
	return m.userList[userId] != nil
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
		if socket := m.userList[rs.UserId]; socket != nil {
			socket.Emit(event, data)
		}
	}
}

func (m *UserMgr) sendMsg(userId int, event string, msgdata interface{}) {
	socket := m.userList[userId]
	if socket == nil {
		return
	}
	socket.Emit(event, msgdata)
}
