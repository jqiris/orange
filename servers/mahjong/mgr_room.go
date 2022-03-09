package mahjong

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/utils"
	"github.com/jqiris/orange/database"
	"github.com/jqiris/orange/model"
	"github.com/jqiris/orange/protos"
)

var (
	DiFen     = []int32{1, 2, 5}
	MaxFan    = []int32{3, 4, 5}
	JuShu     = []int32{2, 8}
	JuShuCost = []int32{2, 3}
)

type RoomMgr struct {
	UserLocation  map[int64]*protos.UserLocation //用户位置信息
	Rooms         map[string]*protos.MjRoom      //房间map
	CreatingRooms map[string]bool                //创建标记
	TotalRooms    int32                          //房间总数
	lockLocation  sync.RWMutex
	lockRoom      sync.RWMutex
	lockCreating  sync.RWMutex
}

func NewRoomMgr() *RoomMgr {
	return &RoomMgr{
		UserLocation:  make(map[int64]*protos.UserLocation),
		Rooms:         make(map[string]*protos.MjRoom),
		CreatingRooms: make(map[string]bool),
		TotalRooms:    0,
	}
}
func (m *RoomMgr) getCreatingRoom(roomId string) bool {
	m.lockCreating.RLock()
	defer m.lockCreating.RUnlock()
	if v, ok := m.CreatingRooms[roomId]; ok {
		return v
	}
	return false
}
func (m *RoomMgr) saveCreatingRoom(roomId string, data bool) {
	m.lockCreating.Lock()
	defer m.lockCreating.Unlock()
	m.CreatingRooms[roomId] = data
}
func (m *RoomMgr) delCreatingRoom(roomId string) {
	m.lockCreating.Lock()
	defer m.lockCreating.Unlock()
	delete(m.CreatingRooms, roomId)
}

func (m *RoomMgr) getRoom(roomId string) *protos.MjRoom {
	m.lockRoom.RLock()
	defer m.lockRoom.RUnlock()
	if v, ok := m.Rooms[roomId]; ok {
		return v
	}
	return nil
}

func (m *RoomMgr) saveRoom(roomId string, data *protos.MjRoom) {
	m.lockRoom.Lock()
	defer m.lockRoom.Unlock()
	m.Rooms[roomId] = data
}
func (m *RoomMgr) delRoom(roomId string) {
	m.lockRoom.Lock()
	defer m.lockRoom.Unlock()
	delete(m.Rooms, roomId)
}

func (m *RoomMgr) getLocation(userId int64) *protos.UserLocation {
	m.lockLocation.RLock()
	defer m.lockLocation.RUnlock()
	if v, ok := m.UserLocation[userId]; ok {
		return v
	}
	return nil
}

func (m *RoomMgr) getUserLocations() map[int64]*protos.UserLocation {
	m.lockLocation.RLock()
	defer m.lockLocation.RUnlock()
	return m.UserLocation
}

func (m *RoomMgr) saveLocation(userId int64, data *protos.UserLocation) {
	m.lockLocation.Lock()
	defer m.lockLocation.Unlock()
	m.UserLocation[userId] = data
}

func (m *RoomMgr) delLocation(userId int64) {
	m.lockLocation.Lock()
	defer m.lockLocation.Unlock()
	delete(m.UserLocation, userId)
}

func (m *RoomMgr) getUserSeat(userId int64) int32 {
	if v := m.getLocation(userId); v != nil {
		return v.SeatIndex
	}
	return -1
}

func (m *RoomMgr) createRoom(userId int64, conf model.GameConf, gems int32, serverId, serverIp string, port int32) (errcode int32, roomId string) {
	errcode = 1
	if len(conf.Type) < 1 {
		return
	}

	if conf.Difen < 0 || int(conf.Difen) > len(DiFen) {
		return
	}
	if conf.Zimo < 0 || conf.Zimo > 2 {
		return
	}
	if conf.Zuidafanshu < 0 || int(conf.Zuidafanshu) > len(MaxFan) {
		return
	}
	if conf.Jushuxuanze < 0 || int(conf.Jushuxuanze) > len(JuShu) {
		return
	}
	cost := JuShuCost[conf.Jushuxuanze]
	if cost > gems {
		errcode = 2222
		return
	}
	errcode, roomId = m.fnCreate(userId, conf, serverId, serverIp, port)
	return
}

