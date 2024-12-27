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
        "/auth/login": {
            "post": {
                "description": "\"You can either login as a USER or an ADMIN. Case-sensitive roles\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth_Login"
                ],
                "parameters": [
                    {
                        "description": "Auth_Login",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/models.LoginResponse"
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
                            "$ref": "#/definitions/models.RegisterInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/models.RegisterResponse"
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
                            "$ref": "#/definitions/models.SendEmailOtpInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/models.SendEmailResponse"
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
                            "$ref": "#/definitions/models.VerifyOtpInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/models.VerifyOtpResponse"
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
        "/orders": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order_GetAll"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Header must be set for valid response",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/orders.OrderGetAllResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order_Create"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Header must be set for valid response",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Order_Create",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/orders.OrderCreateInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/orders.OrderCreateResponse"
                        }
                    }
                }
            }
        },
        "/orders/cancel/{id}": {
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order_UpdateStatus"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of order",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Order_UpdateStatus",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/orders.OrderUpdateStatusInput"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Header must be set for valid response",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/orders.OrderCancelResponse"
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
        "/orders/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
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
                    },
                    {
                        "type": "string",
                        "description": "Header must be set for valid response",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/orders.OrderGetOneResponse"
                            }
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
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
                    },
                    {
                        "type": "string",
                        "description": "Header must be set for valid response",
                        "name": "Authorization",
                        "in": "header",
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
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
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
                    },
                    {
                        "type": "string",
                        "description": "Header must be set for valid response",
                        "name": "Authorization",
                        "in": "header",
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product_GetAll"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Header must be set for valid response",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/products.ProductGetAllResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product_Create"
                ],
                "parameters": [
                    {
                        "description": "Product_Create",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/products.ProductCreateInput"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Header must be set for valid response",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/products.ProductCreateResponse"
                        }
                    }
                }
            }
        },
        "/products/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
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
                    },
                    {
                        "type": "string",
                        "description": "Header must be set for valid response",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/products.ProductGetOneResponse"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
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
                    },
                    {
                        "type": "string",
                        "description": "Header must be set for valid response",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/products.ProductDeleteResponse"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
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
                    },
                    {
                        "type": "string",
                        "description": "Header must be set for valid response",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Product_Update",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/products.ProductUpdateInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/products.ProductUpdateResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.LoginInput": {
            "type": "object",
            "required": [
                "email",
                "password",
                "role"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string",
                    "enum": [
                        "ADMIN",
                        "USER"
                    ]
                }
            }
        },
        "models.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "models.RegisterInput": {
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
        "models.RegisterResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.SendEmailOtpInput": {
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
        "models.SendEmailResponse": {
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
        "models.VerifyOtpInput": {
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
        "models.VerifyOtpResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "orders.OrderCancelResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "orders.OrderCreateInput": {
            "type": "object",
            "properties": {
                "orderId": {
                    "type": "integer"
                },
                "productId": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "orders.OrderCreateResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "orders.OrderGetAllResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "product1": {
                    "type": "string"
                },
                "product2": {
                    "type": "string"
                },
                "product3": {
                    "type": "string"
                }
            }
        },
        "orders.OrderGetOneResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "productId": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "orders.OrderUpdateStatusInput": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "enum": [
                        "CANCEL"
                    ]
                }
            }
        },
        "products.ProductCreateInput": {
            "type": "object",
            "required": [
                "description",
                "price",
                "stock",
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "imageUrl": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "stock": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "products.ProductCreateResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "imageUrl": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "products.ProductDeleteResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "products.ProductGetAllResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "products.ProductGetOneResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "price": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "products.ProductUpdateInput": {
            "type": "object",
            "required": [
                "description",
                "price",
                "stock",
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "imageUrl": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "stock": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "products.ProductUpdateResponse": {
            "type": "object",
            "properties": {
                "message": {
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
