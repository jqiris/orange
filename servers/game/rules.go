package game

type GameMgr interface {
	setReady(userId int)
	hasBegan(roomId int)
	dissolveRequest(roomId, userId int)
	dissolveAgree(roomId, userId int, agree bool)
	doDissolve(roomId int)
	huanSanZhang(userId int, p1, p2, p3 int)
	dingQue(userId int, que int)
	chuPai(userId, pai int)
	peng(userId int)
	gang(userId, pai int)
	hu(userId int)
	guo(userId int)
}
