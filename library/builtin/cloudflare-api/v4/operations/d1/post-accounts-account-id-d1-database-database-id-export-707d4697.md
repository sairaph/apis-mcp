---
title: Export D1 Database as SQL
page_id: operation-post-accounts-account-id-d1-database-database-id-export-e24219aa
path: operations/d1
description: |-
    Returns a URL where the SQL contents of your D1 can be downloaded. Note: this process may take
    some time for larger DBs, during which your D1 will be unavailable to serve queries. To avoid
    blocking your DB unnecessarily, an in-progress export must be continually polled or will automatically cancel.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/d1/database/{database_id}/export
operation_ids:
    - d1-export-database
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Export D1 Database as SQL

`POST /accounts/{account_id}/d1/database/{database_id}/export`

Operation ID: `d1-export-database`

Returns a URL where the SQL contents of your D1 can be downloaded. Note: this process may take
some time for larger DBs, during which your D1 will be unavailable to serve queries. To avoid
blocking your DB unnecessarily, an in-progress export must be continually polled or will automatically cancel.

## Definition

```yaml
{"operationId": "d1-export-database", "summary": "Export D1 Database as SQL", "description": "Returns a URL where the SQL contents of your D1 can be downloaded. Note: this process may take\nsome time for larger DBs, during which your D1 will be unavailable to serve queries. To avoid\nblocking your DB unnecessarily, an in-progress export must be continually polled or will automatically cancel.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/d1_account-identifier"}}, {"name": "database_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/d1_database-identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"current_bookmark": {"description": "To poll an in-progress export, provide the current bookmark (returned by your first polling response)", "type": "string", "x-auditable": true}, "dump_options": {"type": "object", "properties": {"no_data": {"description": "Export only the table definitions, not their contents", "type": "boolean", "x-auditable": true}, "no_schema": {"description": "Export only each table's contents, not its definition", "type": "boolean", "x-auditable": true}, "tables": {"description": "Filter the export to just one or more tables. Passing an empty array is the same as not passing anything and means: export all tables.", "type": "array", "items": {"type": "string", "x-auditable": true}}}}, "output_format": {"description": "Specifies that you will poll this endpoint until the export completes", "type": "string", "enum": ["polling"], "x-auditable": true}}, "required": ["output_format"]}}}}, "responses": {"200": {"description": "Polled successfully, task no longer running (errored or complete)", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/d1_api-response-common"}, {"properties": {"result": {"type": "object", "properties": {"at_bookmark": {"description": "The current time-travel bookmark for your D1, used to poll for updates. Will not change for the duration of the export task.", "type": "string", "x-auditable": true}, "error": {"description": "Only present when status = 'error'. Contains the error message.", "type": "string", "x-auditable": true}, "messages": {"description": "Logs since the last time you polled", "type": "array", "items": {"type": "string", "x-auditable": true}}, "result": {"description": "Only present when status = 'complete'", "type": "object", "properties": {"filename": {"description": "The generated SQL filename.", "type": "string", "x-auditable": true}, "signed_url": {"description": "The URL to download the exported SQL. Available for one hour.", "type": "string"}}}, "status": {"type": "string", "enum": ["complete", "error"], "x-auditable": true}, "success": {"type": "boolean", "x-auditable": true}, "type": {"type": "string", "enum": ["export"], "x-auditable": true}}}}}]}}}}, "202": {"description": "Polled successfully, task is currently running", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/d1_api-response-common"}, {"properties": {"result": {"type": "object", "properties": {"at_bookmark": {"description": "The current time-travel bookmark for your D1, used to poll for updates. Will not change for the duration of the export task.", "type": "string", "x-auditable": true}, "messages": {"description": "Logs since the last time you polled", "type": "array", "items": {"type": "string", "x-auditable": true}}, "status": {"type": "string", "enum": ["active"], "x-auditable": true}, "success": {"type": "boolean", "x-auditable": true}, "type": {"type": "string", "enum": ["export"], "x-auditable": true}}}}}]}}}}, "4XX": {"description": "Poll failed (API error)", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/d1_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["D1"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "d1", "x-fern-sdk-method-name": "export", "x-forge-globals": [{"alias": "v", "description": "Show version number", "name": "version", "required": {"default": false}, "type": "boolean"}, {"description": "Run as if Wrangler was started in the specified directory instead of the current working directory", "name": "cwd", "required": {"default": "."}, "type": "string"}, {"alias": "c", "description": "Path to Wrangler configuration file", "name": "config", "required": {"default": "./wrangler.toml"}, "type": "string"}, {"alias": "e", "description": "Environment to use for operations, and for selecting .env and .dev.vars files", "name": "env", "required": {"default": "dev"}, "type": "string"}, {"description": "Path to an .env file to load", "name": "env-file", "required": {"default": "./.env"}, "type": "string"}, {"alias": "x-provision", "description": "Experimental: Enable automatic resource provisioning", "name": "experimental-provision", "required": {"default": true}, "type": "boolean"}, {"alias": "x-auto-create", "description": "Automatically provision draft bindings with new resources", "name": "experimental-auto-create", "required": {"default": true}, "type": "boolean"}], "x-forge-hidden": true}
```
