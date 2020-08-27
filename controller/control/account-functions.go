package control

import (
	"Licencia-First-Attempt/model/existence"
	"errors"
)

func RegisterEmployer(emp existence.Employer) error {
	if !DB.DoesEmployerExistWithUsername(emp.Username) {
		if !DB.DoesEmployerExistWithEmail(emp.Email) {
			return DB.InsertEmployer(emp)
		}
		return errors.New("duplicate email: " + emp.Username)
	}
	return errors.New("duplicate username: " + emp.Username)
}
