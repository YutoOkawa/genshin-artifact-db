package service

type MockGetArtifactService struct {
	MockArtifact         *ArtifactDTO
	MockGetArtifactError error
}

func (s *MockGetArtifactService) GetArtifact(id string) (*ArtifactDTO, error) {
	return s.MockArtifact, s.MockGetArtifactError
}

type MockGetArtifactsByTypeService struct {
	MockArtifacts               []*ArtifactDTO
	MockGetArtifactsByTypeError error
}

func (s *MockGetArtifactsByTypeService) GetArtifactsByType(artifactType string) ([]*ArtifactDTO, error) {
	return s.MockArtifacts, s.MockGetArtifactsByTypeError
}

type MockGetArtifactsBySetService struct {
	MockArtifacts              []*ArtifactDTO
	MockGetArtifactsBySetError error
}

func (s *MockGetArtifactsBySetService) GetArtifactsBySet(artifactSet string) ([]*ArtifactDTO, error) {
	return s.MockArtifacts, s.MockGetArtifactsBySetError
}

type MockGetArtifactByTypeAndSetService struct {
	MockArtifacts                    []*ArtifactDTO
	MockGetArtifactByTypeAndSetError error
}

func (s *MockGetArtifactByTypeAndSetService) GetArtifactByTypeAndSet(artifactType, artifactSet string) ([]*ArtifactDTO, error) {
	return s.MockArtifacts, s.MockGetArtifactByTypeAndSetError
}