func (m *RoomMgr) fnCreate(userId int64, conf model.GameConf, serverId, ip string, port int32) (int32, string) {
	roomId := m.generateRoomId()
	room := m.getRoom(roomId)
	if room != nil || m.getCreatingRoom(roomId) {
		return m.fnCreate(userId, conf, serverId, ip, port)
	} else {
		m.saveCreatingRoom(roomId, true)
		if _, err := database.GetMjRoomById(roomId); err == nil {
			m.delCreatingRoom(roomId)
			return m.fnCreate(userId, conf, serverId, ip, port)
		} else {
			createTime := time.Now().Unix()
			roomInfo := &protos.MjRoom{
				Uuid:       "",
				GameType:   conf.Type,
				RoomId:     roomId,
				NumOfGames: 0,
				CreateTime: createTime,
				NextButton: 0,
				Seats:      make([]*protos.MjSeat, 4),
				Conf: &protos.MjRoomConf{
					Type:        conf.Type,
					Zimo:        conf.Zimo,
					Jiangdui:    conf.Jiangdui,
					Hsz:         conf.Huansanzhang,
					Dianganghua: conf.Dianganghua,
					Menqing:     conf.Menqing,
					Tiandihu:    conf.Tiandihu,
					BaseScore:   DiFen[conf.Difen],
					MaxFan:      MaxFan[conf.Zuidafanshu],
					MaxGames:    JuShu[conf.Jushuxuanze],
					Creator:     userId,
				},
			}
			for i := 0; i < 4; i++ {
				roomInfo.Seats[i] = &protos.MjSeat{
					Userid:      0,
					Score:       0,
					Name:        "",
					Ready:       false,
					Seatindex:   int32(i),
					NumZiMo:     0,
					NumJiePao:   0,
					NumDianPao:  0,
					NumAnGang:   0,
					NumMingGang: 0,
					NumChaJiao:  0,
				}
			}
			bs, _ := json.Marshal(roomInfo.Conf)
			uuid := fmt.Sprintf("%v%v", createTime, roomId)
			if err = database.CreateMjRoom(&model.MahjongRoom{
				UUID:       uuid,
				RoomID:     roomId,
				GameType:   conf.Type,
				BaseInfo:   string(bs),
				CreateTime: createTime,
				ServerID:   serverId,
				Landlord:   userId,
			}); err != nil {
				logger.Error(err)
				return 3, ""
			} else {
				roomInfo.Uuid = uuid
				m.saveRoom(roomId, roomInfo)
				m.TotalRooms++
				return 0, roomId
			}
		}
	}
}

func (m *RoomMgr) generateRoomId() string {
	var rstr []string
	for i := 0; i < 6; i++ {
		s := utils.IntToString(rand.Intn(10))
		rstr = append(rstr, s)
	}
	return strings.Join(rstr, "")
}

func (m *RoomMgr) getUserRoom(userId int64) string {
	if v := m.getLocation(userId); v != nil {
		return v.RoomId
	}
	return ""
}

func (m *RoomMgr) enterRoom(roomId string, userId int64, userName string) int32 {
	fnTakeSeat := func(room *protos.MjRoom) int32 {
		if m.getUserRoom(userId) == roomId {
			return 0 //已存在
		}
		for i := 0; i < 4; i++ {
			seat := room.Seats[i]
			if seat.Userid <= 0 {
				seat.Userid = userId
				seat.Name = userName
				m.saveLocation(userId, &protos.UserLocation{
					RoomId:    roomId,
					SeatIndex: int32(i),
				})
				err := database.UpdateMjSeatInfo(roomId, int32(i), seat.Userid, "", seat.Name)
				if err != nil {
					logger.Error(err)
				}
				return 0 //正常
			}
		}
		return 1 //房间已满
	}
	room := m.getRoom(roomId)
	if room != nil {
		return fnTakeSeat(room)
	} else {
		res, err := database.GetMjRoomById(roomId)
		if err != nil {
			logger.Error(err)
			return 2 //找不到房间
		}
		room, err := m.constructRoomFromDb(res)
		if err != nil {
			logger.Error(err)
			return 3 //创建房间失败
		}
		return fnTakeSeat(room)
	}
}

