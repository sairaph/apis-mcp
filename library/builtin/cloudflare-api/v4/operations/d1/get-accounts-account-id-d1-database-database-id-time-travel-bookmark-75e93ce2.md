---
title: Get D1 database bookmark
page_id: operation-get-accounts-account-id-d1-database-database-id-time-travel-bookmark-94aad3d6
path: operations/d1
description: |-
    Retrieves the current bookmark, or the nearest bookmark at or before a provided timestamp.
    Bookmarks can be used with the restore endpoint to revert the database to a previous point in time.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/d1/database/{database_id}/time_travel/bookmark
operation_ids:
    - d1-time-travel-get-bookmark
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get D1 database bookmark

`GET /accounts/{account_id}/d1/database/{database_id}/time_travel/bookmark`

Operation ID: `d1-time-travel-get-bookmark`

Retrieves the current bookmark, or the nearest bookmark at or before a provided timestamp.
Bookmarks can be used with the restore endpoint to revert the database to a previous point in time.

## Definition

```yaml
{"operationId": "d1-time-travel-get-bookmark", "summary": "Get D1 database bookmark", "description": "Retrieves the current bookmark, or the nearest bookmark at or before a provided timestamp.\nBookmarks can be used with the restore endpoint to revert the database to a previous point in time.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/d1_account-identifier"}}, {"name": "database_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/d1_database-identifier"}}, {"name": "timestamp", "in": "query", "description": "An optional ISO 8601 timestamp. If provided, returns the nearest available bookmark at or before this timestamp. If omitted, returns the current bookmark.", "schema": {"$ref": "#/components/schemas/d1_time-travel-timestamp"}}], "responses": {"200": {"description": "Bookmark retrieved successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/d1_api-response-common"}, {"properties": {"result": {"type": "object", "properties": {"bookmark": {"$ref": "#/components/schemas/d1_time-travel-bookmark"}}}}}]}}}}, "4XX": {"description": "Failed to retrieve bookmark", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/d1_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["D1"], "x-api-token-group": ["D1 Read", "D1 Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "d1.time-travel", "x-fern-sdk-method-name": "get-bookmark", "x-forge-epilogue": "This command acts on remote D1 Databases.", "x-forge-globals": [{"alias": "v", "description": "Show version number", "name": "version", "required": {"default": false}, "type": "boolean"}, {"description": "Run as if Wrangler was started in the specified directory instead of the current working directory", "name": "cwd", "required": {"default": "."}, "type": "string"}, {"alias": "c", "description": "Path to Wrangler configuration file", "name": "config", "required": {"default": "./wrangler.toml"}, "type": "string"}, {"alias": "e", "description": "Environment to use for operations, and for selecting .env and .dev.vars files", "name": "env", "required": {"default": "dev"}, "type": "string"}, {"description": "Path to an .env file to load", "name": "env-file", "required": {"default": "./.env"}, "type": "string"}, {"alias": "x-provision", "description": "Experimental: Enable automatic resource provisioning", "name": "experimental-provision", "required": {"default": true}, "type": "boolean"}, {"alias": "x-auto-create", "description": "Automatically provision draft bindings with new resources", "name": "experimental-auto-create", "required": {"default": true}, "type": "boolean"}], "x-forge-hidden": true}
```
