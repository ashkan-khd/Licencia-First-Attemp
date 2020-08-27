package model

import (
	"Licencia-First-Attempt/model/existence"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Database struct {
	db *pg.DB
}

func NewDb() *Database {
	dbc := dbConnection{
		username: "postgres",
		password: "mbsoli1743399413",
	}
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     dbc.username,
		Password: dbc.password,
		Database: "Licencia-First",
	})
	return &Database{db}
}

type dbConnection struct {
	username string
	password string
}

func (db *Database) InitEmployerTable() error {
	options := orm.CreateTableOptions{
		IfNotExists: true,
	}
	err := db.db.CreateTable(&existence.Employer{}, options)
	return err
}
