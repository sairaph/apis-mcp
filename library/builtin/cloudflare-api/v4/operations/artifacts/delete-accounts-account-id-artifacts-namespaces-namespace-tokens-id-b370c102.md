---
title: Revoke a token
page_id: operation-delete-accounts-account-id-artifacts-namespaces-namespace-tokens-id-38a53eae
path: operations/artifacts
description: Revokes an Artifacts repository token.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/artifacts/namespaces/{namespace}/tokens/{id}
operation_ids:
    - artifacts_tokens_revoke
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Revoke a token

`DELETE /accounts/{account_id}/artifacts/namespaces/{namespace}/tokens/{id}`

Operation ID: `artifacts_tokens_revoke`

Revokes an Artifacts repository token.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "description": "Cloudflare account ID.", "required": true, "schema": {"type": "string"}}]
```

## Definition

```yaml
{"operationId": "artifacts_tokens_revoke", "summary": "Revoke a token", "description": "Revokes an Artifacts repository token.", "parameters": [{"name": "namespace", "in": "path", "description": "Artifacts namespace name.", "required": true, "schema": {"type": "string", "pattern": "^[a-zA-Z0-9][a-zA-Z0-9._-]*$"}}, {"name": "id", "in": "path", "description": "Token ID. Must match /^[0-9a-z]{16}$/.", "required": true, "schema": {"type": "string", "pattern": "^[0-9a-z]{16}$"}}], "responses": {"200": {"description": "Token revoked.", "content": {"application/json": {"schema": {}}}}, "401": {"description": "Authentication required.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}, "minItems": 1}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}}, "result": {"type": "object", "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["result", "success", "errors", "messages"]}}}}, "403": {"description": "Insufficient permissions.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}, "minItems": 1}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}}, "result": {"type": "object", "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["result", "success", "errors", "messages"]}}}}, "404": {"description": "Resource not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}, "minItems": 1}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}}, "result": {"type": "object", "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["result", "success", "errors", "messages"]}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}, "minItems": 1}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "integer", "maximum": 9007199254740991, "minimum": -9007199254740991}, "documentation_url": {"type": "string", "format": "uri"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object"}}, "result": {"type": "object", "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["result", "success", "errors", "messages"]}}}}}, "security": [{"api_token": []}], "tags": ["Artifacts"], "x-api-token-group": ["Artifacts Edit"], "x-fern-availability": "beta", "x-fern-sdk-group-name": "artifacts.namespaces", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
