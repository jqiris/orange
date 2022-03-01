package game

type Seat struct {
	UserId      int    `json:"userId"`
	Score       int    `json:"score"`
	Name        string `json:"name"`
	Ready       bool   `json:"ready"`
	SeatIndex   int    `json:"seatIndex"`
	NumZiMo     int    `json:"numZiMo"`
	NumJiePao   int    `json:"numJiePao"`
	NumDianPao  int    `json:"numDianPao"`
	NumAnGang   int    `json:"numAnGang"`
	NumMingGang int    `json:"numMingGang"`
	NumChaJiao  int    `json:"numChaJiao"`
	Ip          string `json:"ip"`
	Online      bool   `json:"online"`
}
