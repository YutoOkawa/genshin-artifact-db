package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/YutoOkawa/genshin-artifact-db/pkg/repository"
	"github.com/YutoOkawa/genshin-artifact-db/pkg/service"

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

func TestGetArtifactsByType(t *testing.T) {
	testArtifacts := []*service.ArtifactDTO{
		{
			Set:   "test-set-1",
			Type:  "test-type-1",
			Level: 20,
			PrimaryStat: service.StatusDTO{
				Type:  "test-type-1",
				Value: 0,
			},
			SubStat: []service.StatusDTO{
				{
					Type:  "test-sub-type-1",
					Value: 0,
				},
			},
		},
	}

	tests := []struct {
		name string

		mockArtifacts               []*service.ArtifactDTO
		mockGetArtifactsByTypeError error

		artifactType string

		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "ShouldGetArtifactsByTypeSuccessfully",

			mockArtifacts: testArtifacts,

			artifactType: "test-type",

			expectedStatusCode: 200,
			expectedResponse: func() string {
				response, _ := json.Marshal(testArtifacts)
				return string(response)
			}(),
		},
		{
			name: "ShouldReturnErrorWhenArtifactsNotFound",

			mockArtifacts:               nil,
			mockGetArtifactsByTypeError: repository.ErrArtifactNotFound,

			artifactType: "non-existent-type",

			expectedStatusCode: 404,
			expectedResponse:   `{"error":"artifact not found"}`,
		},
		{
			name: "ShouldReturnErrorWhenGetArtifactsByTypeFails",

			mockArtifacts:               nil,
			mockGetArtifactsByTypeError: errors.New("internal server error"),

			artifactType: "test-type",

			expectedStatusCode: 500,
			expectedResponse:   `{"error":"Internal server error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &service.MockGetArtifactsByTypeService{
				MockArtifacts:               tt.mockArtifacts,
				MockGetArtifactsByTypeError: tt.mockGetArtifactsByTypeError,
			}
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.GET("/artifacts/type/:type", GetArtifactsByType(service))

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/artifacts/type/%s", tt.artifactType), nil)
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

func TestGetArtifactsBySet(t *testing.T) {
	testArtifact := &service.ArtifactDTO{
		Set:   "test-set",
		Type:  "test-type",
		Level: 0,
		PrimaryStat: service.StatusDTO{
			Type:  "test-type",
			Value: 0,
		},
		SubStat: []service.StatusDTO{
			{
				Type:  "test-type",
				Value: 0,
			},
		},
	}

	tests := []struct {
		name string

		mockGetArtifactsBySetResponse []*service.ArtifactDTO
		mockGetArtifactsBySetError    error

		artifactSet string

		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "ShouldGetArtifactsBySetSuccessfully",

			mockGetArtifactsBySetResponse: []*service.ArtifactDTO{testArtifact},

			artifactSet: "test-set",

			expectedStatusCode: 200,
			expectedResponse: func() string {
				response, _ := json.Marshal([]*service.ArtifactDTO{testArtifact})
				return string(response)
			}(),
		},
		{
			name: "ShouldReturnErrorWhenArtifactsNotFound",

			mockGetArtifactsBySetResponse: nil,
			mockGetArtifactsBySetError:    repository.ErrArtifactNotFound,

			artifactSet: "non-existent-set",

			expectedStatusCode: 404,
			expectedResponse:   `{"error":"artifact not found"}`,
		},
		{
			name: "ShouldReturnErrorWhenGetArtifactsBySetFails",

			mockGetArtifactsBySetResponse: nil,
			mockGetArtifactsBySetError:    errors.New("internal server error"),

			artifactSet: "test-set",

			expectedStatusCode: 500,
			expectedResponse:   `{"error":"Internal server error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &service.MockGetArtifactsBySetService{
				MockArtifacts:              tt.mockGetArtifactsBySetResponse,
				MockGetArtifactsBySetError: tt.mockGetArtifactsBySetError,
			}
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.GET("/artifacts/set/:set", GetArtifactsBySet(service))

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/artifacts/set/%s", tt.artifactSet), nil)
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

func TestGetArtifactsByTypeAndSet(t *testing.T) {
	testArtifact := &service.ArtifactDTO{
		Set:   "test-set",
		Type:  "test-type",
		Level: 0,
		PrimaryStat: service.StatusDTO{
			Type:  "test-type",
			Value: 0,
		},
		SubStat: []service.StatusDTO{
			{
				Type:  "test-type",
				Value: 0,
			},
		},
	}

	tests := []struct {
		name string

		mockGetArtifactByTypeAndSetResponse []*service.ArtifactDTO
		mockGetArtifactByTypeAndSetError    error

		artifactType string
		artifactSet  string

		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "ShouldGetArtifactByTypeAndSetSuccessfully",

			mockGetArtifactByTypeAndSetResponse: []*service.ArtifactDTO{testArtifact},

			artifactType: "test-type",
			artifactSet:  "test-set",

			expectedStatusCode: 200,
			expectedResponse: func() string {
				response, _ := json.Marshal([]*service.ArtifactDTO{testArtifact})
				return string(response)
			}(),
		},
		{
			name: "ShouldReturnErrorWhenArtifactNotFound",

			mockGetArtifactByTypeAndSetResponse: nil,
			mockGetArtifactByTypeAndSetError:    repository.ErrArtifactNotFound,

			artifactType: "non-existent-type",
			artifactSet:  "non-existent-set",

			expectedStatusCode: 404,
			expectedResponse:   `{"error":"artifact not found"}`,
		},
		{
			name: "ShouldReturnErrorWhenGetArtifactByTypeAndSetFails",

			mockGetArtifactByTypeAndSetResponse: nil,
			mockGetArtifactByTypeAndSetError:    errors.New("internal server error"),

			artifactType: "test-type",
			artifactSet:  "test-set",

			expectedStatusCode: 500,
			expectedResponse:   `{"error":"Internal server error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &service.MockGetArtifactByTypeAndSetService{
				MockArtifacts:                    tt.mockGetArtifactByTypeAndSetResponse,
				MockGetArtifactByTypeAndSetError: tt.mockGetArtifactByTypeAndSetError,
			}
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.GET("/artifacts/type/:type/set/:set", GetArtifacts(service))

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/artifacts/type/%s/set/%s", tt.artifactType, tt.artifactSet), nil)
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

func TestCreateArtifact(t *testing.T) {
	testCreateArtifactRequestParam := CreateArtifactRequestParam{
		ArtifactSet: "Gladiator's Finale",
		Type:        "FLOWER",
		Level:       0,
		PrimaryStat: StatRequestParam{
			Type:  "ATK_PERCENT",
			Value: 0,
		},
		Substats: []StatRequestParam{
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
		createArtifactRequestParamByte []byte

		// THEN
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "ShouldCreateArtifactSuccessfully",

			createArtifactRequestParamByte: func() []byte {
				body, _ := json.Marshal(testCreateArtifactRequestParam)
				return body
			}(),

			expectedStatusCode: 201,
			expectedResponse:   `{"message":"Artifact created successfully"}`,
		},
		{
			name: "ShouldReturnErrorWhenCreateArtifactRequestParamIsInvalid",

			createArtifactRequestParamByte: []byte(`invalid`),

			expectedStatusCode: 400,
			expectedResponse:   `{"error":"Invalid request body"}`,
		},
		{
			name: "ShouldReturnErrorWhenArtifactSaverFails",

			createArtifactRequestParamByte: func() []byte {
				body, _ := json.Marshal(testCreateArtifactRequestParam)
				return body
			}(),

			mockArtifactSaverError: errors.New("artifact saver error"),

			expectedStatusCode: 500,
			expectedResponse:   `{"error":"Internal server error: artifact saver error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &service.MockCreateArtifactService{
				MockCreateArtifactError: tt.mockArtifactSaverError,
			}

			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.POST("/artifacts", CreateArtifact(service))

			w := httptest.NewRecorder()

			req := httptest.NewRequest("POST", "/artifacts", bytes.NewBuffer(tt.createArtifactRequestParamByte))
			req.Header.Set("Content-Type", "application/json")
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
