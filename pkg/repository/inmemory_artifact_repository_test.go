package repository

import (
	"genshin-artifact-db/pkg/entity"
	"testing"
)

func TestInMemoryArtifactRepositoryGetArtifactByIO(t *testing.T) {
	tests := []struct {
		name string

		mockArtifact *entity.Artifact

		expectedError bool
	}{
		{
			name: "ShouldInMemoryArtifactRepositoryGetArtifactByIDSuccessfully",

			mockArtifact: &entity.Artifact{
				ID: "test-id",
			},

			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := InMemoryArtifactRepository{
				artifacts: map[string]*entity.Artifact{
					"test-id": tt.mockArtifact,
				},
			}

			artifact, err := repo.GetArtifactByID("test-id")
			if artifact == nil {
				t.Errorf("expected artifact to be nil, got %v", artifact)
			}

			if (err != nil) != tt.expectedError {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err != nil)
			}
		})
	}
}

func TestInMemoryArtifactRepositoryGetArtifactByTypeAndSet(t *testing.T) {
	tests := []struct {
		name string

		mockArtifacts map[string]*entity.Artifact

		artifactType entity.ArtifactType
		artifactSet  entity.ArtifactSet

		expectedGotArtifactsLength int
		expectedError              bool
	}{
		{
			name: "ShouldInMemoryArtifactRepositoryGetArtifactByTypeAndSetSuccessfully",

			mockArtifacts: map[string]*entity.Artifact{
				"test-id-1": {
					ID:          "test-id-1",
					ArtifactSet: entity.ARTIFACT_SET_GLADIATORS_FINALOFFERING,
					Type:        entity.ARTIFACT_TYPE_FLOWER,
				},
				"test-id-2": {
					ID:          "test-id-2",
					ArtifactSet: entity.ARTIFACT_SET_GLADIATORS_FINALOFFERING,
					Type:        entity.ARTIFACT_TYPE_FLOWER,
				},
			},

			artifactType: entity.ARTIFACT_TYPE_FLOWER,
			artifactSet:  entity.ARTIFACT_SET_GLADIATORS_FINALOFFERING,

			expectedGotArtifactsLength: 2,
			expectedError:              false,
		},
		{
			name: "ShouldInMemoryArtifactRepositoryGetArtifactByTypeAndSetSuccessfullyWithAnotherSet",

			mockArtifacts: map[string]*entity.Artifact{
				"test-id-1": {
					ID:          "test-id-1",
					ArtifactSet: entity.ARTIFACT_SET_WANDERERS_TROUPE,
					Type:        entity.ARTIFACT_TYPE_FLOWER,
				},
				"test-id-2": {
					ID:          "test-id-2",
					ArtifactSet: entity.ARTIFACT_SET_GLADIATORS_FINALOFFERING,
					Type:        entity.ARTIFACT_TYPE_FLOWER,
				},
			},

			artifactType: entity.ARTIFACT_TYPE_FLOWER,
			artifactSet:  entity.ARTIFACT_SET_GLADIATORS_FINALOFFERING,

			expectedGotArtifactsLength: 1,
			expectedError:              false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := InMemoryArtifactRepository{
				artifacts: tt.mockArtifacts,
			}

			result, err := repo.GetArtifactByTypeAndSet(tt.artifactType, tt.artifactSet)
			if len(result) != tt.expectedGotArtifactsLength {
				t.Errorf("expected %d artifacts, got %d", tt.expectedGotArtifactsLength, len(result))
			}

			if (err != nil) != tt.expectedError {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err != nil)
			}
		})
	}
}
