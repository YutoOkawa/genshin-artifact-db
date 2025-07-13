package entity

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewPrimaryStat(t *testing.T) {
	testPrimaryStat := &PrimaryStat{
		Type:  "ATK_PERCENT",
		Value: 0.1,
	}

	tests := []struct {
		name string

		// WHEN
		statType string
		value    float64

		// THEN
		expectedPrimaryStat *PrimaryStat
		expectedError       error
	}{
		{
			name: "ShouldNewPrimaryStatSuccessfully",

			statType: "ATK_PERCENT",
			value:    0.1,

			expectedPrimaryStat: testPrimaryStat,
			expectedError:       nil,
		},
		{
			name: "ShouldReturnErrorWhenInvalidPrimaryStatType",

			statType: "INVALID_TYPE",
			value:    0.1,

			expectedPrimaryStat: nil,
			expectedError:       ErrInvalidPrimaryStatType,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			primaryStat, err := NewPrimaryStat(tt.statType, tt.value)

			if diff := cmp.Diff(tt.expectedPrimaryStat, primaryStat); diff != "" {
				t.Errorf("NewPrimaryStat() mismatch (-want +got):\n%s", diff)
			}

			if !errors.Is(err, tt.expectedError) {
				t.Errorf("NewPrimaryStat() error = %v, expectedError %v", err, tt.expectedError)
			}
		})
	}
}

func TestNewSubStat(t *testing.T) {
	testSubStat := &Substat{
		Type:  "ATK_PERCENT",
		Value: 0.1,
	}

	tests := []struct {
		name string

		// WHEN
		statType string
		value    float64

		// THEN
		expectedSubStat *Substat
		expectedError   error
	}{
		{
			name: "ShouldNewSubStatSuccessfully",

			statType: "ATK_PERCENT",
			value:    0.1,

			expectedSubStat: testSubStat,
			expectedError:   nil,
		},
		{
			name: "ShoudlReturnErrorWhenInvalidSubstatType",

			statType: "INVALID_TYPE",
			value:    0.1,

			expectedSubStat: nil,
			expectedError:   ErrInvalidSubstatType,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			subStat, err := NewSubstat(tt.statType, tt.value)

			if diff := cmp.Diff(tt.expectedSubStat, subStat); diff != "" {
				t.Errorf("NewSubstat() mismatch (-want +got):\n%s", diff)
			}

			if !errors.Is(err, tt.expectedError) {
				t.Errorf("NewSubstat() error = %v, expectedError %v", err, tt.expectedError)
			}
		})
	}
}

func TestNewArtifact(t *testing.T) {
	testArtifact := &Artifact{
		ID:          "test-id",
		ArtifactSet: ARTIFACT_SET_BLOODSTAINED_CHIVALRY,
		Type:        ARTIFACT_TYPE_FLOWER,
		Level:       0,
		PrimaryStat: PrimaryStat{Type: "ATK_PERCENT", Value: 0.1},
		Substats:    []Substat{{Type: "ATK_PERCENT", Value: 0.1}},
	}

	tests := []struct {
		name string

		// WHEN
		id           string
		artifactSet  string
		artifactType string
		level        int
		primaryStat  PrimaryStat
		substats     []Substat

		// THEN
		expectedArtifact *Artifact
		expectedError    error
	}{
		{
			name: "ShouldNewArtifactSuccessfully",

			id:           "test-id",
			artifactSet:  string(ARTIFACT_SET_BLOODSTAINED_CHIVALRY),
			artifactType: string(ARTIFACT_TYPE_FLOWER),
			level:        0,
			primaryStat: PrimaryStat{
				Type:  "ATK_PERCENT",
				Value: 0.1,
			},
			substats: []Substat{
				{
					Type:  "ATK_PERCENT",
					Value: 0.1,
				},
			},

			expectedArtifact: testArtifact,
			expectedError:    nil,
		},
		{
			name: "ShouldReturnErrorWhenInvalidArtifactID",

			id: "",

			expectedError: ErrInvalidArtifactID,
		},
		{
			name: "ShouldReturnErrorWhenInvalidArtifactSet",

			id:          "test-id",
			artifactSet: "INVALID_SET",

			expectedError: ErrInvalidArtifactSet,
		},
		{
			name: "ShouldReturnErrorWhenInvalidArtifactType",

			id:           "test-id",
			artifactSet:  string(ARTIFACT_SET_BLOODSTAINED_CHIVALRY),
			artifactType: "INVALID_TYPE",

			expectedError: ErrInvalidArtifactType,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			artifact, err := NewArtifact(
				tt.id,
				tt.artifactSet,
				tt.artifactType,
				tt.level,
				tt.primaryStat,
				tt.substats,
			)

			if diff := cmp.Diff(tt.expectedArtifact, artifact); diff != "" {
				t.Errorf("NewArtifact() mismatch (-want +got):\n%s", diff)
			}

			if !errors.Is(err, tt.expectedError) {
				t.Errorf("NewArtifact() error = %v, expectedError %v", err, tt.expectedError)
			}
		})
	}
}
