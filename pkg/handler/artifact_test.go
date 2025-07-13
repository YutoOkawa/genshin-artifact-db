package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"genshin-artifact-db/pkg/repository"
	"genshin-artifact-db/pkg/service"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
)

func TestGetArtifact(t *testing.T) {
	testArtifactDTO := &service.ArtifactDTO{
		Set:   "test-set",
		Type:  "test-type",
		Level: 20,
		PrimaryStat: service.StatusDTO{
			Type:  "test-type",
			Value: 0,
		},
		SubStat: []service.StatusDTO{
			{
				Type:  "test-sub-type",
				Value: 0,
			},
		},
	}

	tests := []struct {
		name string

		// GIVEN
		mockArtifactDTO   *service.ArtifactDTO
		mockArtifactError error

		// WHEN
		artifactID string

		// THEN
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "ShouldGetArtifactSuccessfully",

			mockArtifactDTO: testArtifactDTO,

			artifactID: "test-id",

			expectedStatusCode: 200,
			expectedResponse: func() string {
				response, _ := json.Marshal(testArtifactDTO)
				return string(response)
			}(),
		},
		{
			name: "ShouldReturnErrorWhenArtifactNotFound",

			mockArtifactDTO:   nil,
			mockArtifactError: repository.ErrArtifactNotFound,

			artifactID: "non-existent-id",

			expectedStatusCode: 404,
			expectedResponse:   `{"error":"artifact not found"}`,
		},
		{
			name: "ShouldReturnErrorWhenGetArtifactByIDFails",

			mockArtifactDTO:   nil,
			mockArtifactError: errors.New("internal server error"),

			artifactID: "test-id",

			expectedStatusCode: 500,
			expectedResponse:   `{"error":"Internal server error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			artifactService := &service.MockGetArtifactService{
				MockArtifact:         tt.mockArtifactDTO,
				MockGetArtifactError: tt.mockArtifactError,
			}
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.GET("/artifact/:id", GetArtifact(artifactService))

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/artifact/%s", tt.artifactID), nil)
			r.ServeHTTP(w, req)

			if w.Code != tt.expectedStatusCode {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatusCode, w.Code)
			}

			if diff := cmp.Diff(tt.expectedResponse, w.Body.String()); diff != "" {
				t.Errorf("Response mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
