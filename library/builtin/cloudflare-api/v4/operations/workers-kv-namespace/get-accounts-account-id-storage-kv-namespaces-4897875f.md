---
title: List Namespaces
page_id: operation-get-accounts-account-id-storage-kv-namespaces-2831ab5c
path: operations/workers-kv-namespace
description: Returns the namespaces owned by an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/storage/kv/namespaces
operation_ids:
    - workers-kv-namespace-list-namespaces
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Namespaces

`GET /accounts/{account_id}/storage/kv/namespaces`

Operation ID: `workers-kv-namespace-list-namespaces`

Returns the namespaces owned by an account.

## Definition

```yaml
{"operationId": "workers-kv-namespace-list-namespaces", "summary": "List Namespaces", "description": "Returns the namespaces owned by an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_identifier"}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Maximum number of results per page.", "type": "number", "default": 20, "maximum": 1000, "minimum": 1}}, {"name": "order", "in": "query", "schema": {"description": "Field to order results by.", "type": "string", "example": "id", "enum": ["id", "title"]}}, {"name": "direction", "in": "query", "schema": {"description": "Direction to order namespaces.", "type": "string", "example": "asc", "enum": ["asc", "desc"]}}], "responses": {"200": {"description": "List Namespaces response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers-kv_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/workers-kv_namespace"}}}, "type": "object"}]}}}}, "4XX": {"description": "List Namespaces response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers-kv_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers KV Namespace"], "x-api-token-group": ["Workers KV Storage Write", "Workers KV Storage Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "namespaces", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
