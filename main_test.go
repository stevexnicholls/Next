package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stevexnicholls/next/auth"
	"github.com/stevexnicholls/next/models"
	"github.com/stevexnicholls/next/restapi"
	"github.com/stevexnicholls/next/restapi/operations/kv"
	"github.com/go-openapi/swag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const target = "http://localhost:3000/api"

// mocks contains on the restapi.Config dependencies that should be mocked
type mocks struct {
	kv   restapi.MockKvAPI
	backup restapi.MockBackupAPI
}

func (m *mocks) assertExpectations(t *testing.T) {
	m.kv.AssertExpectations(t)
	m.backup.AssertExpectations(t)
}

func TestHTTPHandler(t *testing.T) {
	t.Parallel()

	// declare the test cases
	tests := []struct {
		// name for test case
		name string
		// req is the request that will be tested
		req *http.Request
		// api key that will be added to the request
		apikey string
		// prepare mocks before running the test
		prepare func(*testing.T, *mocks)
		// wantCode is the expected response status code
		wantCode int
		// wantBody is the expected response body
		wantBody []byte
	}{
		{
			name:     "get key value by anonymous should be unauthorized",
			req:      httptest.NewRequest(http.MethodGet, target+"/kv/1", nil),
			wantCode: http.StatusUnauthorized,
		},
		{
			name:   "get key value using a valid api key",
			req:    httptest.NewRequest(http.MethodGet, target+"/kv/1", nil),
			apikey: "abcde",
			prepare: func(t *testing.T, m *mocks) {
				m.kv.On("ValueGet", mock.Anything, mock.Anything).
					Return(&kv.ValueGetOK{Payload: &models.KeyVakye{Key: 1, Value: swag.String("1")}}).
					Once()
			},
			wantCode: http.StatusOK,
			wantBody: []byte(`{"key":1,"name":"2"}`),
		},
		{
			name:     "key value create by anonymous should be unauthorized",
			req:      httptest.NewRequest(http.MethodPut, target+"/kv", bytes.NewBufferString(`{"key":"1"}`)),
			wantCode: http.StatusUnauthorized,
		},
		{
			name:   "key value create by an admin",
			req:    httptest.NewRequest(http.MethodPut, target+"/kv", bytes.NewBufferString(`{"key":"1", "value": "1"}`)),
			apikey: "abcde",
			prepare: func(t *testing.T, m *mocks) {
				m.kv.On("ValueCreate", mock.Anything, mock.Anything).
					Return(&kv.ValueCreateCreated{Payload: &models.KeyValue{Key: 1, Value: 1}}).
					Once()
			},
			wantCode: http.StatusCreated,
			wantBody: []byte(`{"key":1,"value":"1"}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				resp  = httptest.NewRecorder()
				mocks mocks
			)

			h, err := restapi.Handler(restapi.Config{
				KvAPI:     &mocks.kv,
				BackupAPI:   &mocks.backup,
				AuthToken:  auth.Token,
				Authorizer: auth.Request,
				Logger:     t.Logf,
			})
			require.Nil(t, err)

			tt.req.Header.Set("Content-Type", "application/json")
			tt.req.Header.Set("x-api-key", tt.apikey)

			// prepare mocks
			if tt.prepare != nil {
				tt.prepare(t, &mocks)
			}

			h.ServeHTTP(resp, tt.req)

			t.Logf("Got response for request %s %s: %d %s", tt.req.Method, tt.req.URL, resp.Code, resp.Body.String())

			// assert the response expectations
			assert.Equal(t, tt.wantCode, resp.Code)
			if tt.wantBody != nil {
				assert.JSONEq(t, string(tt.wantBody), resp.Body.String())
			}

			mocks.assertExpectations(t)
		})
	}
}