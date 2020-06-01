package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"qask/internal/app/model"
	"qask/internal/app/questions/testwww"
	"qask/internal/app/store/sqlstore/mysql"
	"qask/internal/app/store/teststore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_server_handleGetUser(t *testing.T) {
	db, tearDown := mysql.TestDB(t, "root:12345678@tcp(172.20.0.5)/qask_test")
	tearDown("users")

	store := mysql.New(db)
	questions := testwww.New()
	s := newServer(store, questions)

	user := model.TestUser()
	err := s.store.User().CreateUser(user)
	assert.NoError(t, err)

	tests := []struct {
		name    string
		path    string
		payload interface{}
		expect  int
	}{
		{
			name: "valid request with id",
			path: "/users/id/1",
			payload: map[string]string{
				"from": "telegram",
			},
			expect: http.StatusOK,
		},
		{
			name: "valid request with tgid",
			path: "/users/tgid/100",
			payload: map[string]string{
				"from": "telegram",
			},
			expect: http.StatusOK,
		},
		{
			name: "invalid request with id (user not found)",
			path: "/users/id/2",
			payload: map[string]string{
				"from": "telegram",
			},
			expect: http.StatusNotFound,
		},
		{
			name: "invalid request with tgid (user not found)",
			path: "/users/tgid/101",
			payload: map[string]string{
				"from": "telegram",
			},
			expect: http.StatusNotFound,
		},
		{
			name: "invalid request with id (invalid id)",
			path: "/users/id/100000000000000000000000000000000000000",
			payload: map[string]string{
				"from": "telegram",
			},
			expect: http.StatusBadRequest,
		},
		{
			name: "invalid request with tgid (invalid tgid)",
			path: "/users/tgid/100000000000000000000000000000000000000",
			payload: map[string]string{
				"from": "telegram",
			},
			expect: http.StatusBadRequest,
		},
		{
			name: "invalid request with id (invalid from)",
			path: "/users/id/1",
			payload: map[string]string{
				"from": "telega",
			},
			expect: http.StatusBadRequest,
		},
		{
			name: "invalid request with tgid (invalid from)",
			path: "/users/tgid/100",
			payload: map[string]string{
				"from": "telega",
			},
			expect: http.StatusBadRequest,
		},
		{
			name:    "invalid request with id (nil payload)",
			path:    "/users/id/1",
			payload: nil,
			expect:  http.StatusBadRequest,
		},
		{
			name:    "invalid request with tgid (nil payload)",
			path:    "/users/tgid/100",
			payload: nil,
			expect:  http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			if tt.payload != nil {
				err := json.NewEncoder(b).Encode(tt.payload)
				assert.NoError(t, err)
			}

			req, err := http.NewRequest(http.MethodGet, tt.path, b)
			assert.NoError(t, err)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tt.expect, rec.Code)
		})
	}
}

func Test_server_handleQuestionsGet(t *testing.T) {
	store := teststore.New()
	questions := testwww.New()

	server := newServer(store, questions)
	tests := []struct {
		name string
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server = server
		})
	}
}

func Test_server_handleUsersGet(t *testing.T) {
	store := teststore.New()
	questions := testwww.New()
	s := newServer(store, questions)

	tests := []struct {
		name    string
		payload interface{}
		expect  int
	}{
		{
			name: "valid request",
			payload: map[string]string{
				"from": "telegram",
			},
			expect: http.StatusOK,
		},
		{
			name: "invalid request (invalid from)",
			payload: map[string]string{
				"from": "telega",
			},
			expect: http.StatusBadRequest,
		},
		{
			name:    "invalid request (nil payload)",
			payload: nil,
			expect:  http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			if tt.payload != nil {
				err := json.NewEncoder(b).Encode(tt.payload)
				assert.NoError(t, err)
			}

			req, err := http.NewRequest(http.MethodGet, "/users", b)
			assert.NoError(t, err)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tt.expect, rec.Code)
		})
	}
}

func Test_server_handleUsersPost(t *testing.T) {
	store := teststore.New()
	questions := testwww.New()
	s := newServer(store, questions)

	testCases := []struct {
		name    string
		method  string
		payload interface{}
		expect  int
	}{
		{
			name:   "invalid method put",
			method: http.MethodPut,
			expect: http.StatusMethodNotAllowed,
		},
		{
			name:   "valid user",
			method: http.MethodPost,
			payload: map[string]interface{}{
				"firstName": "Nikita",
				"userName":  "Nikita",
				"tgID":      100,
				"from":      "telegram",
			},
			expect: http.StatusCreated,
		},
		{
			name:   "valid user with exists username",
			method: http.MethodPost,
			payload: map[string]interface{}{
				"firstName": "ValidFirstName",
				"userName":  "Nikita",
				"tgID":      1000,
				"from":      "telegram",
			},
			expect: http.StatusBadRequest,
		},
		{
			name:   "valid user with exists telegram id",
			method: http.MethodPost,
			payload: map[string]interface{}{
				"firstName": "ValidFirstName",
				"userName":  "ValidUserName",
				"tgID":      100,
				"from":      "telegram",
			},
			expect: http.StatusBadRequest,
		},
		{
			name:    "nil payload",
			method:  http.MethodPost,
			payload: nil,
			expect:  http.StatusBadRequest,
		},
		{
			name:    "invalid payload",
			method:  http.MethodPost,
			payload: "invalid payload",
			expect:  http.StatusBadRequest,
		},
		{
			name:   "invalid firstname",
			method: http.MethodPost,
			payload: map[string]interface{}{
				"firstName": "n",
				"userName":  "nikita",
				"tgID":      100,
				"from":      "telegram",
			},
			expect: http.StatusBadRequest,
		},
		{
			name:   "invalid username",
			method: http.MethodPost,
			payload: map[string]interface{}{
				"firstName": "nikita",
				"userName":  "nik",
				"tgID":      100,
				"from":      "telegram",
			},
			expect: http.StatusBadRequest,
		},
		{
			name:   "invalid telegram id(eq 0)",
			method: http.MethodPost,
			payload: map[string]interface{}{
				"firstName": "nikita",
				"userName":  "nikita",
				"tgID":      0,
				"from":      "telegram",
			},
			expect: http.StatusBadRequest,
		},
		{
			name:   "invalid telegram id(less than 0)",
			method: http.MethodPost,
			payload: map[string]interface{}{
				"firstName": "nikita",
				"userName":  "nikita",
				"tgID":      -100,
				"from":      "telegram",
			},
			expect: http.StatusBadRequest,
		},
		{
			name:   "invalid from",
			method: http.MethodPost,
			payload: map[string]interface{}{
				"firstName": "nikita",
				"userName":  "nikita",
				"tgID":      100,
				"from":      "telega",
			},
			expect: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			if tc.payload != nil {
				err := json.NewEncoder(b).Encode(tc.payload)
				assert.NoError(t, err)
			}

			req, err := http.NewRequest(tc.method, "/users", b)
			assert.NoError(t, err)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expect, rec.Code)
		})
	}
}
