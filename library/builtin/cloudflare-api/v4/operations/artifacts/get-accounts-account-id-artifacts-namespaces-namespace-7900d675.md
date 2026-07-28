---
title: Get a namespace
page_id: operation-get-accounts-account-id-artifacts-namespaces-namespace-5a78dc61
path: operations/artifacts
description: Returns an Artifacts namespace summary.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/artifacts/namespaces/{namespace}
operation_ids:
    - artifacts_namespaces_get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a namespace

`GET /accounts/{account_id}/artifacts/namespaces/{namespace}`

Operation ID: `artifacts_namespaces_get`

Returns an Artifacts namespace summary.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "description": "Cloudflare account ID.", "required": true, "schema": {"type": "string"}}]
```

## Definition

```yaml
{"operationId": "artifacts_namespaces_get", "summary": "Get a namespace", "description": "Returns an Artifacts namespace summary.", "parameters": [{"name": "namespace", "in": "path", "description": "Artifacts namespace name.", "required": true, "schema": {"type": "string", "pattern": "^[a-zA-Z0-9][a-zA-Z0-9._-]*$"}}], "responses": {"200": {"description": "Namespace.", "content": {"application/json": {"schema": {}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}, "minItems": 1}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}}, "result": {"type": "object", "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["result", "success", "errors", "messages"]}}}}, "401": {"description": "Authentication required.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}, "minItems": 1}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}}, "result": {"type": "object", "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["result", "success", "errors", "messages"]}}}}, "403": {"description": "Insufficient permissions.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}, "minItems": 1}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}}, "result": {"type": "object", "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["result", "success", "errors", "messages"]}}}}, "404": {"description": "Resource not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}, "minItems": 1}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}}, "result": {"type": "object", "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["result", "success", "errors", "messages"]}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}, "minItems": 1}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}}, "result": {"type": "object", "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["result", "success", "errors", "messages"]}}}}}, "security": [{"api_token": []}], "tags": ["Artifacts"], "x-api-token-group": ["Artifacts Read"], "x-fern-availability": "beta", "x-fern-sdk-group-name": "artifacts.namespaces", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
