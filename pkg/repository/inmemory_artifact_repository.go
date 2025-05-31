package repository

import "genshin-artifact-db/pkg/entity"

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
		return nil, nil
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
	return result, nil
}
