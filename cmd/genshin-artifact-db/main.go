package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/YutoOkawa/genshin-artifact-db/pkg/config"
	"github.com/YutoOkawa/genshin-artifact-db/pkg/handler"
	"github.com/YutoOkawa/genshin-artifact-db/pkg/repository"
	"github.com/YutoOkawa/genshin-artifact-db/pkg/server"
	"github.com/YutoOkawa/genshin-artifact-db/pkg/service"
	"github.com/gin-gonic/gin"
)

func main() {
	configPath := flag.String("config", config.DefaultConfigPath, "設定ファイルのパス")
	portFlag := flag.String("port", "", "サーバーポート (設定ファイルを上書き)")
	dataFlag := flag.String("data", "", "データファイルパス (設定ファイルを上書き)")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if *portFlag != "" {
		cfg.Port = *portFlag
	}
	if *dataFlag != "" {
		cfg.DataFilePath = *dataFlag
	}

	log.Printf("Starting server with config: port=%s, data_file=%s", cfg.Port, cfg.DataFilePath)

	artifactRepository := repository.NewInMemoryArtifactRepository()
	if err := artifactRepository.LoadJSONFile(cfg.DataFilePath); err != nil {
		log.Printf("Warning: Failed to load data file: %v", err)
	}

	getArtifactService := service.NewGetArtifactService(artifactRepository)
	createArtifactService := service.NewUpdateArtifactService(artifactRepository)

	r := gin.Default()
	r.GET("/artifact/:id", handler.GetArtifact(getArtifactService))
	r.GET("/artifacts/type/:type", handler.GetArtifactsByType(getArtifactService))
	r.GET("/artifacts/set/:set", handler.GetArtifactsBySet(getArtifactService))
	r.GET("/artifacts/type/:type/set/:set", handler.GetArtifacts(getArtifactService))

	r.POST("/artifact", handler.CreateArtifact(createArtifactService))

	serve := server.NewServer(cfg.Port, r, 1)
	serverCh := serve.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	for {
		select {
		case <-quit:
			if err := artifactRepository.SaveJSONFile(cfg.DataFilePath); err != nil {
				log.Fatalf("Failed to save artifacts: %v", err)
			}
			serve.Shutdown()
			log.Println("Server shutdown")
			return
		case err := <-serverCh:
			if err != nil {
				log.Fatalf("Server error: %v", err)
			}
			log.Println("Server stopped")
			return
		}
	}

}
