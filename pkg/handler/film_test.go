package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	vkfilms "github.com/bitbox228/vk-films-api"
	"github.com/bitbox228/vk-films-api/pkg/service"
	mockservice "github.com/bitbox228/vk-films-api/pkg/service/mocks"
	"github.com/magiconair/properties/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler_createFilm(t *testing.T) {
	type mockBehaviour func(s *mockservice.MockFilm, actor vkfilms.CreateFilmInput)

	parse, _ := time.Parse(time.DateOnly, "2000-01-01")
	testTable := []struct {
		name                string
		inputBody           string
		inputFilm           vkfilms.CreateFilmInput
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "ok",
			inputBody: `{"name":"test","description":"test description","rating":2,"date":"2000-01-01"}`,
			inputFilm: vkfilms.CreateFilmInput{
				Name:        "test",
				Description: "test description",
				Rating:      2,
				Date:        vkfilms.JsonDate(parse),
			},
			mockBehaviour: func(s *mockservice.MockFilm, film vkfilms.CreateFilmInput) {
				s.EXPECT().Create(film).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			film := mockservice.NewMockFilm(c)
			testCase.mockBehaviour(film, testCase.inputFilm)

			services := &service.Service{Film: film}
			handler := NewHandler(services)

			mux := http.NewServeMux()
			mux.HandleFunc("POST /films", handler.createFilm)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/films",
				bytes.NewBufferString(testCase.inputBody))

			mux.ServeHTTP(w, req)

			var got, expected map[string]interface{}
			json.Unmarshal([]byte(w.Body.String()), &got)
			json.Unmarshal([]byte(testCase.expectedRequestBody), &expected)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, got, expected)
		})
	}
}

func TestHandler_deleteFilm(t *testing.T) {
	type mockBehaviour func(s *mockservice.MockFilm, id int)

	testTable := []struct {
		name                string
		id                  int
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "ok",
			id:   1,
			mockBehaviour: func(s *mockservice.MockFilm, id int) {
				s.EXPECT().Delete(id).Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"status":"ok"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			film := mockservice.NewMockFilm(c)
			testCase.mockBehaviour(film, testCase.id)

			services := &service.Service{Film: film}
			handler := NewHandler(services)

			mux := http.NewServeMux()
			mux.HandleFunc("DELETE /films", handler.deleteFilm)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", fmt.Sprintf("/films?id=%d", testCase.id),
				bytes.NewBufferString(""))

			mux.ServeHTTP(w, req)

			var got, expected map[string]interface{}
			json.Unmarshal([]byte(w.Body.String()), &got)
			json.Unmarshal([]byte(testCase.expectedRequestBody), &expected)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, got, expected)
		})
	}
}
