---
title: Delete D1 Database
page_id: operation-delete-accounts-account-id-d1-database-database-id-94ea69a9
path: operations/d1
description: Deletes the specified D1 database.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/d1/database/{database_id}
operation_ids:
    - d1-delete-database
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete D1 Database

`DELETE /accounts/{account_id}/d1/database/{database_id}`

Operation ID: `d1-delete-database`

Deletes the specified D1 database.

## Definition

```yaml
{"operationId": "d1-delete-database", "summary": "Delete D1 Database", "description": "Deletes the specified D1 database.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/d1_account-identifier"}}, {"name": "database_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/d1_database-identifier"}}], "responses": {"200": {"description": "Delete D1 database response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/d1_api-response-common"}, {"properties": {"result": {"type": "object", "enum": [null], "nullable": true}}}]}}}}, "4XX": {"description": "Delete D1 database response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/d1_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["D1"], "x-api-token-group": ["D1 Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "d1", "x-fern-sdk-method-name": "delete", "x-forge-epilogue": "This command acts on remote D1 Databases.", "x-forge-globals": [{"alias": "v", "description": "Show version number", "name": "version", "required": {"default": false}, "type": "boolean"}, {"description": "Run as if Wrangler was started in the specified directory instead of the current working directory", "name": "cwd", "required": {"default": "."}, "type": "string"}, {"alias": "c", "description": "Path to Wrangler configuration file", "name": "config", "required": {"default": "./wrangler.toml"}, "type": "string"}, {"alias": "e", "description": "Environment to use for operations, and for selecting .env and .dev.vars files", "name": "env", "required": {"default": "dev"}, "type": "string"}, {"description": "Path to an .env file to load", "name": "env-file", "required": {"default": "./.env"}, "type": "string"}, {"alias": "x-provision", "description": "Experimental: Enable automatic resource provisioning", "name": "experimental-provision", "required": {"default": true}, "type": "boolean"}, {"alias": "x-auto-create", "description": "Automatically provision draft bindings with new resources", "name": "experimental-auto-create", "required": {"default": true}, "type": "boolean"}], "x-forge-hidden": true}
```
