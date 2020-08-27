package control

import (
	"Licencia-First-Attempt/controller/control/utils"
	"Licencia-First-Attempt/model/database"
	"Licencia-First-Attempt/model/existence"
	"errors"

	"github.com/gin-gonic/gin"
)

type Control struct {
}

var DB *database.Database

func NewControl() *Control {
	DB = database.NewDb()
	if err := DB.InitEmployerTable(); err != nil {
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
		return utils.RegisterEmployer(employer)
	} else if accountType == "freelancer" {
		//TODO
	} else {
		return errors.New("invalid query: " + accountType)
	}
	return nil
}
