---
title: Restore D1 Database to a bookmark or point in time
page_id: operation-post-accounts-account-id-d1-database-database-id-time-travel-restore-fdc1d768
path: operations/d1
description: Restores a D1 database to a previous point in time either via a bookmark or a timestamp.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/d1/database/{database_id}/time_travel/restore
operation_ids:
    - d1-time-travel-restore
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Restore D1 Database to a bookmark or point in time

`POST /accounts/{account_id}/d1/database/{database_id}/time_travel/restore`

Operation ID: `d1-time-travel-restore`

Restores a D1 database to a previous point in time either via a bookmark or a timestamp.

## Definition

```yaml
{"operationId": "d1-time-travel-restore", "summary": "Restore D1 Database to a bookmark or point in time", "description": "Restores a D1 database to a previous point in time either via a bookmark or a timestamp.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/d1_account-identifier"}}, {"name": "database_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/d1_database-identifier"}}, {"name": "bookmark", "in": "query", "description": "A bookmark to restore the database to. Required if `timestamp` is not provided.", "schema": {"$ref": "#/components/schemas/d1_time-travel-bookmark"}}, {"name": "timestamp", "in": "query", "description": "An ISO 8601 timestamp to restore the database to. Required if `bookmark` is not provided.", "schema": {"$ref": "#/components/schemas/d1_time-travel-timestamp"}}], "responses": {"200": {"description": "Database restored successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/d1_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/d1_time-travel-restore-response"}}}]}}}}, "4XX": {"description": "Restore operation failed", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/d1_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["D1"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "d1.time-travel", "x-fern-sdk-method-name": "restore", "x-forge-epilogue": "This command acts on remote D1 Databases.", "x-forge-globals": [{"alias": "v", "description": "Show version number", "name": "version", "required": {"default": false}, "type": "boolean"}, {"description": "Run as if Wrangler was started in the specified directory instead of the current working directory", "name": "cwd", "required": {"default": "."}, "type": "string"}, {"alias": "c", "description": "Path to Wrangler configuration file", "name": "config", "required": {"default": "./wrangler.toml"}, "type": "string"}, {"alias": "e", "description": "Environment to use for operations, and for selecting .env and .dev.vars files", "name": "env", "required": {"default": "dev"}, "type": "string"}, {"description": "Path to an .env file to load", "name": "env-file", "required": {"default": "./.env"}, "type": "string"}, {"alias": "x-provision", "description": "Experimental: Enable automatic resource provisioning", "name": "experimental-provision", "required": {"default": true}, "type": "boolean"}, {"alias": "x-auto-create", "description": "Automatically provision draft bindings with new resources", "name": "experimental-auto-create", "required": {"default": true}, "type": "boolean"}], "x-forge-hidden": true}
```
