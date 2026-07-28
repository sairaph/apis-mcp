---
title: List queries
page_id: operation-get-accounts-account-id-workers-observability-queries-7c805779
path: operations/saved-queries
description: List saved queries.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/observability/queries
operation_ids:
    - queries.list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List queries

`GET /accounts/{account_id}/workers/observability/queries`

Operation ID: `queries.list`

List saved queries.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "description": "Your Cloudflare account ID.", "required": true, "schema": {"type": "string"}}]
```

## Definition

```yaml
{"operationId": "queries.list", "summary": "List queries", "description": "List saved queries.", "parameters": [{"name": "page", "in": "query", "schema": {"type": "number", "default": 1, "minimum": 1}}, {"name": "perPage", "in": "query", "schema": {"type": "number", "default": 20, "maximum": 50, "minimum": 5}}, {"name": "order", "in": "query", "schema": {"type": "string", "default": "desc", "enum": ["asc", "desc"]}}, {"name": "orderBy", "in": "query", "schema": {"type": "string", "default": "updated", "enum": ["created", "updated"]}}], "responses": {"200": {"description": "Successful request", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string", "enum": ["Successful request"]}}, "required": ["message"], "type": "object"}}, "result": {"type": "array", "items": {"$ref": "#/components/schemas/workers-observability_query"}}, "success": {"type": "boolean", "enum": [true]}}, "required": ["messages", "success", "errors", "result"]}}}}, "401": {"description": "Unauthorized", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"detail": {"type": "string"}, "message": {"type": "string", "enum": ["Unauthorized"]}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["errors", "success", "messages"]}}}}, "404": {"description": "Not found", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"detail": {"type": "string"}, "message": {"type": "string", "enum": ["Not found"]}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["errors", "success", "messages"]}}}}, "500": {"description": "Internal error", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"detail": {"type": "string"}, "message": {"type": "string", "enum": ["Internal error"]}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["errors", "success", "messages"]}}}}}, "tags": ["Saved Queries"], "x-api-token-group": ["Workers Observability Write", "Workers Observability Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "observability.queries", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
