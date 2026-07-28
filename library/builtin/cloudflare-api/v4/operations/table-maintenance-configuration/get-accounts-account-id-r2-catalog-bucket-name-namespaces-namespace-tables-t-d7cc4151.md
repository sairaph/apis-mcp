---
title: Get table maintenance configuration
page_id: operation-get-accounts-account-id-r2-catalog-bucket-name-namespaces-namespace-tabl-9d046a6c
path: operations/table-maintenance-configuration
description: |-
    Retrieve the maintenance configuration for a specific table,
    including compaction settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/r2-catalog/{bucket_name}/namespaces/{namespace}/tables/{table_name}/maintenance-configs
operation_ids:
    - get-table-maintenance-config
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get table maintenance configuration

`GET /accounts/{account_id}/r2-catalog/{bucket_name}/namespaces/{namespace}/tables/{table_name}/maintenance-configs`

Operation ID: `get-table-maintenance-config`

Retrieve the maintenance configuration for a specific table,
including compaction settings.

## Definition

```yaml
{"operationId": "get-table-maintenance-config", "summary": "Get table maintenance configuration", "description": "Retrieve the maintenance configuration for a specific table,\nincluding compaction settings.\n", "parameters": [{"name": "account_id", "in": "path", "description": "Identifies the account.", "required": true, "schema": {"$ref": "#/components/schemas/r2-data-catalog_account-id"}}, {"name": "bucket_name", "in": "path", "description": "Specifies the R2 bucket name.", "required": true, "schema": {"$ref": "#/components/schemas/r2-data-catalog_bucket-name"}}, {"name": "namespace", "in": "path", "description": "The namespace identifier (use %1F as separator for nested namespaces).", "required": true, "schema": {"type": "string"}, "example": "my_namespace%1Fsub_namespace"}, {"name": "table_name", "in": "path", "description": "The table name.", "required": true, "schema": {"type": "string"}, "example": "my_table"}], "responses": {"200": {"description": "Table maintenance configuration retrieved successfully.", "content": {"application/json": {"example": {"errors": [], "messages": [], "result": {"maintenance_config": {"compaction": {"state": "enabled", "target_size_mb": "128"}, "snapshot_expiration": {"max_snapshot_age": "7d", "min_snapshots_to_keep": 100, "state": "enabled"}}}, "success": true}, "schema": {"allOf": [{"$ref": "#/components/schemas/r2-data-catalog_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/r2-data-catalog_table-maintenance-config-response"}}, "type": "object"}]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "401": {"description": "Authentication failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "404": {"description": "Table not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Table Maintenance Configuration"], "x-api-token-group": ["Workers R2 Data Catalog Write", "Workers R2 Data Catalog Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2-data-catalog.namespaces.tables.maintenance-configs", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
