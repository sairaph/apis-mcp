---
title: Get catalog maintenance configuration
page_id: operation-get-accounts-account-id-r2-catalog-bucket-name-maintenance-configs-cb3ae488
path: operations/maintenance-configuration
description: |-
    Retrieve the maintenance configuration for a specific catalog,
    including compaction settings and credential status.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/r2-catalog/{bucket_name}/maintenance-configs
operation_ids:
    - get-maintenance-config
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get catalog maintenance configuration

`GET /accounts/{account_id}/r2-catalog/{bucket_name}/maintenance-configs`

Operation ID: `get-maintenance-config`

Retrieve the maintenance configuration for a specific catalog,
including compaction settings and credential status.

## Definition

```yaml
{"operationId": "get-maintenance-config", "summary": "Get catalog maintenance configuration", "description": "Retrieve the maintenance configuration for a specific catalog,\nincluding compaction settings and credential status.\n", "parameters": [{"name": "account_id", "in": "path", "description": "Identifies the account.", "required": true, "schema": {"$ref": "#/components/schemas/r2-data-catalog_account-id"}}, {"name": "bucket_name", "in": "path", "description": "Specifies the R2 bucket name.", "required": true, "schema": {"$ref": "#/components/schemas/r2-data-catalog_bucket-name"}}], "responses": {"200": {"description": "Maintenance configuration retrieved successfully.", "content": {"application/json": {"example": {"errors": [], "messages": [], "result": {"credential_status": "present", "maintenance_config": {"compaction": {"state": "enabled", "target_size_mb": "128"}, "snapshot_expiration": {"max_snapshot_age": "7d", "min_snapshots_to_keep": 100, "state": "enabled"}}}, "success": true}, "schema": {"allOf": [{"$ref": "#/components/schemas/r2-data-catalog_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/r2-data-catalog_catalog-maintenance-config-response"}}, "type": "object"}]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "401": {"description": "Authentication failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "404": {"description": "Catalog not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Maintenance Configuration"], "x-api-token-group": ["Workers R2 Data Catalog Write", "Workers R2 Data Catalog Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2-data-catalog.maintenance-configs", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
