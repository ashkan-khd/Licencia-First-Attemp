package utils

import (
	"Licencia-First-Attempt/controller/control"
	"Licencia-First-Attempt/model/existence"
	"errors"
)

func RegisterEmployer(emp existence.Employer) error {
	if !control.DB.DoesEmployerExist(emp.Username) {
		return control.DB.InsertEmployer(emp)
	}
	return errors.New("duplicate username: " + emp.Username)
}
