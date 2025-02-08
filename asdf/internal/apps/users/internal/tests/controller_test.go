package users_test

import (
	"net/http"
	"testing"

	rest "asdf/internal/apps/users/internal/adapters/httprouter"
)

// GET /users
// 200, 500

// GET /users/:id
// 200, 404, 422, 500

// POST /users
// 201, 400, 422, 500

// PUT /users/:id
// 200, 400, 422, 500

// DELETE /users/:id
// 200, 404, 422, 500

type handlerTests struct {
	url     string
	tests   []restTc
	handler func(w http.ResponseWriter, r *http.Request)
}

type restTc struct {
	method         string
	body           string
	schema         string
	expectedStatus int
}

func createRESTTestCases(t *testing.T, c *rest.UsersController) []handlerTests {
	t.Helper()

	tc := []handlerTests{
		{
			handler: c.FindUsersH,
			url:     "/users",
			tests: []restTc{
				{
					method:         "GET",
					body:           "",
					schema:         "",
					expectedStatus: 200,
				},
			},
		},
	}

	return tc
}
