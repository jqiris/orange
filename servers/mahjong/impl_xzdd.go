package mahjong

import (
	"encoding/base64"
	"encoding/json"
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/jqiris/orange/constant"
	"github.com/jqiris/orange/database"
	"github.com/jqiris/orange/model"
	"github.com/jqiris/orange/protos"
	"github.com/jqiris/orange/tools"

	"github.com/jqiris/kungfu/v2/logger"
)

type XzddMj struct {
	Games            map[string]*protos.MjGameData
	GameSeatsOfUsers map[int64]*protos.MjSeat
	DissolvingList   []string
	lockGame         sync.RWMutex
	lockSeat         sync.RWMutex
}

func NewXzddMj() *XzddMj {
	return &XzddMj{
		Games:            make(map[string]*protos.MjGameData),
		GameSeatsOfUsers: make(map[int64]*protos.MjSeat),
		DissolvingList:   make([]string, 0),
	}
}

func (m *XzddMj) getGame(roomId string) *protos.MjGameData {
	m.lockGame.RLock()
	defer m.lockGame.RUnlock()
	if v, ok := m.Games[roomId]; ok {
		return v
	}
	return nil
}
func (m *XzddMj) saveGame(roomId string, data *protos.MjGameData) {
	m.lockGame.Lock()
	defer m.lockGame.Unlock()
	m.Games[roomId] = data
}
func (m *XzddMj) delGame(roomId string) {
	m.lockGame.Lock()
	defer m.lockGame.Unlock()
	delete(m.Games, roomId)
}

func (m *XzddMj) getSeat(userId int64) *protos.MjSeat {
	m.lockSeat.RLock()
	defer m.lockSeat.RUnlock()
	if v, ok := m.GameSeatsOfUsers[userId]; ok {
		return v
	}
	return nil
}
func (m *XzddMj) saveSeat(userId int64, data *protos.MjSeat) {
	m.lockSeat.Lock()
	defer m.lockSeat.Unlock()
	m.GameSeatsOfUsers[userId] = data
}

func (m *XzddMj) delSeat(userId int64) {
	m.lockSeat.Lock()
	defer m.lockSeat.Unlock()
	delete(m.GameSeatsOfUsers, userId)
}

func (m *XzddMj) String() string {
	return "四川麻将血战到底"
}
func (m *XzddMj) begin(roomId string) {
	roomInfo := roomMgr.getRoom(roomId)
	if roomInfo == nil {
		return
	}
	seats := roomInfo.Seats
	game := &protos.MjGameData{
		Conf:          roomInfo.Conf,
		Uuid:          roomInfo.Uuid,
		GameIndex:     roomInfo.NumOfGames,
		Button:        roomInfo.NextButton,
		Mahjongs:      make([]int32, 108),
		CurrentIndex:  0,
		GameSeats:     make([]*protos.MjSeat, 4),
		NumOfQue:      0,
		Turn:          0,
		ChuPai:        -1,
		State:         "idle",
		FirstHupai:    -1,
		Yipaoduoxiang: -1,
		Fangpaoshumu:  -1,
		ActionList:    []int32{},
		HupaiList:     []int32{},
		ChupaiCnt:     0,
	}
	roomInfo.NumOfGames++

	for i := 0; i < 4; i++ {
		data := &protos.MjSeat{}
		data.Game = game
		data.Seatindex = int32(i)
		data.Userid = seats[i].Userid
		data.Holds = make([]int32, 0)
		data.Folds = make([]int32, 0)
		data.Angangs = make([]int32, 0)
		data.Wangangs = make([]int32, 0)
		data.Diangangs = make([]int32, 0)
		data.Pengs = make([]int32, 0)
		data.Que = -1
		data.Huanpais = make([]int32, 0)
		data.CountMap = make(map[int32]int32)
		data.TingMap = make(map[int32]*protos.MjTingData)
		data.Pattern = ""
		data.CanGang = false
		data.GangPai = make([]int32, 0)
		data.CanPeng = false
		data.CanHu = false
		data.CanChuPai = false
		data.GuoHuFan = -1
		data.Hued = false
		data.Iszimo = false
		data.IsGangHu = false
		data.Actions = make([]*protos.MjActionData, 0)
		data.Fan = 0
		data.Score = 0
		data.LastFangGangSeat = 0
		//统计信息
		data.NumZiMo = 0
		data.NumJiePao = 0
		data.NumDianPao = 0
		data.NumAnGang = 0
		data.NumMingGang = 0
		data.NumChaJiao = 0
		game.GameSeats[i] = data
		m.saveSeat(data.Userid, data)
	}
	m.saveGame(roomId, game)
	m.shuffle(game)
	m.deal(game)

	numOfMJ := len(game.Mahjongs) - int(game.CurrentIndex)
	huansanzhang := roomInfo.Conf.Hsz

	for i := 0; i < len(seats); i++ {
		st := seats[i]
		holds := game.GameSeats[i].Holds
		logger.Infof("holds:%+v", holds)
		//通知玩家手牌
		userMgr.sendMsg(st.Userid, "game_holds_push", holds)
		//通知还剩多少张牌
		userMgr.sendMsg(st.Userid, "mj_count_push", numOfMJ)
		//通知还剩多少局
		userMgr.sendMsg(st.Userid, "game_num_push", roomInfo.NumOfGames)
		//通知游戏开始
		userMgr.sendMsg(st.Userid, "game_begin_push", game.Button)

		if huansanzhang == true {
			game.State = "huanpai"
			//通知准备换牌
			userMgr.sendMsg(st.Userid, "game_huanpai_push")
		} else {
			game.State = "dingque"
			//通知准备定缺
			userMgr.sendMsg(st.Userid, "game_dingque_push")
		}
	}
}

func (m *XzddMj) deal(game *protos.MjGameData) {
	game.CurrentIndex = 0
	//每人13张 一共 13*4 ＝ 52张 庄家多一张 53张
	seatIndex := game.Button
	for i := 0; i < 52; i++ {
		m.mopai(game, seatIndex)
		seatIndex++
		seatIndex %= 4
	}
	//庄家多摸最后一张
	m.mopai(game, game.Button)

	//当前轮设置为庄家
	game.Turn = game.Button
}

func (m *XzddMj) mopai(game *protos.MjGameData, seatIndex int32) int32 {
	if int(game.CurrentIndex) >= len(game.Mahjongs) {
		return -1
	}
	data := game.GameSeats[seatIndex]
	pai := game.Mahjongs[game.CurrentIndex]
	data.Holds = append(data.Holds, pai)
	data.CountMap[pai]++
	game.CurrentIndex++
	return pai
}

func (m *XzddMj) shuffle(game *protos.MjGameData) {
	mahjongs := game.Mahjongs
	index := 0
	for i := 0; i < 9; i++ {
		for c := 0; c < 4; c++ {
			mahjongs[index] = int32(i)
			index++
		}
	}
	for i := 9; i < 18; i++ {
		for c := 0; c < 4; c++ {
			mahjongs[index] = int32(i)
			index++
		}
	}
	for i := 18; i < 27; i++ {
		for c := 0; c < 4; c++ {
			mahjongs[index] = int32(i)
			index++
		}
	}
	mjLen := len(mahjongs)
	for i := 0; i < mjLen; i++ {
		lastIndex := mjLen - i - 1
		randIndex := int32(math.Floor(float64(rand.Intn(100)*lastIndex) / 100))
		mahjongs[randIndex], mahjongs[lastIndex] = mahjongs[lastIndex], mahjongs[randIndex]
	}
}

