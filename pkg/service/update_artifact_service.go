package service

import (
	"crypto/rand"

	"github.com/YutoOkawa/genshin-artifact-db/pkg/entity"
	"github.com/YutoOkawa/genshin-artifact-db/pkg/repository"
)

type StatCommand struct {
	Type  string
	Value float64
}

type CreateArtifactCommand struct {
	ArtifactSet string
	Type        string
	Level       int
	PrimaryStat StatCommand
	Substats    []StatCommand
}

type CreateArtifactServiceInterface interface {
	CreateArtifact(artifact entity.Artifact) error
}

type UpdateArtifactService struct {
	artifactSaver repository.ArtifactSaver
}

func NewUpdateArtifactService(artifactSaver repository.ArtifactSaver) *UpdateArtifactService {
	return &UpdateArtifactService{
		artifactSaver: artifactSaver,
	}
}

func (s *UpdateArtifactService) CreateArtifact(artifactCommand CreateArtifactCommand) error {
	artifact := entity.NewArtifact(
		rand.Text(),
		entity.ArtifactSet(artifactCommand.ArtifactSet),
		entity.ArtifactType(artifactCommand.Type),
		artifactCommand.Level,
		entity.NewPrimaryStat(
			entity.PrimaryStatType(artifactCommand.PrimaryStat.Type),
			artifactCommand.PrimaryStat.Value,
		),
		func() []entity.Substat {
			substats := make([]entity.Substat, len(artifactCommand.Substats))
			for i, substat := range artifactCommand.Substats {
				substats[i] = entity.NewSubstat(
					entity.SubstatType(substat.Type),
					substat.Value,
				)
			}
			return substats
		}(),
	)

	if err := s.artifactSaver.SaveArtifact(artifact); err != nil {
		return err
	}
	return nil
}
