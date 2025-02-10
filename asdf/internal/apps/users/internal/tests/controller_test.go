package users_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	rest "asdf/internal/apps/users/internal/adapters/httprouter"
	"asdf/internal/dtos"
	"asdf/internal/lib"

	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

// GET /users
// 200, 500

// GET /users/:id
// 200, 404, 422, 500

// PUT /users/:id
// 200, 400, 422, 500

// DELETE /users/:id
// 200, 404, 422, 500

// Authentication tests

type handlerTests struct {
	url    string
	method string
	tests  []restTc
}

type restTc struct {
	body           []byte
	tcName         string
	schema         string
	targetUrl      string
	expectedStatus int
	mockReturns    []any
}

func createInvalidUser(t *testing.T) []byte {
	uj, err := json.Marshal(&dtos.CreateUserParams{
		Username: "",
		Email:    "",
		Password: "",
	})
	assert.Nil(t, err)
	return uj
}

func createRESTTestCases(t *testing.T) []handlerTests {
	t.Helper()

	validUser := createValidUser(t)
	validUserJson, err := json.Marshal(validUser)
	assert.Nil(t, err)

	createdUser := &dtos.User{
		ID:        1,
		Username:  randomUsername(),
		
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tc := []handlerTests{
		{
			url:    "/users",
			method: http.MethodGet,
			// GET /users
			// 200, 200, 500
			tests: []restTc{
				{
					tcName:         "find users empty list",
					schema:         findUsersSchema200,
					expectedStatus: http.StatusOK,
					targetUrl:      "/users",
					mockReturns:    []any{[]*dtos.User{}, nil},
				},
				{
					tcName:         "find users with data",
					schema:         findUsersSchema200,
					expectedStatus: http.StatusOK,
					targetUrl:      "/users",
					mockReturns:    []any{[]*dtos.User{createdUser}, nil},
				},
				{
					tcName:         "internal server error",
					schema:         schema500,
					expectedStatus: http.StatusInternalServerError,
					targetUrl:      "/users",
					mockReturns:    []any{nil, errors.New("")},
				},
			},
		},
		{
			url:    "/users/:id",
			method: http.MethodGet,
			// GET /users/:id
			// 200, 404, 422, 500
			tests: []restTc{
				{
					tcName:         "get user with id",
					schema:         getUserSchema200,
					expectedStatus: http.StatusOK,
					targetUrl:      "/users/1",
					mockReturns:    []any{&dtos.User{ID: 1, Username: "qwer", Email: randomEmail(), CreatedAt: time.Now(), UpdatedAt: time.Now()}, nil},
				},
				{
					tcName:         "record not found",
					schema:         schema404,
					expectedStatus: http.StatusNotFound,
					targetUrl:      "/users/10000",
					mockReturns:    []any{nil, lib.ErrNoRecord},
				},
				{
					tcName:         "invalid id",
					schema:         schemaInvalidID422,
					expectedStatus: http.StatusUnprocessableEntity,
					targetUrl:      "/users/asdf",
					mockReturns:    nil,
				},
				{
					tcName:         "internal server error",
					schema:         schema500,
					expectedStatus: http.StatusInternalServerError,
					targetUrl:      "/users/1",
					mockReturns:    []any{nil, errors.New("")},
				},
			},
		},
		{
			url:    "/users",
			method: http.MethodPost,
			// POST /users
			// 201, 400, 422, 500
			tests: []restTc{
				{
					tcName:         "",
					schema:         schema201,
					expectedStatus: http.StatusCreated,
					body:           validUserJson,
					targetUrl:      "/users",
					mockReturns:    []any{int64(1), nil},
				},
				{
					tcName:         "malformed json",
					schema:         schema400,
					expectedStatus: http.StatusBadRequest,
					targetUrl:      "/users",
					mockReturns:    []any{int64(1), nil},
					body:           []byte(`{"username":}`),
				},
				{
					tcName:         "validation errors",
					schema:         userValidationSchema422,
					expectedStatus: http.StatusUnprocessableEntity,
					targetUrl:      "/users",
					mockReturns: []any{
						int64(-1), lib.ErrInvalidData{
							"username": []string{},
							"email":    []string{},
							"password": []string{},
						},
					},
					body: createInvalidUser(t),
				},
				{
					tcName:         "internal server error",
					schema:         schema500,
					expectedStatus: http.StatusInternalServerError,
					targetUrl:      "/users",
					mockReturns:    []any{int64(-1), errors.New("")},
					body:           validUserJson,
				},
			},
		},
		{
			url:    "/users/:id",
			method: http.MethodPatch,
			// 200, 400, 404, 422, 500
			tests: []restTc{
				{
					tcName:         "successful patch",
					schema:         patchUserSchema200,
					expectedStatus: http.StatusOK,
					body:           validUserJson,
					targetUrl:      "/users/1",
					mockReturns:    []any{nil},
				},
				{
					tcName:         "malformed json",
					schema:         schema400,
					expectedStatus: http.StatusBadRequest,
					targetUrl:      "/users/1",
					mockReturns:    []any{nil},
					body:           []byte(`{"username":}`),
				},
				{
					tcName:         "record not found",
					schema:         schema404,
					expectedStatus: http.StatusNotFound,
					targetUrl:      "/users/10000",
					mockReturns:    []any{lib.ErrNoRecord},
					body:           createInvalidUser(t),
				},
				{
					tcName:         "invalid id",
					schema:         schemaInvalidID422,
					expectedStatus: http.StatusUnprocessableEntity,
					targetUrl:      "/users/-1",
					mockReturns:    []any{nil},
					body:           createInvalidUser(t),
				},
				{
					tcName:         "validation errors",
					schema:         userValidationSchema422,
					expectedStatus: http.StatusUnprocessableEntity,
					targetUrl:      "/users/1",
					mockReturns: []any{
						lib.ErrInvalidData{
							"username": []string{},
							"email":    []string{},
							"password": []string{},
						},
					},
					body: createInvalidUser(t),
				},
				{
					tcName:         "internal server error",
					schema:         schema500,
					expectedStatus: http.StatusInternalServerError,
					targetUrl:      "/users/1",
					mockReturns:    []any{errors.New("")},
					body:           validUserJson,
				},
			},
		},
		{
			url:    "/users/:id",
			method: http.MethodDelete,
			// 200, 400, 404, 422, 500
			tests: []restTc{
				{
					tcName:         "successful delete",
					schema:         deleteUserSchema200,
					expectedStatus: http.StatusOK,
					targetUrl:      "/users/1",
					mockReturns:    []any{nil},
				},
				{
					tcName:         "no record found",
					schema:         schema404,
					expectedStatus: http.StatusNotFound,
					targetUrl:      "/users/10000",
					mockReturns:    []any{lib.ErrNoRecord},
				},
				{
					tcName:         "invalid id",
					schema:         schemaInvalidID422,
					expectedStatus: http.StatusUnprocessableEntity,
					targetUrl:      "/users/-1",
					mockReturns:    []any{nil},
					body:           createInvalidUser(t),
				},
				{
					tcName:         "internal server error",
					schema:         schema500,
					expectedStatus: http.StatusInternalServerError,
					targetUrl:      "/users/1",
					mockReturns:    []any{errors.New("")},
					body:           validUserJson,
				},
			},
		},
	}

	return tc
}

func createNewRequest(t *testing.T, method, url string, body []byte) *http.Request {
	t.Helper()
	if body != nil {
		return httptest.NewRequest(method, url, bytes.NewReader(body))
	}
	return httptest.NewRequest(method, url, nil)
}

func TestUsersController(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, handler := range createRESTTestCases(t) {
		for _, tc := range handler.tests {
			t.Run(fmt.Sprintf("-- SchemaValidation | %s  | %s %s - %d", tc.tcName, handler.method, handler.url, tc.expectedStatus), func(t *testing.T) {
				t.Parallel()
				mockApp := NewMockIUsers(ctrl)
				usersController := rest.NewUsersController(rootApp, mockApp)

				router := httprouter.New()

				switch {
				case handler.method == http.MethodGet && handler.url == "/users":
					router.HandlerFunc(handler.method, handler.url, usersController.FindUsersH)
					if tc.mockReturns != nil {
						mockApp.EXPECT().FindUsers(gomock.Any(), gomock.Any()).Return(tc.mockReturns...).Times(1)
					}
				case handler.method == http.MethodGet && handler.url == "/users/:id":
					router.HandlerFunc(handler.method, handler.url, usersController.GetUserH)
					if tc.mockReturns != nil {
						mockApp.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(tc.mockReturns...).Times(1)
					}

				case handler.method == http.MethodPost && handler.url == "/users":
					router.HandlerFunc(handler.method, handler.url, usersController.CreateUserH)
					if tc.mockReturns != nil {
						mockApp.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(tc.mockReturns...).Times(1)
					}
				case handler.method == http.MethodPatch && handler.url == "/users/:id":
					router.HandlerFunc(handler.method, handler.url, usersController.PatchUserH)
					if tc.mockReturns != nil {
						mockApp.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(tc.mockReturns...).Times(1)
					}

				case handler.method == http.MethodDelete && handler.url == "/users/:id":
					router.HandlerFunc(handler.method, handler.url, usersController.DeleteUserH)
					if tc.mockReturns != nil {
						mockApp.EXPECT().DeleteUser(gomock.Any(), gomock.Any()).Return(tc.mockReturns...).Times(1)
					}
				}

				w := httptest.NewRecorder()
				r := createNewRequest(t, handler.method, tc.targetUrl, tc.body)

				router.ServeHTTP(w, r)

				gotStatus := w.Result().StatusCode
				assert.Equalf(t, tc.expectedStatus, gotStatus, "tc: '%s' | %s %s | want %d, got %d", tc.tcName, handler.method, tc.targetUrl, tc.expectedStatus, gotStatus)

				body, err := io.ReadAll(w.Result().Body)
				assert.Nil(t, err, "want nil")

				validateSchema(t, tc.schema, body)
			})
		}
	}
}
