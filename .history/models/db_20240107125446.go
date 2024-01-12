package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dbType := "sqlite"
	dbType += ":./database.db"

	db, err := gorm.Open(sqlite.Open(dbType), &gorm.Config{})
	if err != nil {
		panic("Error al conectar a la base de datos: " + err.Error())
	}

	err = Migrate(db) // Agrega tus modelos aquí
	if err != nil {
		panic("Error al ejecutar migraciones: " + err.Error())
	}

	DB = db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
