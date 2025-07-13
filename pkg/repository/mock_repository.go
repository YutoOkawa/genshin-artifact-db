package repository

import "genshin-artifact-db/pkg/entity"

type MockArtifactGetter struct {
	GetArtifactByIDResponse *entity.Artifact
	GetArtifactByIDError    error

	GetArtifactByTypeAndSetResponse []*entity.Artifact
	GetArtifactByTypeAndSetError    error
}

func (m *MockArtifactGetter) GetArtifactByID(id string) (*entity.Artifact, error) {
	return m.GetArtifactByIDResponse, m.GetArtifactByIDError
}

func (m *MockArtifactGetter) GetArtifactByTypeAndSet(artifactType entity.ArtifactType, artifactSet entity.ArtifactSet) ([]*entity.Artifact, error) {
	return m.GetArtifactByTypeAndSetResponse, m.GetArtifactByTypeAndSetError
}

type MockArtifactSaver struct {
	SaveArtifactError error
}

func (m *MockArtifactSaver) SaveArtifact(artifact *entity.Artifact) error {
	return m.SaveArtifactError
}
