package db

import (
	"fmt"
	"post-service/pkg/config"

	"github.com/jmoiron/sqlx"
)

func NewDB(cfg config.Config) *sqlx.DB {

	dns := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		cfg.GetString("user"),
		cfg.GetString("password"),
		cfg.GetString("dbname"),
		cfg.GetString("host"),
		cfg.GetInt("port"),
	)

	sqlx.

}