func (m *XzddMj) SetReady(userId int64) {
	roomId := roomMgr.getUserRoom(userId)
	if len(roomId) < 1 {
		return
	}
	roomInfo := roomMgr.getRoom(roomId)
	if roomInfo == nil {
		return
	}
	roomMgr.setReady(userId, true)

	game := m.getGame(roomId)
	if game == nil {
		if len(roomInfo.Seats) == 4 {
			for i := 0; i < 4; i++ {
				st := roomInfo.Seats[i]
				if !st.Ready || !userMgr.isOnline(st.Userid) {
					return
				}
			}
			//4个人到齐了，并且都准备好了，则开始新的一局
			go m.begin(roomId)
		}
	} else {
		numOfMJ := len(game.Mahjongs) - int(game.CurrentIndex)
		data := map[string]any{
			"state":         game.State,
			"numofmj":       numOfMJ,
			"button":        game.Button,
			"turn":          game.Turn,
			"chuPai":        game.ChuPai,
			"huanpaimethod": game.HuanpaiMethod,
		}
		var seats []*protos.MjSeat
		var seatData *protos.MjSeat
		for i := 0; i < 4; i++ {
			sd := game.GameSeats[i]
			st := &protos.MjSeat{
				Userid:    sd.Userid,
				Folds:     sd.Folds,
				Angangs:   sd.Angangs,
				Diangangs: sd.Diangangs,
				Wangangs:  sd.Wangangs,
				Pengs:     sd.Pengs,
				Que:       sd.Que,
				Hued:      sd.Hued,
				Iszimo:    sd.Iszimo,
			}
			if sd.Userid == userId {
				st.Holds = sd.Holds
				st.Huanpais = sd.Huanpais
				seatData = sd
			} else {
				st.Huanpais = sd.Huanpais
			}
			seats = append(seats, st)
		}
		data["seats"] = seats
		//同步整个信息给客户端
		userMgr.sendMsg(userId, "game_sync_push", data)
		m.sendOperations(game, seatData, game.ChuPai)
	}
}

func (m *XzddMj) sendOperations(game *protos.MjGameData, seatData *protos.MjSeat, pai int32) {
	if m.hasOperations(seatData) {
		if pai == -1 {
			pai = seatData.Holds[len(seatData.Holds)-1]
		}
		data := map[string]any{
			"pai":     pai,
			"hu":      seatData.CanHu,
			"peng":    seatData.CanPeng,
			"gang":    seatData.CanGang,
			"gangpai": seatData.GangPai,
			"si":      seatData.Seatindex,
		}
		//如果可以有操作，则进行操作
		userMgr.sendMsg(seatData.Userid, "game_action_push", data)
	} else {
		userMgr.sendMsg(seatData.Userid, "game_action_push")
	}
}

func (m *XzddMj) hasOperations(seatData *protos.MjSeat) bool {
	if seatData.CanGang || seatData.CanPeng || seatData.CanHu {
		return true
	}
	return false
}

func (m *XzddMj) HasBegan(roomId string) bool {
	game := m.getGame(roomId)
	if game != nil {
		return true
	}
	roomInfo := roomMgr.getRoom(roomId)
	if roomInfo != nil {
		return roomInfo.NumOfGames > 0
	}
	return false
}

func (m *XzddMj) DissolveRequest(roomId string, userId int64) *protos.MjRoom {
	roomInfo := roomMgr.getRoom(roomId)
	if roomInfo == nil {
		return nil
	}
	if roomInfo.Dr != nil {
		return nil
	}
	seatIndex := roomMgr.getUserSeat(userId)
	if seatIndex == -1 {
		return nil
	}
	dr := &protos.MjDissolveData{
		EndTime: time.Now().UnixMilli() + 30000,
		States:  []bool{false, false, false, false},
	}
	dr.States[seatIndex] = true
	roomInfo.Dr = dr
	m.DissolvingList = append(m.DissolvingList, roomId)
	return roomInfo
}

func (m *XzddMj) DissolveAgree(roomId string, userId int64, agree bool) *protos.MjRoom {
	roomInfo := roomMgr.getRoom(roomId)
	if roomInfo == nil {
		return nil
	}
	if roomInfo.Dr == nil {
		return nil
	}
	seatIndex := roomMgr.getUserSeat(userId)
	if seatIndex == -1 {
		return nil
	}
	if agree {
		roomInfo.Dr.States[seatIndex] = true
	} else {
		roomInfo.Dr = nil
		idx := tools.IndexOf(m.DissolvingList, roomId)
		if idx != -1 {
			m.DissolvingList = tools.SliceDel(m.DissolvingList, idx, 1)
		}
	}
	return roomInfo
}

func (m *XzddMj) DoDissolve(roomId string) {
	roomInfo := roomMgr.getRoom(roomId)
	if roomInfo == nil {
		return
	}
	game := m.getGame(roomId)
	go m.doGameOver(game, roomInfo.Seats[0].Userid, true)
}

func (m *XzddMj) HuanSanZhang(userId int64, p1 int32, p2 int32, p3 int32) {
	seatData := m.getSeat(userId)
	if seatData == nil {
		logger.Error("can't find user game data.")
		return
	}
	game := seatData.Game
	if game.State != "huanpai" {
		logger.Errorf("can't recv DingQue when game.state == %v", game.State)
		return
	}
	if len(seatData.Huanpais) > 0 {
		logger.Error("player has done this action")
		return
	}
	if seatData.CountMap[p1] == 0 {
		return
	}
	seatData.CountMap[p1]--

	if seatData.CountMap[p2] == 0 {
		seatData.CountMap[p1]++
		return
	}
	seatData.CountMap[p2]--

	if seatData.CountMap[p3] == 0 {
		seatData.CountMap[p1]++
		seatData.CountMap[p2]++
		return
	}

	seatData.CountMap[p1]++
	seatData.CountMap[p2]++

	seatData.Huanpais = []int32{p1, p2, p3}

	for _, p := range seatData.Huanpais {
		idx := tools.IndexOf(seatData.Holds, p)
		seatData.Holds = tools.SliceDel(seatData.Holds, idx, 1)
		seatData.CountMap[p]--
	}
	userMgr.sendMsg(seatData.Userid, "game_holds_push", seatData.Holds)

	for _, sd := range game.GameSeats {
		if sd == seatData {
			var rd = map[string]any{
				"si":       seatData.Userid,
				"Huanpais": seatData.Huanpais,
			}
			userMgr.sendMsg(sd.Userid, "huanpai_notify", rd)
		} else {
			var rd = map[string]any{
				"si":       seatData.Userid,
				"Huanpais": []int32{},
			}
			userMgr.sendMsg(sd.Userid, "huanpai_notify", rd)
		}
	}

	//如果还有未换牌的玩家，则继承等待
	for _, gs := range game.GameSeats {
		if len(gs.Huanpais) == 0 {
			return
		}
	}

	//换牌函数
	var fn = func(s1 *protos.MjSeat, huanjin []int32) {
		for i := 0; i < len(huanjin); i++ {
			var p = huanjin[i]
			s1.Holds = append(s1.Holds, p)
			s1.CountMap[p]++
		}
	}

	//开始换牌
	var f = rand.Intn(100)
	var s = game.GameSeats
	var huanpaiMethod int32 = 0
	//对家换牌
	if f < 33 {
		fn(s[0], s[2].Huanpais)
		fn(s[1], s[3].Huanpais)
		fn(s[2], s[0].Huanpais)
		fn(s[3], s[1].Huanpais)
		huanpaiMethod = 0
	} else if f < 66 {
		//换下家的牌
		fn(s[0], s[1].Huanpais)
		fn(s[1], s[2].Huanpais)
		fn(s[2], s[3].Huanpais)
		fn(s[3], s[0].Huanpais)
		huanpaiMethod = 1
	} else {
		//换上家的牌
		fn(s[0], s[3].Huanpais)
		fn(s[1], s[0].Huanpais)
		fn(s[2], s[1].Huanpais)
		fn(s[3], s[2].Huanpais)
		huanpaiMethod = 2
	}

	var rd = map[string]any{
		"method": huanpaiMethod,
	}
	game.HuanpaiMethod = huanpaiMethod

	game.State = "dingque"
	for _, gs := range s {
		Userid := gs.Userid
		userMgr.sendMsg(Userid, "game_huanpai_over_push", rd)

		userMgr.sendMsg(Userid, "game_holds_push", gs.Holds)
		//通知准备定缺
		userMgr.sendMsg(Userid, "game_dingque_push")
	}
}

func (m *XzddMj) constructGameBaseInfo(game *protos.MjGameData) {
	var baseInfo = map[string]any{
		"type":     game.Conf.Type,
		"button":   game.Button,
		"index":    game.GameIndex,
		"mahjongs": game.Mahjongs,
	}
	seats := make([][]int32, 4)
	for i := 0; i < 4; i++ {
		seats[i] = game.GameSeats[i].Holds
	}
	baseInfo["game_seats"] = seats
	bs, _ := json.Marshal(baseInfo)
	game.BaseInfoJson = string(bs)
}

