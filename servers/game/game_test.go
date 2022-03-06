package game

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestStructToJson(t *testing.T) {
	data := &Room{}
	bs, _ := json.Marshal(data)
	fmt.Println(string(bs))
}
