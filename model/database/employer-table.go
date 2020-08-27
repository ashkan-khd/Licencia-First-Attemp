package database

import "Licencia-First-Attempt/model/existence"

func (db Database) DoesEmployerExist(username string) bool {
	resultSet := &[]existence.Employer{}
	_ = db.db.Model(resultSet).Where("username = ?", username).Select()
	return len(*resultSet) != 0
}

func (db *Database) InsertEmployer(emp existence.Employer) error {
	_, err := db.db.Model(emp).Insert()
	return err
}
