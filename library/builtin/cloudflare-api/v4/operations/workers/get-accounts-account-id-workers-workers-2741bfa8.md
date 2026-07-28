---
title: List Workers
page_id: operation-get-accounts-account-id-workers-workers-5dee0272
path: operations/workers
description: List all Workers for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/workers
operation_ids:
    - listWorkers
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Workers

`GET /accounts/{account_id}/workers/workers`

Operation ID: `listWorkers`

List all Workers for an account.

## Definition

```yaml
{"operationId": "listWorkers", "summary": "List Workers", "description": "List all Workers for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "page", "in": "query", "description": "Current page.", "schema": {"type": "integer", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "description": "Items per-page.", "schema": {"type": "integer", "default": 10, "maximum": 100, "minimum": 1}}, {"name": "order_by", "in": "query", "description": "Property to sort results by.", "schema": {"type": "string", "default": "deployed_on", "enum": ["deployed_on", "updated_on", "created_on", "name"]}}, {"name": "order", "in": "query", "description": "Sort direction.", "schema": {"type": "string", "default": "desc", "enum": ["asc", "desc"]}}], "responses": {"200": {"description": "List Workers success.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/workers_Worker"}}}, "required": ["result"], "type": "object"}]}}}}, "401": {"$ref": "#/components/responses/workers_ErrorAuth"}, "500": {"$ref": "#/components/responses/workers_ErrorInternalServer"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.beta.workers", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
