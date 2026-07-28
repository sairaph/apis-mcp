---
title: Get D1 Database
page_id: operation-get-accounts-account-id-d1-database-database-id-a2f83fb2
path: operations/d1
description: Returns the specified D1 database.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/d1/database/{database_id}
operation_ids:
    - d1-get-database
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get D1 Database

`GET /accounts/{account_id}/d1/database/{database_id}`

Operation ID: `d1-get-database`

Returns the specified D1 database.

## Definition

```yaml
{"operationId": "d1-get-database", "summary": "Get D1 Database", "description": "Returns the specified D1 database.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/d1_account-identifier"}}, {"name": "database_id", "in": "path", "required": true, "schema": {"oneOf": [{"$ref": "#/components/schemas/d1_database-identifier"}, {"$ref": "#/components/schemas/d1_database-name"}]}}, {"name": "fields", "in": "query", "description": "Comma-separated list of fields to include in the response. When omitted,\nall fields are returned.\n", "schema": {"type": "array", "items": {"enum": ["uuid", "name", "created_at", "version", "jurisdiction", "num_tables", "file_size", "running_in_region", "read_replication"], "type": "string"}}, "example": "uuid,name,version", "explode": false, "style": "form"}], "responses": {"200": {"description": "Database details response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/d1_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/d1_database-details-response"}}}]}}}}, "4XX": {"description": "Database details response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/d1_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["D1"], "x-api-token-group": ["D1 Read", "D1 Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "d1", "x-fern-sdk-method-name": "get", "x-forge-globals": [{"alias": "v", "description": "Show version number", "name": "version", "required": {"default": false}, "type": "boolean"}, {"description": "Run as if Wrangler was started in the specified directory instead of the current working directory", "name": "cwd", "required": {"default": "."}, "type": "string"}, {"alias": "c", "description": "Path to Wrangler configuration file", "name": "config", "required": {"default": "./wrangler.toml"}, "type": "string"}, {"alias": "e", "description": "Environment to use for operations, and for selecting .env and .dev.vars files", "name": "env", "required": {"default": "dev"}, "type": "string"}, {"description": "Path to an .env file to load", "name": "env-file", "required": {"default": "./.env"}, "type": "string"}, {"alias": "x-provision", "description": "Experimental: Enable automatic resource provisioning", "name": "experimental-provision", "required": {"default": true}, "type": "boolean"}, {"alias": "x-auto-create", "description": "Automatically provision draft bindings with new resources", "name": "experimental-auto-create", "required": {"default": true}, "type": "boolean"}], "x-forge-hidden": true}
```
