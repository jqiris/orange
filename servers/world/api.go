package world

import (
	"github.com/gin-gonic/gin"
	"github.com/jqiris/kungfu/v2/discover"
	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/orange/model"
)

func Login(c *gin.Context) model.Msg {
	serverList := discover.GetServerList()
	msg := model.Msg{
		Data: map[string]interface{}{
			"serverList": serverList,
		},
	}
	logger.Info("world login")
	return msg
}
