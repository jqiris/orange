package tests

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/jqiris/kungfu/v2/config"
	"github.com/jqiris/kungfu/v2/coroutines"
	"github.com/jqiris/kungfu/v2/discover"
	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/stores"
	"github.com/jqiris/kungfu/v2/utils"
	"github.com/jqiris/orange/constant"
	"github.com/jqiris/orange/protos"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("../config.json")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	//frame init
	frameCfg := viper.Get("frame")
	if err := config.InitFrameConf(frameCfg); err != nil {
		logger.Fatal(err)
	}
	//init discover
	discover.InitDiscoverer(config.GetDiscoverConf())

	//init stores
	stores.InitStoreKeeper(config.GetStoresConf())
}

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

var codeMap sync.Map

func GenerateCode(key string, low, high int) error {
	for i := 0; i < 100000; i++ {
		code := utils.NextRoomCode(key, low, high)
		if _, ok := codeMap.Load(code); ok {
			return fmt.Errorf("GenerateCode code exist:%v", code)
		}
		codeMap.Store(code, 1)
	}
	return nil
}
func TestRoomCodeGenerate(t *testing.T) {
	w := coroutines.NewWaitCoroutines()
	w.AddCoroutine(func(wg *sync.WaitGroup) {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		GenerateCode(constant.MahjongRoomCode, constant.MahjongRoomCodeLow, constant.MahjongRoomCodeHigh)
	})
	w.AddCoroutine(func(wg *sync.WaitGroup) {
		defer wg.Done()
		time.Sleep(1000 * time.Millisecond)
		GenerateCode(constant.MahjongRoomCode, constant.MahjongRoomCodeLow, constant.MahjongRoomCodeHigh)
	})
	w.Wait()
	logger.Info("GenerateCode success")
	codeMap.Range(func(key, value interface{}) bool {
		code := key.(int)
		logger.Infof("code:%v", code)
		return true
	})
}
