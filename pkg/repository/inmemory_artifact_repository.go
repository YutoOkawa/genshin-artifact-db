package repository

import (
	"errors"
	"genshin-artifact-db/pkg/entity"
)

var ErrArtifactNotFound = errors.New("artifact not found")

type InMemoryArtifactRepository struct {
	artifacts map[string]*entity.Artifact
}

func NewInMemoryArtifactRepository() *InMemoryArtifactRepository {
	return &InMemoryArtifactRepository{
		artifacts: make(map[string]*entity.Artifact),
	}
}

func (repo *InMemoryArtifactRepository) GetArtifactByID(id string) (*entity.Artifact, error) {
	artifact, exists := repo.artifacts[id]
	if !exists {
		return nil, ErrArtifactNotFound
	}
	return artifact, nil
}

func (repo *InMemoryArtifactRepository) GetArtifactByTypeAndSet(artifactType entity.ArtifactType, artifactSet entity.ArtifactSet) ([]*entity.Artifact, error) {
	var result []*entity.Artifact
	for _, artifact := range repo.artifacts {
		if artifact.Type == artifactType && artifact.ArtifactSet == artifactSet {
			result = append(result, artifact)
		}
	}

	if len(result) == 0 {
		return nil, ErrArtifactNotFound
	}

	return result, nil
}
