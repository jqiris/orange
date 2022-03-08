package mahjong

import (
	"fmt"
	"time"

	"github.com/jqiris/kungfu/v2/utils"
	"github.com/jqiris/orange/constant"
)

type TokenData struct {
	UserId   int64 `json:"user_id"`
	Time     int64 `json:"time"`
	LifeTime int64 `json:"life_time"`
}

type TokenMgr struct {
	tokens map[string]TokenData
	users  map[int64]string
}

func NewTokenMgr() *TokenMgr {
	return &TokenMgr{
		tokens: make(map[string]TokenData),
		users:  make(map[int64]string),
	}
}

func (m *TokenMgr) createToken(userId, lifeTime int64) string {
	if token, ok := m.users[userId]; ok {
		m.delToken(token)
	}
	nowTime := time.Now().Unix()
	token := utils.Md5(fmt.Sprintf("%v%v%v", userId, constant.TOKEN_KEY, nowTime))
	m.tokens[token] = TokenData{
		UserId:   userId,
		Time:     nowTime,
		LifeTime: lifeTime,
	}
	m.users[userId] = token
	return token
}

func (m *TokenMgr) getToken(userId int64) string {
	return m.users[userId]
}

func (m *TokenMgr) delToken(token string) {
	if info, ok := m.tokens[token]; ok {
		delete(m.tokens, token)
		delete(m.users, info.UserId)
	}
}
func (m *TokenMgr) isTokenVaild(token string) bool {
	if info, ok := m.tokens[token]; ok {
		if info.Time+info.LifeTime < time.Now().Unix() {
			return false
		}
		return true
	}
	return false
}

func (m *TokenMgr) getUserId(token string) int64 {
	if v, ok := m.tokens[token]; ok {
		return v.UserId
	}
	return 0
}
