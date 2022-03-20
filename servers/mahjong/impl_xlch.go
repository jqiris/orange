package mahjong

import (
	"encoding/base64"
	"encoding/json"
	"math/rand"
	"time"

	"github.com/jqiris/orange/constant"
	"github.com/jqiris/orange/database"
	"github.com/jqiris/orange/model"
	"github.com/jqiris/orange/tools"
	"github.com/jqiris/saki/card"
	"github.com/jqiris/saki/wall"

	"sync"

	"github.com/jqiris/kungfu/v2/logger"
)

type XlchMj struct {
	Games            map[string]*model.MjGameData
	GameSeatsOfUsers map[int64]*model.MjSeat
	DissolvingList   []string
	lockGame         sync.RWMutex
	lockSeat         sync.RWMutex
}

func NewXlchMj() *XlchMj {
	return &XlchMj{
		Games:            make(map[string]*model.MjGameData),
		GameSeatsOfUsers: make(map[int64]*model.MjSeat),
		DissolvingList:   make([]string, 0),
	}
}
func (m *XlchMj) getGame(roomId string) *model.MjGameData {
	m.lockGame.RLock()
	defer m.lockGame.RUnlock()
	if v, ok := m.Games[roomId]; ok {
		return v
	}
	return nil
}
func (m *XlchMj) saveGame(roomId string, data *model.MjGameData) {
	m.lockGame.Lock()
	defer m.lockGame.Unlock()
	m.Games[roomId] = data
}
func (m *XlchMj) delGame(roomId string) {
	m.lockGame.Lock()
	defer m.lockGame.Unlock()
	delete(m.Games, roomId)
}

func (m *XlchMj) getSeat(userId int64) *model.MjSeat {
	m.lockSeat.RLock()
	defer m.lockSeat.RUnlock()
	if v, ok := m.GameSeatsOfUsers[userId]; ok {
		return v
	}
	return nil
}
func (m *XlchMj) saveSeat(userId int64, data *model.MjSeat) {
	m.lockSeat.Lock()
	defer m.lockSeat.Unlock()
	m.GameSeatsOfUsers[userId] = data
}

func (m *XlchMj) delSeat(userId int64) {
	m.lockSeat.Lock()
	defer m.lockSeat.Unlock()
	delete(m.GameSeatsOfUsers, userId)
}

func (m *XlchMj) String() string {
	return "四川麻将血流成河"
}