func (m *XzddMj) DingQue(userId int64, que int32) {
	seatData := m.getSeat(userId)
	if seatData == nil {
		logger.Error("can't find user game data.")
		return
	}
	game := seatData.Game
	if game.State != "dingque" {
		logger.Errorf("can't recv DingQue when game.state == %v", game.State)
		return
	}
	if seatData.Que < 0 {
		game.NumOfQue++
	}
	seatData.Que = que

	//检查玩家可以做的动作
	//如果4个人都定缺了，通知庄家出牌
	if game.NumOfQue == 4 {
		go m.doDingQue(game, seatData)
	} else {
		userMgr.broadcastInRoom("game_dingque_notify_push", seatData.Userid, seatData.Userid, true)
	}
}

func (m *XzddMj) doDingQue(game *protos.MjGameData, seatData *protos.MjSeat) {
	m.constructGameBaseInfo(game)
	arr := []int32{1, 1, 1, 1}
	for i, gs := range game.GameSeats {
		arr[i] = gs.Que
	}
	userMgr.broadcastInRoom("game_dingque_finish_push", arr, seatData.Userid, true)
	userMgr.broadcastInRoom("game_playing_push", nil, seatData.Userid, true)

	//进行听牌检查
	for _, gs := range game.GameSeats {
		var duoyu int32 = -1
		if len(gs.Holds) == 14 {
			isPop := false
			isPop, gs.Holds, duoyu = tools.SlicePop(gs.Holds)
			if isPop {
				gs.CountMap[duoyu] -= 1
			}
		}
		m.checkCanTingPai(game, gs)
		if duoyu >= 0 {
			gs.Holds = append(gs.Holds, duoyu)
			gs.CountMap[duoyu]++
		}
	}

	turnSeat := game.GameSeats[game.Turn]
	game.State = "playing"
	//通知玩家出牌方
	turnSeat.CanChuPai = true
	userMgr.broadcastInRoom("game_chupai_push", turnSeat.Userid, turnSeat.Userid, true)
	//检查是否可以暗杠或者胡
	//直杠
	m.checkCanAnGang(game, turnSeat)
	//检查胡 用最后一张来检查
	m.checkCanHu(game, turnSeat, turnSeat.Holds[len(turnSeat.Holds)-1])
	//通知前端
	m.sendOperations(game, turnSeat, game.ChuPai)
	// logger.Infof("doDingQue,seat:%+v,turn:%+v", game.GameSeats[game.Turn], game.Turn)
}

func (m *XzddMj) ChuPai(userId int64, pai int32) {
	seatData := m.getSeat(userId)
	if seatData == nil {
		logger.Error("can't find the user game data")
		return
	}
	// logger.Infof("ChuPai userId: %d, pai: %d,seatData:%+v", userId, pai, seatData)
	game := seatData.Game
	if game.Turn != seatData.Seatindex {
		logger.Error("not your turn")
		return
	}
	if seatData.Hued {
		logger.Error("you have already hued,no kidding plz.")
		return
	}
	if seatData.CanChuPai == false {
		logger.Error("no need chupai")
		return
	}
	if m.hasOperations(seatData) {
		logger.Error("plz Guo before you chupai")
		return
	}
	index := tools.IndexOf(seatData.Holds, pai)
	if index == -1 {
		logger.Errorf("can't find the mj,holds:%+v,pai:%+v", seatData.Holds, pai)
		return
	}
	seatData.CanChuPai = false
	game.ChupaiCnt++
	seatData.GuoHuFan = -1
	seatData.Holds = tools.SliceDel(seatData.Holds, index, 1)
	seatData.CountMap[pai]--
	game.ChuPai = pai
	m.recordGameAction(game, game.Turn, constant.ActionChupai, pai)
	m.checkCanTingPai(game, seatData)
	userMgr.broadcastInRoom("game_chupai_notify_push", map[string]any{
		"userId": seatData.Userid,
		"pai":    pai,
	}, seatData.Userid, true)
	if v := seatData.TingMap[game.ChuPai]; v != nil {
		seatData.GuoHuFan = v.Fan
	}
	//检查是否有人要胡，要碰 要杠
	hasActions := false
	for i, ddd := range game.GameSeats {
		if int(game.Turn) == i {
			continue
		}
		if ddd.Hued {
			continue
		}
		m.checkCanHu(game, ddd, pai)
		if seatData.LastFangGangSeat == -1 {
			if v, ok := seatData.TingMap[pai]; ok && ddd.CanHu && ddd.GuoHuFan >= 0 && v.Fan <= ddd.GuoHuFan {
				ddd.CanHu = false
				userMgr.sendMsg(ddd.Userid, "guohu_push")
			}
		}
		m.checkCanPeng(game, ddd, pai)
		m.checkCanDianGang(game, ddd, pai)
		if m.hasOperations(ddd) {
			m.sendOperations(game, ddd, game.ChuPai)
			hasActions = true
		}
	}
	//如果没有人有操作，则向下一家发牌，并通知他出牌
	if !hasActions {
		time.AfterFunc(500*time.Millisecond, func() {
			userMgr.broadcastInRoom("guo_notify_push", map[string]any{"userId": seatData.Userid, "pai": game.ChuPai}, seatData.Userid, true)
			seatData.Folds = append(seatData.Folds, game.ChuPai)
			game.ChuPai = -1
			m.moveToNextUser(game, -1)
			m.doUserMoPai(game)
		})
	}
}

func (m *XzddMj) checkCanDianGang(game *protos.MjGameData, seatData *protos.MjSeat, targetPai int32) {
	//检查玩家手上的牌
	//如果没有牌了，则不能再杠
	if len(game.Mahjongs) <= int(game.CurrentIndex) {
		return
	}
	if m.getMJType(targetPai) == seatData.Que {
		return
	}
	count := seatData.CountMap[targetPai]
	if count >= 3 {
		seatData.CanGang = true
		seatData.GangPai = append(seatData.GangPai, targetPai)
		return
	}
}
func (m *XzddMj) checkCanPeng(game *protos.MjGameData, seatData *protos.MjSeat, targetPai int32) {
	if m.getMJType(targetPai) == seatData.Que {
		return
	}
	count := seatData.CountMap[targetPai]
	if count >= 2 {
		seatData.CanPeng = true
	}
}

func (m *XzddMj) checkCanTingPai(game *protos.MjGameData, seatData *protos.MjSeat) {
	seatData.TingMap = make(map[int32]*protos.MjTingData)
	//检查手上的牌是不是已打缺，如果未打缺，则不进行判定
	for _, pai := range seatData.Holds {
		if m.getMJType(pai) == seatData.Que {
			return
		}
	}
	//检查是否是七对 前提是没有碰，也没有杠 ，即手上拥有13张牌
	if len(seatData.Holds) == 13 {
		//有5对牌
		var danPai int32 = -1
		pairCount := 0
		for k, c := range seatData.CountMap {
			if c == 2 || c == 3 {
				pairCount++
			} else if c == 4 {
				pairCount += 2
			}
			if c == 1 || c == 3 {
				//如果已经有单牌了，表示不止一张单牌，并没有下叫。直接闪
				if danPai >= 0 {
					break
				}
				danPai = k
			}
		}
		//检查是否有6对 并且单牌是不是目标牌
		if pairCount == 6 {
			//七对只能和一张，就是手上那张单牌
			//七对的番数＝ 2番+N个4个牌（即龙七对）
			seatData.TingMap[danPai] = &protos.MjTingData{
				Fan:     2,
				Pattern: "7pairs",
			}
		}
	}
	//检查是否是对对胡  由于四川麻将没有吃，所以只需要检查手上的牌
	//对对胡叫牌有两种情况
	//1、N坎 + 1张单牌
	//2、N-1坎 + 两对牌
	singleCount := 0
	pairCount := 0
	colCount := 0
	var arr []int32
	for k, c := range seatData.CountMap {
		if c == 1 {
			singleCount++
			arr = append(arr, k)
		} else if c == 2 {
			pairCount++
			arr = append(arr, k)
		} else if c == 3 {
			colCount++
		} else if c == 4 {
			//手上有4个一样的牌，在四川麻将中是和不了对对胡的 随便加点东西
			singleCount++
			pairCount += 2
		}
	}
	if (pairCount == 2 && singleCount == 0) || (pairCount == 0 && singleCount == 1) {
		for _, p := range arr {
			if seatData.TingMap[p] == nil {
				seatData.TingMap[p] = &protos.MjTingData{
					Fan:     1,
					Pattern: "duidui",
				}
			}
		}
	}
	//检查是不是平胡
	if seatData.Que != 0 {
		checkTingPai(seatData, 0, 9)
	}

	if seatData.Que != 1 {
		checkTingPai(seatData, 9, 18)
	}

	if seatData.Que != 2 {
		checkTingPai(seatData, 18, 27)
	}
}

