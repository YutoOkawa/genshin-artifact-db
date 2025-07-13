package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/YutoOkawa/genshin-artifact-db/pkg/handler"
	"github.com/YutoOkawa/genshin-artifact-db/pkg/repository"
	"github.com/YutoOkawa/genshin-artifact-db/pkg/server"
	"github.com/YutoOkawa/genshin-artifact-db/pkg/service"
	"github.com/gin-gonic/gin"
)

func main() {
	artifactRepository := repository.NewInMemoryArtifactRepository()
	artifactRepository.LoadJSONFile("artifacts.json")

	getArtifactService := service.NewGetArtifactService(artifactRepository)
	createArtifactService := service.NewUpdateArtifactService(artifactRepository)

	r := gin.Default()
	r.GET("/artifact/:id", handler.GetArtifact(getArtifactService))
	r.GET("/artifacts/type/:type", handler.GetArtifactsByType(getArtifactService))
	r.GET("/artifacts/set/:set", handler.GetArtifactsBySet(getArtifactService))
	r.GET("/artifacts/type/:type/set/:set", handler.GetArtifacts(getArtifactService))

	r.POST("/artifact", handler.CreateArtifact(createArtifactService))

	serve := server.NewServer(":8080", r, 1)
	serverCh := serve.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	for {
		select {
		case <-quit:
			if err := artifactRepository.SaveJSONFile("artifacts.json"); err != nil {
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
