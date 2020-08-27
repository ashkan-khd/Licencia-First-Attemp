package routing

import (
	"Licencia-First-Attempt/controller/control"

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

	router.server.Run(":" + router.port)
	return nil
}