func (m *XzddMj) getMJType(id int32) int32 {
	if id >= 0 && id < 9 {
		//筒
		return 0
	} else if id >= 9 && id < 18 {
		//条
		return 1
	} else if id >= 18 && id < 27 {
		//万
		return 2
	}
	return -1
}

func (m *XzddMj) recordGameAction(game *protos.MjGameData, si, action, pai int32) {
	game.ActionList = append(game.ActionList, si, action)
	if pai > -1 {
		game.ActionList = append(game.ActionList, pai)
	}
}

func (m *XzddMj) Peng(userId int64) {
	logger.Infof("Peng,userId:%v", userId)
	seatData := m.getSeat(userId)
	if seatData == nil {
		logger.Error("can't find user game data.")
		return
	}
	game := seatData.Game
	//如果是他出的牌，则忽略
	if game.Turn == seatData.Seatindex {
		logger.Error("it's your turn.")
		return
	}

	//如果没有碰的机会，则不能再碰
	if seatData.CanPeng == false {
		logger.Error("seatData.Peng == false")
		return
	}

	//和的了，就不要再来了
	if seatData.Hued {
		logger.Error("you have already hued. no kidding plz.")
		return
	}

	//如果有人可以胡牌，则需要等待
	i := game.Turn
	for {
		i = (i + 1) % 4
		if i == game.Turn {
			break
		} else {
			var ddd = game.GameSeats[i]
			if ddd.CanHu && i != seatData.Seatindex {
				return
			}
		}
	}

	m.clearAllOptions(game, nil)

	//验证手上的牌的数目
	pai := game.ChuPai
	c := seatData.CountMap[pai]
	if c < 2 {
		logger.Infof("pai:%v,count:%v", pai, c)
		logger.Info(seatData.Holds)
		logger.Info("lack of mj.")
		return
	}

	//进行碰牌处理
	//扣掉手上的牌
	//从此人牌中扣除
	for i := 0; i < 2; i++ {
		var index = tools.IndexOf(seatData.Holds, pai)
		if index == -1 {
			logger.Info("can't find mj.")
			return
		}
		seatData.Holds = tools.SliceDel(seatData.Holds, index, 1)
		seatData.CountMap[pai]--
	}
	seatData.Pengs = append(seatData.Pengs, pai)
	game.ChuPai = -1
	m.recordGameAction(game, seatData.Seatindex, constant.ActionPeng, pai)
	//广播通知其它玩家
	userMgr.broadcastInRoom("peng_notify_push", map[string]any{"userid": seatData.Userid, "pai": pai}, seatData.Userid, true)
	//碰的玩家打牌
	m.moveToNextUser(game, seatData.Seatindex)
	//广播通知玩家出牌方
	seatData.CanChuPai = true
	userMgr.broadcastInRoom("game_chupai_push", seatData.Userid, seatData.Userid, true)
}

func (m *XzddMj) moveToNextUser(game *protos.MjGameData, nextSeat int32) {
	game.Fangpaoshumu = 0
	//找到下一个没有和牌的玩家
	if nextSeat == -1 {
		for {
			game.Turn++
			game.Turn %= 4
			var turnSeat = game.GameSeats[game.Turn]
			if turnSeat.Hued == false {
				return
			}
		}
	} else {
		game.Turn = nextSeat
	}
}

func (m *XzddMj) clearAllOptions(game *protos.MjGameData, seatData *protos.MjSeat) {
	fnClear := func(sd *protos.MjSeat) {
		sd.CanPeng = false
		sd.CanGang = false
		sd.GangPai = make([]int32, 0)
		sd.CanHu = false
		sd.LastFangGangSeat = -1
	}
	if seatData != nil {
		fnClear(seatData)
	} else {
		game.QiangGangContext = nil
		for _, seat := range game.GameSeats {
			fnClear(seat)
		}
	}
}

func (m *XzddMj) Gang(userId int64, pai int32) {
	seatData := m.getSeat(userId)
	if seatData == nil {
		logger.Error("can't find user game data.")
		return
	}
	seatIndex := seatData.Seatindex
	game := seatData.Game

	//如果没有杠的机会，则不能再杠
	if seatData.CanGang == false {
		logger.Error("seatData.Gang == false")
		return
	}

	//和的了，就不要再来了
	if seatData.Hued {
		logger.Error("you have already hued. no kidding plz.")
		return
	}

	if tools.IndexOf(seatData.GangPai, pai) == -1 {
		logger.Error("the given pai can't be ganged.")
		return
	}

	//如果有人可以胡牌，则需要等待
	i := game.Turn
	for {
		i = (i + 1) % 4
		if i == game.Turn {
			break
		} else {
			var ddd = game.GameSeats[i]
			if ddd.CanHu && i != seatData.Seatindex {
				return
			}
		}
	}

	numOfCnt := seatData.CountMap[pai]

	gangtype := ""
	//弯杠 去掉碰牌
	if numOfCnt == 1 {
		gangtype = "wangang"
	} else if numOfCnt == 3 {
		gangtype = "diangang"
	} else if numOfCnt == 4 {
		gangtype = "angang"
	} else {
		logger.Error("invalid pai count.")
		return
	}

	game.ChuPai = -1
	m.clearAllOptions(game, nil)
	seatData.CanChuPai = false

	userMgr.broadcastInRoom("hangang_notify_push", seatIndex, seatData.Userid, true)

	//如果是弯杠，则需要检查是否可以抢杠
	turnSeat := game.GameSeats[game.Turn]
	if numOfCnt == 1 {
		var canQiangGang = m.checkCanQiangGang(game, turnSeat, seatData, pai)
		if canQiangGang {
			return
		}
	}
	m.doGang(game, turnSeat, seatData, gangtype, numOfCnt, pai)
}

func (m *XzddMj) doGang(game *protos.MjGameData, turnSeat, seatData *protos.MjSeat, gangtype string, numOfCnt, pai int32) {
	seatIndex := seatData.Seatindex
	gameTurn := turnSeat.Seatindex

	isZhuanShouGang := false
	if gangtype == "wangang" {
		idx := tools.IndexOf(seatData.Pengs, pai)
		if idx >= 0 {
			seatData.Pengs = tools.SliceDel(seatData.Pengs, idx, 1)
		}

		//如果最后一张牌不是杠的牌，则认为是转手杠
		if seatData.Holds[len(seatData.Holds)-1] != pai {
			isZhuanShouGang = true
		}
	}
	//进行碰牌处理
	//扣掉手上的牌
	//从此人牌中扣除
	for i := 0; i < int(numOfCnt); i++ {
		index := tools.IndexOf(seatData.Holds, pai)
		if index == -1 {
			logger.Info(seatData.Holds)
			logger.Info("can't find mj.")
			return
		}
		seatData.Holds = tools.SliceDel(seatData.Holds, index, 1)
		seatData.CountMap[pai]--
	}

	m.recordGameAction(game, seatData.Seatindex, constant.ActionGang, pai)

	//记录下玩家的杠牌
	if gangtype == "angang" {
		seatData.Angangs = append(seatData.Angangs, pai)
		var ac = m.recordUserAction(game, seatData, "angang")
		ac.Score = game.Conf.BaseScore * 2
	} else if gangtype == "diangang" {
		seatData.Diangangs = append(seatData.Diangangs, pai)
		var ac = m.recordUserAction(game, seatData, "diangang", gameTurn)
		ac.Score = game.Conf.BaseScore * 2
		var fs = turnSeat
		m.recordUserAction(game, fs, "fanggang", seatIndex)
	} else if gangtype == "wangang" {
		seatData.Wangangs = append(seatData.Wangangs, pai)
		if isZhuanShouGang == false {
			var ac = m.recordUserAction(game, seatData, "wangang")
			ac.Score = game.Conf.BaseScore
		} else {
			m.recordUserAction(game, seatData, "zhuanshougang")
		}
	}

	m.checkCanTingPai(game, seatData)
	//通知其他玩家，有人杠了牌
	userMgr.broadcastInRoom("gang_notify_push", map[string]any{"userid": seatData.Userid, "pai": pai, "gangtype": gangtype}, seatData.Userid, true)

	//变成自己的轮子
	m.moveToNextUser(game, seatIndex)
	//再次摸牌
	m.doUserMoPai(game)

	//只能放在这里。因为过手就会清除杠牌标记
	seatData.LastFangGangSeat = gameTurn
}

