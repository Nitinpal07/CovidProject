// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package echosimple

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
  "swagger": "2.0",
  "info": {
    "description": "Covid-19 Tracker by Nitin Pal",
    "version": "1.0.0",
    "title": "Covid-19 Tracker",
    "termsOfService": "http://swagger.io/terms/",
    "contact": {
      "email": "nitinpaldev@gmail.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    }
  },
  "host": "go-covidtracker.herokuapp.com",
  "basePath": "/",
  "tags": [
    {
      "name": "Covid Tracker",
      "description": "Get Covid-19 Updates in India",
      "externalDocs": {
        "description": "Find out more",
        "url": "http://swagger.io"
      }
    }
  ],
  "schemes": [
    "https",
    "http"
  ],
  "paths": {
    "/covidcases": {
      "get": {
        "summary": "updates covid19 cases in the mongodb .",
        "description": "UPDATE COVID CASES",
        "produces": [
          "application/json"
        ],
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/getCases": {
      "get": {
        "summary": "get covid cases in the state obtained through gps coordinated provided by user .",
        "parameters": [
          {
            "in": "query",
            "name": "lat",
            "description": "latitude value",
            "type": "number",
            "required": true
          },
          {
            "in": "query",
            "name": "lng",
            "description": "longitude value",
            "type": "number",
            "required": true
          }
        ],
        "description": "GET COVID CASES",
        "produces": [
          "application/json"
        ],
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    }
  },
  "securityDefinitions": {
    "petstore_auth": {
      "type": "oauth2",
      "authorizationUrl": "http://petstore.swagger.io/oauth/dialog",
      "flow": "implicit",
      "scopes": {
        "write:pets": "modify pets in your account",
        "read:pets": "read your pets"
      }
    },
    "api_key": {
      "type": "apiKey",
      "name": "api_key",
      "in": "header"
    }
  },
  "definitions": {
    "result": {
      "type": "object",
      "properties": {
        "state": {
          "type": "string",
          "format": "string"
        },
        "activecases": {
          "type": "integer"
        },
        "lastupdatedtime": {
          "type": "string",
          "format": "string"
        },
        "totalcases": {
          "type": "integer"
        }
      },
      "xml": {
        "name": "Result"
      }
    }
  },
  "externalDocs": {
    "description": "Find out more about Swagger",
    "url": "http://swagger.io"
  }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}