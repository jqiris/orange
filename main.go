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
)

func main() {
	//设置时间随机种子
	rand.Seed(time.Now().UnixNano())
	//配置解析
	cfg := flag.String("conf", "config.json", "locate the config file")
	flag.Parse()

	if err := config.InitConf("config.json"); err != nil {
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
