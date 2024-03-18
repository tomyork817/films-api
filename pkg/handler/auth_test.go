package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	vkfilms "github.com/bitbox228/vk-films-api"
	"github.com/bitbox228/vk-films-api/pkg/service"
	mockservice "github.com/bitbox228/vk-films-api/pkg/service/mocks"
	"github.com/magiconair/properties/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehaviour func(s *mockservice.MockAuthorization, user vkfilms.User)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           vkfilms.User
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "ok",
			inputBody: `{"username":"test","password":"test","role":"user"}`,
			inputUser: vkfilms.User{
				Name:     "test",
				Password: "test",
				Role:     vkfilms.USER,
			},
			mockBehaviour: func(s *mockservice.MockAuthorization, user vkfilms.User) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1,"role":"user"}`,
		}, {
			name:                "empty fields",
			inputBody:           `{"password":"test","role":"user"}`,
			mockBehaviour:       func(s *mockservice.MockAuthorization, user vkfilms.User) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"not all required fields are filled in"}`,
		}, {
			name:      "service failure",
			inputBody: `{"username":"test","password":"test","role":"user"}`,
			inputUser: vkfilms.User{
				Name:     "test",
				Password: "test",
				Role:     vkfilms.USER,
			},
			mockBehaviour: func(s *mockservice.MockAuthorization, user vkfilms.User) {
				s.EXPECT().CreateUser(user).Return(1, errors.New("service failure"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mockservice.NewMockAuthorization(c)
			testCase.mockBehaviour(auth, testCase.inputUser)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			mux := http.NewServeMux()
			mux.HandleFunc("POST /sign-up", handler.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
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
