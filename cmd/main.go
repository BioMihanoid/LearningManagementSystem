package main

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/config"
	"github.com/BioMihanoid/LearningManagementSystem/internal/handlers"
	"github.com/BioMihanoid/LearningManagementSystem/internal/repository"
	"github.com/BioMihanoid/LearningManagementSystem/internal/repository/postgres"
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
	"github.com/BioMihanoid/LearningManagementSystem/pkg"
)

func main() {
	cfg := config.ParseConfig()

	// TODO: init logger: logrus

	db, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		panic(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handlers.NewHandler(services)

	srv := new(pkg.Server)
	if err = srv.Start(cfg, handler.InitRoutes()); err != nil {
	}
}
