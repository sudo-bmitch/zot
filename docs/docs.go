// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-12-11 12:03:05.055900322 -0800 PST m=+0.052058015

package docs

import (
	"bytes"
	"encoding/json"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "APIs for Open Container Initiative Distribution Specification",
        "title": "Open Container Initiative Distribution Specification",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "v0.1.0-dev"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v2/": {
            "get": {
                "description": "Check if this API version is supported",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Check API support",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v2/_catalog": {
            "get": {
                "description": "List all image repositories",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "List image repositories",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.RepositoryList"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v2/{name}/blobs/uploads": {
            "post": {
                "description": "Create a new image blob/layer upload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create image blob/layer upload",
                "parameters": [
                    {
                        "type": "string",
                        "description": "repository name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "accepted",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "Location": {
                                "type": "string",
                                "description": "/v2/{name}/blobs/uploads/{uuid}"
                            },
                            "Range": {
                                "type": "string",
                                "description": "bytes=0-0"
                            }
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v2/{name}/blobs/uploads/{uuid}": {
            "get": {
                "description": "Get an image's blob/layer upload given a uuid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get image blob/layer upload",
                "parameters": [
                    {
                        "type": "string",
                        "description": "repository name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "upload uuid",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "no content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update and finish an image's blob/layer upload given a digest",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update image blob/layer upload",
                "parameters": [
                    {
                        "type": "string",
                        "description": "repository name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "upload uuid",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "blob/layer digest",
                        "name": "digest",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an image's blob/layer given a digest",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete image blob/layer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "repository name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "upload uuid",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "Resume an image's blob/layer upload given an uuid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Resume image blob/layer upload",
                "parameters": [
                    {
                        "type": "string",
                        "description": "repository name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "upload uuid",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "accepted",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "Location": {
                                "type": "string",
                                "description": "/v2/{name}/blobs/uploads/{uuid}"
                            },
                            "Range": {
                                "type": "string",
                                "description": "bytes=0-128"
                            }
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "416": {
                        "description": "range not satisfiable",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v2/{name}/blobs/{digest}": {
            "get": {
                "description": "Get an image's blob/layer given a digest",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/vnd.oci.image.layer.v1.tar+gzip"
                ],
                "summary": "Get image blob/layer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "repository name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "blob/layer digest",
                        "name": "digest",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.ImageManifest"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an image's blob/layer given a digest",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete image blob/layer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "repository name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "blob/layer digest",
                        "name": "digest",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "accepted",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "head": {
                "description": "Check an image's blob/layer given a digest",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Check image blob/layer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "repository name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "blob/layer digest",
                        "name": "digest",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.ImageManifest"
                        },
                        "headers": {
                            "api.DistContentDigestKey": {
                                "type": "object",
                                "description": "OK"
                            }
                        }
                    }
                }
            }
        },
        "/v2/{name}/manifests/{reference}": {
            "get": {
                "description": "Get an image's manifest given a reference or a digest",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/vnd.oci.image.manifest.v1+json"
                ],
                "summary": "Get image manifest",
                "parameters": [
                    {
                        "type": "string",
                        "description": "repository name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "image reference or digest",
                        "name": "reference",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.ImageManifest"
                        },
                        "headers": {
                            "api.DistContentDigestKey": {
                                "type": "object",
                                "description": "OK"
                            }
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an image's manifest given a reference or a digest",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update image manifest",
                "parameters": [
                    {
                        "type": "string",
                        "description": "repository name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "image reference or digest",
                        "name": "reference",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an image's manifest given a reference or a digest",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete image manifest",
                "parameters": [
                    {
                        "type": "string",
                        "description": "repository name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "image reference or digest",
                        "name": "reference",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "head": {
                "description": "Check an image's manifest given a reference or a digest",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Check image manifest",
                "parameters": [
                    {
                        "type": "string",
                        "description": "repository name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "image reference or digest",
                        "name": "reference",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "api.DistContentDigestKey": {
                                "type": "object",
                                "description": "OK"
                            }
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v2/{name}/tags/list": {
            "get": {
                "description": "List all image tags in a repository",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "List image tags",
                "parameters": [
                    {
                        "type": "string",
                        "description": "test",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.ImageTags"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.ImageManifest": {
            "type": "object"
        },
        "api.ImageTags": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "api.RepositoryList": {
            "type": "object",
            "properties": {
                "repositories": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        }
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
var SwaggerInfo = swaggerInfo{ Schemes: []string{}}

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface {}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}