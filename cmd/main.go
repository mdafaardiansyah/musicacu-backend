package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mdafaardiansyah/musicacu-backend/internal/configs"
	membershipsHandler "github.com/mdafaardiansyah/musicacu-backend/internal/handler/memberships"
	tracksHandler "github.com/mdafaardiansyah/musicacu-backend/internal/handler/tracks"
	"github.com/mdafaardiansyah/musicacu-backend/internal/models/memberships"
	"github.com/mdafaardiansyah/musicacu-backend/internal/models/trackactivities"
	membershipsRepo "github.com/mdafaardiansyah/musicacu-backend/internal/repository/memberships"
	"github.com/mdafaardiansyah/musicacu-backend/internal/repository/spotify"
	membershipsSvc "github.com/mdafaardiansyah/musicacu-backend/internal/service/memberships"
	"github.com/mdafaardiansyah/musicacu-backend/internal/service/tracks"
	"github.com/mdafaardiansyah/musicacu-backend/pkg/httpclient"
	"github.com/mdafaardiansyah/musicacu-backend/pkg/internalsql"
	"log"
	"net/http"
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
	db.AutoMigrate(&trackactivities.TrackActivity{})

	r := gin.Default()

	httpClient := httpclient.NewClient(&http.Client{})
	spotifyOutbound := spotify.NewSpotifyOutbound(cfg, httpClient)

	membershipRepo := membershipsRepo.NewRepository(db)

	membershipSvc := membershipsSvc.NewService(cfg, membershipRepo)
	trackSvc := tracks.NewService(spotifyOutbound)

	membershipHandler := membershipsHandler.NewHandler(r, membershipSvc)
	membershipHandler.RegisterRoute()

	tracksHandler := tracksHandler.NewHandler(r, trackSvc)
	tracksHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
