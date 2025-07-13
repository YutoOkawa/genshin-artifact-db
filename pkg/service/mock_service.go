package service

type MockGetArtifactService struct {
	MockArtifact         *ArtifactDTO
	MockGetArtifactError error
}

func (s *MockGetArtifactService) GetArtifact(id string) (*ArtifactDTO, error) {
	return s.MockArtifact, s.MockGetArtifactError
}
