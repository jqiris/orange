package mahjong

import "github.com/jqiris/orange/protos"

func checkTingPai(seatData *protos.MjSeat, begin, end int32) {
	for i := begin; i < end; i++ {
		//如果这牌已经在和了，就不用检查了
		if seatData.TingMap[i] != nil {
			continue
		}
		//将牌加入到计数中
		old := seatData.CountMap[i]
		seatData.CountMap[i]++

		seatData.Holds = append(seatData.Holds, i)
		//逐个判定手上的牌
		var ret = checkCanHu(seatData)
		if ret {
			//平胡 0番
			seatData.TingMap[i] = &protos.MjTingData{
				Pattern: "normal",
				Fan:     0,
			}
		}

		//搞完以后，撤消刚刚加的牌
		seatData.CountMap[i] = old
		seatData.Holds = seatData.Holds[:len(seatData.Holds)-1]
	}
}

func checkCanHu(seatData *protos.MjSeat) bool {
	for k, c := range seatData.CountMap {
		if c < 2 {
			continue
		}
		//如果当前牌大于等于２，则将它选为将牌
		seatData.CountMap[k] -= 2
		//逐个判定剩下的牌是否满足　３Ｎ规则,一个牌会有以下几种情况
		//1、0张，则不做任何处理
		//2、2张，则只可能是与其它牌形成匹配关系
		//3、3张，则可能是单张形成 A-2,A-1,A  A-1,A,A+1  A,A+1,A+2，也可能是直接成为一坎
		//4、4张，则只可能是一坎+单张
		ret := checkSingle(seatData)
		seatData.CountMap[k] += 2
		if ret {
			//kanzi.push(k);
			//kanzi.push(k);
			//console.log(kanzi);
			return true
		}
	}
	return false
}

func checkSingle(seatData *protos.MjSeat) bool {
	holds := seatData.Holds
	var selected int32 = -1
	var c int32 = 0
	for _, pai := range holds {
		c = seatData.CountMap[pai]
		if c != 0 {
			selected = pai
			break
		}
	}
	//如果没有找到剩余牌，则表示匹配成功了
	if selected == -1 {
		return true
	}
	//否则，进行匹配
	if c == 3 {
		//直接作为一坎
		seatData.CountMap[selected] = 0
		var ret = checkSingle(seatData)
		//立即恢复对数据的修改
		seatData.CountMap[selected] = c
		if ret == true {
			return true
		}
	} else if c == 4 {
		//直接作为一坎
		seatData.CountMap[selected] = 1
		var ret = checkSingle(seatData)
		//立即恢复对数据的修改
		seatData.CountMap[selected] = c
		//如果作为一坎能够把牌匹配完，直接返回TRUE。
		if ret == true {
			return true
		}
	}

	//按单牌处理
	return matchSingle(seatData, selected)
}

func matchSingle(seatData *protos.MjSeat, selected int32) bool {
	//分开匹配 A-2,A-1,A
	matched := true
	v := selected % 9
	if v < 2 {
		matched = false
	} else {
		for i := 0; i < 3; i++ {
			var t int32 = selected - 2 + int32(i)
			var cc = seatData.CountMap[t]
			if cc == 0 {
				matched = false
				break
			}
		}
	}

	//匹配成功，扣除相应数值
	if matched {
		seatData.CountMap[selected-2]--
		seatData.CountMap[selected-1]--
		seatData.CountMap[selected]--
		var ret = checkSingle(seatData)
		seatData.CountMap[selected-2]++
		seatData.CountMap[selected-1]++
		seatData.CountMap[selected]++
		if ret == true {
			return true
		}
	}

	//分开匹配 A-1,A,A + 1
	matched = true
	if v < 1 || v > 7 {
		matched = false
	} else {
		for i := 0; i < 3; i++ {
			var t int32 = selected - 1 + int32(i)
			var cc = seatData.CountMap[t]
			if cc == 0 {
				matched = false
				break
			}
		}
	}

	//匹配成功，扣除相应数值
	if matched {
		seatData.CountMap[selected-1]--
		seatData.CountMap[selected]--
		seatData.CountMap[selected+1]--
		var ret = checkSingle(seatData)
		seatData.CountMap[selected-1]++
		seatData.CountMap[selected]++
		seatData.CountMap[selected+1]++
		if ret == true {
			return true
		}
	}

	//分开匹配 A,A+1,A + 2
	matched = true
	if v > 6 {
		matched = false
	} else {
		for i := 0; i < 3; i++ {
			var t int32 = selected + int32(i)
			var cc = seatData.CountMap[t]
			if cc == 0 {
				matched = false
				break
			}
		}
	}

	//匹配成功，扣除相应数值
	if matched {
		seatData.CountMap[selected]--
		seatData.CountMap[selected+1]--
		seatData.CountMap[selected+2]--
		var ret = checkSingle(seatData)
		seatData.CountMap[selected]++
		seatData.CountMap[selected+1]++
		seatData.CountMap[selected+2]++
		if ret == true {
			return true
		}
	}
	return false
}
