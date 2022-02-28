package mj

var (
	MjMap map[string]ChessMj
)

type ChessMj interface {
	String() string
	GetTotalNums() int
	GetMjType(id int) int
}

func init() {
	MjMap = make(map[string]ChessMj)
	MjMap["xlch"] = NewXlchMj() //四川麻将血流成河
}
