package main

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/config"
	"github.com/BioMihanoid/LearningManagementSystem/internal/handlers"
	"github.com/BioMihanoid/LearningManagementSystem/internal/repository"
	"github.com/BioMihanoid/LearningManagementSystem/internal/repository/postgres"
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
	"github.com/BioMihanoid/LearningManagementSystem/pkg"
	"github.com/sirupsen/logrus"
)

// @title Learning Management System API
// @version 1.0
// @description API Server for Learning Management System Application

// @host localhost:8002
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	cfg := config.ParseConfig()

	db, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		logrus.Panic(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handlers.NewHandler(services)

	srv := new(pkg.Server)
	if err = srv.Start(cfg, handler.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}