func (m *XzddMj) doUserMoPai(game *protos.MjGameData) {
	game.ChuPai = -1
	turnSeat := game.GameSeats[game.Turn]
	turnSeat.LastFangGangSeat = -1
	turnSeat.GuoHuFan = -1
	pai := m.mopai(game, game.Turn)
	//牌摸完了，结束
	if pai == -1 {
		go m.doGameOver(game, turnSeat.Userid)
		return
	} else {
		numOfMJ := len(game.Mahjongs) - int(game.CurrentIndex)
		userMgr.broadcastInRoom("mj_count_push", numOfMJ, turnSeat.Userid, true)
	}

	m.recordGameAction(game, game.Turn, constant.ActionMopai, pai)

	//通知前端新摸的牌
	userMgr.sendMsg(turnSeat.Userid, "game_mopai_push", pai)
	//检查是否可以暗杠或者胡
	//检查胡，直杠，弯杠
	m.checkCanAnGang(game, turnSeat)
	m.checkCanWanGang(game, turnSeat)

	//检查看是否可以和
	m.checkCanHu(game, turnSeat, pai)

	//广播通知玩家出牌方
	turnSeat.CanChuPai = true
	userMgr.broadcastInRoom("game_chupai_push", turnSeat.Userid, turnSeat.Userid, true)

	//通知玩家做对应操作
	m.sendOperations(game, turnSeat, game.ChuPai)
}

func (m *XzddMj) doGameOver(game *protos.MjGameData, userId int64, args ...bool) {
	forceEnd := false
	if len(args) > 0 {
		forceEnd = args[0]
	}
	roomId := roomMgr.getUserRoom(userId)
	if len(roomId) < 1 {
		return
	}
	roomInfo := roomMgr.getRoom(roomId)
	if roomInfo == nil {
		return
	}
	var results []map[string]any
	dbresult := []int32{0, 0, 0, 0}
	fnNoticeResult := func(isEnd bool) {
		var endInfo []map[string]any
		if isEnd {
			for _, rs := range roomInfo.Seats {
				endInfo = append(endInfo, map[string]any{
					"numzimo":      rs.NumZiMo,
					"numjiepao":    rs.NumJiePao,
					"numdianpao":   rs.NumDianPao,
					"numangang":    rs.NumAnGang,
					"numminggang":  rs.NumMingGang,
					"numchadajiao": rs.NumChaJiao,
				})
			}
		}
		userMgr.broadcastInRoom("game_over_push", map[string]any{"results": results, "endinfo": endInfo}, userId, true)
		//如果局数已够，则进行整体结算，并关闭房间
		if isEnd {
			time.AfterFunc(1500*time.Millisecond, func() {
				if roomInfo.NumOfGames > 1 {
					m.storeHistory(roomInfo)
				}
				userMgr.kickAllInRoom(roomId)
				roomMgr.destroy(roomId)
				err := database.ArchiveMjActions(roomInfo.GameType, roomInfo.Uuid)
				if err != nil {
					logger.Error(err)
				} else {
					err = database.DeleteMjActions(roomInfo.Uuid)
					if err != nil {
						logger.Error(err)
					}
				}
			})
		}
	}

	if game != nil {
		if !forceEnd {
			m.calculateResult(game, roomInfo)
		}
		for i, rs := range roomInfo.Seats {
			sd := game.GameSeats[i]
			rs.Ready = false
			rs.Score += sd.Score
			rs.NumZiMo += sd.NumZiMo
			rs.NumJiePao += sd.NumJiePao
			rs.NumDianPao += sd.NumDianPao
			rs.NumAnGang += sd.NumAnGang
			rs.NumMingGang += sd.NumMingGang
			rs.NumChaJiao += sd.NumChaJiao
			userRT := map[string]any{
				"userId":     sd.Userid,
				"pengs":      sd.Pengs,
				"actions":    map[string]any{},
				"wangangs":   sd.Wangangs,
				"diangangs":  sd.Diangangs,
				"angangs":    sd.Angangs,
				"numofgen":   sd.Numofgen,
				"holds":      sd.Holds,
				"fan":        sd.Fan,
				"score":      sd.Score,
				"totalscore": rs.Score,
				"qingyise":   sd.Qingyise,
				"pattern":    sd.Pattern,
				"isganghu":   sd.IsGangHu,
				"menqing":    sd.IsMenQing,
				"zhongzhang": sd.IsZhongZhang,
				"jingouhu":   sd.IsJinGouHu,
				"haidihu":    sd.IsHaiDiHu,
				"tianhu":     sd.IsTianHu,
				"dihu":       sd.IsDiHu,
				"huorder":    tools.IndexOf(game.HupaiList, int32(i)),
			}
			var actions []map[string]any
			for _, ac := range sd.Actions {
				actions = append(actions, map[string]any{"type": ac.Type})
			}
			userRT["actions"] = actions
			results = append(results, userRT)

			dbresult[i] = sd.Score
			m.delSeat(sd.Userid)
		}
		m.delGame(roomId)
		old := roomInfo.NextButton
		if game.Yipaoduoxiang >= 0 {
			roomInfo.NextButton = game.Yipaoduoxiang
		} else if game.FirstHupai >= 0 {
			roomInfo.NextButton = game.FirstHupai
		} else {
			roomInfo.NextButton = (game.Turn + 1) % 4
		}
		if old != roomInfo.NextButton {
			err := database.UpdateMjRoom(roomId, map[string]any{"next_button": roomInfo.NextButton})
			if err != nil {
				logger.Error(err)
			}
		}
	}
	if forceEnd || game == nil {
		fnNoticeResult(true)
	} else {
		err := m.storeGame(game)
		if err != nil {
			logger.Error(err)
		}
		err = database.UpdateMjAction(roomInfo.Uuid, game.GameIndex, map[string]any{
			"result":         tools.Stringify(dbresult),
			"action_records": tools.Stringify(game.ActionList),
		})
		if err != nil {
			logger.Error(err)
		}
		err = database.UpdateMjRoom(roomId, map[string]any{"num_of_turns": roomInfo.NumOfGames})
		if err != nil {
			logger.Error(err)
		}
		if roomInfo.NumOfGames == 1 {
			var cost int32 = 2
			if roomInfo.Conf.MaxGames == 8 {
				cost = 3
			}
			if err := database.CostGems(userId, cost); err != nil {
				logger.Error(err)
			}
		}
		isEnd := roomInfo.NumOfGames >= roomInfo.Conf.MaxGames
		fnNoticeResult(isEnd)
	}
}

func (m *XzddMj) storeGame(game *protos.MjGameData) error {
	return database.CreateMjAction(&model.MahjongAction{
		RoomUUID:   game.Uuid,
		GameType:   game.Conf.Type,
		GameIndex:  game.GameIndex,
		BaseInfo:   game.BaseInfoJson,
		CreateTime: time.Now().Unix(),
	})
}

