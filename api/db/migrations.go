package db

import (
	"api/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func RunMigrations(cfg *config.Config) error {
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.Username, cfg.Password, cfg.Server, cfg.DBName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	dir := "migrations"

	err = goose.Up(db, dir)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
