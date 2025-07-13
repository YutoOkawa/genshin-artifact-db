package repository

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/YutoOkawa/genshin-artifact-db/pkg/entity"
)

var (
	ErrArtifactNotFound      = errors.New("artifact not found")
	ErrArtifactAlreadyExists = errors.New("artifact already exists")
	ErrArtifactIsNil         = errors.New("artifact is nil")
	ErrArtifactIDIsEmpty     = errors.New("artifact ID is empty")
)

type InMemoryArtifactRepository struct {
	Artifacts map[string]*entity.Artifact
}

func NewInMemoryArtifactRepository() *InMemoryArtifactRepository {
	return &InMemoryArtifactRepository{
		Artifacts: make(map[string]*entity.Artifact),
	}
}

func (repo *InMemoryArtifactRepository) GetArtifactByID(id string) (*entity.Artifact, error) {
	if id == "" {
		return nil, ErrArtifactIDIsEmpty
	}

	artifact, exists := repo.Artifacts[id]
	if !exists {
		return nil, ErrArtifactNotFound
	}
	return artifact, nil
}

func (repo *InMemoryArtifactRepository) GetArtifactByTypeAndSet(artifactType entity.ArtifactType, artifactSet entity.ArtifactSet) ([]*entity.Artifact, error) {
	var result []*entity.Artifact
	for _, artifact := range repo.Artifacts {
		if artifact.Type == artifactType && artifact.ArtifactSet == artifactSet {
			result = append(result, artifact)
		}
	}

	if len(result) == 0 {
		return nil, ErrArtifactNotFound
	}

	return result, nil
}

func (repo *InMemoryArtifactRepository) GetArtifactByType(artifactType entity.ArtifactType) ([]*entity.Artifact, error) {
	var result []*entity.Artifact
	for _, artifact := range repo.Artifacts {
		if artifact.Type == artifactType {
			result = append(result, artifact)
		}
	}

	if len(result) == 0 {
		return nil, ErrArtifactNotFound
	}

	return result, nil
}

func (repo *InMemoryArtifactRepository) GetArtifactBySet(artifactSet entity.ArtifactSet) ([]*entity.Artifact, error) {
	var result []*entity.Artifact
	for _, artifact := range repo.Artifacts {
		if artifact.ArtifactSet == artifactSet {
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

	if _, exists := repo.Artifacts[artifact.ID]; exists {
		return ErrArtifactAlreadyExists
	}

	repo.Artifacts[artifact.ID] = artifact
	return nil
}

func (repo *InMemoryArtifactRepository) DeleteArtifactByID(id string) error {
	if id == "" {
		return ErrArtifactIDIsEmpty
	}

	if _, exists := repo.Artifacts[id]; !exists {
		return ErrArtifactNotFound
	}

	delete(repo.Artifacts, id)
	return nil
}

type artifacts struct {
	Artifacts map[string]*entity.Artifact `json:"artifacts"`
}

func (repo *InMemoryArtifactRepository) SaveJSONFile(filename string) error {
	var ArtifactData artifacts
	ArtifactData.Artifacts = repo.Artifacts

	artifactBytes, err := json.Marshal(ArtifactData)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, artifactBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (repo *InMemoryArtifactRepository) LoadJSONFile(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	var ArtifactData artifacts
	err = json.Unmarshal(file, &ArtifactData)
	if err != nil {
		return err
	}
	repo.Artifacts = ArtifactData.Artifacts
	return nil
}
