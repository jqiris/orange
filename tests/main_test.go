package tests

import (
	"fmt"
	"testing"

	"github.com/jqiris/kungfu/v2/utils"
)

func TestTokenKey(t *testing.T) {
	uid, nickname, tokenKey := 1001, "jason", "orangeToken"
	sign := fmt.Sprintf("%v|%v|%v", uid, nickname, tokenKey)
	sign = utils.Md5(sign)
	fmt.Println(sign)
}
