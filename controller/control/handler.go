package control

import (
	"Licencia-First-Attempt/model"
	"Licencia-First-Attempt/model/existence"
	"errors"

	"github.com/gin-gonic/gin"
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

//QUERIES : "account-type"
func (controller *Control) Register(ctx *gin.Context) error {
	accountType := ctx.Query("account-type")
	if accountType == "employer" {
		employer := existence.Employer{}
		if err := ctx.ShouldBindJSON(&employer); err != nil {
			return err
		}

	} else if accountType == "freelancer" {
		//TODO
	} else {
		return errors.New("Invalid query : " + accountType)
	}
	return nil
}
