package sinks

import (
	"database/sql"
	"datapipe/data"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type PlanetScaleClient struct {
	db *sql.DB
}

func NewPlanetScaleClient(dsn string) *PlanetScaleClient {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to PlanetScale!")
	return &PlanetScaleClient{db}
}

func (c *PlanetScaleClient) Store(record data.Record) {
	panic("implement me...")
}
