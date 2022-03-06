package game

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestStructToJson(t *testing.T) {
	data := &Room{
		Uuid:       "",
		Id:         0,
		NumOfGames: 0,
		CreateTime: 0,
		NextButton: 0,
		Seats:      []*Seat{},
		Conf:       RoomConf{},
		Dr:         &DissolveData{},
	}
	bs, _ := json.Marshal(data)
	fmt.Println(string(bs))
}
