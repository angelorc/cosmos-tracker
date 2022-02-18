// Package swagger GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package swagger

import "github.com/swaggo/swag"

const docTemplate_swagger = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/bitsong/{address}": {
            "get": {
                "description": "Get bitsong balances by address.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bitsong"
                ],
                "summary": "Get bitsong balances by address.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bitsong address to query",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.BalancesResponse"
                        }
                    }
                }
            }
        },
        "/osmosis/{address}": {
            "get": {
                "description": "Get osmosis balances by address.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "osmosis"
                ],
                "summary": "Get osmosis balances by address.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Osmosis address to query",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.BalancesResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "server.BalancesResponse": {
            "type": "object",
            "properties": {
                "available": {
                    "type": "object",
                    "additionalProperties": true
                },
                "delegations": {
                    "type": "object",
                    "additionalProperties": true
                },
                "rewards": {
                    "type": "object",
                    "additionalProperties": true
                },
                "totals": {
                    "type": "object",
                    "additionalProperties": true
                }
            }
        }
    }
}`

// SwaggerInfo_swagger holds exported Swagger Info so clients can modify it
var SwaggerInfo_swagger = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Cosmos Tracker Server API",
	Description:      "The cosmos tracker rest server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo_swagger.InstanceName(), SwaggerInfo_swagger)
}
