package game

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/utils"
	"github.com/jqiris/orange/constant"
	"github.com/jqiris/orange/model"
)

//router
func (s *GameServer) HttpRouter(app *gin.Engine) {
	app.GET("/get_server_info", s.GetServerInfo)
	app.GET("/create_room", s.CreateRoom)
	app.GET("/enter_room", s.EnterRoom)
	app.GET("/is_room_runing", s.IsRoomRuning)
}

//向大厅服定时心跳
func (s *GameServer) update() {
	ticker := time.NewTicker(time.Duration(constant.HTTP_TICK_TIME) * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			s.gameServerInfo["load"] = []string{utils.IntToString(roomMgr.getTotallRooms())}
			url := fmt.Sprintf("http://%v:%v/register_gs?%v", constant.HALL_IP, constant.HALL_PORT, s.gameServerInfo.Encode())
			resp, err := http.Get(url)
			if err != nil {
				logger.Error(err)
				continue
			}
			logger.Infof("game http heart:%v", resp.StatusCode)
		}
	}
}

func (s *GameServer) Error(c *gin.Context, msg string, args ...int) {
	errcode := 1
	if len(args) > 0 {
		errcode = args[0]
	}
	c.JSON(200, map[string]interface{}{"errcode": errcode, "errmsg": msg})
}

func (s *GameServer) Success(c *gin.Context, data map[string]interface{}) {
	data["errcode"] = 0
	data["errmsg"] = "ok"
	c.JSON(200, data)
}

//获取服务信息
func (s *GameServer) GetServerInfo(c *gin.Context) {
	serverId, sign := c.Query("serverid"), c.Query("sign")
	if serverId != constant.SERVER_ID || len(sign) < 1 {
		s.Error(c, "invalid parameters")
		return
	}
	md5 := utils.Md5(serverId + constant.ROOM_PRI_KEY)
	if md5 != sign {
		s.Error(c, "sign check failed")
		return
	}
	locations := roomMgr.getUserLocations()
	var arr []int
	locations.Range(func(key, value any) bool {
		if v, ok := value.(*UserLocation); ok {
			arr = append(arr, key.(int))
			arr = append(arr, v.RoomId)
		}
		return true
	})
	s.Success(c, map[string]interface{}{"userroominfo": arr})
}

//创建房间
func (s *GameServer) CreateRoom(c *gin.Context) {
	userId, sign, gems, conf := utils.StringToInt(c.Query("userid")), c.Query("sign"), c.Query("gems"), c.Query("conf")
	if userId < 1 || len(sign) < 1 || len(conf) < 1 {
		s.Error(c, "Invalid parameters")
		return
	}
	md5 := utils.Md5(c.Query("userid") + conf + gems + constant.ROOM_PRI_KEY)
	if md5 != sign {
		s.Error(c, "sign check failed")
		return
	}
	var cfg model.GameConf
	err := json.Unmarshal([]byte(conf), &cfg)
	if err != nil {
		s.Error(c, err.Error())
		return
	}
	errcode, roomId := roomMgr.createRoom(userId, cfg, gems, c.ClientIP(), constant.CLIENT_PORT)
	if errcode != 0 || roomId < 1 {
		s.Error(c, "create failed")
		return
	}
	s.Success(c, map[string]interface{}{"roomid": roomId})
}

//进入房间
func (s *GameServer) EnterRoom(c *gin.Context) {
	userId, name, roomId, sign := utils.StringToInt(c.Query("userid")), c.Query("name"), utils.StringToInt(c.Query("roomid")), c.Query("sign")
	if userId < 1 || roomId < 1 || len(sign) < 1 {
		s.Error(c, "invalid parameters")
		return
	}
	md5 := utils.Md5(fmt.Sprintf("%v%v%v%v", userId, name, roomId, constant.ROOM_PRI_KEY))
	if md5 != sign {
		s.Error(c, "sign check failed")
		return
	}
	ret := roomMgr.enterRoom(roomId, userId, name)
	if ret != 0 {
		switch ret {
		case 1:
			s.Error(c, "room is full")
			return
		case 2:
			s.Error(c, "can't find room")
			return
		case 3:
			s.Error(c, "build room from db failed")
			return
		}
	}
	token := tokenMgr.createToken(userId, 5000)
	s.Success(c, map[string]interface{}{"token": token})
}

//房间是否运行
func (s *GameServer) IsRoomRuning(c *gin.Context) {
	roomId, sign := utils.StringToInt(c.Query("roomid")), c.Query("sign")
	if roomId < 1 || len(sign) < 1 {
		s.Error(c, "invalid parameters")
		return
	}
	md5 := utils.Md5(fmt.Sprintf("%v%v", roomId, constant.ROOM_PRI_KEY))
	if md5 != sign {
		s.Error(c, "sign check failed", 2)
		return
	}
	s.Success(c, map[string]interface{}{"runing": true})
}
