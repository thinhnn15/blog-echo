package router

import (
	"github.com/labstack/echo"
	"github.com/thinhnn15/blog-echo/handler"
)

type API struct {
	Echo *echo.Echo
	PostHandler handler.PostHandler
}

func(api *API) SetupRouter(){
	api.Echo.POST("/posts/new", api.PostHandler.HandleNewPost)
}
