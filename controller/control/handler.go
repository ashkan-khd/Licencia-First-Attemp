package control

import (
	"Licencia-First-Attempt/controller/control/utils/users"
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
	err := DB.Initialize()
	if err != nil {
		panic(err)
	}
	return &Control{}
}

/*
json{
username
email
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
		return users.RegisterEmployer(employer, DB)
	} else if accountType == "freelancer" {
		//TODO
	} else {
		return errors.New("invalid query: " + accountType)
	}
	return nil
}
