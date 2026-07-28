---
title: Get R2 catalog details
page_id: operation-get-accounts-account-id-r2-catalog-bucket-name-8328504a
path: operations/r2-catalog-management
description: |-
    Retrieve detailed information about a specific R2 catalog by bucket name.
    Returns catalog status, maintenance configuration, and credential status.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/r2-catalog/{bucket_name}
operation_ids:
    - get-catalog-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get R2 catalog details

`GET /accounts/{account_id}/r2-catalog/{bucket_name}`

Operation ID: `get-catalog-details`

Retrieve detailed information about a specific R2 catalog by bucket name.
Returns catalog status, maintenance configuration, and credential status.

## Definition

```yaml
{"operationId": "get-catalog-details", "summary": "Get R2 catalog details", "description": "Retrieve detailed information about a specific R2 catalog by bucket name.\nReturns catalog status, maintenance configuration, and credential status.\n", "parameters": [{"name": "account_id", "in": "path", "description": "Identifies the account.", "required": true, "schema": {"$ref": "#/components/schemas/r2-data-catalog_account-id"}}, {"name": "bucket_name", "in": "path", "description": "Specifies the R2 bucket name.", "required": true, "schema": {"$ref": "#/components/schemas/r2-data-catalog_bucket-name"}}], "responses": {"200": {"description": "R2 catalog details.", "content": {"application/json": {"example": {"errors": [], "messages": [], "result": {"bucket": "analytics-bucket", "credential_status": "present", "id": "550e8400-e29b-41d4-a716-446655440000", "maintenance_config": {"compaction": {"state": "enabled", "target_size_mb": "128"}, "snapshot_expiration": {"max_snapshot_age": "7d", "min_snapshots_to_keep": 100, "state": "enabled"}}, "name": "account123_analytics-bucket", "status": "active"}, "success": true}, "schema": {"allOf": [{"$ref": "#/components/schemas/r2-data-catalog_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/r2-data-catalog_catalog"}}, "type": "object"}]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "401": {"description": "Authentication failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "404": {"description": "Catalog not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["R2 Catalog Management"], "x-api-token-group": ["Workers R2 Data Catalog Write", "Workers R2 Data Catalog Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2-data-catalog", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
