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

func TestHandler_createActor(t *testing.T) {
	type mockBehaviour func(s *mockservice.MockActor, actor vkfilms.Actor)

	parse, _ := time.Parse(time.DateOnly, "2000-01-01")
	testTable := []struct {
		name                string
		inputBody           string
		inputActor          vkfilms.Actor
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "ok",
			inputBody: `{"name":"test","sex":"male","birthday":"2000-01-01"}`,
			inputActor: vkfilms.Actor{
				Name:     "test",
				Sex:      vkfilms.MALE,
				Birthday: vkfilms.JsonDate(parse),
			},
			mockBehaviour: func(s *mockservice.MockActor, actor vkfilms.Actor) {
				s.EXPECT().Create(actor).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			actor := mockservice.NewMockActor(c)
			testCase.mockBehaviour(actor, testCase.inputActor)

			services := &service.Service{Actor: actor}
			handler := NewHandler(services)

			mux := http.NewServeMux()
			mux.HandleFunc("POST /actors", handler.createActor)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/actors",
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

func TestHandler_deleteActor(t *testing.T) {
	type mockBehaviour func(s *mockservice.MockActor, id int)

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
			mockBehaviour: func(s *mockservice.MockActor, id int) {
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

			actor := mockservice.NewMockActor(c)
			testCase.mockBehaviour(actor, testCase.id)

			services := &service.Service{Actor: actor}
			handler := NewHandler(services)

			mux := http.NewServeMux()
			mux.HandleFunc("DELETE /actors", handler.deleteActor)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", fmt.Sprintf("/actors?id=%d", testCase.id),
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

func TestHandler_updateActor(t *testing.T) {
	type mockBehaviour func(s *mockservice.MockActor, id int, actor vkfilms.UpdateActorInput)

	parse, _ := time.Parse(time.DateOnly, "2000-01-01")
	name := "test"
	date := vkfilms.JsonDate(parse)
	testTable := []struct {
		name                string
		inputBody           string
		id                  int
		inputActor          vkfilms.UpdateActorInput
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "ok",
			inputBody: `{"name":"test","birthday":"2000-01-01"}`,
			inputActor: vkfilms.UpdateActorInput{
				Name:     &name,
				Birthday: &date,
			},
			mockBehaviour: func(s *mockservice.MockActor, id int, actor vkfilms.UpdateActorInput) {
				s.EXPECT().Update(id, actor).Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"status":"ok"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			actor := mockservice.NewMockActor(c)
			testCase.mockBehaviour(actor, testCase.id, testCase.inputActor)

			services := &service.Service{Actor: actor}
			handler := NewHandler(services)

			mux := http.NewServeMux()
			mux.HandleFunc("PUT /actors", handler.updateActor)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", fmt.Sprintf("/actors?id=%d", testCase.id),
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
