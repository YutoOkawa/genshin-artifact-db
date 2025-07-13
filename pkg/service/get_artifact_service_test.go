package service

import (
	"fmt"
	"genshin-artifact-db/pkg/entity"
	"testing"

	"github.com/google/go-cmp/cmp"
)

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

func TestGetArtifactServiceGetArtifactByID(t *testing.T) {
	testArtifact := &entity.Artifact{
		ID:          "test-id",
		ArtifactSet: "test-set",
		Type:        "test-type",
		Level:       0,
		PrimaryStat: entity.PrimaryStat{
			Type:  "test-type",
			Value: 0,
		},
		Substats: []entity.Substat{
			{
				Type:  "test-type",
				Value: 0,
			},
		},
	}

	testArtifactDTO := ArtifactDTO{
		Set:   "test-set",
		Type:  "test-type",
		Level: 0,
		PrimaryStat: StatusDTO{
			Type:  "test-type",
			Value: 0,
		},
		SubStat: []StatusDTO{
			{
				Type:  "test-type",
				Value: 0,
			},
		},
	}
	tests := []struct {
		name string

		mockGetArtifactByIDResponse           *entity.Artifact
		mockGetArtifactByIDError              error
		mockGetArtifactByTypeAndSetRepository map[string]*entity.Artifact
		mockGetArtifactByTypeAndSetError      error

		artifactID string

		expectedArtifact *ArtifactDTO
		expectedError    bool
	}{
		{
			name: "ShouldGetArtifactByIDSuccessfully",

			// GIVEN
			mockGetArtifactByIDResponse: testArtifact,

			// WHEN
			artifactID: "test-id",

			// THEN
			expectedArtifact: &testArtifactDTO,
			expectedError:    false,
		},
		{
			name: "ShouldReturnErrorWhenGetArtifactByIDFails",

			// GIVEN
			mockGetArtifactByIDError: fmt.Errorf("GetArtifactByID error"),

			// WHEN
			artifactID: "non-existent-id",

			// THEN
			expectedArtifact: nil,
			expectedError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MockArtifactGetter{
				GetArtifactByIDResponse: tt.mockGetArtifactByIDResponse,
				GetArtifactByIDError:    tt.mockGetArtifactByIDError,
			}
			service := GetArtifactService{
				arrifactGetter: repo,
			}
			result, err := service.GetArtifactByID(tt.artifactID)

			if diff := cmp.Diff(tt.expectedArtifact, result); diff != "" {
				t.Errorf("GetArtifactByID() mismatch (-want +got):\n%s", diff)
			}

			if (err != nil) != tt.expectedError {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err != nil)
			}
		})
	}
}

func TestGetArtifactServicegetArtifactByTypeAndSet(t *testing.T) {
	testArtifact := &entity.Artifact{
		ID:          "test-id",
		ArtifactSet: "test-set",
		Type:        "test-type",
		Level:       0,
		PrimaryStat: entity.PrimaryStat{
			Type:  "test-type",
			Value: 0,
		},
		Substats: []entity.Substat{
			{
				Type:  "test-type",
				Value: 0,
			},
		},
	}
	testArtifact2 := &entity.Artifact{
		ID:          "test-id-2",
		ArtifactSet: "test-set",
		Type:        "test-type",
		Level:       0,
		PrimaryStat: entity.PrimaryStat{
			Type:  "test-type-2",
			Value: 0,
		},
		Substats: []entity.Substat{
			{
				Type:  "test-type-2",
				Value: 0,
			},
		},
	}

	testArtifactDTO := &ArtifactDTO{
		Set:   "test-set",
		Type:  "test-type",
		Level: 0,
		PrimaryStat: StatusDTO{
			Type:  "test-type",
			Value: 0,
		},
		SubStat: []StatusDTO{
			{
				Type:  "test-type",
				Value: 0,
			},
		},
	}

	testArtifactDTO2 := &ArtifactDTO{
		Set:   "test-set",
		Type:  "test-type",
		Level: 0,
		PrimaryStat: StatusDTO{
			Type:  "test-type-2",
			Value: 0,
		},
		SubStat: []StatusDTO{
			{
				Type:  "test-type-2",
				Value: 0,
			},
		},
	}

	tests := []struct {
		name string

		mockGetArtifactByTypeAndSetResponse []*entity.Artifact
		mockGetArtifactByTypeAndSetError    error

		expectedArtifacts []*ArtifactDTO
		expectedError     bool
	}{
		{
			name: "ShouldGetArtifactByTypeAndSetSuccessfully",

			mockGetArtifactByTypeAndSetResponse: []*entity.Artifact{testArtifact},

			expectedArtifacts: []*ArtifactDTO{testArtifactDTO},
			expectedError:     false,
		},
		{
			name: "ShouldGetAArtifactByTypeAndSetSuccessfully",

			mockGetArtifactByTypeAndSetResponse: []*entity.Artifact{testArtifact, testArtifact2},

			expectedArtifacts: []*ArtifactDTO{testArtifactDTO, testArtifactDTO2},
			expectedError:     false,
		},
		{
			name: "ShouldReturnErrorWhenGetArtifactByTypeAndSetFails",

			mockGetArtifactByTypeAndSetError: fmt.Errorf("GetArtifactByTypeAndSet error"),

			expectedArtifacts: nil,
			expectedError:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MockArtifactGetter{
				GetArtifactByTypeAndSetResponse: tt.mockGetArtifactByTypeAndSetResponse,
				GetArtifactByTypeAndSetError:    tt.mockGetArtifactByTypeAndSetError,
			}
			service := GetArtifactService{
				arrifactGetter: repo,
			}
			result, err := service.GetArtifactByTypeAndSet("test-type", "test-set")

			if diff := cmp.Diff(tt.expectedArtifacts, result); diff != "" {
				t.Errorf("GetArtifactByTypeAndSet() mismatch (-want +got):\n%s", diff)
			}

			if (err != nil) != tt.expectedError {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err != nil)
			}
		})
	}
}
