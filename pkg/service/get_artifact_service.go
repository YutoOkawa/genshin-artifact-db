package service

import (
	"errors"
	"genshin-artifact-db/pkg/entity"
	"genshin-artifact-db/pkg/repository"
)

var (
	ErrArtifactIDIsEmpty = errors.New("artifact ID is empty")
)

type GetArtifactServiceInterface interface {
	GetArtifactByID(id string) (*entity.Artifact, error)
	GetArtifactByTypeAndSet(artifactType entity.ArtifactType, artifactSet entity.ArtifactSet) ([]*entity.Artifact, error)
}

type GetArtifactService struct {
	arrifactGetter repository.ArtifactGetter
}

func NewGetArtifactService(arrifactGetter repository.ArtifactGetter) *GetArtifactService {
	return &GetArtifactService{
		arrifactGetter: arrifactGetter,
	}
}

func (s *GetArtifactService) GetArtifactByID(id string) (*entity.Artifact, error) {
	return s.arrifactGetter.GetArtifactByID(id)
}
