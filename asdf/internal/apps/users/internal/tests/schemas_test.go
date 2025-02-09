package users_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

func validateSchema(t *testing.T, schema string, resp []byte) {
	t.Helper()

	schemaLoader := gojsonschema.NewStringLoader(schema)
	documentLoader := gojsonschema.NewBytesLoader(resp)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	assert.Nil(t, err)

	assert.Equalf(t, true, result.Valid(), "schema validation failed, got %s", string(resp))

	if !result.Valid() {
		for _, desc := range result.Errors() {
			t.Errorf("- %s\n", desc)
		}
	}
}



const deleteUserSchema200 string = `
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "patchUserSchema200",
  "type": "object",
  "properties": {
    "message": {
      "type": "string",
      "pattern": "deleted successfully"
    }
  },
  "required": [
    "message"
  ]
}`

const patchUserSchema200 string = `
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "patchUserSchema200",
  "type": "object",
  "properties": {
    "message": {
      "type": "string",
      "pattern": "patched successfully"
    }
  },
  "required": [
    "message"
  ]
}`

const userValidationSchema422 string = `
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "userValidationSchema422",
  "type": "object",
  "properties": {
    "errors": {
      "type": "object",
      "properties": {
        "username": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "password": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "email": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      },
      "required": [
        "username",
        "password",
        "email"
      ]
    }
  },
  "required": [
    "errors"
  ]
}
`

const schema201 string = `
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "schema201",
  "type": "object",
  "properties": {
    "user_id": {
      "type": "number",
       "minimum": 1
    }
  },
  "required": [
    "user_id"
  ]
}
`

const schema400 string = `
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "schema400",
  "type": "object",
  "properties": {
    "error": {
      "type": "string",
      "minLength": 1
    }
  },
  "required": [
    "error"
  ]
}
`

const schemaInvalidID422 string = `
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "schema422",
  "type": "object",
  "properties": {
    "error": {
      "type": "string",
      "pattern": "invalid parameter: id"
    }
  },
  "required": [
    "error"
  ]
}
`

const schema404 string = `
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "schema404",
  "type": "object",
  "properties": {
    "error": {
      "type": "string",
      "pattern": "The requested resource could not be found"
    }
  },
  "required": [
    "error"
  ]
}`

const getUserSchema200 string = `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "getUserSchema200",
  "type": "object",
  "properties": {
    "user": {
      "type": "object",
      "properties": {
        "id": {
          "type": "number"
        },
        "username": {
          "type": "string"
        },
        "email": {
          "type": "string"
        },
        "created_at": {
          "type": "string"
        },
        "updated_at": {
          "type": "string"
        }
      },
      "required": [
        "id",
        "username",
        "email",
        "created_at",
        "updated_at"
      ]
    }
  },
  "required": [
    "user"
  ]
}`

const schema500 string = `
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "schema500",
  "type": "object",
  "properties": {
    "error": {
      "type": "string",
      "pattern": "The server encountered a problem and could not process your request"
    }
  },
  "required": [
    "error"
  ]
}`

const findUsersSchema200 string = `
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "findUsersSchema",
  "type": "object",
  "properties": {
    "users": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "id": {
            "type": "number"
          },
          "username": {
            "type": "string"
          },
          "email": {
            "type": "string"
          },
          "created_at": {
            "type": "string"
          },
          "updated_at": {
            "type": "string"
          }
        },
        "required": [
          "id",
          "username",
          "email",
          "created_at",
          "updated_at"
        ]
      }
    }
  },
  "required": [
    "users"
  ]
}
`
