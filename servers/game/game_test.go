package game

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestStructToJson(t *testing.T) {
	data := &UserLocation{
		RoomId:    0,
		SeatIndex: 0,
	}
	bs, _ := json.Marshal(data)
	fmt.Println(string(bs))
}
