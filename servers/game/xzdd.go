package game

import (
	"math"
	"math/rand"

	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/orange/constant"
	"github.com/jqiris/orange/tools"
)

type XzddMj struct {
	games            map[int]*GameData
	gameSeatsOfUsers map[int]*Seat
}

func NewXzddMj() *XzddMj {
	return &XzddMj{
		games:            make(map[int]*GameData),
		gameSeatsOfUsers: make(map[int]*Seat),
	}
}

func (m *XzddMj) String() string {
	return "四川麻将血战到底"
}
func (m *XzddMj) begin(roomId int) {
	roomInfo := roomMgr.getRoom(roomId)
	if roomInfo == nil {
		return
	}
	seats := roomInfo.Seats
	game := &GameData{
		Conf:          roomInfo.Conf,
		RoomInfo:      roomInfo,
		GameIndex:     roomInfo.NumOfGames,
		Button:        roomInfo.NextButton,
		Mahjongs:      make([]int, 108),
		CurrentIndex:  0,
		GameSeats:     make([]*Seat, 4),
		NumOfQue:      0,
		Turn:          0,
		ChuPai:        -1,
		State:         "idle",
		FirstHupai:    -1,
		Yipaoduoxiang: -1,
		Fangpaoshumu:  -1,
		ActionList:    []int{},
		HupaiList:     []int{},
		ChupaiCnt:     0,
	}
	roomInfo.NumOfGames++

	for i := 0; i < 4; i++ {
		data := &Seat{}
		data.Game = game
		data.SeatIndex = i
		data.UserId = seats[i].UserId
		data.Holds = make([]int, 0)
		data.Folds = make([]int, 0)
		data.Angangs = make([]int, 0)
		data.Wangangs = make([]int, 0)
		data.Pengs = make([]int, 0)
		data.Que = -1
		data.Huanpais = make([]int, 0)
		data.CountMap = make(map[int]int)
		data.TingMap = make(map[int]*TingData)
		data.Pattern = ""
		data.CanGang = false
		data.GangPai = make([]int, 0)
		data.CanPeng = false
		data.CanHu = false
		data.CanChuPai = false
		data.GuoHuFan = -1
		data.Hued = false
		data.Iszimo = false
		data.IsGangHu = false
		data.Actions = make([]int, 0)
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
		m.gameSeatsOfUsers[data.UserId] = data
	}
	m.games[roomId] = game
	m.shuffle(game)
	m.deal(game)

	numOfMJ := len(game.Mahjongs) - game.CurrentIndex
	huansanzhang := roomInfo.Conf.Hsz

	for i := 0; i < len(seats); i++ {
		st := seats[i]
		//通知玩家手牌
		userMgr.sendMsg(st.UserId, "game_holds_push", game.GameSeats[i].Holds)
		//通知还剩多少张牌
		userMgr.sendMsg(st.UserId, "mj_count_push", numOfMJ)
		//通知还剩多少局
		userMgr.sendMsg(st.UserId, "game_num_push", roomInfo.NumOfGames)
		//通知游戏开始
		userMgr.sendMsg(st.UserId, "game_begin_push", game.Button)

		if huansanzhang == true {
			game.State = "huanpai"
			//通知准备换牌
			userMgr.sendMsg(st.UserId, "game_huanpai_push", nil)
		} else {
			game.State = "dingque"
			//通知准备定缺
			userMgr.sendMsg(st.UserId, "game_dingque_push", nil)
		}
	}
}