func (m *RoomMgr) constructRoomFromDb(dbdata *model.MahjongRoom) (*protos.MjRoom, error) {
	var cfg *protos.MjRoomConf
	if err := json.Unmarshal([]byte(dbdata.BaseInfo), &cfg); err != nil {
		return nil, err
	}
	roomInfo := &protos.MjRoom{
		Uuid:       dbdata.UUID,
		RoomId:     dbdata.RoomID,
		GameType:   dbdata.GameType,
		NumOfGames: dbdata.NumOfTurns,
		CreateTime: dbdata.CreateTime,
		NextButton: dbdata.NextButton,
		Seats:      make([]*protos.MjSeat, 4),
		Conf:       cfg,
	}
	roomId := dbdata.RoomID
	for i := 0; i < 4; i++ {
		s := &protos.MjSeat{}
		switch i {
		case 0:
			s.Userid = dbdata.UserID0
			s.Score = dbdata.UserScore0
			s.Name = dbdata.UserName0
		case 1:
			s.Userid = dbdata.UserID1
			s.Score = dbdata.UserScore1
			s.Name = dbdata.UserName1
		case 2:
			s.Userid = dbdata.UserID2
			s.Score = dbdata.UserScore2
			s.Name = dbdata.UserName2
		case 3:
			s.Userid = dbdata.UserID3
			s.Score = dbdata.UserScore3
			s.Name = dbdata.UserName3
		}
		s.Ready = false
		s.Seatindex = int32(i)
		s.NumZiMo = 0
		s.NumJiePao = 0
		s.NumDianPao = 0
		s.NumAnGang = 0
		s.NumMingGang = 0
		s.NumChaJiao = 0
		if s.Userid > 0 {
			m.saveLocation(s.Userid, &protos.UserLocation{
				RoomId:    roomId,
				SeatIndex: int32(i),
			})
		}
		roomInfo.Seats[i] = s
	}
	m.saveRoom(roomId, roomInfo)
	m.TotalRooms++
	return roomInfo, nil
}

func (m *RoomMgr) setReady(userId int64, value bool) {
	roomId := m.getUserRoom(userId)
	if len(roomId) < 1 {
		return
	}
	room := m.getRoom(roomId)
	if room == nil {
		return
	}
	seatIndex := m.getUserSeat(userId)
	if seatIndex < 0 {
		return
	}
	s := room.Seats[seatIndex]
	s.Ready = value
}

func (m *RoomMgr) getTotallRooms() int32 {
	return m.TotalRooms
}

func (m *RoomMgr) destroy(roomId string) {
	roomInfo := m.getRoom(roomId)
	if roomInfo == nil {
		return
	}
	for _, seat := range roomInfo.Seats {
		userId := seat.Userid
		if userId > 0 {
			m.delLocation(userId)
			err := database.UpdateUser(userId, map[string]any{"room_id": ""})
			if err != nil {
				logger.Error(err)
			}
		}
	}
	m.delRoom(roomId)
	m.TotalRooms--
	err := database.ArchiveMjRoom(roomInfo.GameType, roomId)
	if err != nil {
		logger.Error(err)
	} else {
		err = database.DeleteMjRoom(roomId)
		if err != nil {
			logger.Error(err)
		}
	}
}

func (m *RoomMgr) isCreator(roomId string, userId int64) bool {
	roomInfo := m.getRoom(roomId)
	if roomInfo == nil {
		return false
	}
	return roomInfo.Conf.Creator == userId
}

func (m *RoomMgr) exitRoom(userId int64) {
	location := m.getLocation(userId)
	if location == nil {
		return
	}
	roomId := location.RoomId
	seatIndex := location.SeatIndex
	room := m.getRoom(roomId)
	m.delLocation(userId)
	if room == nil || seatIndex == -1 {
		return
	}
	seat := room.Seats[seatIndex]
	seat.Userid = 0
	seat.Name = ""
	numOfPlayers := 0
	for _, gs := range room.Seats {
		if gs.Userid > 0 {
			numOfPlayers++
		}
	}
	err := database.UpdateUser(userId, map[string]any{"room_id": ""})
	if err != nil {
		logger.Error(err)
	}
	if numOfPlayers == 0 {
		m.destroy(roomId)
	}
}
