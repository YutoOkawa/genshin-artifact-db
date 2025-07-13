package service

import (
	"errors"
	"genshin-artifact-db/pkg/entity"
	"genshin-artifact-db/pkg/repository"
)

var (
	ErrArtifactIDIsEmpty = errors.New("artifact ID is empty")
)

type StatusDTO struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

type ArtifactDTO struct {
	Set         string      `json:"set"`
	Type        string      `json:"type"`
	Level       int         `json:"level"`
	PrimaryStat StatusDTO   `json:"primary_stat"`
	SubStat     []StatusDTO `json:"sub_stat"`
}

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
