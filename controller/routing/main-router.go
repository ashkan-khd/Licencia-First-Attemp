package routing

import "github.com/gin-gonic/gin"

type Listener interface {
	Listen() error
}

type router struct {
	port   string
	server *gin.Engine
}

func NewRouter(port string) Listener {
	var listener Listener = &router{port, gin.New()}
	return listener
}

func (router *router) Listen() error {

	router.server.Run(":" + router.port)
	return nil
}
