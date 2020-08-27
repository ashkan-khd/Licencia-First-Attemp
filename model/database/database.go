package database

import (
	"Licencia-First-Attempt/model/existence"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

var (
	dbc = dbConnection{
		username: "ashka",
		password: "a124578",
	}

	options = &orm.CreateTableOptions{
		IfNotExists: true,
	}
)

type Initializable interface {
	Initialize() error
}

type dbConnection struct {
	username string
	password string
}

type Database struct {
	db *pg.DB
}

func NewDb() *Database {
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     dbc.username,
		Password: dbc.password,
		Database: "Licencia-First",
	})
	return &Database{db}
}

func (db *Database) Initialize() error {

	if err := db.initEmployerTable(); err != nil {
		return err
	}

	if err := db.initFieldTable(); err != nil {
		return err
	}

	return nil
}

func (db *Database) initEmployerTable() error {
	err := db.db.CreateTable(&existence.Employer{}, options)
	return err
}

func (db *Database) initFieldTable() error {
	err := db.db.CreateTable(&existence.Field{}, options)
	return err
}

func (db *Database) addDefaultFields() error {
	return nil
}
