package hall

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jqiris/kungfu/v2/base"
	"github.com/jqiris/kungfu/v2/launch"
	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/rpc"
	"github.com/jqiris/kungfu/v2/treaty"
	"github.com/jqiris/orange/constant"
)

type HallServer struct {
	*base.ServerHttp
}

func (h *HallServer) HandleSelfEvent(req *rpc.MsgRpc) []byte {
	resp, err := h.DealMsg(rpc.CodeTypeProto, h.Rpc, req)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return resp
}
func (h *HallServer) HandleBroadcastEvent(req *rpc.MsgRpc) []byte {
	return nil
}

func HallServerCreator(s *treaty.Server) (rpc.ServerEntity, error) {
	//http handler
	app := gin.Default()
	server := &HallServer{
		ServerHttp: base.NewServerHttp(s, app),
	}
	server.SelfEventHandler = server.HandleSelfEvent
	server.BroadcastEventHandler = server.HandleBroadcastEvent
	//msg handler register
	// server.Register(int32(protos.MsgId_MsgChan), server.ChanResp)
	//cors
	app.Use(cors.Default())
	//router
	server.Router(app)
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
