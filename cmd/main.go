package main

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/config"
	"github.com/BioMihanoid/LearningManagementSystem/internal/repository"
	"github.com/BioMihanoid/LearningManagementSystem/internal/repository/postgres"
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
)

func main() {
	// TODO: parse config: godotenv

	// TODO: init logger: logrus

	db, err := postgres.NewPostgresDB(config.Config{})
	if err != nil {
		panic(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	_ = services
}