func (m *XzddMj) deal(game *GameData) {
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

func (m *XzddMj) mopai(game *GameData, seatIndex int) int {
	if game.CurrentIndex >= len(game.Mahjongs) {
		return -1
	}
	data := game.GameSeats[seatIndex]
	mahjongs := data.Holds
	pai := game.Mahjongs[game.CurrentIndex]
	mahjongs = append(mahjongs, pai)
	data.CountMap[pai]++
	game.CurrentIndex++
	return pai
}

func (m *XzddMj) shuffle(game *GameData) {
	mahjongs := game.Mahjongs
	index := 0
	for i := 0; i < 9; i++ {
		for c := 0; c < 4; c++ {
			mahjongs[index] = i
			index++
		}
	}
	for i := 9; i < 18; i++ {
		for c := 0; c < 4; c++ {
			mahjongs[index] = i
			index++
		}
	}
	for i := 18; i < 27; i++ {
		for c := 0; c < 4; c++ {
			mahjongs[index] = i
			index++
		}
	}
	mjLen := len(mahjongs)
	for i := 0; i < mjLen; i++ {
		lastIndex := mjLen - i - 1
		index := int(math.Floor(float64(rand.Intn(100)*lastIndex) / 100))
		mahjongs[index], mahjongs[lastIndex] = mahjongs[lastIndex], mahjongs[index]
	}
}

func (m *XzddMj) setReady(userId int) {
	roomId := roomMgr.getUserRoom(userId)
	if roomId < 1 {
		return
	}
	roomInfo := roomMgr.getRoom(roomId)
	if roomInfo == nil {
		return
	}
	roomMgr.setReady(userId, true)

	game := m.games[roomId]
	if game == nil {
		if len(roomInfo.Seats) == 4 {
			for i := 0; i < 4; i++ {
				st := roomInfo.Seats[i]
				if !st.Ready || !userMgr.isOnline(st.UserId) {
					return
				}
			}
			//4个人到齐了，并且都准备好了，则开始新的一局
			m.begin(roomId)
		}
	} else {
		numOfMJ := len(game.Mahjongs) - game.CurrentIndex
		data := map[string]interface{}{
			"state":         game.State,
			"numofmj":       numOfMJ,
			"button":        game.Button,
			"turn":          game.Turn,
			"chuPai":        game.ChuPai,
			"huanpaimethod": game.HuanpaiMethod,
		}
		var seats []*Seat
		var seatData *Seat
		for i := 0; i < 4; i++ {
			sd := game.GameSeats[i]
			st := &Seat{
				UserId:    sd.UserId,
				Folds:     sd.Folds,
				Angangs:   sd.Angangs,
				Diangangs: sd.Diangangs,
				Wangangs:  sd.Wangangs,
				Pengs:     sd.Pengs,
				Que:       sd.Que,
				Hued:      sd.Hued,
				Iszimo:    sd.Iszimo,
			}
			if sd.UserId == userId {
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

func (m *XzddMj) sendOperations(game *GameData, seatData *Seat, pai int) {
	if m.hasOperations(seatData) {
		if pai == -1 {
			pai = seatData.Holds[len(seatData.Holds)-1]
		}
		data := map[string]interface{}{
			"pai":     pai,
			"hu":      seatData.CanHu,
			"peng":    seatData.CanPeng,
			"gang":    seatData.CanGang,
			"gangpai": seatData.GangPai,
			"si":      seatData.SeatIndex,
		}
		//如果可以有操作，则进行操作
		userMgr.sendMsg(seatData.UserId, "game_action_push", data)
	} else {
		userMgr.sendMsg(seatData.UserId, "game_action_push", nil)
	}
}

func (m *XzddMj) hasOperations(seatData *Seat) bool {
	if seatData.CanGang || seatData.CanPeng || seatData.CanHu {
		return true
	}
	return false
}

func (m *XzddMj) hasBegan(roomId int) bool {
	game := m.games[roomId]
	if game != nil {
		return true
	}
	roomInfo := roomMgr.getRoom(roomId)
	if roomInfo != nil {
		return roomInfo.NumOfGames > 0
	}
	return false
}

func (m *XzddMj) dissolveRequest(roomId int, userId int) {
	panic("not implemented") // TODO: Implement
}

func (m *XzddMj) dissolveAgree(roomId int, userId int, agree bool) {
	panic("not implemented") // TODO: Implement
}

func (m *XzddMj) doDissolve(roomId int) {
	panic("not implemented") // TODO: Implement
}

func (m *XzddMj) huanSanZhang(userId int, p1 int, p2 int, p3 int) {
	panic("not implemented") // TODO: Implement
}

func (m *XzddMj) dingQue(userId int, que int) {
	panic("not implemented") // TODO: Implement
}

func (m *XzddMj) chuPai(userId int, pai int) {
	seatData := m.gameSeatsOfUsers[userId]
	if seatData == nil {
		logger.Error("can't find the user game data")
		return
	}

	game := seatData.Game
	if game.Turn != seatData.SeatIndex {
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
		logger.Error("plz guo before you chupai")
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
	m.recordGameAction(game, game.Turn, constant.ACTION_CHUPAI, pai)

}

func (m *XzddMj) checkCanTingPai(game *GameData, seatData *Seat) {
	seatData.TingMap = make(map[int]*TingData)
	//检查手上的牌是不是已打缺，如果未打缺，则不进行判定
	for _, pai := range seatData.Holds {
		if m.getMJType(pai) == seatData.Que {
			return
		}
	}
	//检查是否是七对 前提是没有碰，也没有杠 ，即手上拥有13张牌
	if len(seatData.Holds) == 13 {
		//有5对牌
		danPai := -1
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
			seatData.TingMap[danPai] = &TingData{
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
	var arr []int
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
				seatData.TingMap[p] = &TingData{
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

func (m *XzddMj) getMJType(id int) int {
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

func (m *XzddMj) recordGameAction(game *GameData, si, action, pai int) {
	game.ActionList = append(game.ActionList, si, action)
	if pai > -1 {
		game.ActionList = append(game.ActionList, pai)
	}
}

func (m *XzddMj) peng(userId int) {
	seatData := m.gameSeatsOfUsers[userId]
	if seatData == nil {
		logger.Error("can't find user game data.")
		return
	}
	game := seatData.Game
	//如果是他出的牌，则忽略
	if game.Turn == seatData.SeatIndex {
		logger.Error("it's your turn.")
		return
	}

	//如果没有碰的机会，则不能再碰
	if seatData.CanPeng == false {
		logger.Error("seatData.peng == false")
		return
	}

	//和的了，就不要再来了
	if seatData.Hued {
		logger.Error("you have already hued. no kidding plz.")
		return
	}

	//如果有人可以胡牌，则需要等待
	var i = game.Turn
	for {
		var i = (i + 1) % 4
		if i == game.Turn {
			break
		} else {
			var ddd = game.GameSeats[i]
			if ddd.CanHu && i != seatData.SeatIndex {
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
	m.recordGameAction(game, seatData.SeatIndex, constant.ACTION_PENG, pai)
	//广播通知其它玩家
	userMgr.broadcastInRoom("peng_notify_push", map[string]interface{}{"userid": seatData.UserId, "pai": pai}, seatData.UserId, true)
	//碰的玩家打牌
	m.moveToNextUser(game, seatData.SeatIndex)
	//广播通知玩家出牌方
	seatData.CanChuPai = true
	userMgr.broadcastInRoom("game_chupai_push", seatData.UserId, seatData.UserId, true)
}

func (m *XzddMj) moveToNextUser(game *GameData, nextSeat int) {
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

func (m *XzddMj) clearAllOptions(game *GameData, seatData *Seat) {
	fnClear := func(sd *Seat) {
		sd.CanPeng = false
		sd.CanGang = false
		sd.GangPai = make([]int, 0)
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

func (m *XzddMj) gang(userId int, pai int) {
	seatData := m.gameSeatsOfUsers[userId]
	if seatData == nil {
		logger.Error("can't find user game data.")
		return
	}
	seatIndex := seatData.SeatIndex
	game := seatData.Game

	//如果没有杠的机会，则不能再杠
	if seatData.CanGang == false {
		logger.Error("seatData.gang == false")
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
	var i = game.Turn
	for {
		var i = (i + 1) % 4
		if i == game.Turn {
			break
		} else {
			var ddd = game.GameSeats[i]
			if ddd.CanHu && i != seatData.SeatIndex {
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

	userMgr.broadcastInRoom("hangang_notify_push", seatIndex, seatData.UserId, true)

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

func (m *XzddMj) doGang(game *GameData, turnSeat, seatData *Seat, gangtype string, numOfCnt, pai int) {
	seatIndex := seatData.SeatIndex
	gameTurn := turnSeat.SeatIndex

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
	for i := 0; i < numOfCnt; i++ {
		index := tools.IndexOf(seatData.Holds, pai)
		if index == -1 {
			logger.Info(seatData.Holds)
			logger.Info("can't find mj.")
			return
		}
		seatData.Holds = tools.SliceDel(seatData.Holds, index, 1)
		seatData.CountMap[pai]--
	}

	m.recordGameAction(game, seatData.SeatIndex, constant.ACTION_GANG, pai)

	//记录下玩家的杠牌
	if gangtype == "angang" {
		seatData.Angangs = append(seatData.Angangs, pai)
		var ac = m.recordUserAction(game, seatData, "angang")
		ac["score"] = game.Conf.BaseScore * 2
	} else if gangtype == "diangang" {
		seatData.Diangangs = append(seatData.Diangangs, pai)
		var ac = m.recordUserAction(game, seatData, "diangang", gameTurn)
		ac["score"] = game.Conf.BaseScore * 2
		var fs = turnSeat
		m.recordUserAction(game, fs, "fanggang", seatIndex)
	} else if gangtype == "wangang" {
		seatData.Wangangs = append(seatData.Wangangs, pai)
		if isZhuanShouGang == false {
			var ac = m.recordUserAction(game, seatData, "wangang")
			ac["score"] = game.Conf.BaseScore
		} else {
			m.recordUserAction(game, seatData, "zhuanshougang")
		}
	}

	m.checkCanTingPai(game, seatData)
	//通知其他玩家，有人杠了牌
	userMgr.broadcastInRoom("gang_notify_push", map[string]interface{}{"userid": seatData.UserId, "pai": pai, "gangtype": gangtype}, seatData.UserId, true)

	//变成自己的轮子
	m.moveToNextUser(game, seatIndex)
	//再次摸牌
	m.doUserMoPai(game)

	//只能放在这里。因为过手就会清除杠牌标记
	seatData.LastFangGangSeat = gameTurn
}

func (m *XzddMj) doUserMoPai(game *GameData) {
	game.ChuPai = -1
	turnSeat := game.GameSeats[game.Turn]
	turnSeat.LastFangGangSeat = -1
	turnSeat.GuoHuFan = -1
	pai := m.mopai(game, game.Turn)
	//牌摸完了，结束
	if pai == -1 {
		m.doGameOver(game, turnSeat.UserId)
		return
	} else {
		numOfMJ := len(game.Mahjongs) - game.CurrentIndex
		userMgr.broadcastInRoom("mj_count_push", numOfMJ, turnSeat.UserId, true)
	}

	m.recordGameAction(game, game.Turn, constant.ACTION_MOPAI, pai)

	//通知前端新摸的牌
	userMgr.sendMsg(turnSeat.UserId, "game_mopai_push", pai)
	//检查是否可以暗杠或者胡
	//检查胡，直杠，弯杠
	m.checkCanAnGang(game, turnSeat)
	m.checkCanWanGang(game, turnSeat)

	//检查看是否可以和
	m.checkCanHu(game, turnSeat, pai)

	//广播通知玩家出牌方
	turnSeat.CanChuPai = true
	userMgr.broadcastInRoom("game_chupai_push", turnSeat.UserId, turnSeat.UserId, true)

	//通知玩家做对应操作
	m.sendOperations(game, turnSeat, game.ChuPai)
}

func (m *XzddMj) doGameOver(game *GameData, UserId int, forceEnd ...bool) {
	//todo doGameOver
}

func (m *XzddMj) recordUserAction(game *GameData, seatData *Seat, typ string, target ...interface{}) map[string]interface{} {
	//todo recordUserAction
	return nil
}

//检查是否可以暗杠
func (m *XzddMj) checkCanAnGang(game *GameData, seatData *Seat) {
	//如果没有牌了，则不能再杠
	if len(game.Mahjongs) <= game.CurrentIndex {
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
func (m *XzddMj) checkCanWanGang(game *GameData, seatData *Seat) {
	//如果没有牌了，则不能再杠
	if len(game.Mahjongs) <= game.CurrentIndex {
		return
	}
	for _, pai := range seatData.Pengs {
		if seatData.CountMap[pai] == 1 {
			seatData.CanGang = true
			seatData.GangPai = append(seatData.GangPai, pai)
		}
	}
}

func (m *XzddMj) checkCanQiangGang(game *GameData, turnSeat *Seat, seatData *Seat, pai int) bool {
	hasActions := false
	for i := 0; i < len(game.GameSeats); i++ {
		//杠牌者不检查
		if seatData.SeatIndex == i {
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
		game.QiangGangContext = &QiangGangData{
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

func (m *XzddMj) checkCanHu(game *GameData, seatData *Seat, targetPai int) {
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

func (m *XzddMj) hu(userId int) {
	panic("not implemented") // TODO: Implement
}

func (m *XzddMj) guo(userId int) {
	seatData := m.gameSeatsOfUsers[userId]
	if seatData == nil {
		logger.Error("can't find user game data.")
		return
	}
	seatIndex := seatData.SeatIndex
	game := seatData.Game
	if !m.hasOperations(seatData) {
		logger.Error("no need guo.")
		return
	}

	doNothing := game.ChuPai == -1 && game.Turn == seatIndex
	userMgr.sendMsg(seatData.UserId, "guo_result", nil)
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
		uid := game.GameSeats[game.Turn].UserId
		userMgr.broadcastInRoom("guo_notify_push", map[string]interface{}{"userId": uid, "pai": game.ChuPai}, seatData.UserId, true)
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
