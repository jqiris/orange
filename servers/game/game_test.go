package game

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jqiris/orange/protos"
)

func TestStructToJson(t *testing.T) {
	data := &protos.UserLocation{
		RoomId:    0,
		SeatIndex: 0,
	}
	bs, _ := json.Marshal(data)
	fmt.Println(string(bs))
}
