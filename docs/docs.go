// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
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
        "/companies": {
            "get": {
                "description": "List all Companies",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "List Companies",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ListCompaniesResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/company": {
            "post": {
                "description": "Create a new company",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Create Company",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.CreateCompanyRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/handler.CreateCompanyResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.CreateCompanyRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "addressCity": {
                    "type": "string"
                },
                "addressComplement": {
                    "type": "string"
                },
                "addressNumber": {
                    "type": "string"
                },
                "addressState": {
                    "type": "string"
                },
                "addressZipCode": {
                    "type": "string"
                },
                "cnpj": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "handler.CreateCompanyResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.CompanyResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.ErrorResponse": {
            "type": "object",
            "properties": {
                "errorCode": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.ListCompaniesResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.CompanyResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "schemas.CompanyResponse": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "addressCity": {
                    "type": "string"
                },
                "addressComplement": {
                    "type": "string"
                },
                "addressNumber": {
                    "type": "string"
                },
                "addressState": {
                    "type": "string"
                },
                "addressZipCode": {
                    "type": "string"
                },
                "cnpj": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.UserResponse"
                    }
                }
            }
        },
        "schemas.UserResponse": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "companyId": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
