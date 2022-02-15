package world

import (
	"github.com/gin-gonic/gin"
	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/orange/model"
)

func Login(c *gin.Context) model.Msg {
	msg := model.Msg{}
	logger.Info("world login")
	return msg
}
