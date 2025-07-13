package service

import (
	"genshin-artifact-db/pkg/entity"
	"genshin-artifact-db/pkg/repository"
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
	GetArtifact(id string) (*ArtifactDTO, error)
}

type GetArtifactsServiceInterface interface {
	GetArtifactByTypeAndSet(artifactType entity.ArtifactType, artifactSet entity.ArtifactSet) ([]*ArtifactDTO, error)
}

type GetArtifactService struct {
	arrifactGetter repository.ArtifactGetter
}

func NewGetArtifactService(arrifactGetter repository.ArtifactGetter) *GetArtifactService {
	return &GetArtifactService{
		arrifactGetter: arrifactGetter,
	}
}

func (s *GetArtifactService) GetArtifact(id string) (*ArtifactDTO, error) {
	artifact, err := s.arrifactGetter.GetArtifactByID(id)
	if err != nil {
		return nil, err
	}

	artifactDTO := &ArtifactDTO{
		Set:   string(artifact.ArtifactSet),
		Type:  string(artifact.Type),
		Level: artifact.Level,
		PrimaryStat: StatusDTO{
			Type:  string(artifact.PrimaryStat.Type),
			Value: artifact.PrimaryStat.Value,
		},
	}

	artifactDTO.SubStat = make([]StatusDTO, 0, len(artifact.Substats))
	for _, subStat := range artifact.Substats {
		subStatDTO := StatusDTO{
			Type:  string(subStat.Type),
			Value: subStat.Value,
		}
		artifactDTO.SubStat = append(artifactDTO.SubStat, subStatDTO)
	}

	return artifactDTO, nil
}

func (s *GetArtifactService) GetArtifactByTypeAndSet(artifactType entity.ArtifactType, artifactSet entity.ArtifactSet) ([]*ArtifactDTO, error) {
	artifacts, err := s.arrifactGetter.GetArtifactByTypeAndSet(artifactType, artifactSet)
	if err != nil {
		return nil, err
	}

	artifactDTOs := make([]*ArtifactDTO, 0, len(artifacts))
	for _, artifact := range artifacts {
		artifactDTO := &ArtifactDTO{
			Set:   string(artifact.ArtifactSet),
			Type:  string(artifact.Type),
			Level: artifact.Level,
			PrimaryStat: StatusDTO{
				Type:  string(artifact.PrimaryStat.Type),
				Value: artifact.PrimaryStat.Value,
			},
		}

		artifactDTO.SubStat = make([]StatusDTO, 0, len(artifact.Substats))
		for _, subStat := range artifact.Substats {
			subStatDTO := StatusDTO{
				Type:  string(subStat.Type),
				Value: subStat.Value,
			}
			artifactDTO.SubStat = append(artifactDTO.SubStat, subStatDTO)
		}

		artifactDTOs = append(artifactDTOs, artifactDTO)
	}

	return artifactDTOs, nil
}

func (s *GetArtifactService) GetArtifactByType(artifactType entity.ArtifactType) ([]*ArtifactDTO, error) {
	artifacts, err := s.arrifactGetter.GetArtifactByType(artifactType)
	if err != nil {
		return nil, err
	}

	artifactDTOs := make([]*ArtifactDTO, 0, len(artifacts))
	for _, artifact := range artifacts {
		artifactDTO := &ArtifactDTO{
			Set:   string(artifact.ArtifactSet),
			Type:  string(artifact.Type),
			Level: artifact.Level,
			PrimaryStat: StatusDTO{
				Type:  string(artifact.PrimaryStat.Type),
				Value: artifact.PrimaryStat.Value,
			},
		}

		artifactDTO.SubStat = make([]StatusDTO, 0, len(artifact.Substats))
		for _, subStat := range artifact.Substats {
			subStatDTO := StatusDTO{
				Type:  string(subStat.Type),
				Value: subStat.Value,
			}
			artifactDTO.SubStat = append(artifactDTO.SubStat, subStatDTO)
		}

		artifactDTOs = append(artifactDTOs, artifactDTO)
	}

	return artifactDTOs, nil
}
