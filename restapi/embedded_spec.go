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
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "API",
    "version": "0.0.1"
  },
  "host": "127.0.0.1:52521",
  "basePath": "/",
  "paths": {
    "/": {
      "get": {
        "description": "get the hello world",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/account/all": {
      "get": {
        "description": "dump all accounts in network",
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Error"
          }
        }
      }
    },
    "/account/my": {
      "get": {
        "description": "get the available accounts",
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Error"
          }
        }
      }
    },
    "/account@{num}": {
      "get": {
        "description": "get the detail of one account",
        "parameters": [
          {
            "type": "integer",
            "description": "number of account",
            "name": "num",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Error"
          }
        }
      }
    },
    "/account@{num}/balance": {
      "get": {
        "description": "get the balance of one account",
        "parameters": [
          {
            "type": "integer",
            "description": "number of account",
            "name": "num",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Error"
          }
        }
      }
    },
    "/balance/my": {
      "get": {
        "description": "get all my accounts",
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Error"
          }
        }
      }
    },
    "/block/{hash}": {
      "get": {
        "description": "get the block by hash",
        "parameters": [
          {
            "type": "string",
            "description": "hash of the block",
            "name": "hash",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Error"
          }
        }
      }
    },
    "/block@{height}": {
      "get": {
        "description": "get the block by height",
        "parameters": [
          {
            "type": "integer",
            "description": "height of the block",
            "name": "height",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Error"
          }
        }
      }
    },
    "/node/{addr}": {
      "post": {
        "description": "add remote node",
        "parameters": [
          {
            "type": "string",
            "description": "libp2p style addr of remote node",
            "name": "addr",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Error"
          }
        }
      }
    },
    "/tx/{hash}": {
      "get": {
        "description": "get the tx by hash",
        "parameters": [
          {
            "type": "string",
            "description": "hash of the tx",
            "name": "hash",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Error"
          }
        }
      }
    },
    "/txpool/check/{hash}": {
      "get": {
        "description": "check the tx in hash",
        "parameters": [
          {
            "type": "string",
            "description": "hash of the tx",
            "name": "hash",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Error"
          }
        }
      }
    },
    "/txpool/send/{rawTx}": {
      "post": {
        "description": "send tx into network",
        "parameters": [
          {
            "type": "string",
            "description": "raw of the tx",
            "name": "rawTx",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Error"
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "API",
    "version": "0.0.1"
  },
  "host": "127.0.0.1:52521",
  "basePath": "/",
  "paths": {
    "/": {
      "get": {
        "description": "get the hello world",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/account/all": {
      "get": {
        "description": "dump all accounts in network",
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Error"
          }
        }
      }
    },
    "/account/my": {
      "get": {
        "description": "get the available accounts",
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Error"
          }
        }
      }
    },
    "/account@{num}": {
      "get": {
        "description": "get the detail of one account",
        "parameters": [
          {
            "type": "integer",
            "description": "number of account",
            "name": "num",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Error"
          }
        }
      }
    },
    "/account@{num}/balance": {
      "get": {
        "description": "get the balance of one account",
        "parameters": [
          {
            "type": "integer",
            "description": "number of account",
            "name": "num",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Error"
          }
        }
      }
    },
    "/balance/my": {
      "get": {
        "description": "get all my accounts",
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Error"
          }
        }
      }
    },
    "/block/{hash}": {
      "get": {
        "description": "get the block by hash",
        "parameters": [
          {
            "type": "string",
            "description": "hash of the block",
            "name": "hash",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Error"
          }
        }
      }
    },
    "/block@{height}": {
      "get": {
        "description": "get the block by height",
        "parameters": [
          {
            "type": "integer",
            "description": "height of the block",
            "name": "height",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Error"
          }
        }
      }
    },
    "/node/{addr}": {
      "post": {
        "description": "add remote node",
        "parameters": [
          {
            "type": "string",
            "description": "libp2p style addr of remote node",
            "name": "addr",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Error"
          }
        }
      }
    },
    "/tx/{hash}": {
      "get": {
        "description": "get the tx by hash",
        "parameters": [
          {
            "type": "string",
            "description": "hash of the tx",
            "name": "hash",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Error"
          }
        }
      }
    },
    "/txpool/check/{hash}": {
      "get": {
        "description": "check the tx in hash",
        "parameters": [
          {
            "type": "string",
            "description": "hash of the tx",
            "name": "hash",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Error"
          }
        }
      }
    },
    "/txpool/send/{rawTx}": {
      "post": {
        "description": "send tx into network",
        "parameters": [
          {
            "type": "string",
            "description": "raw of the tx",
            "name": "rawTx",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Error"
          }
        }
      }
    }
  }
}`))
}