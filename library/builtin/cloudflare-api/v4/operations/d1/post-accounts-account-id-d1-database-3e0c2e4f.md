---
title: Create D1 Database
page_id: operation-post-accounts-account-id-d1-database-6702908b
path: operations/d1
description: Returns the created D1 database.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/d1/database
operation_ids:
    - d1-create-database
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create D1 Database

`POST /accounts/{account_id}/d1/database`

Operation ID: `d1-create-database`

Returns the created D1 database.

## Definition

```yaml
{"operationId": "d1-create-database", "summary": "Create D1 Database", "description": "Returns the created D1 database.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/d1_account-identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"jurisdiction": {"$ref": "#/components/schemas/d1_jurisdiction"}, "name": {"$ref": "#/components/schemas/d1_database-name"}, "primary_location_hint": {"$ref": "#/components/schemas/d1_primary-location-hint"}, "read_replication": {"$ref": "#/components/schemas/d1_read-replication-details-for-request"}}, "required": ["name"]}}}}, "responses": {"200": {"description": "Returns the created D1 database's metadata", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/d1_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/d1_database-details-response"}}}]}}}}, "4XX": {"description": "Database details response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/d1_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["D1"], "x-api-token-group": ["D1 Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "d1", "x-fern-sdk-method-name": "create", "x-forge-epilogue": "This command acts on remote D1 Databases.", "x-forge-globals": [{"alias": "v", "description": "Show version number", "name": "version", "required": {"default": false}, "type": "boolean"}, {"description": "Run as if Wrangler was started in the specified directory instead of the current working directory", "name": "cwd", "required": {"default": "."}, "type": "string"}, {"alias": "c", "description": "Path to Wrangler configuration file", "name": "config", "required": {"default": "./wrangler.toml"}, "type": "string"}, {"alias": "e", "description": "Environment to use for operations, and for selecting .env and .dev.vars files", "name": "env", "required": {"default": "dev"}, "type": "string"}, {"description": "Path to an .env file to load", "name": "env-file", "required": {"default": "./.env"}, "type": "string"}, {"alias": "x-provision", "description": "Experimental: Enable automatic resource provisioning", "name": "experimental-provision", "required": {"default": true}, "type": "boolean"}, {"alias": "x-auto-create", "description": "Automatically provision draft bindings with new resources", "name": "experimental-auto-create", "required": {"default": true}, "type": "boolean"}], "x-forge-hidden": true}
```
