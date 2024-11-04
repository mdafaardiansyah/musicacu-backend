package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/mdafaardiansyah/musicacu-backend/internal/configs"
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

	r := gin.Default()

	r.Run(cfg.Service.Port)
}
