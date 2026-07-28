---
title: List D1 Databases
page_id: operation-get-accounts-account-id-d1-database-9837a0c7
path: operations/d1
description: Returns a list of D1 databases.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/d1/database
operation_ids:
    - d1-list-databases
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List D1 Databases

`GET /accounts/{account_id}/d1/database`

Operation ID: `d1-list-databases`

Returns a list of D1 databases.

## Definition

```yaml
{"operationId": "d1-list-databases", "summary": "List D1 Databases", "description": "Returns a list of D1 databases.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/d1_account-identifier"}}, {"name": "name", "in": "query", "schema": {"description": "a database name to search for.", "type": "string"}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Number of items per page.", "type": "number", "default": 1000, "maximum": 10000, "minimum": 10}}], "responses": {"200": {"description": "List D1 databases response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/d1_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/d1_database-response"}}, "result_info": {"type": "object", "properties": {"count": {"description": "Total number of results for the requested service", "type": "number", "example": 1}, "page": {"description": "Current page within paginated list of results", "type": "number", "example": 1}, "per_page": {"description": "Number of results per page of results", "type": "number", "example": 20}, "total_count": {"description": "Total results available without any search parameters", "type": "number", "example": 2000}}}}}]}}}}, "4XX": {"description": "List D1 databases response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/d1_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["D1"], "x-api-token-group": ["D1 Read", "D1 Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "d1", "x-fern-sdk-method-name": "list", "x-forge-epilogue": "This command acts on remote D1 Databases.", "x-forge-globals": [{"alias": "v", "description": "Show version number", "name": "version", "required": {"default": false}, "type": "boolean"}, {"description": "Run as if Wrangler was started in the specified directory instead of the current working directory", "name": "cwd", "required": {"default": "."}, "type": "string"}, {"alias": "c", "description": "Path to Wrangler configuration file", "name": "config", "required": {"default": "./wrangler.toml"}, "type": "string"}, {"alias": "e", "description": "Environment to use for operations, and for selecting .env and .dev.vars files", "name": "env", "required": {"default": "dev"}, "type": "string"}, {"description": "Path to an .env file to load", "name": "env-file", "required": {"default": "./.env"}, "type": "string"}, {"alias": "x-provision", "description": "Experimental: Enable automatic resource provisioning", "name": "experimental-provision", "required": {"default": true}, "type": "boolean"}, {"alias": "x-auto-create", "description": "Automatically provision draft bindings with new resources", "name": "experimental-auto-create", "required": {"default": true}, "type": "boolean"}], "x-forge-hidden": true}
```
