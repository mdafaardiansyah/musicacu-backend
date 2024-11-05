package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mdafaardiansyah/musicacu-backend/internal/configs"
	membershipsHandler "github.com/mdafaardiansyah/musicacu-backend/internal/handler/memberships"
	"github.com/mdafaardiansyah/musicacu-backend/internal/models/memberships"
	membershipsRepo "github.com/mdafaardiansyah/musicacu-backend/internal/repository/memberships"
	membershipsSvc "github.com/mdafaardiansyah/musicacu-backend/internal/service/memberships"
	"github.com/mdafaardiansyah/musicacu-backend/pkg/internalsql"
	"log"
)

func main() {
	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{
			"./configs/",
			"./internal/configs/",
		}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatalf("error initializing configs: %+v\n", err)
	}
	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("error connecting to database %+v\n", err)
	}

	db.AutoMigrate(&memberships.User{})

	r := gin.Default()

	membershipRepo := membershipsRepo.NewRepository(db)

	membershipSvc := membershipsSvc.NewService(cfg, membershipRepo)

	membershipHandler := membershipsHandler.NewHandler(r, membershipSvc)
	membershipHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
