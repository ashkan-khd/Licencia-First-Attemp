package routing

import (
	"Licencia-First-Attempt/controller/control"
	"Licencia-First-Attempt/view"

	"github.com/gin-gonic/gin"
)

type Listener interface {
	Listen() error
}

type router struct {
	port       string
	server     *gin.Engine
	controller *control.Control
}

func NewRouter(port string) Listener {

	var listener Listener = &router{port, gin.New(), control.NewControl()}
	return listener
}

func (router *router) Listen() error {
	router.server.POST("/register", func(context *gin.Context) {
		view.RespondRegister(context, router.controller.Register(context))
	})

	router.server.Run(":" + router.port)
	return nil
}
