package service

import (
	"errors"
	"testing"

	"github.com/YutoOkawa/genshin-artifact-db/pkg/repository"
)

func TestUpdateArtifactServiceCreateArtifact(t *testing.T) {
	testArtifactCommand := CreateArtifactCommand{
		ArtifactSet: "Gladiator",
		Type:        "FLOWER",
		Level:       0,
		PrimaryStat: StatCommand{
			Type:  "ATK_PERCENT",
			Value: 0,
		},
		Substats: []StatCommand{
			{
				Type:  "ATK_PERCENT",
				Value: 0,
			},
		},
	}

	tests := []struct {
		name string

		// GIVEN
		mockArtifactSaverError error

		// WHEN
		artifactCommand CreateArtifactCommand

		// THEN
		expectedError bool
	}{
		{
			name: "ShouldCreateArtifactSuccessfully",

			artifactCommand: testArtifactCommand,

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

			err := service.CreateArtifact(tt.artifactCommand)
			if (err != nil) != tt.expectedError {
				t.Errorf("CreateArtifact() error = %v, expectedError %v", err, tt.expectedError)
			}
		})
	}
}