func (m *XzddMj) calculateResult(game *protos.MjGameData, roomInfo *protos.MjRoom) {
	isNeedChaDaJiao := m.needChaDaJiao(game)
	if isNeedChaDaJiao {
		m.chaJiao(game)
	}
	baseScore := game.Conf.BaseScore
	numOfHued := 0
	for _, sd := range game.GameSeats {
		if sd.Hued {
			numOfHued++
		}
	}
	for _, sd := range game.GameSeats {
		sd.NumAnGang = int32(len(sd.Angangs))
		sd.NumMingGang = int32(len(sd.Wangangs) + len(sd.Diangangs))
		//对所有胡牌的玩家进行统计
		if m.isTinged(sd) {
			//统计自己的番子和分数
			//基础番(平胡0番，对对胡1番、七对2番) + 清一色2番 + 杠+1番
			//杠上花+1番，杠上炮+1番 抢杠胡+1番，金钩胡+1番，海底胡+1番
			var fan = sd.Fan
			if m.isQingYiSe(sd) {
				sd.Qingyise = true
				fan += 2
			}

			var numOfGangs = len(sd.Diangangs) + len(sd.Wangangs) + len(sd.Angangs)
			for _, pai := range sd.Pengs {
				if sd.CountMap[pai] == 1 {
					numOfGangs++
				}
			}
			for _, c := range sd.CountMap {
				if c == 4 {
					numOfGangs++
				}
			}
			sd.Numofgen = int32(numOfGangs)

			//金钩胡
			if len(sd.Holds) == 1 || len(sd.Holds) == 2 {
				fan += 1
				sd.IsJinGouHu = true
			}

			if sd.IsHaiDiHu {
				fan += 1
			}

			if game.Conf.Tiandihu {
				if sd.IsTianHu {
					fan += 3
				} else if sd.IsDiHu {
					fan += 2
				}
			}

			var isjiangdui = false
			if game.Conf.Jiangdui {
				if sd.Pattern == "7pairs" {
					if sd.Numofgen > 0 {
						sd.Numofgen -= 1
						sd.Pattern = "l7pairs"
						isjiangdui = m.isJiangDui(sd)
						if isjiangdui {
							sd.Pattern = "j7paris"
							fan += 2
						} else {
							fan += 1
						}
					}
				} else if sd.Pattern == "duidui" {
					isjiangdui = m.isJiangDui(sd)
					if isjiangdui {
						sd.Pattern = "jiangdui"
						fan += 2
					}
				}
			}

			if game.Conf.Menqing {
				//不是将对，才检查中张
				if !isjiangdui {
					sd.IsZhongZhang = m.isZhongZhang(sd)
					if sd.IsZhongZhang {
						fan += 1
					}
				}

				sd.IsMenQing = m.isMenQing(sd)
				if sd.IsMenQing {
					fan += 1
				}
			}

			fan += sd.Numofgen
			if sd.IsGangHu {
				fan += 1
			}
			if sd.IsQiangGangHu {
				fan += 1
			}

			//收杠钱
			var additonalscore int32 = 0
			for a := 0; a < len(sd.Actions); a++ {
				var ac = sd.Actions[a]
				if ac.Type == "fanggang" {
					var ts = game.GameSeats[ac.Targets[0]]
					//检查放杠的情况，如果目标没有和牌，且没有叫牌，则不算 用于优化前端显示
					if isNeedChaDaJiao && (ts.Hued) == false && (m.isTinged(ts) == false) {
						ac.State = "nop"
					}
				} else if ac.Type == "angang" || ac.Type == "wangang" || ac.Type == "diangang" {
					if ac.State != "nop" {
						var acscore = ac.Score
						additonalscore += int32(len(ac.Targets)) * acscore * baseScore
						//扣掉目标方的分
						for t := 0; t < len(ac.Targets); t++ {
							var six = ac.Targets[t]
							game.GameSeats[six].Score -= acscore * baseScore
						}
					}
				} else if ac.Type == "maozhuanyu" {
					//对于呼叫转移，如果对方没有叫牌，表示不得行
					if m.isTinged(ac.Owner) {
						//如果
						var ref = ac.Ref
						var acscore = ref.Score
						var total = int32(len(ref.Targets)) * acscore * baseScore
						additonalscore += total
						//扣掉目标方的分
						if ref.PayTimes == 0 {
							for t := 0; t < len(ref.Targets); t++ {
								var six = ref.Targets[t]
								game.GameSeats[six].Score -= acscore * baseScore
							}
						} else {
							//如果已经被扣过一次了，则由杠牌这家赔
							ac.Owner.Score -= total
						}
						ref.PayTimes++
						ac.Owner = nil
						ac.Ref = nil
					}
				} else if ac.Type == "zimo" || ac.Type == "hu" || ac.Type == "ganghua" || ac.Type == "dianganghua" || ac.Type == "gangpaohu" || ac.Type == "qiangganghu" || ac.Type == "chadajiao" {
					var extraScore int32 = 0
					if ac.Iszimo {
						if game.Conf.Zimo == 0 {
							//自摸加底
							extraScore = baseScore
						}
						if game.Conf.Zimo == 1 {
							fan += 1
						} else {
							//nothing.
						}
						sd.NumZiMo++
					} else {
						if ac.Type != "chadajiao" {
							sd.NumJiePao++
						}
					}

					var score = m.computeFanScore(game, fan) + extraScore
					sd.Score += score * int32(len(ac.Targets))

					for t := 0; t < len(ac.Targets); t++ {
						var six = ac.Targets[t]
						var td = game.GameSeats[six]
						td.Score -= score
						if td != sd {
							if ac.Type == "chadajiao" {
								td.NumChaJiao++
							} else if !ac.Iszimo {
								td.NumDianPao++
							}
						}
					}
				}
			}
			if fan > game.Conf.MaxFan {
				fan = game.Conf.MaxFan
			}
			//一定要用 += 。 因为此时的sd.score可能是负的
			sd.Score += additonalscore
			if len(sd.Pattern) > 0 {
				sd.Fan = fan
			}
		} else {
			var arr []int32
			for a := len(sd.Actions) - 1; a >= 0; a-- {
				var ac = sd.Actions[a]
				if ac.Type == "angang" || ac.Type == "wangang" || ac.Type == "diangang" {
					//如果3家都胡牌，则需要结算。否则认为是查叫
					if numOfHued < 3 {
						arr = append(arr, int32(a))
					} else {
						if ac.State != "nop" {
							var acscore = ac.Score
							sd.Score += int32(len(ac.Targets)) * acscore * baseScore
							//扣掉目标方的分
							for t := 0; t < len(ac.Targets); t++ {
								var six = ac.Targets[t]
								game.GameSeats[six].Score -= acscore * baseScore
							}
						}
					}
				}
			}
			if len(arr) > 0 {
				for _, a := range arr {
					sd.Actions = append(sd.Actions[:a], sd.Actions[a+1:]...)
				}
			}
		}
	}
}

func (m *XzddMj) computeFanScore(game *protos.MjGameData, fan int32) int32 {
	if fan > game.Conf.MaxFan {
		fan = game.Conf.MaxFan
	}
	return (1 << fan) * game.Conf.BaseScore
}

func (m *XzddMj) isMenQing(gameSeatData *protos.MjSeat) bool {
	return len(gameSeatData.Pengs)+len(gameSeatData.Wangangs)+len(gameSeatData.Diangangs) == 0
}

func (m *XzddMj) isZhongZhang(gameSeatData *protos.MjSeat) bool {
	fn := func(arr []int32) bool {
		for i := 0; i < len(arr); i++ {
			var pai = arr[i]
			if pai == 0 || pai == 8 || pai == 9 || pai == 17 || pai == 18 || pai == 26 {
				return false
			}
		}
		return true
	}

	if fn(gameSeatData.Pengs) == false {
		return false
	}
	if fn(gameSeatData.Angangs) == false {
		return false
	}
	if fn(gameSeatData.Diangangs) == false {
		return false
	}
	if fn(gameSeatData.Wangangs) == false {
		return false
	}
	if fn(gameSeatData.Holds) == false {
		return false
	}
	return true
}