func (m *XlchMj) deal(game *model.MjGameData) {
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

func (m *XlchMj) mopai(game *model.MjGameData, seatIndex int) int {
	if game.Wall.IsAllDrawn() {
		return -1
	}
	data := game.GameSeats[seatIndex]
	pai := game.Wall.ForwardDraw()
	data.Holds.AddTile(pai, 1)
	return pai
}

//是否需要查大叫(有两家以上未胡，且有人没有下叫)
func (m *XlchMj) needChaDaJiao(game *model.MjGameData) bool {
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
	//如果没有任何一个人叫牌，则不需要查叫
	if numOfTinged == 0 {
		return false
	}

	//如果都听牌了，也不需要查叫
	if numOfUntinged == 0 {
		return false
	}
	return true
}

func (m *XlchMj) chaJiao(game *model.MjGameData) {
	arr := m.findUnTingedPlayers(game)
	if len(arr) == 0 {
		return
	}
	for i, ts := range game.GameSeats {
		//如果听牌了，则未叫牌的人要给钱
		if m.isTinged(ts) {
			cur := m.findMaxFanTingPai(ts)
			ts.Huinfo = append(ts.Huinfo, &model.MjHuData{
				Ishupai:  true,
				Pai:      cur.Pai,
				Action:   "chadajiao",
				Fan:      cur.Fan,
				Pattern:  cur.Pattern,
				Numofgen: m.getNumOfGen(ts),
			})
			for _, j := range arr {
				st := game.GameSeats[j]
				st.Huinfo = append(st.Huinfo, &model.MjHuData{
					Action: "beichadajiao",
					Target: int(i),
					Index:  int(len(ts.Huinfo) - 1),
				})
			}
		}
	}
}

func (m *XlchMj) isTinged(seatData *model.MjSeat) bool {
	return len(seatData.TingMap) > 0
}

func (m *XlchMj) isMenQing(gameSeatData *model.MjSeat) bool {
	return len(gameSeatData.Pengs)+len(gameSeatData.Wangangs)+len(gameSeatData.Diangangs) == 0
}

func (m *XlchMj) isZhongZhang(gameSeatData *model.MjSeat) bool {
	fn := func(arr []int) bool {
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
	if fn(gameSeatData.Holds.GetList()) == false {
		return false
	}
	return true
}

func (m *XlchMj) isJiangDui(gameSeatData *model.MjSeat) bool {
	fn := func(arr []int) bool {
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
	if fn(gameSeatData.Holds.GetList()) == false {
		return false
	}
	return true
}

func (m *XlchMj) isQingYiSe(gameSeatData *model.MjSeat) bool {
	typ := gameSeatData.Holds.GetIndexType(0)
	//检查手上的牌
	if m.isSameType(typ, gameSeatData.Holds.GetList()) == false {
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

func (m *XlchMj) isSameType(typ int, arr []int) bool {
	for _, pai := range arr {
		t := card.GetMjType(pai)
		if typ != -1 && typ != t {
			return false
		}
		typ = t
	}
	return true
}

func (m *XlchMj) moveToNextUser(game *model.MjGameData, nextSeat int) {
	game.Fangpaoshumu = 0
	//找到下一个没有和牌的玩家
	if nextSeat == -1 {
		game.Turn++
		game.Turn %= 4
		return
	} else {
		game.Turn = nextSeat
	}
}

func (m *XlchMj) findUnTingedPlayers(game *model.MjGameData) []int {
	var arr []int
	for i, ts := range game.GameSeats {
		//如果没有胡，且没有听牌
		if !ts.Hued && !m.isTinged(ts) {
			arr = append(arr, int(i))
		}
	}
	return arr
}

func (m *XlchMj) findMaxFanTingPai(ts *model.MjSeat) *model.MjTingData {
	//找出最大番
	var cur *model.MjTingData
	for k, tpai := range ts.TingMap {
		if cur == nil || tpai.Fan > cur.Fan {
			cur = tpai
			cur.Pai = k
		}
	}
	return cur
}

func (m *XlchMj) getNumOfGen(seatData *model.MjSeat) int {
	numOfGangs := int(len(seatData.Diangangs) + len(seatData.Wangangs) + len(seatData.Angangs))
	for _, pai := range seatData.Pengs {
		if seatData.Holds.GetTileCnt(pai) == 1 {
			numOfGangs++
		}
	}
	gangs := seatData.Holds.GetNumTiles(4)
	numOfGangs += len(gangs)
	return numOfGangs
}

func (m *XlchMj) hasOperations(seatData *model.MjSeat) bool {
	if seatData.CanGang || seatData.CanPeng || seatData.CanHu {
		return true
	}
	return false
}

func (m *XlchMj) sendOperations(game *model.MjGameData, seatData *model.MjSeat, pai int) {
	if m.hasOperations(seatData) {
		if pai == -1 {
			pai = seatData.Holds.Peek()
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

func (m *XlchMj) calculateResult(game *model.MjGameData, roomInfo *model.MjRoom) {
	isNeedChaDaJiao := m.needChaDaJiao(game)
	if isNeedChaDaJiao {
		m.chaJiao(game)
	}
	baseScore := game.Conf.BaseScore
	for _, sd := range game.GameSeats {
		//对所有胡牌的玩家进行统计
		if m.isTinged(sd) {
			//收杠钱
			var additionalScore int = 0
			for _, ac := range sd.Actions {
				if ac.Type == "fanggang" {
					ts := game.GameSeats[ac.Targets[0]]
					//检查放杠的情况，如果目标没有和牌，且没有叫牌，则不算 用于优化前端显示
					if isNeedChaDaJiao && ts.Hued == false && m.isTinged(ts) == false {
						ac.State = "nop"
					}
				} else if ac.Type == "angang" || ac.Type == "wangang" || ac.Type == "diangang" {
					if ac.State != "nop" {
						acScore := ac.Score
						additionalScore += int(len(ac.Targets)) * acScore * baseScore
						//扣掉目标方的分
						for _, six := range ac.Targets {
							game.GameSeats[six].Score -= acScore * baseScore
						}
					}
				} else if ac.Type == "maozhuanyu" {
					//对于呼叫转移，如果对方没有叫牌，表示不得行
					if m.isTinged(ac.Owner) {
						ref := ac.Ref
						acScore := ref.Score
						total := int(len(ref.Targets)) * acScore * baseScore
						additionalScore += total
						//扣掉目标方的分
						if ref.PayTimes == 0 {
							for _, six := range ref.Targets {
								game.GameSeats[six].Score -= acScore * baseScore
							}
						} else {
							//如果已经被扣过一次了，则由杠牌这家赔
							ac.Owner.Score -= total
						}
						ref.PayTimes++
						ac.Owner = nil
						ac.Ref = nil
					}
				}
			}

			if m.isQingYiSe(sd) {
				sd.Qingyise = true
			}

			if game.Conf.Menqing {
				sd.IsMenQing = m.isMenQing(sd)
			}

			//金钩胡
			listNum := sd.Holds.GetListNum()
			if listNum == 1 || listNum == 2 {
				sd.IsJinGouHu = true
			}
			sd.NumAnGang = int(len(sd.Angangs))
			sd.NumMingGang = int(len(sd.Wangangs) + len(sd.Diangangs))
			//进行胡牌结算
			for _, info := range sd.Huinfo {
				if !info.Ishupai {
					sd.NumDianPao++
					continue
				}
				//统计自己的番子和分数
				//基础番(平胡0番，对对胡1番、七对2番) + 清一色2番 + 杠+1番
				//杠上花+1番，杠上炮+1番 抢杠胡+1番，金钩胡+1番，海底胡+1番
				fan := info.Fan
				sd.Holds.AddTile(info.Pai, 1)
				if sd.Qingyise {
					fan += 2
				}
				//金钩胡
				if sd.IsJinGouHu {
					fan += 1
				}

				if info.IsHaiDiHu {
					fan += 1
				}
				if game.Conf.Tiandihu {
					if info.IsTianHu {
						fan += 3
					} else if info.IsDiHu {
						fan += 2
					}
				}
				isJiangDui := false
				if game.Conf.Jiangdui {
					if info.Pattern == "7pairs" {
						if info.Numofgen > 0 {
							info.Numofgen -= 1
							info.Pattern = "l7pairs"
							isJiangDui = m.isJiangDui(sd)
							if isJiangDui {
								info.Pattern = "j7pairs"
								fan += 2
							} else {
								fan += 1
							}
						}
					} else if info.Pattern == "duidui" {
						isJiangDui = m.isJiangDui(sd)
						if isJiangDui {
							info.Pattern = "jiangdui"
							fan += 2
						}
					}
				}
				if game.Conf.Menqing {
					//不是将对，才检查中张
					if !isJiangDui {
						sd.IsZhongZhang = m.isJiangDui(sd)
						if sd.IsZhongZhang {
							fan += 1
						}
					}
					if sd.IsMenQing {
						fan += 1
					}
				}
				fan += info.Numofgen
				if info.Action == "ganghua" || info.Action == "dianganghua" || info.Action == "gangpaohu" || info.Action == "qiangganghu" {
					fan += 1
				}
				var extraScore int = 0
				if info.Iszimo {
					if game.Conf.Zimo == 0 {
						//自摸加底
						extraScore = baseScore
					} else if game.Conf.Zimo == 1 {
						fan += 1
					} else {
						//nothing
					}
				}
				//和牌的玩家才加这个分
				score := m.computeFanScore(game, fan) + extraScore
				if info.Action == "chadajiao" {
					//收所有没有叫牌的人的钱
					for _, st := range game.GameSeats {
						if !m.isTinged(st) {
							st.Score -= score
							sd.Score += score
							//被查叫次数
							if st != sd {
								st.NumChaJiao++
							}
						}
					}
				} else if info.Iszimo {
					//收所有人的钱
					sd.Score += score * int(len(game.GameSeats))
					for _, st := range game.GameSeats {
						st.Score -= score
					}
					sd.NumZiMo++
				} else {
					//收放炮者的钱
					sd.Score += score
					game.GameSeats[info.Target].Score -= score
					sd.NumJiePao++
				}

				//撤除胡的那张牌
				sd.Holds.Pop()
				if fan > game.Conf.MaxFan {
					fan = game.Conf.MaxFan
				}
				info.Fan = fan
			}
			//一定要用 += 。 因为此时的sd.score可能是负的
			sd.Score += additionalScore
		} else {
			var arr []int
			for a := len(sd.Actions) - 1; a >= 0; a-- {
				var ac = sd.Actions[a]
				if ac.Type == "angang" || ac.Type == "wangang" || ac.Type == "diangang" {
					//如果3家都胡牌，则需要结算。否则认为是查叫
					if isNeedChaDaJiao {
						arr = append(arr, int(a))
					} else {
						if ac.State != "nop" {
							var acscore = ac.Score
							sd.Score += int(len(ac.Targets)) * acscore * baseScore
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

func (m *XlchMj) computeFanScore(game *model.MjGameData, fan int) int {
	if fan > game.Conf.MaxFan {
		fan = game.Conf.MaxFan
	}
	return (1 << fan) * game.Conf.BaseScore
}

func (m *XlchMj) begin(roomId string) {
	roomInfo := roomMgr.getRoom(roomId)
	if roomInfo == nil {
		return
	}
	seats := roomInfo.Seats
	w := wall.NewWall()
	w.SetTiles(card.MahjongCards108)
	game := &model.MjGameData{
		Conf:          roomInfo.Conf,
		Uuid:          roomInfo.Uuid,
		GameIndex:     roomInfo.NumOfGames,
		Button:        roomInfo.NextButton,
		GameSeats:     make([]*model.MjSeat, 4),
		NumOfQue:      0,
		Turn:          0,
		ChuPai:        -1,
		State:         "idle",
		FirstHupai:    -1,
		Yipaoduoxiang: -1,
		Fangpaoshumu:  -1,
		ActionList:    []int{},
		ChupaiCnt:     0,
		Wall:          w,
	}
	roomInfo.NumOfGames++

	for i := 0; i < 4; i++ {
		data := &model.MjSeat{}
		data.Game = game
		data.Seatindex = int(i)
		data.Userid = seats[i].Userid
		data.Holds = card.NewCMap()
		data.Folds = card.NewCMap()
		data.Angangs = make([]int, 0)
		data.Wangangs = make([]int, 0)
		data.Diangangs = make([]int, 0)
		data.Pengs = make([]int, 0)
		data.Que = -1
		data.Huanpais = make([]int, 0)
		data.TingMap = make(map[int]*model.MjTingData)
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
		data.Actions = make([]*model.MjActionData, 0)
		data.Fan = 0
		data.Score = 0
		data.Huinfo = make([]*model.MjHuData, 0)
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
	// m.shuffle(game)
	game.Wall.Shuffle()
	m.deal(game)

	numOfMJ := game.Wall.RemainLength()
	huansanzhang := roomInfo.Conf.Hsz

	for i := 0; i < len(seats); i++ {
		st := seats[i]
		holds := game.GameSeats[i].Holds.GetList()
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

func (m *XlchMj) doGameOver(game *model.MjGameData, userId int64, args ...bool) {
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
	dbresult := []int{0, 0, 0, 0}
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
				"holds":      sd.Holds,
				"score":      sd.Score,
				"totalscore": rs.Score,
				"qingyise":   sd.Qingyise,
				"menqing":    sd.IsMenQing,
				"jingouhu":   sd.IsJinGouHu,
				"huinfo":     sd.Huinfo,
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
			var cost int = 2
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

func (m *XlchMj) checkCanTingPai(game *model.MjGameData, seatData *model.MjSeat) {
	seatData.TingMap = make(map[int]*model.MjTingData)
	//检查手上的牌是不是已打缺，如果未打缺，则不进行判定
	for _, pai := range seatData.Holds.GetUnique() {
		if card.IsSameType(pai, seatData.Que) {
			return
		}
	}
	//检查是否是七对 前提是没有碰，也没有杠 ，即手上拥有13张牌
	if seatData.Holds.GetListNum() == 13 {
		//有5对牌
		var danPai int = -1
		pairCount := 0
		for k, c := range seatData.Holds.GetTileMap() {
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
			seatData.TingMap[danPai] = &model.MjTingData{
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
	for k, c := range seatData.Holds.GetTileMap() {
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
				seatData.TingMap[p] = &model.MjTingData{
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
func (m *XlchMj) storeGame(game *model.MjGameData) error {
	return database.CreateMjAction(&model.MahjongAction{
		RoomUUID:   game.Uuid,
		GameType:   game.Conf.Type,
		GameIndex:  game.GameIndex,
		BaseInfo:   game.BaseInfoJson,
		CreateTime: time.Now().Unix(),
	})
}

func (m *XlchMj) storeHistory(roomInfo *model.MjRoom) {
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

func (m *XlchMj) storeSingleHistory(userId int64, history map[string]any) {
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
func (m *XlchMj) SetReady(userId int64) {
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
		numOfMJ := game.Wall.RemainLength()
		data := map[string]any{
			"state":   game.State,
			"numofmj": numOfMJ,
			"button":  game.Button,
			"turn":    game.Turn,
			"chuPai":  game.ChuPai,
		}
		var seats []*model.MjSeat
		var seatData *model.MjSeat
		for i := 0; i < 4; i++ {
			sd := game.GameSeats[i]
			st := &model.MjSeat{
				Userid:    sd.Userid,
				Folds:     sd.Folds,
				Angangs:   sd.Angangs,
				Diangangs: sd.Diangangs,
				Wangangs:  sd.Wangangs,
				Pengs:     sd.Pengs,
				Que:       sd.Que,
				Hued:      sd.Hued,
				Iszimo:    sd.Iszimo,
				Huinfo:    sd.Huinfo,
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

func (m *XlchMj) HasBegan(roomId string) bool {
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

func (m *XlchMj) DissolveRequest(roomId string, userId int64) *model.MjRoom {
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
	dr := &model.MjDissolveData{
		EndTime: time.Now().UnixMilli() + 30000,
		States:  []bool{false, false, false, false},
	}
	dr.States[seatIndex] = true
	roomInfo.Dr = dr
	m.DissolvingList = append(m.DissolvingList, roomId)
	return roomInfo
}

func (m *XlchMj) DissolveAgree(roomId string, userId int64, agree bool) *model.MjRoom {
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

func (m *XlchMj) DoDissolve(roomId string) {
	roomInfo := roomMgr.getRoom(roomId)
	if roomInfo == nil {
		return
	}
	game := m.getGame(roomId)
	go m.doGameOver(game, roomInfo.Seats[0].Userid, true)
}

func (m *XlchMj) HuanSanZhang(userId int64, p1, p2, p3 int) {
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
	tileMaps := seatData.Holds.GetTileMap()
	if tileMaps[p1] == 0 || tileMaps[p2] == 0 || tileMaps[p3] == 0 {
		return
	}

	seatData.Huanpais = []int{p1, p2, p3}

	for _, p := range seatData.Huanpais {
		seatData.Holds.DelTile(p, 1)
	}
	userMgr.sendMsg(seatData.Userid, "game_holds_push", seatData.Holds.GetList())

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
				"Huanpais": []int{},
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
	var fn = func(s1 *model.MjSeat, huanjin []int) {
		for i := 0; i < len(huanjin); i++ {
			var p = huanjin[i]
			s1.Holds.AddTile(p, 1)
		}
	}

	//开始换牌
	var f = rand.Intn(100)
	var s = game.GameSeats
	var huanpaiMethod int = 0
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
		userid := gs.Userid
		userMgr.sendMsg(userid, "game_huanpai_over_push", rd)

		userMgr.sendMsg(userid, "game_holds_push", gs.Holds.GetList())
		//通知准备定缺
		userMgr.sendMsg(userid, "game_dingque_push")
	}
}

func (m *XlchMj) DingQue(userId int64, que int) {
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
func (m *XlchMj) constructGameBaseInfo(game *model.MjGameData) {
	var baseInfo = map[string]any{
		"type":     game.Conf.Type,
		"button":   game.Button,
		"index":    game.GameIndex,
		"mahjongs": game.Wall.GetTiles(),
	}
	seats := make([][]int, 4)
	for i := 0; i < 4; i++ {
		seats[i] = game.GameSeats[i].Holds.GetList()
	}
	baseInfo["game_seats"] = seats
	bs, _ := json.Marshal(baseInfo)
	game.BaseInfoJson = string(bs)
}
func (m *XlchMj) doDingQue(game *model.MjGameData, seatData *model.MjSeat) {
	m.constructGameBaseInfo(game)
	arr := []int{1, 1, 1, 1}
	for i, gs := range game.GameSeats {
		arr[i] = gs.Que
	}
	userMgr.broadcastInRoom("game_dingque_finish_push", arr, seatData.Userid, true)
	userMgr.broadcastInRoom("game_playing_push", nil, seatData.Userid, true)

	//进行听牌检查
	for _, gs := range game.GameSeats {
		var duoyu int = -1
		listNum := gs.Holds.GetListNum()
		if listNum == 14 {
			duoyu = gs.Holds.Pop()
		}
		m.checkCanTingPai(game, gs)
		if duoyu >= 0 {
			gs.Holds.AddTile(duoyu, 1)
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
	m.checkCanHu(game, turnSeat, turnSeat.Holds.Peek())
	//通知前端
	m.sendOperations(game, turnSeat, game.ChuPai)
	// logger.Infof("doDingQue,seat:%+v,turn:%+v", game.GameSeats[game.Turn], game.Turn)
}

//检查是否可以暗杠
func (m *XlchMj) checkCanAnGang(game *model.MjGameData, seatData *model.MjSeat) {
	//如果没有牌了，则不能再杠
	if game.Wall.IsAllDrawn() {
		return
	}

	for pai, c := range seatData.Holds.GetTileMap() {
		if card.GetMjType(pai) != seatData.Que && c == 4 {
			seatData.CanGang = true
			seatData.GangPai = append(seatData.GangPai, pai)
		}
	}
}

//检查是否可以弯杠(自己摸起来的时候)
func (m *XlchMj) checkCanWanGang(game *model.MjGameData, seatData *model.MjSeat) {
	//如果没有牌了，则不能再杠
	if game.Wall.IsAllDrawn() {
		return
	}
	for _, pai := range seatData.Pengs {
		if seatData.Holds.GetTileCnt(pai) == 1 {
			seatData.CanGang = true
			seatData.GangPai = append(seatData.GangPai, pai)
		}
	}
}

func (m *XlchMj) checkCanHu(game *model.MjGameData, seatData *model.MjSeat, targetPai int) {
	game.LastHuPaiSeat = -1
	if card.GetMjType(targetPai) == seatData.Que {
		return
	}
	seatData.CanHu = false
	for k, _ := range seatData.TingMap {
		if targetPai == k {
			seatData.CanHu = true
		}
	}
}

func (m *XlchMj) checkCanDianGang(game *model.MjGameData, seatData *model.MjSeat, targetPai int) {
	//检查玩家手上的牌
	//如果没有牌了，则不能再杠
	if game.Wall.IsAllDrawn() {
		return
	}
	if card.GetMjType(targetPai) == seatData.Que {
		return
	}
	count := seatData.Holds.GetTileCnt(targetPai)
	if count >= 3 {
		seatData.CanGang = true
		seatData.GangPai = append(seatData.GangPai, targetPai)
		return
	}
}
func (m *XlchMj) checkCanPeng(game *model.MjGameData, seatData *model.MjSeat, targetPai int) {
	if card.GetMjType(targetPai) == seatData.Que {
		return
	}
	if seatData.Holds.GetTileCnt(targetPai) >= 2 {
		seatData.CanPeng = true
	}
}

func (m *XlchMj) ChuPai(userId int64, pai int) {
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
	if seatData.CanChuPai == false {
		logger.Error("no need chupai")
		return
	}
	if m.hasOperations(seatData) {
		logger.Error("plz Guo before you chupai")
		return
	}
	//如果是胡了的人，则只能打最后一张牌
	if seatData.Hued {
		if seatData.Holds.Peek() != pai {
			logger.Error("only deal last one when hued.")
			return
		}
	}
	cnt := seatData.Holds.GetTileCnt(pai)
	if cnt == 0 {
		logger.Errorf("can't find the mj,holds:%+v,pai:%+v", seatData.Holds, pai)
		return
	}
	seatData.CanChuPai = false
	game.ChupaiCnt++
	seatData.GuoHuFan = -1
	seatData.Holds.DelTile(pai, 1)
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

		//未胡牌的才检查杠和碰
		if !ddd.Hued {
			m.checkCanPeng(game, ddd, pai)
			m.checkCanDianGang(game, ddd, pai)
		}
		m.checkCanHu(game, ddd, pai)
		if seatData.LastFangGangSeat == -1 {
			if v, ok := seatData.TingMap[pai]; ok && ddd.CanHu && ddd.GuoHuFan >= 0 && v.Fan <= ddd.GuoHuFan {
				ddd.CanHu = false
				userMgr.sendMsg(ddd.Userid, "guohu_push")
			}
		}

		if m.hasOperations(ddd) {
			m.sendOperations(game, ddd, game.ChuPai)
			hasActions = true
		}
	}
	//如果没有人有操作，则向下一家发牌，并通知他出牌
	if !hasActions {
		time.AfterFunc(500*time.Millisecond, func() {
			userMgr.broadcastInRoom("guo_notify_push", map[string]any{"userId": seatData.Userid, "pai": game.ChuPai}, seatData.Userid, true)
			seatData.Folds.AddTile(game.ChuPai, 1)
			game.ChuPai = -1
			m.moveToNextUser(game, -1)
			m.doUserMoPai(game)
		})
	}
}

func (m *XlchMj) Peng(userId int64) {
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
	c := seatData.Holds.GetTileCnt(pai)
	if c < 2 {
		logger.Infof("pai:%v,count:%v", pai, c)
		logger.Info(seatData.Holds)
		logger.Info("lack of mj.")
		return
	}

	//进行碰牌处理
	//扣掉手上的牌
	//从此人牌中扣除
	if !seatData.Holds.DelTile(pai, 2) {
		logger.Info("can't find mj.")
		return
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

func (m *XlchMj) doUserMoPai(game *model.MjGameData) {
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
		numOfMJ := game.Wall.RemainLength()
		userMgr.broadcastInRoom("mj_count_push", numOfMJ, turnSeat.Userid, true)
	}

	m.recordGameAction(game, game.Turn, constant.ActionMopai, pai)

	//通知前端新摸的牌
	userMgr.sendMsg(turnSeat.Userid, "game_mopai_push", pai)
	//检查是否可以暗杠或者胡
	//检查胡，直杠，弯杠
	if !turnSeat.Hued {
		m.checkCanAnGang(game, turnSeat)
	}
	//如果未胡牌，或者摸起来的牌可以杠，才检查弯杠
	if !turnSeat.Hued || turnSeat.Holds.Peek() == pai {
		m.checkCanWanGang(game, turnSeat)
	}
	//检查看是否可以和
	m.checkCanHu(game, turnSeat, pai)

	//广播通知玩家出牌方
	turnSeat.CanChuPai = true
	userMgr.broadcastInRoom("game_chupai_push", turnSeat.Userid, turnSeat.Userid, true)

	//通知玩家做对应操作
	m.sendOperations(game, turnSeat, game.ChuPai)
}

func (m *XlchMj) clearAllOptions(game *model.MjGameData, seatData *model.MjSeat) {
	fnClear := func(sd *model.MjSeat) {
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

func (m *XlchMj) recordGameAction(game *model.MjGameData, si, action, pai int) {
	game.ActionList = append(game.ActionList, si, action)
	if pai > -1 {
		game.ActionList = append(game.ActionList, pai)
	}
}

func (m *XlchMj) Gang(userId int64, pai int) {
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

	numOfCnt := seatData.Holds.GetTileCnt(pai)

	//胡了的，只能直杠
	if numOfCnt != 1 && seatData.Hued {
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

func (m *XlchMj) recordUserAction(game *model.MjGameData, seatData *model.MjSeat, typ string, args ...any) *model.MjActionData {
	var target any
	if len(args) > 0 {
		target = args[0]
	}
	d := &model.MjActionData{Type: typ, Targets: []int{}}
	if target != nil {
		switch v := target.(type) {
		case int:
			d.Targets = append(d.Targets, v)
		case []int:
			d.Targets = v
		default:
			logger.Errorf("unknown target type:%+v", target)
		}
	} else {
		for i, _ := range game.GameSeats {
			idx := int(i)
			//血流成河，所有自摸，暗杠，弯杠，都算三家
			if idx != seatData.Seatindex {
				d.Targets = append(d.Targets, idx)
			}
		}
	}
	seatData.Actions = append(seatData.Actions, d)
	return d
}

func (m *XlchMj) doGang(game *model.MjGameData, turnSeat, seatData *model.MjSeat, gangtype string, numOfCnt, pai int) {
	seatIndex := seatData.Seatindex
	gameTurn := turnSeat.Seatindex

	isZhuanShouGang := false
	if gangtype == "wangang" {
		idx := tools.IndexOf(seatData.Pengs, pai)
		if idx >= 0 {
			seatData.Pengs = tools.SliceDel(seatData.Pengs, idx, 1)
		}

		//如果最后一张牌不是杠的牌，则认为是转手杠
		if seatData.Holds.Peek() != pai {
			isZhuanShouGang = true
		}
	}
	//进行碰牌处理
	//扣掉手上的牌
	//从此人牌中扣除
	if !seatData.Holds.DelTile(pai, numOfCnt) {
		logger.Info("can't find mj.")
		return
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
func (m *XlchMj) checkCanQiangGang(game *model.MjGameData, turnSeat *model.MjSeat, seatData *model.MjSeat, pai int) bool {
	hasActions := false
	for i := 0; i < len(game.GameSeats); i++ {
		//杠牌者不检查
		if seatData.Seatindex == int(i) {
			continue
		}
		ddd := game.GameSeats[i]

		m.checkCanHu(game, ddd, pai)
		if ddd.CanHu {
			m.sendOperations(game, ddd, pai)
			hasActions = true
		}
	}
	if hasActions {
		game.QiangGangContext = &model.MjQiangGangData{
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

func (m *XlchMj) Hu(userId int64) {
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

	//标记为和牌
	seatData.Hued = true
	var hupai = game.ChuPai
	var isZimo = false

	var turnSeat = game.GameSeats[game.Turn]
	huData := &model.MjHuData{
		Ishupai:       true,
		Pai:           -1,
		Action:        "",
		IsGangHu:      false,
		IsQiangGangHu: false,
		Iszimo:        false,
		Target:        -1,
		Fan:           0,
		Pattern:       "",
		IsHaiDiHu:     false,
		IsTianHu:      false,
		IsDiHu:        false,
	}
	huData.Numofgen = m.getNumOfGen(seatData)
	huData.IsGangHu = turnSeat.LastFangGangSeat >= 0
	seatData.Huinfo = append(seatData.Huinfo, huData)
	var notify int = -1

	if game.QiangGangContext != nil {
		var gangSeat = game.QiangGangContext.SeatData
		hupai = game.QiangGangContext.Pai
		notify = hupai
		huData.Iszimo = false
		huData.Action = "qiangganghu"
		huData.IsQiangGangHu = true
		huData.Target = gangSeat.Seatindex
		huData.Pai = hupai
		m.recordGameAction(game, seatIndex, constant.ActionHu, hupai)
		game.QiangGangContext.IsValid = false

		if gangSeat.Holds.DelTile(hupai, 1) {
			userMgr.sendMsg(gangSeat.Userid, "game_holds_push", gangSeat.Holds.GetList())
		}
		gangSeat.Huinfo = append(gangSeat.Huinfo, &model.MjHuData{
			Action: "beiqianggang",
			Target: seatData.Seatindex,
			Index:  int(len(seatData.Huinfo) - 1),
		})
	} else if game.ChuPai == -1 {
		hupai = seatData.Holds.Pop()
		notify = hupai
		huData.Pai = hupai
		if huData.IsGangHu {
			if turnSeat.LastFangGangSeat == seatIndex {
				huData.Action = "ganghua"
				huData.Iszimo = true
			} else {
				var diangganghuaZimo = game.Conf.Dianganghua == 1
				huData.Action = "dianganghua"
				huData.Iszimo = diangganghuaZimo
				huData.Target = turnSeat.LastFangGangSeat
			}
		} else {
			huData.Action = "zimo"
			huData.Iszimo = true
		}

		isZimo = true
		m.recordGameAction(game, seatIndex, constant.ActionZimo, hupai)
	} else {
		notify = game.ChuPai
		huData.Pai = hupai
		var at = "hu"
		//炮胡
		if turnSeat.LastFangGangSeat >= 0 {
			at = "gangpaohu"
		}

		huData.Action = at
		huData.Iszimo = false
		huData.Target = game.Turn

		//毛转雨
		if turnSeat.LastFangGangSeat >= 0 {
			for i := len(turnSeat.Actions) - 1; i >= 0; i-- {
				var t = turnSeat.Actions[i]
				if t.Type == "diangang" || t.Type == "wangang" || t.Type == "angang" {
					t.State = "nop"
					t.PayTimes = 0

					var nac = &model.MjActionData{
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
		if at == "gangpaohu" {
			at = "gangpao"
		} else {
			at = "fangpao"
		}
		fs.Huinfo = append(fs.Huinfo, &model.MjHuData{
			Action: at,
			Target: seatData.Seatindex,
			Index:  int(len(seatData.Huinfo) - 1),
		})

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
	huData.Fan = ti.Fan
	huData.Pattern = ti.Pattern
	huData.Iszimo = isZimo
	//如果是最后一张牌，则认为是海底胡
	huData.IsHaiDiHu = game.Wall.IsAllDrawn()

	if game.Conf.Tiandihu {
		if game.ChupaiCnt == 0 && game.Button == seatData.Seatindex && game.ChuPai == -1 {
			huData.IsTianHu = true
		} else if game.ChupaiCnt == 1 && game.Turn == game.Button && game.Button != seatData.Seatindex && game.ChuPai != -1 {
			huData.IsDiHu = true
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

func (m *XlchMj) Guo(userId int64) {
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
		seatData.Folds.AddTile(game.ChuPai, 1)
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
	RegGameMgr("xlch", NewXlchMj())
}
