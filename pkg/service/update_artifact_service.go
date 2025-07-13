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
	CreateArtifact(artifactCommand CreateArtifactCommand) error
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
	primaryStat, err := entity.NewPrimaryStat(
		artifactCommand.PrimaryStat.Type,
		artifactCommand.PrimaryStat.Value,
	)
	if err != nil {
		return err
	}

	subStats := make([]entity.Substat, 0, len(artifactCommand.Substats))
	for _, substat := range artifactCommand.Substats {
		subStat, err := entity.NewSubstat(substat.Type, substat.Value)
		if err != nil {
			return err
		}
		subStats = append(subStats, *subStat)
	}

	artifact, err := entity.NewArtifact(
		rand.Text(),
		artifactCommand.ArtifactSet,
		artifactCommand.Type,
		artifactCommand.Level,
		*primaryStat,
		subStats,
	)

	if err != nil {
		return err
	}

	if err := s.artifactSaver.SaveArtifact(artifact); err != nil {
		return err
	}
	return nil
}
