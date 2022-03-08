package mahjong

import (
	"encoding/json"
	"fmt"
	"github.com/jqiris/orange/tools"
	"testing"

	"github.com/jqiris/orange/protos"
)

func TestStructToJson(t *testing.T) {
	data := &protos.UserLocation{
		RoomId:    "",
		SeatIndex: 0,
	}
	bs, _ := json.Marshal(data)
	fmt.Println(string(bs))
}

func TestSlicePop(t *testing.T) {
	var a []int32
	b, i := tools.SlicePop(a)
	fmt.Println(b, i)
}
