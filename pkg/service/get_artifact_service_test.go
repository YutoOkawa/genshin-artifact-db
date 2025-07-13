package service

import (
	"errors"
	"genshin-artifact-db/pkg/entity"
	"genshin-artifact-db/pkg/repository"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetArtifactServiceGetArtifactByID(t *testing.T) {
	testArtifact := &entity.Artifact{
		ID: "test-id",
	}
	tests := []struct {
		name             string
		mockArtifacts    map[string]*entity.Artifact
		artifactID       string
		expectedArtifact *entity.Artifact
		expectedError    error
	}{
		{
			name: "ShouldGetArtifactByIDSuccessfully",
			mockArtifacts: map[string]*entity.Artifact{
				"test-id": testArtifact,
			},
			artifactID:       "test-id",
			expectedArtifact: testArtifact,
			expectedError:    nil,
		},
		{
			name:             "ShouldReturnErrorWhenArtifactNotFound",
			mockArtifacts:    map[string]*entity.Artifact{},
			artifactID:       "non-existent-id",
			expectedArtifact: nil,
			expectedError:    repository.ErrArtifactNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repository.InMemoryArtifactRepository{
				Artifacts: tt.mockArtifacts,
			}
			service := GetArtifactService{
				arrifactGetter: &repo,
			}
			result, err := service.GetArtifactByID(tt.artifactID)

			if diff := cmp.Diff(tt.expectedArtifact, result); diff != "" {
				t.Errorf("GetArtifactByID() mismatch (-want +got):\n%s", diff)
			}

			if !errors.Is(err, tt.expectedError) {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
			}
		})
	}
}