func (m *XzddMj) isJiangDui(gameSeatData *protos.MjSeat) bool {
	fn := func(arr []int32) bool {
		for i := 0; i < len(arr); i++ {
			var pai = arr[i]
			if pai != 1 && pai != 4 && pai != 7 && pai != 9 && pai != 13 && pai != 16 && pai != 18 && pai != 21 && pai != 25 {
				return false
			}
		}
		return true
	}

	if fn(gameSeatData.Pengs) == false {
		return false
	}
	if fn(gameSeatData.Angangs) == false {
		return false
	}
	if fn(gameSeatData.Diangangs) == false {
		return false
	}
	if fn(gameSeatData.Wangangs) == false {
		return false
	}
	if fn(gameSeatData.Holds) == false {
		return false
	}
	return true
}

func (m *XzddMj) isQingYiSe(gameSeatData *protos.MjSeat) bool {
	typ := m.getMJType(gameSeatData.Holds[0])
	//检查手上的牌
	if m.isSameType(typ, gameSeatData.Holds) == false {
		return false
	}
	//检查杠下的牌
	if m.isSameType(typ, gameSeatData.Angangs) == false {
		return false
	}
	if m.isSameType(typ, gameSeatData.Wangangs) == false {
		return false
	}
	if m.isSameType(typ, gameSeatData.Diangangs) == false {
		return false
	}

	//检查碰牌
	if m.isSameType(typ, gameSeatData.Pengs) == false {
		return false
	}
	return true
}

func (m *XzddMj) isSameType(typ int32, arr []int32) bool {
	for _, pai := range arr {
		t := m.getMJType(pai)
		if typ != -1 && typ != t {
			return false
		}
		typ = t
	}
	return true
}

func (m *XzddMj) storeHistory(roomInfo *protos.MjRoom) {
	seats := roomInfo.Seats
	var history = map[string]any{
		"uuid": roomInfo.Uuid,
		"id":   roomInfo.RoomId,
		"time": roomInfo.CreateTime,
	}
	nseats := make([]map[string]any, len(seats))
	for i, rs := range seats {
		nseats[i] = map[string]any{
			"userid": rs.Userid,
			"name":   base64.StdEncoding.EncodeToString([]byte(rs.Name)),
			"score":  rs.Score,
		}
	}
	history["seats"] = nseats
	for _, s := range seats {
		m.storeSingleHistory(s.Userid, history)
	}
}

func (m *XzddMj) storeSingleHistory(userId int64, history map[string]any) {
	user, err := database.GetUserById(userId)
	if err != nil {
		logger.Error(err)
		return
	}
	var res []map[string]any
	if len(user.History) > 0 {
		err = json.Unmarshal([]byte(user.History), &res)
		if err != nil {
			logger.Error(err)
			return
		}
	}
	if len(res) >= 10 {
		res = res[1:]
	}
	res = append(res, history)
	err = database.UpdateUser(userId, map[string]any{"history": tools.Stringify(res)})
	if err != nil {
		logger.Error(err)
	}
}

func (m *XzddMj) chaJiao(game *protos.MjGameData) {
	arr := m.findUnTingedPlayers(game)
	for _, ts := range game.GameSeats {
		//如果没有胡，但是听牌了，则未叫牌的人要给钱
		if !ts.Hued && m.isTinged(ts) {
			cur := m.findMaxFanTingPai(ts)
			ts.Fan = cur.Fan
			ts.Pattern = cur.Pattern
			m.recordUserAction(game, ts, "chadajiao", arr)
		}
	}
}
func (m *XzddMj) findMaxFanTingPai(ts *protos.MjSeat) *protos.MjTingData {
	//找出最大番
	var cur *protos.MjTingData
	for _, tpai := range ts.TingMap {
		if cur == nil || tpai.Fan > cur.Fan {
			cur = tpai
		}
	}
	return cur
}

func (m *XzddMj) findUnTingedPlayers(game *protos.MjGameData) []int32 {
	var arr []int32
	for i, ts := range game.GameSeats {
		//如果没有胡，且没有听牌
		if !ts.Hued && !m.isTinged(ts) {
			arr = append(arr, int32(i))
			m.recordUserAction(game, ts, "beichadajiao", -1)
		}
	}
	return arr
}

//是否需要查大叫(有两家以上未胡，且有人没有下叫)
func (m *XzddMj) needChaDaJiao(game *protos.MjGameData) bool {
	numOfHued := 0
	numOfTinged := 0
	numOfUntinged := 0
	for _, ts := range game.GameSeats {
		if ts.Hued {
			numOfHued++
			numOfTinged++
		} else if m.isTinged(ts) {
			numOfTinged++
		} else {
			numOfUntinged++
		}
	}
	//如果三家都胡牌了，不需要查叫
	if numOfHued == 3 {
		return false
	}

	//如果没有任何一个人叫牌，也没有任何一个胡牌，则不需要查叫
	if numOfTinged == 0 {
		return false
	}

	//如果都听牌了，也不需要查叫
	if numOfUntinged == 0 {
		return false
	}
	return true
}

func (m *XzddMj) isTinged(seatData *protos.MjSeat) bool {
	return len(seatData.TingMap) > 0
}

func (m *XzddMj) recordUserAction(game *protos.MjGameData, seatData *protos.MjSeat, typ string, args ...any) *protos.MjActionData {
	var target any
	if len(args) > 0 {
		target = args[0]
	}
	d := &protos.MjActionData{Type: typ, Targets: []int32{}}
	if target != nil {
		switch v := target.(type) {
		case int32:
			d.Targets = append(d.Targets, v)
		case []int32:
			d.Targets = v
		default:
			logger.Errorf("unknown target type:%+v", target)
		}
	} else {
		for i, s := range game.GameSeats {
			idx := int32(i)
			if idx != seatData.Seatindex && s.Hued == false {
				d.Targets = append(d.Targets, idx)
			}
		}
	}
	seatData.Actions = append(seatData.Actions, d)
	return d
}

//检查是否可以暗杠
func (m *XzddMj) checkCanAnGang(game *protos.MjGameData, seatData *protos.MjSeat) {
	//如果没有牌了，则不能再杠
	if len(game.Mahjongs) <= int(game.CurrentIndex) {
		return
	}

	for pai, c := range seatData.CountMap {
		if m.getMJType(pai) != seatData.Que && c == 4 {
			seatData.CanGang = true
			seatData.GangPai = append(seatData.GangPai, pai)
		}
	}
}

//检查是否可以弯杠(自己摸起来的时候)
func (m *XzddMj) checkCanWanGang(game *protos.MjGameData, seatData *protos.MjSeat) {
	//如果没有牌了，则不能再杠
	if len(game.Mahjongs) <= int(game.CurrentIndex) {
		return
	}
	for _, pai := range seatData.Pengs {
		if seatData.CountMap[pai] == 1 {
			seatData.CanGang = true
			seatData.GangPai = append(seatData.GangPai, pai)
		}
	}
}

func (m *XzddMj) checkCanQiangGang(game *protos.MjGameData, turnSeat *protos.MjSeat, seatData *protos.MjSeat, pai int32) bool {
	hasActions := false
	for i := 0; i < len(game.GameSeats); i++ {
		//杠牌者不检查
		if seatData.Seatindex == int32(i) {
			continue
		}
		ddd := game.GameSeats[i]
		//已经和牌的不再检查
		if ddd.Hued {
			continue
		}

		m.checkCanHu(game, ddd, pai)
		if ddd.CanHu {
			m.sendOperations(game, ddd, pai)
			hasActions = true
		}
	}
	if hasActions {
		game.QiangGangContext = &protos.MjQiangGangData{
			TurnSeat: turnSeat,
			SeatData: seatData,
			Pai:      pai,
			IsValid:  true,
		}
	} else {
		game.QiangGangContext = nil
	}
	return game.QiangGangContext != nil
}

func (m *XzddMj) checkCanHu(game *protos.MjGameData, seatData *protos.MjSeat, targetPai int32) {
	game.LastHuPaiSeat = -1
	if m.getMJType(targetPai) == seatData.Que {
		return
	}
	seatData.CanHu = false
	for k, _ := range seatData.TingMap {
		if targetPai == k {
			seatData.CanHu = true
		}
	}
}

