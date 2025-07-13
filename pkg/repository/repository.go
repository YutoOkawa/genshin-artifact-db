package repository

import "github.com/YutoOkawa/genshin-artifact-db/pkg/entity"

type ArtifactGetter interface {
	GetArtifactByID(id string) (*entity.Artifact, error)
	GetArtifactByTypeAndSet(artifactType entity.ArtifactType, artifactSet entity.ArtifactSet) ([]*entity.Artifact, error)
	GetArtifactByType(artifactType entity.ArtifactType) ([]*entity.Artifact, error)
	GetArtifactBySet(artifactSet entity.ArtifactSet) ([]*entity.Artifact, error)
}

type ArtifactSaver interface {
	SaveArtifact(artifact *entity.Artifact) error
}

type ArtifactDeleter interface {
	DeleteArtifactByID(id string) error
}
