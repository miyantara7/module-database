package main

import (
	"fmt"
	"log"

	"github.com/vins7/module-database/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Conn *gorm.DB

	tables []interface{}
}

func NewDB(cfg interface{}, tables ...interface{}) *DB {

	config, ok := cfg.(*entity.Database)
	if !ok {
		log.Fatal("config not match !")
	}

	pg := fmt.Sprintf("host= %v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta",
		config.Host,
		config.Username,
		config.Password,
		config.Dbname,
		config.Port)

	db, err := gorm.Open(postgres.Open(pg))
	if err != nil {
		log.Fatalf("error connect db ,%s", err.Error())
	}

	return &DB{
		Conn:   db,
		tables: tables,
	}
}

func (d *DB) MigrateSchema() {
	if d.tables != nil {
		d.Conn.AutoMigrate(d.tables...)
	}
}
