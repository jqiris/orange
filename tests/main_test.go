package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jqiris/kungfu/v2/utils"
	"github.com/jqiris/orange/protos"
)

func TestTokenKey(t *testing.T) {
	uid, nickname, tokenKey := 1001, "jason", "orangeToken"
	sign := fmt.Sprintf("%v|%v|%v", uid, nickname, tokenKey)
	sign = utils.Md5(sign)
	fmt.Println(sign)
}

func TestJsonProtoOneof(t *testing.T) {
	req := &protos.MsgRequest{
		MsgId: protos.MsgId_MsgOffline,
		Any: &protos.MsgRequest_OfflineRequest{
			OfflineRequest: &protos.OfflineRequest{
				Uid: 1001,
			},
		},
	}
	bs, err := json.Marshal(req)
	fmt.Println(string(bs), err)
}
