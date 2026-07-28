---
title: Enable R2 bucket as a catalog
page_id: operation-post-accounts-account-id-r2-catalog-bucket-name-enable-a9687afd
path: operations/r2-catalog-management
description: |-
    Enable an R2 bucket as an Apache Iceberg catalog. This operation creates
    the necessary catalog infrastructure and activates the bucket for storing
    Iceberg metadata and data files.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/r2-catalog/{bucket_name}/enable
operation_ids:
    - enable-catalog
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Enable R2 bucket as a catalog

`POST /accounts/{account_id}/r2-catalog/{bucket_name}/enable`

Operation ID: `enable-catalog`

Enable an R2 bucket as an Apache Iceberg catalog. This operation creates
the necessary catalog infrastructure and activates the bucket for storing
Iceberg metadata and data files.

## Definition

```yaml
{"operationId": "enable-catalog", "summary": "Enable R2 bucket as a catalog", "description": "Enable an R2 bucket as an Apache Iceberg catalog. This operation creates\nthe necessary catalog infrastructure and activates the bucket for storing\nIceberg metadata and data files.\n", "parameters": [{"name": "account_id", "in": "path", "description": "Identifies the account.", "required": true, "schema": {"$ref": "#/components/schemas/r2-data-catalog_account-id"}}, {"name": "bucket_name", "in": "path", "description": "Specifies the R2 bucket name to enable as catalog.", "required": true, "schema": {"$ref": "#/components/schemas/r2-data-catalog_bucket-name"}}], "responses": {"200": {"description": "Catalog enabled successfully.", "content": {"application/json": {"example": {"errors": [], "messages": [], "result": {"id": "550e8400-e29b-41d4-a716-446655440000", "name": "account123_my-bucket"}, "success": true}, "schema": {"allOf": [{"$ref": "#/components/schemas/r2-data-catalog_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/r2-data-catalog_catalog-activation-response"}}, "type": "object"}]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "401": {"description": "Authentication failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "404": {"description": "R2 bucket not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "409": {"description": "Catalog already enabled.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["R2 Catalog Management"], "x-api-token-group": ["Workers R2 Data Catalog Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2-data-catalog", "x-fern-sdk-method-name": "enable", "x-forge-hidden": true}
```
