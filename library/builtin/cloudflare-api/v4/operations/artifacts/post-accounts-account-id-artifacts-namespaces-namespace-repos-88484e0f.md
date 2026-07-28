---
title: Create a repository
page_id: operation-post-accounts-account-id-artifacts-namespaces-namespace-repos-24464478
path: operations/artifacts
description: Creates a Git-compatible Artifacts repository in a namespace.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/artifacts/namespaces/{namespace}/repos
operation_ids:
    - artifacts_repos_create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a repository

`POST /accounts/{account_id}/artifacts/namespaces/{namespace}/repos`

Operation ID: `artifacts_repos_create`

Creates a Git-compatible Artifacts repository in a namespace.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "description": "Cloudflare account ID.", "required": true, "schema": {"type": "string"}}]
```

## Definition

```yaml
{"operationId": "artifacts_repos_create", "summary": "Create a repository", "description": "Creates a Git-compatible Artifacts repository in a namespace.", "parameters": [{"name": "namespace", "in": "path", "description": "Artifacts namespace name.", "required": true, "schema": {"type": "string", "pattern": "^[a-zA-Z0-9][a-zA-Z0-9._-]*$"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {}}}}, "responses": {"201": {"description": "Repository created.", "content": {"application/json": {"schema": {}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}, "minItems": 1}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}}, "result": {"type": "object", "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["result", "success", "errors", "messages"]}}}}, "401": {"description": "Authentication required.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}, "minItems": 1}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}}, "result": {"type": "object", "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["result", "success", "errors", "messages"]}}}}, "403": {"description": "Insufficient permissions.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}, "minItems": 1}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}}, "result": {"type": "object", "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["result", "success", "errors", "messages"]}}}}, "409": {"description": "Operation conflict.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}, "minItems": 1}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}}, "result": {"type": "object", "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["result", "success", "errors", "messages"]}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}, "minItems": 1}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}}, "result": {"type": "object", "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["result", "success", "errors", "messages"]}}}}}, "security": [{"api_token": []}], "tags": ["Artifacts"], "x-api-token-group": ["Artifacts Edit"], "x-fern-availability": "beta", "x-fern-sdk-group-name": "artifacts.namespaces", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
