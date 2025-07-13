package repository

import (
	"errors"
	"genshin-artifact-db/pkg/entity"
	"testing"
)

func TestInMemoryArtifactRepositoryGetArtifactByID(t *testing.T) {
	testArtifact := &entity.Artifact{
		ID: "test-id",
	}
	tests := []struct {
		name string

		mockArtifact map[string]*entity.Artifact

		artifactID string

		expectedArtifact *entity.Artifact
		expectedError    bool
	}{
		{
			name: "ShouldInMemoryArtifactRepositoryGetArtifactByIDSuccessfully",

			mockArtifact: map[string]*entity.Artifact{
				"test-id": testArtifact,
			},

			artifactID: "test-id",

			expectedArtifact: testArtifact,
			expectedError:    false,
		},
		{
			name: "ShouldInMemoryArtifactRepositoryReturnErrorWhenArtifactNotFound",

			mockArtifact: map[string]*entity.Artifact{
				"test-id": testArtifact,
			},

			artifactID: "non-existent-id",

			expectedArtifact: nil,
			expectedError:    true,
		},
		{
			name: "ShouldInMemoryArtifactRepositoryReturnErrorWhenArtifactIDIsEmpty",

			mockArtifact: map[string]*entity.Artifact{
				"test-id": testArtifact,
			},

			artifactID: "",

			expectedArtifact: nil,
			expectedError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := InMemoryArtifactRepository{
				Artifacts: tt.mockArtifact,
			}

			artifact, err := repo.GetArtifactByID(tt.artifactID)
			if artifact != tt.expectedArtifact {
				t.Errorf("expected artifact: %v, got: %v", tt.expectedArtifact, artifact)
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
		{
			name: "ShouldInMemoryArtifactRepositoryReturnErrorWhenNoArtifactsFound",

			mockArtifacts: map[string]*entity.Artifact{
				"test-id-1": {
					ID:          "test-id-1",
					ArtifactSet: entity.ARTIFACT_SET_WANDERERS_TROUPE,
					Type:        entity.ARTIFACT_TYPE_FLOWER,
				},
			},

			artifactType: entity.ARTIFACT_TYPE_PLUME,
			artifactSet:  entity.ARTIFACT_SET_GLADIATORS_FINALOFFERING,

			expectedGotArtifactsLength: 0,
			expectedError:              true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := InMemoryArtifactRepository{
				Artifacts: tt.mockArtifacts,
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

func TestInMemoryArtifactRepositorySaveArtifact(t *testing.T) {
	tests := []struct {
		name string

		mockArtifacts map[string]*entity.Artifact

		artifact *entity.Artifact

		expectedError error
	}{
		{
			name: "ShouldInMemoryArtifactRepositorySaveArtifactSuccessfully",

			mockArtifacts: map[string]*entity.Artifact{},

			artifact: &entity.Artifact{
				ID: "new-id",
			},

			expectedError: nil,
		},
		{
			name: "ShouldInMemoryArtifactRepositoryReturnErrorWhenArtifactIDAlreadyExists",

			mockArtifacts: map[string]*entity.Artifact{
				"existing-id": {
					ID: "existing-id",
				},
			},

			artifact: &entity.Artifact{
				ID: "existing-id",
			},

			expectedError: ErrArtifactAlreadyExists,
		},
		{
			name: "ShouldInMemoryArtifactRepositoryReturnErrorWhenArtifactIsNil",

			mockArtifacts: map[string]*entity.Artifact{},

			artifact: nil,

			expectedError: ErrArtifactIsNil,
		},
		{
			name: "ShouldInMemoryArtifactRepositoryReturnErrorWhenArtifactIDIsEmpty",

			mockArtifacts: map[string]*entity.Artifact{},

			artifact: &entity.Artifact{
				ID: "",
			},

			expectedError: ErrArtifactIDIsEmpty,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := InMemoryArtifactRepository{
				Artifacts: tt.mockArtifacts,
			}

			err := repo.SaveArtifact(tt.artifact)

			if !errors.Is(err, tt.expectedError) {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
			}
		})
	}
}

func TestInMemoryArtifactRepositoryDeleteArtifactByID(t *testing.T) {
	tests := []struct {
		name string

		mockArtifacts map[string]*entity.Artifact

		artifactID string

		expectedError error
	}{
		{
			name: "ShouldInMemoryArtifactRepositoryDeleteArtifactByIDSuccessfully",

			mockArtifacts: map[string]*entity.Artifact{
				"test-id": {
					ID: "test-id",
				},
			},

			artifactID: "test-id",

			expectedError: nil,
		},
		{
			name: "ShouldInMemoryArtifactRepositoryReturnErrorWhenArtifactNotFound",

			mockArtifacts: map[string]*entity.Artifact{
				"test-id": {
					ID: "test-id",
				},
			},

			artifactID: "non-existent-id",

			expectedError: ErrArtifactNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := InMemoryArtifactRepository{
				Artifacts: tt.mockArtifacts,
			}

			err := repo.DeleteArtifactByID(tt.artifactID)

			if !errors.Is(err, tt.expectedError) {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
			}
		})
	}
}
