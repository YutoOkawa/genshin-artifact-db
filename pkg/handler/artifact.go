package handler

import (
	"errors"
	"genshin-artifact-db/pkg/repository"
	"genshin-artifact-db/pkg/service"

	"github.com/gin-gonic/gin"
)

func GetArtifact(service service.GetArtifactServiceInterface) func(c *gin.Context) {
	return func(c *gin.Context) {
		artifactID := c.Param("id")

		artifact, err := service.GetArtifact(artifactID)
		if err != nil {
			if errors.Is(err, repository.ErrArtifactNotFound) {
				c.JSON(404, gin.H{"error": err.Error()})
				return
			} else {
				c.JSON(500, gin.H{"error": "Internal server error"})
				return
			}
		}

		c.JSON(200, artifact)
	}
}

func GetArtifactsByType(service service.GetArtifactsByTypeServiceInterface) func(c *gin.Context) {
	return func(c *gin.Context) {
		artifactType := c.Param("type")

		artifacts, err := service.GetArtifactsByType(artifactType)
		if err != nil {
			if errors.Is(err, repository.ErrArtifactNotFound) {
				c.JSON(404, gin.H{"error": err.Error()})
				return
			} else {
				c.JSON(500, gin.H{"error": "Internal server error"})
				return
			}
		}
		c.JSON(200, artifacts)
	}
}

func GetArtifactsBySet(service service.GetArtifactsBySetServiceInterface) func(c *gin.Context) {
	return func(c *gin.Context) {
		artifactSet := c.Param("set")

		artifacts, err := service.GetArtifactsBySet(artifactSet)
		if err != nil {
			if errors.Is(err, repository.ErrArtifactNotFound) {
				c.JSON(404, gin.H{"error": err.Error()})
				return
			} else {
				c.JSON(500, gin.H{"error": "Internal server error"})
				return
			}
		}
		c.JSON(200, artifacts)
	}
}

func GetArtifacts(service service.GetArtifactsServiceInterface) {}

func CreateArtifact(service service.CreateArtifactServiceInterface) {}
