package mahjong

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jqiris/orange/model"
	"github.com/jqiris/orange/tools"
)

func TestStructToJson(t *testing.T) {
	data := &model.UserLocation{
		RoomId:    "",
		SeatIndex: 0,
	}
	bs, _ := json.Marshal(data)
	fmt.Println(string(bs))
}

func TestSlicePop(t *testing.T) {
	var a []int32
	isPop, b, i := tools.SlicePop(a)
	fmt.Println(isPop, b, i)
}
