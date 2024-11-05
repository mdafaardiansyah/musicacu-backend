package memberships

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mdafaardiansyah/musicacu-backend/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_SignUp(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockservice(ctrlMock)

	tests := []struct {
		name               string
		mockFn             func()
		expectedStatusCode int
	}{
		{
			name: "success",
			mockFn: func() {
				mockSvc.EXPECT().SignUp(memberships.SignUpRequest{
					Email:    "test@gmail.com",
					Username: "testusername",
					Password: "password",
				}).Return(nil)
			},
			expectedStatusCode: 201,
		},
		{
			name: "failed",
			mockFn: func() {
				mockSvc.EXPECT().SignUp(memberships.SignUpRequest{
					Email:    "test@gmail.com",
					Username: "testusername",
					Password: "password",
				}).Return(errors.New("username or email exists"))
			},
			expectedStatusCode: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			api := gin.New()

			h := &Handler{
				Engine:  api,
				service: mockSvc,
			}
			h.RegisterRoute()
			w := httptest.NewRecorder()

			endpoint := `/memberships/sign_up`
			model := memberships.SignUpRequest{
				Email:    "test@gmail.com",
				Username: "testusername",
				Password: "password",
			}

			val, err := json.Marshal(model)
			assert.NoError(t, err)

			body := bytes.NewReader(val)
			req, err := http.NewRequest(http.MethodPost, endpoint, body)
			assert.NoError(t, err)
			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}