func (m *XzddMj) Hu(userId int64) {
	seatData := m.getSeat(userId)
	if seatData == nil {
		logger.Error("can't find user game data.")
		return
	}
	var seatIndex = seatData.Seatindex
	var game = seatData.Game

	//如果他不能和牌，那和个啥啊
	if seatData.CanHu == false {
		logger.Info("invalid request.")
		return
	}

	//和的了，就不要再来了
	if seatData.Hued {
		logger.Info("you have already hued. no kidding plz.")
		return
	}

	//标记为和牌
	seatData.Hued = true
	var hupai = game.ChuPai
	var isZimo = false

	var turnSeat = game.GameSeats[game.Turn]
	seatData.IsGangHu = turnSeat.LastFangGangSeat >= 0
	var notify int32 = -1

	if game.QiangGangContext != nil {
		var gangSeat = game.QiangGangContext.SeatData
		hupai = game.QiangGangContext.Pai
		notify = hupai
		var ac = m.recordUserAction(game, seatData, "qiangganghu", gangSeat.Seatindex)
		ac.Iszimo = false
		m.recordGameAction(game, seatIndex, constant.ActionHu, hupai)
		seatData.IsQiangGangHu = true
		game.QiangGangContext.IsValid = false

		idx := tools.IndexOf(gangSeat.Holds, hupai)
		if idx != -1 {
			gangSeat.Holds = tools.SliceDel(gangSeat.Holds, idx, 1)
			gangSeat.CountMap[hupai]--
			userMgr.sendMsg(gangSeat.Userid, "game_holds_push", gangSeat.Holds)
		}
		//将牌添加到玩家的手牌列表，供前端显示
		seatData.Holds = append(seatData.Holds, hupai)
		seatData.CountMap[hupai]++

		m.recordUserAction(game, gangSeat, "beiqianggang", seatIndex)
	} else if game.ChuPai == -1 {
		hupai = seatData.Holds[len(seatData.Holds)-1]
		notify = -1
		if seatData.IsGangHu {
			if turnSeat.LastFangGangSeat == seatIndex {
				var ac = m.recordUserAction(game, seatData, "ganghua")
				ac.Iszimo = true
			} else {
				var diangganghuaZimo = game.Conf.Dianganghua == 1
				if diangganghuaZimo {
					var ac = m.recordUserAction(game, seatData, "dianganghua")
					ac.Iszimo = true
				} else {
					var ac = m.recordUserAction(game, seatData, "dianganghua", turnSeat.LastFangGangSeat)
					ac.Iszimo = false
				}
			}
		} else {
			var ac = m.recordUserAction(game, seatData, "zimo")
			ac.Iszimo = true
		}

		isZimo = true
		m.recordGameAction(game, seatIndex, constant.ActionZimo, hupai)
	} else {
		notify = game.ChuPai
		//将牌添加到玩家的手牌列表，供前端显示
		seatData.Holds = append(seatData.Holds, game.ChuPai)
		seatData.CountMap[game.ChuPai]++

		logger.Info(seatData.Holds)

		var at = "hu"
		//炮胡
		if turnSeat.LastFangGangSeat >= 0 {
			at = "gangpaohu"
		}

		var ac = m.recordUserAction(game, seatData, at, game.Turn)
		ac.Iszimo = false

		//毛转雨
		if turnSeat.LastFangGangSeat >= 0 {
			for i := len(turnSeat.Actions) - 1; i >= 0; i-- {
				var t = turnSeat.Actions[i]
				if t.Type == "diangang" || t.Type == "wangang" || t.Type == "angang" {
					t.State = "nop"
					t.PayTimes = 0

					var nac = &protos.MjActionData{
						Type:  "maozhuanyu",
						Owner: turnSeat,
						Ref:   t,
					}
					seatData.Actions = append(seatData.Actions, nac)
					break
				}
			}
		}

		//记录玩家放炮信息
		fs := game.GameSeats[game.Turn]
		m.recordUserAction(game, fs, "fangpao", seatIndex)

		m.recordGameAction(game, seatIndex, constant.ActionHu, hupai)

		game.Fangpaoshumu++

		if game.Fangpaoshumu > 1 {
			game.Yipaoduoxiang = seatIndex
		}
	}

	if game.FirstHupai < 0 {
		game.FirstHupai = seatIndex
	}

	//保存番数
	ti := seatData.TingMap[hupai]
	seatData.Fan = ti.Fan
	seatData.Pattern = ti.Pattern
	seatData.Iszimo = isZimo
	//如果是最后一张牌，则认为是海底胡
	seatData.IsHaiDiHu = int(game.CurrentIndex) == len(game.Mahjongs)
	game.HupaiList = append(game.HupaiList, seatData.Seatindex)

	if game.Conf.Tiandihu {
		if game.ChupaiCnt == 0 && game.Button == seatData.Seatindex && game.ChuPai == -1 {
			seatData.IsTianHu = true
		} else if game.ChupaiCnt == 1 && game.Turn == game.Button && game.Button != seatData.Seatindex && game.ChuPai != -1 {
			seatData.IsDiHu = true
		}
	}

	m.clearAllOptions(game, seatData)

	//通知前端，有人和牌了
	userMgr.broadcastInRoom("hu_push", map[string]any{"seatindex": seatIndex, "iszimo": isZimo, "hupai": notify}, seatData.Userid, true)

	//
	if game.LastHuPaiSeat == -1 {
		game.LastHuPaiSeat = seatIndex
	} else {
		var lp = (game.LastFangGangSeat - game.Turn + 4) % 4
		var cur = (seatData.Seatindex - game.Turn + 4) % 4
		if cur > lp {
			game.LastHuPaiSeat = seatData.Seatindex
		}
	}

	//如果只有一家没有胡，则结束
	var numOfHued = 0
	for _, ddd := range game.GameSeats {
		if ddd.Hued {
			numOfHued++
		}
	}
	//和了三家
	if numOfHued == 3 {
		go m.doGameOver(game, seatData.Userid)
		return
	}

	//清空所有非胡牌操作
	for _, ddd := range game.GameSeats {
		ddd.CanPeng = false
		ddd.CanGang = false
		ddd.CanChuPai = false
		m.sendOperations(game, ddd, hupai)
	}
	//如果还有人可以胡牌，则等待
	for _, ddd := range game.GameSeats {
		if ddd.CanHu {
			return
		}
	}

	//和牌的下家继续打
	m.clearAllOptions(game, nil)
	game.Turn = game.LastHuPaiSeat
	m.moveToNextUser(game, -1)
	m.doUserMoPai(game)
}

func (m *XzddMj) Guo(userId int64) {
	seatData := m.getSeat(userId)
	if seatData == nil {
		logger.Error("can't find user game data.")
		return
	}
	seatIndex := seatData.Seatindex
	game := seatData.Game
	if !m.hasOperations(seatData) {
		logger.Error("no need Guo.")
		return
	}

	doNothing := game.ChuPai == -1 && game.Turn == seatIndex
	userMgr.sendMsg(seatData.Userid, "guo_result")
	m.clearAllOptions(game, seatData)

	//这里还要处理过胡的情况
	if game.ChuPai >= 0 && seatData.CanHu {
		seatData.GuoHuFan = seatData.TingMap[game.ChuPai].Fan
	}

	if doNothing {
		return
	}

	//如果还有人可以操作，则等待
	for i := 0; i < len(game.GameSeats); i++ {
		var ddd = game.GameSeats[i]
		if m.hasOperations(ddd) {
			return
		}
	}

	//如果是已打出的牌，则需要通知。
	if game.ChuPai >= 0 {
		uid := game.GameSeats[game.Turn].Userid
		userMgr.broadcastInRoom("guo_notify_push", map[string]any{"Userid": uid, "pai": game.ChuPai}, seatData.Userid, true)
		seatData.Folds = append(seatData.Folds, game.ChuPai)
		game.ChuPai = -1
	}
	qiangGangContext := game.QiangGangContext
	//清除所有的操作
	m.clearAllOptions(game, nil)

	if qiangGangContext != nil && qiangGangContext.IsValid {
		m.doGang(game, qiangGangContext.TurnSeat, qiangGangContext.SeatData, "wangang", 1, qiangGangContext.Pai)
	} else {
		//下家摸牌
		m.moveToNextUser(game, -1)
		m.doUserMoPai(game)
	}
}

func init() {
	RegGameMgr("xzdd", NewXzddMj())
}