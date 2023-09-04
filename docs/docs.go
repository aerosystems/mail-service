// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Artem Kostenko",
            "url": "https://github.com/aerosystems"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "https://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/feedback": {
            "post": {
                "description": "Send feedback",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "feedback"
                ],
                "summary": "Send feedback",
                "parameters": [
                    {
                        "description": "feedback request",
                        "name": "feedbackRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/WebServer.FeedbackRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/WebServer.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/WebServer.ErrResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/WebServer.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/WebServer.ErrResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "WebServer.ErrResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "WebServer.FeedbackRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "WebServer.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.1",
	Host:             "gw.verifire.com/mail",
	BasePath:         "/",
	Schemes:          []string{"https"},
	Title:            "Mail Service",
	Description:      "A part of microservice infrastructure, who responsible for sending emails",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
