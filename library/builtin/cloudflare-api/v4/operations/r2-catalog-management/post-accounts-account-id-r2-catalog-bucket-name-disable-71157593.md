---
title: Disable R2 catalog
page_id: operation-post-accounts-account-id-r2-catalog-bucket-name-disable-9a6010fa
path: operations/r2-catalog-management
description: |-
    Disable an R2 bucket as a catalog. This operation deactivates the catalog
    but preserves existing metadata and data files. The catalog can be
    re-enabled later.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/r2-catalog/{bucket_name}/disable
operation_ids:
    - disable-catalog
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Disable R2 catalog

`POST /accounts/{account_id}/r2-catalog/{bucket_name}/disable`

Operation ID: `disable-catalog`

Disable an R2 bucket as a catalog. This operation deactivates the catalog
but preserves existing metadata and data files. The catalog can be
re-enabled later.

## Definition

```yaml
{"operationId": "disable-catalog", "summary": "Disable R2 catalog", "description": "Disable an R2 bucket as a catalog. This operation deactivates the catalog\nbut preserves existing metadata and data files. The catalog can be\nre-enabled later.\n", "parameters": [{"name": "account_id", "in": "path", "description": "Identifies the account.", "required": true, "schema": {"$ref": "#/components/schemas/r2-data-catalog_account-id"}}, {"name": "bucket_name", "in": "path", "description": "Specifies the R2 bucket name to disable as catalog.", "required": true, "schema": {"$ref": "#/components/schemas/r2-data-catalog_bucket-name"}}], "responses": {"204": {"description": "Catalog disabled successfully."}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "401": {"description": "Authentication failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "404": {"description": "Catalog not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["R2 Catalog Management"], "x-api-token-group": ["Workers R2 Data Catalog Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2-data-catalog", "x-fern-sdk-method-name": "disable", "x-forge-hidden": true, "x-forge-require-confirmation": "This operation turns off the data catalog for the bucket. Existing namespace or table state remains, but writes stop."}
```
