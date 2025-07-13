package service

import (
	"genshin-artifact-db/pkg/entity"
	"genshin-artifact-db/pkg/repository"
)

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

func (s *UpdateArtifactService) CreateArtifact(artifact entity.Artifact) error {
	if err := s.artifactSaver.SaveArtifact(&artifact); err != nil {
		return err
	}
	return nil
}
