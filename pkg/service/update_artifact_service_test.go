package service

import (
	"errors"
	"testing"

	"github.com/YutoOkawa/genshin-artifact-db/pkg/entity"
	"github.com/YutoOkawa/genshin-artifact-db/pkg/repository"
)

func TestUpdateArtifactServiceCreateArtifact(t *testing.T) {
	testArtifact := entity.Artifact{
		ID: "test-id",
	}

	tests := []struct {
		name string

		// GIVEN
		mockArtifactSaverError error

		// WHEN
		artifact entity.Artifact

		// THEN
		expectedError bool
	}{
		{
			name: "ShouldCreateArtifactSuccessfully",

			artifact: testArtifact,

			expectedError: false,
		},
		{
			name: "ShouldReturnErrorWhenArtifactSaverFails",

			mockArtifactSaverError: errors.New("artifact saver error"),

			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockArtifactSaver := &repository.MockArtifactSaver{
				SaveArtifactError: tt.mockArtifactSaverError,
			}

			service := UpdateArtifactService{
				artifactSaver: mockArtifactSaver,
			}

			err := service.CreateArtifact(tt.artifact)
			if (err != nil) != tt.expectedError {
				t.Errorf("CreateArtifact() error = %v, expectedError %v", err, tt.expectedError)
			}
		})
	}
}
