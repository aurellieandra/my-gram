package infrastructure

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// INTERFACE
type GormPostgres interface {
	GetConnection() *gorm.DB
}

// STRUCT
type gormPostgresImpl struct {
	master *gorm.DB
}

// NEW GORM POSTGRES
func NewGormPostgres() GormPostgres {
	return &gormPostgresImpl{
		master: connect(),
	}
}

// CONNECT
func connect() *gorm.DB{
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "admin123"
	dbname := "my-gram"

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

//  GORM POSTGRES IMPL
func (g *gormPostgresImpl) GetConnection() *gorm.DB {
	return g.master
}