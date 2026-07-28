---
title: List Namespaces
page_id: operation-get-accounts-account-id-workers-durable-objects-namespaces-c1d0f53c
path: operations/durable-objects-namespace
description: Returns the Durable Object namespaces owned by an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/durable_objects/namespaces
operation_ids:
    - durable-objects-namespace-list-namespaces
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Namespaces

`GET /accounts/{account_id}/workers/durable_objects/namespaces`

Operation ID: `durable-objects-namespace-list-namespaces`

Returns the Durable Object namespaces owned by an account.

## Definition

```yaml
{"operationId": "durable-objects-namespace-list-namespaces", "summary": "List Namespaces", "description": "Returns the Durable Object namespaces owned by an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "page", "in": "query", "description": "Current page.", "schema": {"type": "integer", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "description": "Items per-page.", "schema": {"type": "integer", "default": 20, "maximum": 1000, "minimum": 1}}], "responses": {"200": {"description": "List Namespaces response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/workers_namespace"}}}, "type": "object"}]}}}}, "4XX": {"description": "List Namespaces response failure.", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/workers_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/workers_namespace"}}}, "type": "object"}]}, {"$ref": "#/components/schemas/workers_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Durable Objects Namespace"], "x-api-token-group": ["Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "durable-objects.namespaces", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
