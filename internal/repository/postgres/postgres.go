package postgres

import (
	"database/sql"
	"fmt"
	"github.com/BioMihanoid/LearningManagementSystem/internal/config"
)

func NewPostgresDB(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.User, cfg.Dbname, cfg.Pass, "disable"),
	)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
