---
title: Search Workers
page_id: operation-get-accounts-account-id-workers-scripts-search-2909cc06
path: operations/worker-script
description: Search for Workers in an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/scripts-search
operation_ids:
    - worker-script-search-workers
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Search Workers

`GET /accounts/{account_id}/workers/scripts-search`

Operation ID: `worker-script-search-workers`

Search for Workers in an account.

## Definition

```yaml
{"operationId": "worker-script-search-workers", "summary": "Search Workers", "description": "Search for Workers in an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "name", "in": "query", "schema": {"description": "Worker name to search for. Both exact and partial matches are returned.", "type": "string", "example": "my-worker"}}, {"name": "id", "in": "query", "schema": {"description": "Worker ID (also called tag) to search for. Only exact matches are returned.", "type": "string", "example": "bdf3567828824b74aadd550004cf4913"}}, {"name": "order_by", "in": "query", "schema": {"description": "Property to sort results by. Results are sorted in ascending order.", "type": "string", "example": "created_on", "default": "name", "enum": ["created_on", "modified_on", "name"]}}, {"name": "page", "in": "query", "description": "Current page.", "schema": {"type": "integer", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "description": "Items per page.", "schema": {"type": "integer", "default": 10, "maximum": 100, "minimum": 1}}], "responses": {"200": {"description": "Search Workers success.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"properties": {"created_on": {"$ref": "#/components/schemas/workers_created_on"}, "environment_is_default": {"description": "Whether the environment is the default environment.", "type": "boolean", "example": true}, "environment_name": {"description": "Name of the environment.", "type": "string", "example": "production"}, "id": {"$ref": "#/components/schemas/workers_identifier"}, "modified_on": {"$ref": "#/components/schemas/workers_modified_on"}, "script_name": {"$ref": "#/components/schemas/workers_script_name"}, "service_name": {"description": "Name of the service.", "type": "string", "example": "my-service"}}, "required": ["id", "script_name", "created_on", "modified_on"], "type": "object"}}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Search Workers failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Script"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.scripts", "x-fern-sdk-method-name": "search", "x-forge-hidden": true}
```
