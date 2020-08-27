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
)

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

func (db *Database) InitEmployerTable() error {
	options := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	err := db.db.CreateTable(&existence.Employer{}, options)
	return err
}
