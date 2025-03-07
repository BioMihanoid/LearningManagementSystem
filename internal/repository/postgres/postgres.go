package postgres

import (
	"database/sql"
	"fmt"
	"github.com/BioMihanoid/LearningManagementSystem/internal/config"
	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Dbname, cfg.DB.Pass, cfg.DB.Sslmode),
	)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
