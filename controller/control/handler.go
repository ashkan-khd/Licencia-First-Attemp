package control

import (
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
	DB.Initialize()
	return &Control{}
}

/*
json{
username
password
}
QUERIES : "account-type"
*/
func (controller *Control) Register(ctx *gin.Context) error {
	accountType := ctx.Query("account-type")
	if accountType == "employer" {
		employer := existence.Employer{}
		if err := ctx.ShouldBindJSON(&employer); err != nil {
			return err
		}
		return RegisterEmployer(employer)
	} else if accountType == "freelancer" {
		//TODO
	} else {
		return errors.New("invalid query: " + accountType)
	}
	return nil
}
