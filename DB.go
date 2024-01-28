package main

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"log"
)

const (
	DBhost = "localhost:5432"
	//DBport     = 5432
	DBuser     = "codler"
	DBpassword = "codler"
	DBname     = "codler"
)

var database = &pg.DB{}

func connectDB() {
	database = pg.Connect(&pg.Options{
		Addr:     DBhost,
		User:     DBuser,
		Password: DBpassword,
		Database: DBname,
	})
	//defer database.Close()

	err := createSchema(database)
	if err != nil {
		panic(err)
	}

	log.Print("Database Connected!")
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*User)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			panic(err)
			return err
		}
	}
	log.Print("Scheme Created!")
	return nil
}
