package hall

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jqiris/kungfu/v2/launch"
	"github.com/jqiris/kungfu/v2/plugin"
	"github.com/jqiris/kungfu/v2/rpc"
	"github.com/jqiris/kungfu/v2/treaty"
	"github.com/jqiris/orange/constant"
)

type HallServer struct {
	*rpc.ServerBase
}

func HallServerCreator(s *treaty.Server) (rpc.ServerEntity, error) {

	server := &HallServer{
		ServerBase: rpc.NewServerBase(s),
	}
	//http handler
	app := gin.Default()
	app.Use(cors.Default())
	server.Router(app)
	//reg plugin
	plug := plugin.NewServerHttp(app)
	server.AddPlugin(plug)
	return server, nil
}

func init() {
	launch.RegisterCreator(constant.HallServer, HallServerCreator)
}

//router
func (h *HallServer) Router(app *gin.Engine) {
	app.GET("/login", h.Login)
	app.GET("/create_user", h.CreateUser)
	app.GET("/create_private_room", h.CreatePrivateRoom)
	app.GET("/enter_private_room", h.EnterPrivateRoom)
	app.GET("/get_history_list", h.GetHistoryList)
	app.GET("/get_games_of_room", h.GetGamesOfRoom)
	app.GET("/get_detail_of_game", h.GetDetailOfGame)
	app.GET("/get_user_status", h.GetUserStatus)
	app.GET("/get_message", h.GetMessage)
}
