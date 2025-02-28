// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/login": {
            "post": {
                "description": "Login Users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users-Users"
                ],
                "summary": "Login Users",
                "parameters": [
                    {
                        "description": "Login Users",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/login.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/login.LoginResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "security": [
                    {
                        "JWTBearer": []
                    }
                ],
                "description": "Logout Users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users-Users"
                ],
                "summary": "Logout Users",
                "parameters": [
                    {
                        "description": "Logut Users",
                        "name": "logout",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/logout.LogoutRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/logout.LogoutResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/movies": {
            "post": {
                "security": [
                    {
                        "JWTBearer": []
                    }
                ],
                "description": "Create Movie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Movie"
                ],
                "summary": "Create Movie",
                "parameters": [
                    {
                        "description": "Create Movie",
                        "name": "movie",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/create.CreateMovieRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/create.CreateMovieResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/movies/:uuid": {
            "delete": {
                "security": [
                    {
                        "JWTBearer": []
                    }
                ],
                "description": "Unvote Movie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users-Movie"
                ],
                "summary": "Unvote Movie",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Movie UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Unvote Movie",
                        "name": "vote",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/unvote.UnVoteMovieRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/unvote.UnVoteMovieResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/movies/list": {
            "get": {
                "security": [
                    {
                        "JWTBearer": []
                    }
                ],
                "description": "Search Movie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users-Movie"
                ],
                "summary": "Search Movie",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Per Page",
                        "name": "per_page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Title",
                        "name": "title",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Description",
                        "name": "description",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Artists",
                        "name": "artists",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Genres",
                        "name": "genres",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/getmoviesbyfilter.GetMovieByFilterResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/movies/most-viewed": {
            "get": {
                "security": [
                    {
                        "JWTBearer": []
                    }
                ],
                "description": "Most Viewed Movie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Movie"
                ],
                "summary": "Most Viewed Movie",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/getmostviewedmovie.GetMostViewedMovieRequestResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/movies/track": {
            "post": {
                "security": [
                    {
                        "JWTBearer": []
                    }
                ],
                "description": "Track Movie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users-Movie"
                ],
                "summary": "Track Movie",
                "parameters": [
                    {
                        "description": "Unvote Movie",
                        "name": "vote",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/createorupdate.CreateOrUpdateViewershipRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/createorupdate.CreateOrUpdateViewershipResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/movies/upload": {
            "post": {
                "security": [
                    {
                        "JWTBearer": []
                    }
                ],
                "description": "Upload Movie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Movie"
                ],
                "summary": "Upload Movie",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Movie File",
                        "name": "movie",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/movies/vote": {
            "post": {
                "security": [
                    {
                        "JWTBearer": []
                    }
                ],
                "description": "Vote Movie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users-Movie"
                ],
                "summary": "Vote Movie",
                "parameters": [
                    {
                        "description": "Vote Movie",
                        "name": "vote",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vote.VoteMovieRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/vote.VoteMovieResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/movies/votes/list": {
            "get": {
                "security": [
                    {
                        "JWTBearer": []
                    }
                ],
                "description": "Unvote Movie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users-Movie"
                ],
                "summary": "Unvote Movie",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/getvotedmoviesbyuser.GetVotedMovieByUserResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/movies/{uuid}": {
            "put": {
                "security": [
                    {
                        "JWTBearer": []
                    }
                ],
                "description": "Update Movie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Movie"
                ],
                "summary": "Update Movie",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Movie UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Movie",
                        "name": "movie",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/update.UpdateMovieRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/update.UpdateMovieResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/users/sign-up": {
            "post": {
                "description": "Sign Up Users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users-Users"
                ],
                "summary": "Sign Up Users",
                "parameters": [
                    {
                        "description": "Sign Up Admin",
                        "name": "sign-up",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/create.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/create.CreateUserResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/users/sign-up/admin": {
            "post": {
                "description": "Sign Up Admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-users"
                ],
                "summary": "Sign Up Admin",
                "parameters": [
                    {
                        "description": "Sign Up Admin",
                        "name": "sign-up",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/create.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/create.CreateUserResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "create.CreateMovieRequest": {
            "type": "object",
            "required": [
                "artists",
                "description",
                "duration",
                "genres",
                "title",
                "url"
            ],
            "properties": {
                "artists": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "description": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "genres": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "create.CreateMovieResponse": {
            "type": "object",
            "properties": {
                "uuid": {
                    "type": "string"
                }
            }
        },
        "create.CreateUserRequest": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 16,
                    "minLength": 8
                }
            }
        },
        "create.CreateUserResponse": {
            "type": "object",
            "properties": {
                "uuid": {
                    "type": "string"
                }
            }
        },
        "createorupdate.CreateOrUpdateViewershipRequest": {
            "type": "object",
            "required": [
                "movie_uuid",
                "user_uuid",
                "watching_duration"
            ],
            "properties": {
                "created_by": {
                    "type": "string"
                },
                "movie_uuid": {
                    "type": "string"
                },
                "user_uuid": {
                    "type": "string"
                },
                "watching_duration": {
                    "type": "integer"
                }
            }
        },
        "createorupdate.CreateOrUpdateViewershipResponse": {
            "type": "object",
            "required": [
                "movie_uuid",
                "user_uuid"
            ],
            "properties": {
                "movie_uuid": {
                    "type": "string"
                },
                "user_uuid": {
                    "type": "string"
                }
            }
        },
        "getmostviewedmovie.GetMostViewedMovieRequestResponse": {
            "type": "object",
            "properties": {
                "artists": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "description": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "genres": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                },
                "viewed": {
                    "type": "integer"
                }
            }
        },
        "getmoviesbyfilter.GetMovieByFilterResponse": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/getmoviesbyfilter.MovieResponse"
                    }
                },
                "pagination": {
                    "$ref": "#/definitions/helper.PageInfo"
                }
            }
        },
        "getmoviesbyfilter.MovieResponse": {
            "type": "object",
            "properties": {
                "artists": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "description": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "genres": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "getvotedmoviesbyuser.GetVotedMovieByUserResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/getvotedmoviesbyuser.MovieResponse"
                    }
                }
            }
        },
        "getvotedmoviesbyuser.MovieResponse": {
            "type": "object",
            "properties": {
                "artists": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "description": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "genres": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "helper.BaseResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "helper.PageInfo": {
            "type": "object",
            "properties": {
                "current_page": {
                    "type": "integer"
                },
                "per_page": {
                    "type": "integer"
                },
                "total_data": {
                    "type": "integer"
                },
                "total_page": {
                    "type": "integer"
                }
            }
        },
        "login.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 16,
                    "minLength": 8
                }
            }
        },
        "login.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "logout.LogoutRequest": {
            "type": "object",
            "required": [
                "token"
            ],
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "logout.LogoutResponse": {
            "type": "object",
            "properties": {
                "user_uuid": {
                    "type": "string"
                }
            }
        },
        "unvote.UnVoteMovieRequest": {
            "type": "object",
            "required": [
                "movie_uuid",
                "user_uuid"
            ],
            "properties": {
                "movie_uuid": {
                    "type": "string"
                },
                "user_uuid": {
                    "type": "string"
                }
            }
        },
        "unvote.UnVoteMovieResponse": {
            "type": "object",
            "properties": {
                "uuid": {
                    "type": "string"
                }
            }
        },
        "update.UpdateMovieRequest": {
            "type": "object",
            "required": [
                "artists",
                "description",
                "duration",
                "genres",
                "title",
                "url"
            ],
            "properties": {
                "artists": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "description": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "genres": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "update.UpdateMovieResponse": {
            "type": "object",
            "properties": {
                "artists": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "description": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "genres": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "vote.VoteMovieRequest": {
            "type": "object",
            "required": [
                "movie_uuid",
                "user_uuid"
            ],
            "properties": {
                "movie_uuid": {
                    "type": "string"
                },
                "user_uuid": {
                    "type": "string"
                }
            }
        },
        "vote.VoteMovieResponse": {
            "type": "object",
            "properties": {
                "uuid": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
