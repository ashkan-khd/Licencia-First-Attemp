package control

import (
	"Licencia-First-Attempt/model/existence"
	"errors"
)

func RegisterEmployer(emp existence.Employer) error {
	if !DB.DoesEmployerExist(emp.Username) {
		return DB.InsertEmployer(emp)
	}
	return errors.New("duplicate username: " + emp.Username)
}
