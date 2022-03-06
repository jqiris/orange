package game

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestStructToJson(t *testing.T) {
	data := &ActionData{
		Type:     "",
		Owner:    &Seat{},
		State:    "",
		PayTimes: 0,
		Ref:      &ActionData{},
		Targets:  []int{1},
		Score:    0,
		Iszimo:   false,
	}
	bs, _ := json.Marshal(data)
	fmt.Println(string(bs))
}
