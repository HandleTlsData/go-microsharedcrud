// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https",
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "CRUD operation over Entities.",
    "title": "Microshared-CRUD",
    "version": "1.0.0"
  },
  "paths": {
    "/entity": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "entity"
        ],
        "summary": "Add a new Entity if such Name is not exists. Update otherwise",
        "operationId": "entityStore",
        "parameters": [
          {
            "x-exportParamName": "Body",
            "description": "Entity object that needs to be stored",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Entity"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "405": {
            "description": "Invalid input"
          }
        }
      }
    },
    "/entity/{entityName}": {
      "get": {
        "description": "Finds full entity by it's Name field",
        "produces": [
          "application/json"
        ],
        "tags": [
          "entity"
        ],
        "summary": "Returns a single Entity by given Name",
        "operationId": "entityGet",
        "parameters": [
          {
            "type": "string",
            "x-exportParamName": "EntityName",
            "description": "Name of entity that need to be returned",
            "name": "entityName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Entity"
              }
            }
          },
          "404": {
            "description": "No entity was found"
          }
        }
      },
      "delete": {
        "description": "Delete full entity by it's Name field",
        "produces": [
          "application/json"
        ],
        "tags": [
          "entity"
        ],
        "summary": "Delete a single Entity by given Name",
        "operationId": "entityDelete",
        "parameters": [
          {
            "type": "string",
            "x-exportParamName": "EntityName",
            "description": "Name of entity that need to be returned",
            "name": "entityName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "404": {
            "description": "No entity was found"
          }
        }
      }
    }
  },
  "definitions": {
    "Entity": {
      "type": "object",
      "properties": {
        "Description": {
          "type": "string"
        },
        "ID": {
          "type": "integer",
          "format": "int64"
        },
        "Name": {
          "type": "string"
        },
        "Status": {
          "type": "string"
        }
      },
      "xml": {
        "name": "Entity"
      },
      "example": {
        "Description": "Description",
        "ID": 0,
        "Name": "Name",
        "Status": "Status"
      }
    }
  },
  "tags": [
    {
      "description": "CRUD entities",
      "name": "entity"
    }
  ],
  "externalDocs": {
    "description": "Find out more about Swagger",
    "url": "http://swagger.io"
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https",
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "CRUD operation over Entities.",
    "title": "Microshared-CRUD",
    "version": "1.0.0"
  },
  "paths": {
    "/entity": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "entity"
        ],
        "summary": "Add a new Entity if such Name is not exists. Update otherwise",
        "operationId": "entityStore",
        "parameters": [
          {
            "x-exportParamName": "Body",
            "description": "Entity object that needs to be stored",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Entity"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "405": {
            "description": "Invalid input"
          }
        }
      }
    },
    "/entity/{entityName}": {
      "get": {
        "description": "Finds full entity by it's Name field",
        "produces": [
          "application/json"
        ],
        "tags": [
          "entity"
        ],
        "summary": "Returns a single Entity by given Name",
        "operationId": "entityGet",
        "parameters": [
          {
            "type": "string",
            "x-exportParamName": "EntityName",
            "description": "Name of entity that need to be returned",
            "name": "entityName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Entity"
              }
            }
          },
          "404": {
            "description": "No entity was found"
          }
        }
      },
      "delete": {
        "description": "Delete full entity by it's Name field",
        "produces": [
          "application/json"
        ],
        "tags": [
          "entity"
        ],
        "summary": "Delete a single Entity by given Name",
        "operationId": "entityDelete",
        "parameters": [
          {
            "type": "string",
            "x-exportParamName": "EntityName",
            "description": "Name of entity that need to be returned",
            "name": "entityName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "404": {
            "description": "No entity was found"
          }
        }
      }
    }
  },
  "definitions": {
    "Entity": {
      "type": "object",
      "properties": {
        "Description": {
          "type": "string"
        },
        "ID": {
          "type": "integer",
          "format": "int64"
        },
        "Name": {
          "type": "string"
        },
        "Status": {
          "type": "string"
        }
      },
      "xml": {
        "name": "Entity"
      },
      "example": {
        "Description": "Description",
        "ID": 0,
        "Name": "Name",
        "Status": "Status"
      }
    }
  },
  "tags": [
    {
      "description": "CRUD entities",
      "name": "entity"
    }
  ],
  "externalDocs": {
    "description": "Find out more about Swagger",
    "url": "http://swagger.io"
  }
}`))
}