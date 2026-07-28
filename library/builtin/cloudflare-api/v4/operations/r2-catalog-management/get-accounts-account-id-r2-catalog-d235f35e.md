---
title: List R2 catalogs
page_id: operation-get-accounts-account-id-r2-catalog-f83c8ab8
path: operations/r2-catalog-management
description: |-
    Returns a list of R2 buckets that have been enabled as Apache Iceberg catalogs
    for the specified account. Each catalog represents an R2 bucket configured
    to store Iceberg metadata and data files.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/r2-catalog
operation_ids:
    - list-catalogs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List R2 catalogs

`GET /accounts/{account_id}/r2-catalog`

Operation ID: `list-catalogs`

Returns a list of R2 buckets that have been enabled as Apache Iceberg catalogs
for the specified account. Each catalog represents an R2 bucket configured
to store Iceberg metadata and data files.

## Definition

```yaml
{"operationId": "list-catalogs", "summary": "List R2 catalogs", "description": "Returns a list of R2 buckets that have been enabled as Apache Iceberg catalogs\nfor the specified account. Each catalog represents an R2 bucket configured\nto store Iceberg metadata and data files.\n", "parameters": [{"name": "account_id", "in": "path", "description": "Identifies the account.", "required": true, "schema": {"$ref": "#/components/schemas/r2-data-catalog_account-id"}}], "responses": {"200": {"description": "List of R2 catalogs.", "content": {"application/json": {"example": {"errors": [], "messages": [], "result": {"warehouses": [{"bucket": "analytics-bucket", "id": "550e8400-e29b-41d4-a716-446655440000", "maintenance_config": {"compaction": {"state": "enabled", "target_size_mb": "128"}, "snapshot_expiration": {"max_snapshot_age": "7d", "min_snapshots_to_keep": 100, "state": "enabled"}}, "name": "account123_analytics-bucket", "status": "active"}, {"bucket": "logs-bucket", "id": "6ba7b810-9dad-11d1-80b4-00c04fd430c8", "maintenance_config": {"compaction": {"state": "disabled", "target_size_mb": "128"}, "snapshot_expiration": {"max_snapshot_age": "7d", "min_snapshots_to_keep": 100, "state": "disabled"}}, "name": "account123_logs-bucket", "status": "inactive"}]}, "success": true}, "schema": {"allOf": [{"$ref": "#/components/schemas/r2-data-catalog_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/r2-data-catalog_catalog-list"}}, "type": "object"}]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "401": {"description": "Authentication failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["R2 Catalog Management"], "x-api-token-group": ["Workers R2 Data Catalog Write", "Workers R2 Data Catalog Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2-data-catalog", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
