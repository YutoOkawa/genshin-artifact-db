package repository

import (
	"errors"
	"genshin-artifact-db/pkg/entity"
)

var (
	ErrArtifactNotFound      = errors.New("artifact not found")
	ErrArtifactAlreadyExists = errors.New("artifact already exists")
	ErrArtifactIsNil         = errors.New("artifact is nil")
	ErrArtifactIDIsEmpty     = errors.New("artifact ID is empty")
)

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

func (repo *InMemoryArtifactRepository) SaveArtifact(artifact *entity.Artifact) error {
	if artifact == nil {
		return ErrArtifactIsNil
	}

	if artifact.ID == "" {
		return ErrArtifactIDIsEmpty
	}

	if _, exists := repo.artifacts[artifact.ID]; exists {
		return ErrArtifactAlreadyExists
	}

	repo.artifacts[artifact.ID] = artifact
	return nil
}

func (repo *InMemoryArtifactRepository) DeleteArtifactByID(id string) error {
	if id == "" {
		return ErrArtifactIDIsEmpty
	}

	if _, exists := repo.artifacts[id]; !exists {
		return ErrArtifactNotFound
	}

	delete(repo.artifacts, id)
	return nil
}
