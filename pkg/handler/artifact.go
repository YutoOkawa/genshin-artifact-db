package handler

import (
	"errors"

	"github.com/YutoOkawa/genshin-artifact-db/pkg/repository"
	"github.com/YutoOkawa/genshin-artifact-db/pkg/service"

	"github.com/gin-gonic/gin"
)

type StatRequestParam struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

type CreateArtifactRequestParam struct {
	ArtifactSet string             `json:"artifact_set"`
	Type        string             `json:"type"`
	Level       int                `json:"level"`
	PrimaryStat StatRequestParam   `json:"primary_stat"`
	Substats    []StatRequestParam `json:"substats"`
}

func GetArtifact(artifactService service.GetArtifactServiceInterface) func(c *gin.Context) {
	return func(c *gin.Context) {
		artifactID := c.Param("id")

		artifact, err := artifactService.GetArtifact(artifactID)
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

func GetArtifactsByType(artifactService service.GetArtifactsByTypeServiceInterface) func(c *gin.Context) {
	return func(c *gin.Context) {
		artifactType := c.Param("type")

		artifacts, err := artifactService.GetArtifactsByType(artifactType)
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

func GetArtifactsBySet(artifactService service.GetArtifactsBySetServiceInterface) func(c *gin.Context) {
	return func(c *gin.Context) {
		artifactSet := c.Param("set")

		artifacts, err := artifactService.GetArtifactsBySet(artifactSet)
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

func GetArtifacts(artifactService service.GetArtifactsServiceInterface) func(c *gin.Context) {
	return func(c *gin.Context) {
		artifactType := c.Param("type")
		artifactSet := c.Param("set")

		artifacts, err := artifactService.GetArtifactsByTypeAndSet(artifactType, artifactSet)
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

func CreateArtifact(artifactService service.CreateArtifactServiceInterface) func(c *gin.Context) {
	return func(c *gin.Context) {
		var createArtifactRequestParam CreateArtifactRequestParam
		if err := c.ShouldBindJSON(&createArtifactRequestParam); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		artifactCommand := service.CreateArtifactCommand{
			ArtifactSet: createArtifactRequestParam.ArtifactSet,
			Type:        createArtifactRequestParam.Type,
			Level:       createArtifactRequestParam.Level,
			PrimaryStat: service.StatCommand{
				Type:  createArtifactRequestParam.PrimaryStat.Type,
				Value: createArtifactRequestParam.PrimaryStat.Value,
			},
			Substats: make([]service.StatCommand, len(createArtifactRequestParam.Substats)),
		}

		for i, substat := range createArtifactRequestParam.Substats {
			artifactCommand.Substats[i] = service.StatCommand{
				Type:  substat.Type,
				Value: substat.Value,
			}
		}

		if err := artifactService.CreateArtifact(artifactCommand); err != nil {
			c.JSON(500, gin.H{"error": "Internal server error"})
			return
		}

		c.JSON(201, gin.H{"message": "Artifact created successfully"})
	}
}
