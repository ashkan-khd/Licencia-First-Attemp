package control

import (
	"Licencia-First-Attempt/model"
)

type Control struct {
	db *model.Database
}

func NewControl() *Control {
	db := model.NewDb()
	if err := db.InitEmployerTable(); err != nil {
		panic(err)
	}
	return &Control{}
}
