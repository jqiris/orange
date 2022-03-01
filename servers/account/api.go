package account

import (
	"github.com/gin-gonic/gin"
	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/utils"
	"github.com/jqiris/orange/constant"
	"github.com/jqiris/orange/database"
	"github.com/jqiris/orange/model"
	"github.com/spf13/viper"
)

//注册
func (s *AccountServer) Register(c *gin.Context) {
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

//返回版本号
func (s *AccountServer) GetVersion(c *gin.Context) {
	resp := map[string]interface{}{
		"version": viper.GetString("version"),
	}
	c.JSON(200, resp)
}

//返回服务信息
func (s *AccountServer) GetServerinfo(c *gin.Context) {
	resp := map[string]interface{}{
		"version": viper.GetString("version"),
		"hall":    constant.HALL_ADDR,
		"appweb":  "http://fir.im/2f17",
	}
	c.JSON(200, resp)
}

//游客账号
func (s *AccountServer) Guest(c *gin.Context) {
	account := "guest_" + c.Query("account")
	sign := utils.Md5(account + c.ClientIP() + constant.ACCOUNT_PRI_KEY)
	resp := map[string]interface{}{
		"errcode":  0,
		"errmsg":   "ok",
		"account":  account,
		"halladdr": constant.HALL_ADDR,
		"sign":     sign,
	}
	c.JSON(200, resp)
}

//验证授权
func (s *AccountServer) Auth(c *gin.Context) {
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
	sign := utils.Md5(account + c.ClientIP() + constant.ACCOUNT_PRI_KEY)
	resp := map[string]interface{}{
		"errcode": 0,
		"errmsg":  "ok",
		"account": account,
		sign:      sign,
	}
	c.JSON(200, resp)
}

//用户基本信息
func (s *AccountServer) BaseInfo(c *gin.Context) {
	userId := utils.StringToInt(c.Query("userid"))
	user, err := database.GetUserById(userId)
	if err != nil {
		c.JSON(200, Msg{Errcode: 1, ErrMsg: "invalid userid"})
		return
	}
	resp := map[string]interface{}{
		"errcode":    0,
		"errmsg":     "ok",
		"name":       user.Name,
		"sex":        user.Sex,
		"headimgurl": user.Headimg,
	}
	c.JSON(200, resp)
}
