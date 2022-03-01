package game

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/utils"
	"github.com/jqiris/orange/database"
	"github.com/jqiris/orange/model"
)

var (
	DI_FEN      = []int{1, 2, 5}
	MAX_FAN     = []int{3, 4, 5}
	JU_SHU      = []int{4, 8}
	JU_SHU_COST = []int{2, 3}
)

type RoomMgr struct {
	userLocation  map[int]*UserLocation //用户位置信息
	rooms         map[int]*Room         //房间map
	creatingRooms map[int]bool          //创建标记
	totalRooms    int                   //房间总数

}

func NewRoomMgr() *RoomMgr {
	return &RoomMgr{
		userLocation:  make(map[int]*UserLocation),
		rooms:         make(map[int]*Room),
		creatingRooms: make(map[int]bool),
		totalRooms:    0,
	}
}

func (m *RoomMgr) getRoom(roomId int) *Room {
	return m.rooms[roomId]
}

func (m *RoomMgr) getUserSeat(userId int) int {
	if v, ok := m.userLocation[userId]; ok {
		return v.SeatIndex
	}
	return -1
}

func (m *RoomMgr) getUserLocations() map[int]*UserLocation {
	return m.userLocation
}

func (m *RoomMgr) createRoom(userId int, conf model.GameConf, gems, serverIp string, port int) (errcode int, roomId int) {
	errcode = 1
	if len(conf.Type) < 1 {
		return
	}

	if conf.Difen < 0 || conf.Difen > len(DI_FEN) {
		return
	}
	if conf.Zimo < 0 || conf.Zimo > 2 {
		return
	}
	if conf.Zuidafanshu < 0 || conf.Zuidafanshu > len(MAX_FAN) {
		return
	}
	if conf.Jushuxuanze < 0 || conf.Jushuxuanze > len(JU_SHU) {
		return
	}
	diamond := utils.StringToInt(gems)
	cost := JU_SHU_COST[conf.Jushuxuanze]
	if cost > diamond {
		errcode = 2222
		return
	}
	errcode, roomId = m.fnCreate(userId, conf, serverIp, port)
	return
}

func (m *RoomMgr) fnCreate(userId int, conf model.GameConf, ip string, port int) (int, int) {
	roomId := m.generateRoomId()
	if m.rooms[roomId] != nil || m.creatingRooms[roomId] {
		return m.fnCreate(userId, conf, ip, port)
	} else {
		m.creatingRooms[roomId] = true
		if _, err := database.GetRoomById(roomId); err == nil {
			delete(m.creatingRooms, roomId)
			return m.fnCreate(userId, conf, ip, port)
		} else {
			createTime := time.Now().Unix()
			roomInfo := &Room{
				Uuid:       "",
				Id:         roomId,
				NumOfGames: 0,
				CreateTime: createTime,
				NextButton: 0,
				Seats:      make([]Seat, 4),
				Conf: RoomConf{
					Type:        conf.Type,
					Zimo:        conf.Zimo,
					Jiangdui:    conf.Jiangdui,
					Hsz:         conf.Huansanzhang,
					Dianganghua: conf.Dianganghua,
					Menqing:     conf.Menqing,
					Tiandihu:    conf.Tiandihu,
					BaseScore:   DI_FEN[conf.Difen],
					MaxFan:      MAX_FAN[conf.Zuidafanshu],
					MaxGames:    JU_SHU[conf.Jushuxuanze],
					Creator:     userId,
				},
			}
			if conf.Type == "xlch" {
				roomInfo.GameMgr = NewXlchMj()
			} else {
				roomInfo.GameMgr = NewXzddMj()
			}
			for i := 0; i < 4; i++ {
				roomInfo.Seats[i] = Seat{
					UserId:      0,
					Score:       0,
					Name:        "",
					Ready:       false,
					SeatIndex:   i,
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
			if err := database.CreateRoom(&model.TRoom{
				UUID:       uuid,
				ID:         roomId,
				BaseInfo:   string(bs),
				CreateTime: createTime,
				IP:         ip,
				Port:       port,
			}); err != nil {
				logger.Error(err)
				return 3, 0
			} else {
				roomInfo.Uuid = uuid
				m.rooms[roomId] = roomInfo
				m.totalRooms++
				return 0, roomId
			}
		}
	}
}

func (m *RoomMgr) generateRoomId() int {
	var rstr []string
	for i := 0; i < 6; i++ {
		s := utils.IntToString(rand.Intn(10))
		rstr = append(rstr, s)
	}
	return utils.StringToInt(strings.Join(rstr, ""))
}

func (m *RoomMgr) getUserRoom(userId int) int {
	if v, ok := m.userLocation[userId]; ok {
		return v.RoomId
	}
	return 0
}

func (m *RoomMgr) enterRoom(roomId, userId int, userName string) int {
	fnTakeSeat := func(room *Room) int {
		if m.getUserRoom(userId) == roomId {
			return 0 //已存在
		}
		for i := 0; i < 4; i++ {
			seat := room.Seats[i]
			if seat.UserId <= 0 {
				seat.UserId = userId
				seat.Name = userName
				m.userLocation[userId] = &UserLocation{
					RoomId:    roomId,
					SeatIndex: i,
				}
				err := database.UpdateSeatInfo(roomId, i, seat.UserId, "", seat.Name)
				if err != nil {
					logger.Error(err)
				}
				return 0 //正常
			}
		}
		return 1 //房间已满
	}
	room := m.rooms[roomId]
	if room != nil {
		return fnTakeSeat(room)
	} else {
		res, err := database.GetRoomById(roomId)
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

func (m *RoomMgr) constructRoomFromDb(dbdata *model.TRoom) (*Room, error) {
	var cfg RoomConf
	if err := json.Unmarshal([]byte(dbdata.BaseInfo), &cfg); err != nil {
		return nil, err
	}
	roomInfo := &Room{
		Uuid:       dbdata.UUID,
		Id:         dbdata.ID,
		NumOfGames: dbdata.NumOfTurns,
		CreateTime: dbdata.CreateTime,
		NextButton: dbdata.NextButton,
		Seats:      make([]Seat, 4),
		Conf:       cfg,
	}
	if cfg.Type == "xlch" {
		roomInfo.GameMgr = NewXlchMj()
	} else {
		roomInfo.GameMgr = NewXzddMj()
	}
	roomId := dbdata.ID
	for i := 0; i < 4; i++ {
		s := Seat{}
		switch i {
		case 0:
			s.UserId = dbdata.UserID0
			s.Score = dbdata.UserScore0
			s.Name = dbdata.UserName0
		case 1:
			s.UserId = dbdata.UserID1
			s.Score = dbdata.UserScore1
			s.Name = dbdata.UserName1
		case 2:
			s.UserId = dbdata.UserID2
			s.Score = dbdata.UserScore2
			s.Name = dbdata.UserName2
		case 3:
			s.UserId = dbdata.UserID3
			s.Score = dbdata.UserScore3
			s.Name = dbdata.UserName3
		}
		s.Ready = false
		s.SeatIndex = i
		s.NumZiMo = 0
		s.NumJiePao = 0
		s.NumDianPao = 0
		s.NumAnGang = 0
		s.NumMingGang = 0
		s.NumChaJiao = 0
		if s.UserId > 0 {
			m.userLocation[s.UserId] = &UserLocation{
				RoomId:    roomId,
				SeatIndex: i,
			}
		}
		roomInfo.Seats[i] = s
	}
	m.rooms[roomId] = roomInfo
	m.totalRooms++
	return roomInfo, nil
}
