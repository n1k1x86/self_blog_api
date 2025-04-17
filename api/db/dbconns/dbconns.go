package dbconns

import (
	"api/config"
	"database/sql"
	"fmt"
)

func ConnectToBlogDBConfig(cfg config.BlogDBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.Username, cfg.Password, cfg.Server, cfg.DBName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
