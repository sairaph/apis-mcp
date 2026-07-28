---
title: Update table maintenance configuration
page_id: operation-post-accounts-account-id-r2-catalog-bucket-name-namespaces-namespace-tab-fc018c05
path: operations/table-maintenance-configuration
description: |-
    Update the maintenance configuration for a specific table. This allows you to
    enable or disable compaction and adjust target file sizes for optimization.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/r2-catalog/{bucket_name}/namespaces/{namespace}/tables/{table_name}/maintenance-configs
operation_ids:
    - update-table-maintenance-config
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update table maintenance configuration

`POST /accounts/{account_id}/r2-catalog/{bucket_name}/namespaces/{namespace}/tables/{table_name}/maintenance-configs`

Operation ID: `update-table-maintenance-config`

Update the maintenance configuration for a specific table. This allows you to
enable or disable compaction and adjust target file sizes for optimization.

## Definition

```yaml
{"operationId": "update-table-maintenance-config", "summary": "Update table maintenance configuration", "description": "Update the maintenance configuration for a specific table. This allows you to\nenable or disable compaction and adjust target file sizes for optimization.\n", "parameters": [{"name": "account_id", "in": "path", "description": "Identifies the account.", "required": true, "schema": {"$ref": "#/components/schemas/r2-data-catalog_account-id"}}, {"name": "bucket_name", "in": "path", "description": "Specifies the R2 bucket name.", "required": true, "schema": {"$ref": "#/components/schemas/r2-data-catalog_bucket-name"}}, {"name": "namespace", "in": "path", "description": "The namespace identifier (use %1F as separator for nested namespaces).", "required": true, "schema": {"type": "string"}, "example": "my_namespace%1Fsub_namespace"}, {"name": "table_name", "in": "path", "description": "The table name.", "required": true, "schema": {"type": "string"}, "example": "my_table"}], "requestBody": {"required": true, "content": {"application/json": {"example": {"compaction": {"state": "enabled", "target_size_mb": "256"}, "snapshot_expiration": {"max_snapshot_age": "14d", "min_snapshots_to_keep": 5, "state": "enabled"}}, "schema": {"$ref": "#/components/schemas/r2-data-catalog_table-maintenance-update-request"}}}}, "responses": {"200": {"description": "Table maintenance configuration updated successfully.", "content": {"application/json": {"example": {"errors": [], "messages": [], "result": {"compaction": {"state": "enabled", "target_size_mb": "256"}, "snapshot_expiration": {"max_snapshot_age": "14d", "min_snapshots_to_keep": 5, "state": "enabled"}}, "success": true}, "schema": {"allOf": [{"$ref": "#/components/schemas/r2-data-catalog_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/r2-data-catalog_table-maintenance-config"}}, "type": "object"}]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "401": {"description": "Authentication failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "404": {"description": "Table not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Table Maintenance Configuration"], "x-api-token-group": ["Workers R2 Data Catalog Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2-data-catalog.namespaces.tables.maintenance-configs", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
