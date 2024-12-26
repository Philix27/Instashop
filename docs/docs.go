// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "instashop@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth_Register"
                ],
                "parameters": [
                    {
                        "description": "Auth_Register",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.RegisterInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/auth.RegisterResponse"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/types.ErrMsg"
                        }
                    }
                }
            }
        },
        "/auth/send-otp": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth_SendEmailOtp"
                ],
                "parameters": [
                    {
                        "description": "Auth_SendEmailOtp",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.SendEmailOtpInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/auth.SendEmailResponse"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/types.ErrMsg"
                        }
                    }
                }
            }
        },
        "/auth/verify-otp": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth_VerifyOtp"
                ],
                "parameters": [
                    {
                        "description": "Auth_VerifyOtp",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.VerifyOtpInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/auth.VerifyOtpResponse"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/types.ErrMsg"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "tags": [
                    "Auth_Login"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/auth.LoginDto"
                            }
                        }
                    }
                }
            }
        },
        "/orders/{id}": {
            "get": {
                "tags": [
                    "Order_GetOne"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of order",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "Order_Create"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "Order_Delete"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of order",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "tags": [
                    "Order_Update"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of order",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/product/{id}": {
            "delete": {
                "tags": [
                    "Product_Delete"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of product",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "Product update",
                "tags": [
                    "Product_Update"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of product",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "tags": [
                    "Product_GetAll"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "Product_Create"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/products/{id}": {
            "get": {
                "tags": [
                    "Product_GetOne"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of product",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.LoginDto": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "auth.RegisterInput": {
            "type": "object",
            "required": [
                "confirmPassword",
                "email",
                "password",
                "token"
            ],
            "properties": {
                "confirmPassword": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "auth.RegisterResponse": {
            "type": "object",
            "required": [
                "confirmPassword",
                "email",
                "password",
                "token"
            ],
            "properties": {
                "confirmPassword": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "auth.SendEmailOtpInput": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "auth.SendEmailResponse": {
            "type": "object",
            "properties": {
                "otp": {
                    "description": "todo: remove otp",
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "auth.VerifyOtpInput": {
            "type": "object",
            "required": [
                "email",
                "otp",
                "token"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "otp": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "auth.VerifyOtpResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "types.ErrMsg": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:1323",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Instashop Swagger API",
	Description:      "This is a simple e-commerce server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
