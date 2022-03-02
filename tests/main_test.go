package tests

import (
	"encoding/json"
	"fmt"
	"math/rand"
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
	"github.com/jqiris/orange/tools"
	"github.com/spf13/viper"
)

func init() {
	rand.Seed(time.Now().UnixNano())
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

	//logger
	loggerCfg := viper.GetStringMap("logger")
	initLogger(loggerCfg)
}

func initLogger(cfg map[string]interface{}) {
	zipDay := time.Duration(cfg["zip_day"].(float64)*24) * time.Hour
	options := []logger.Option{
		logger.WithLogLevel(cfg["log_level"].(string)),
		logger.WithOutType(cfg["out_type"].(string)),
		logger.WithLogDir(cfg["log_dir"].(string)),
		logger.WithLogName(cfg["log_name"].(string)),
		logger.WithLogDump(cfg["log_dump"].(bool)),
		logger.WithLogRuntime(cfg["log_runtime"].(bool)),
		logger.WithStdColor(cfg["std_color"].(bool)),
		logger.WithZipDuration(zipDay),
	}
	lg, cf := logger.NewLogger(options...)
	logger.SetLogger(lg, cf)
}

func TestTokenKey(t *testing.T) {
	uid, nickname, tokenKey := 1002, "david", "orangeToken"
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

func GenerateCode(name, key string, low, high, max int) error {
	for i := 0; i < 100000; i++ {
		code := utils.NextRoomCode(key, low, high)
		prefix := utils.RangeRand(10, 99)
		roomCode := int(prefix)*1000000 + code
		if roomCode > (100000000 - 1) {
			return fmt.Errorf("big than max %v GenerateCode roomCode:%v, prefix:%v,code:%v", name, roomCode, prefix, code)
		}
		if _, ok := codeMap.Load(roomCode); ok {
			// max++
			// if max < 5 {
			// 	return GenerateCode(name, key, low, high, max)
			// }
			return fmt.Errorf("%v GenerateCode code exist:%v", name, roomCode)
		}
		codeMap.Store(code, 1)
		logger.Infof("%v GenerateCode:%v", name, roomCode)
	}
	return nil
}
func TestRoomCodeGenerate(t *testing.T) {
	w := coroutines.NewWaitCoroutines()
	w.AddCoroutine(func(wg *sync.WaitGroup) {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		if err := GenerateCode("线程1", constant.MahjongRoomCode, constant.MahjongRoomCodeLow, constant.MahjongRoomCodeHigh, 0); err != nil {
			logger.Error(err)
		}
	})
	w.AddCoroutine(func(wg *sync.WaitGroup) {
		defer wg.Done()
		time.Sleep(200 * time.Millisecond)
		if err := GenerateCode("线程2", constant.MahjongRoomCode, constant.MahjongRoomCodeLow, constant.MahjongRoomCodeHigh, 0); err != nil {
			logger.Error(err)
		}
	})
	w.AddCoroutine(func(wg *sync.WaitGroup) {
		defer wg.Done()
		time.Sleep(300 * time.Millisecond)
		if err := GenerateCode("线程3", constant.MahjongRoomCode, constant.MahjongRoomCodeLow, constant.MahjongRoomCodeHigh, 0); err != nil {
			logger.Error(err)
		}
	})
	w.AddCoroutine(func(wg *sync.WaitGroup) {
		defer wg.Done()
		time.Sleep(400 * time.Millisecond)
		if err := GenerateCode("线程4", constant.MahjongRoomCode, constant.MahjongRoomCodeLow, constant.MahjongRoomCodeHigh, 0); err != nil {
			logger.Error(err)
		}
	})
	w.AddCoroutine(func(wg *sync.WaitGroup) {
		defer wg.Done()
		time.Sleep(500 * time.Millisecond)
		if err := GenerateCode("线程5", constant.MahjongRoomCode, constant.MahjongRoomCodeLow, constant.MahjongRoomCodeHigh, 0); err != nil {
			logger.Error(err)
		}
	})
	w.Wait()
	logger.Info("GenerateCode success")
}

func TestSliceDelete(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := tools.SliceDel(arr, 6, 2)
	fmt.Println(res)
}
