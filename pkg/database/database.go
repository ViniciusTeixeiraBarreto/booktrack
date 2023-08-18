package database

import (
	"context"
	"log"
	"time"
	"web-api/pkg/database/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() *gorm.DB {
	str := "host=localhost port=5432 user=postgres dbname=books sslmode=disable password=postgres"

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil {
		log.Fatal("error: ", err)
	}

	db = database
	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)

	return db
}

func CloseConn() error {
	config, err := db.DB()
	if err != nil {
		return err
	}

	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}

//TODO remove it
func GetDatabase() *gorm.DB {
	return db
}

type connectionKey struct{}

func SetConnection(ctx context.Context, conn *gorm.DB) context.Context {
	if conn == nil {
		conn = db
	}

	conn.WithContext(ctx)
	return context.WithValue(ctx, connectionKey{}, conn)
}

func GetConnection(ctx context.Context) *gorm.DB {
	if ctx == nil {
		return db
	}
	conn, has := ctx.Value(connectionKey{}).(*gorm.DB)

	if !has {
		return db
	}
	return conn
}
