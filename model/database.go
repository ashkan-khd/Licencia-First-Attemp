package model

import "github.com/go-pg/pg"

type database struct {
	db *pg.DB
}

func NewDb() *database {
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "ashka",
		Password: "a124578",
		Database: "Licencia-First",
	})
	return &database{db}
}
