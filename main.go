package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jqiris/kungfu/v2/config"
	"github.com/jqiris/kungfu/v2/discover"
	"github.com/jqiris/kungfu/v2/launch"
	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/stores"
	_ "github.com/jqiris/orange/servers/world"
	"github.com/spf13/viper"
)

func main() {
	//设置时间随机种子
	rand.Seed(time.Now().UnixNano())
	//配置解析
	cfg := flag.String("conf", "config.json", "locate the config file")
	flag.Parse()
	viper.SetConfigFile(*cfg)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	//logger
	loggerCfg := viper.GetStringMap("logger")
	initLogger(loggerCfg)
	//frame init
	frameCfg := viper.Get("frame")
	if err := config.InitFrameConf(frameCfg); err != nil {
		logger.Fatal(err)
	}
	//init discover
	discover.InitDiscoverer(config.GetDiscoverConf())

	//init stores
	stores.InitStoreKeeper(config.GetStoresConf())
	//launch servers
	launch.Startup()
	sg := make(chan os.Signal, 1)
	signal.Notify(sg, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	select {
	case s := <-sg:
		log.Println("server got shutdown signal", s)
	}
	launch.Shutdown()
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
