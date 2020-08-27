package database

import (
	"Licencia-First-Attempt/model/existence"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

var (
	dbc = dbConnection{
		username: "postgres",
		password: "mbsoli1743399413",
	}

	options = &orm.CreateTableOptions{
		IfNotExists: true,
	}
)

type Metadata struct {
	IsFirstInit bool `json:"is_first_init"`
}

type Initializable interface {
	Initialize() error
}

type dbConnection struct {
	username string
	password string
}

type Database struct {
	db   *pg.DB
	meta *Metadata
}

func NewDb() *Database {
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     dbc.username,
		Password: dbc.password,
		Database: "Licencia-First",
	})

	bytes, err := ioutil.ReadFile("model/database/jsons/db-metadata.json")
	if err != nil {
		panic(err)
	}
	meta := &Metadata{}
	err = json.Unmarshal(bytes, &meta)

	return &Database{db, meta}
}

func (db *Database) Initialize() error {

	defer func() {
		if err := db.updateFirstInit(); err != nil {
			panic(err)
		}
	}()

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
	if err != nil {
		return err
	}
	if db.meta.IsFirstInit {
		err = db.addDefaultFields()
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *Database) addDefaultFields() error {
	bytes, err := ioutil.ReadFile("model/database/jsons/default-fields.json")
	if err != nil {
		return err
	}
	fields := []existence.Field{}
	err = json.Unmarshal(bytes, &fields)

	if err != nil {
		return err
	}

	_, err = db.db.Model(&fields).Insert()

	if err != nil {
		return err
	}

	return nil
}

func (db *Database) updateFirstInit() error {

	if err := os.Remove("model/database/jsons/db-metadata.json"); err != nil {
		return nil
	}

	bytes, err := json.Marshal(db.meta)
	if err != nil {
		return err
	}

	file, err := os.Create("model/database/jsons/db-metadata.json")

	if err != nil {
		return err
	}

	if _, err := file.Write(bytes); err != nil {
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	return nil
}
