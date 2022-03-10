package account

import (
	"encoding/base64"
	"fmt"
	"github.com/jqiris/kungfu/v2/discover"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/utils"
	"github.com/jqiris/orange/constant"
	"github.com/jqiris/orange/database"
	"github.com/jqiris/orange/model"
	"github.com/spf13/viper"
)

// Register 注册
func (s *ServerAccount) Register(c *gin.Context) {
	account, password := c.Query("account"), c.Query("password")
	msg := Msg{Errcode: 1, ErrMsg: "account has been used"}
	_, err := database.GetUserByAccount(account)
	if err != nil {
		logger.Error(err)
		c.JSON(200, msg)
	} else {
		if err = database.CreateAccount(&model.TAccount{Account: account, Password: password}); err == nil {
			msg.Errcode = 0
			msg.ErrMsg = "ok"
			c.JSON(200, msg)
		} else {
			logger.Error(err)
		}
	}
	c.JSON(200, msg)
}

// GetVersion 返回版本号
func (s *ServerAccount) GetVersion(c *gin.Context) {
	resp := map[string]any{
		"version": viper.GetString("primary.version"),
	}
	c.JSON(200, resp)
}

// GetServerinfo 返回服务信息
func (s *ServerAccount) GetServerinfo(c *gin.Context) {
	//获取大厅服务器地址
	server := discover.GetServerByType(constant.HallServer, c.ClientIP())
	resp := map[string]any{
		"version": viper.GetString("primary.version"),
		"hall":    fmt.Sprintf("%v:%v", server.ServerIp, server.ClientPort),
		"appweb":  "http://fir.im/2f17",
	}
	c.JSON(200, resp)
}

// Guest 游客账号
func (s *ServerAccount) Guest(c *gin.Context) {
	account := "guest_" + c.Query("account")
	sign := utils.Md5(account + c.ClientIP() + viper.GetString("primary.account_key"))
	server := discover.GetServerByType(constant.HallServer, c.ClientIP())
	resp := map[string]any{
		"errcode":  0,
		"errmsg":   "ok",
		"account":  account,
		"halladdr": fmt.Sprintf("%v:%v", server.ServerIp, server.ClientPort),
		"sign":     sign,
	}
	c.JSON(200, resp)
}

// Auth 验证授权
func (s *ServerAccount) Auth(c *gin.Context) {
	account, password := c.Query("account"), c.Query("password")
	user, err := database.GetAccount(account)
	if err != nil {
		c.JSON(200, Msg{Errcode: 1, ErrMsg: "invalid account"})
		return
	}
	if user.Password != utils.Md5(password) {
		c.JSON(200, Msg{Errcode: 1, ErrMsg: "invalid account"})
		return
	}
	sign := utils.Md5(account + c.ClientIP() + viper.GetString("primary.account_key"))
	resp := map[string]any{
		"errcode": 0,
		"errmsg":  "ok",
		"account": account,
		sign:      sign,
	}
	c.JSON(200, resp)
}

// BaseInfo 用户基本信息
func (s *ServerAccount) BaseInfo(c *gin.Context) {
	userId := utils.StringToInt64(c.Query("userid"))
	user, err := database.GetUserById(userId)
	if err != nil {
		c.JSON(200, Msg{Errcode: 1, ErrMsg: "invalid userid"})
		return
	}
	bs, _ := base64.StdEncoding.DecodeString(user.Name)
	resp := map[string]any{
		"errcode":    0,
		"errmsg":     "ok",
		"name":       string(bs),
		"sex":        user.Sex,
		"headimgurl": user.Headimg,
	}
	c.JSON(200, resp)
}

func (s *ServerAccount) Image(c *gin.Context) {
	imageName := c.Query("url")
	filePath := path.Join(viper.GetString("primary.image_path"), imageName)
	c.File(filePath)
}
