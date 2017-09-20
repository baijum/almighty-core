package controller

import (
	"net/http"
	"time"

	"github.com/fabric8-services/fabric8-wit/app"
	"github.com/fabric8-services/fabric8-wit/application"
	"github.com/goadesign/goa"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// WorkitemRefreshController implements the workitem_refresh resource.
type WorkitemRefreshController struct {
	*goa.Controller
	db     application.DB
	config WorkItemRefreshControllerConfiguration
}

// WorkItemRefreshControllerConfiguration is the configuration for WI refresh
type WorkItemRefreshControllerConfiguration interface {
}

// NewWorkitemRefreshController creates a workitem_refresh controller.
func NewWorkitemRefreshController(service *goa.Service, db application.DB, config WorkItemRefreshControllerConfiguration) *WorkitemRefreshController {
	return &WorkitemRefreshController{Controller: service.NewController("WorkitemRefreshController")}
}

// Refresh runs the refresh action.
func (c *WorkitemRefreshController) Refresh(ctx *app.RefreshWorkitemRefreshContext) error {
	ws, _ := upgrader.Upgrade(ctx.ResponseWriter, ctx.Request, nil)
	go func() {
		time.Sleep(time.Second * 10)
		ws.WriteMessage(websocket.TextMessage, []byte("hello"))
	}()
	//c.RefreshWSHandler(ctx).ServeHTTP(ctx.ResponseWriter, ctx.Request)
	return nil
}

// RefreshWSHandler establishes a websocket connection to run the refresh action.
/*
func (c *WorkitemRefreshController) RefreshWSHandler(ctx *app.RefreshWorkitemRefreshContext) websocket.Handler {
	return func(ws *websocket.Conn) {
		// WorkitemRefreshController_Refresh: start_implement

		// Put your logic here

		// WorkitemRefreshController_Refresh: end_implement
		ws.Write([]byte("refresh workitem_refresh"))
		// Dummy echo websocket server
		io.Copy(ws, ws)
	}
}
*/
