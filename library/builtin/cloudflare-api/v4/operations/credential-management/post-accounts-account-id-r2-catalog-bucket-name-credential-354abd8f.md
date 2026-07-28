---
title: Store catalog credentials
page_id: operation-post-accounts-account-id-r2-catalog-bucket-name-credential-e1030e11
path: operations/credential-management
description: |-
    Store authentication credentials for a catalog. These credentials are used
    to authenticate with R2 storage when performing catalog operations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/r2-catalog/{bucket_name}/credential
operation_ids:
    - store-credentials
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Store catalog credentials

`POST /accounts/{account_id}/r2-catalog/{bucket_name}/credential`

Operation ID: `store-credentials`

Store authentication credentials for a catalog. These credentials are used
to authenticate with R2 storage when performing catalog operations.

## Definition

```yaml
{"operationId": "store-credentials", "summary": "Store catalog credentials", "description": "Store authentication credentials for a catalog. These credentials are used\nto authenticate with R2 storage when performing catalog operations.\n", "parameters": [{"name": "account_id", "in": "path", "description": "Identifies the account.", "required": true, "schema": {"$ref": "#/components/schemas/r2-data-catalog_account-id"}}, {"name": "bucket_name", "in": "path", "description": "Specifies the R2 bucket name.", "required": true, "schema": {"$ref": "#/components/schemas/r2-data-catalog_bucket-name"}}], "requestBody": {"required": true, "content": {"application/json": {"example": {"token": "your-cloudflare-api-token-here"}, "schema": {"$ref": "#/components/schemas/r2-data-catalog_catalog-credential-request"}}}}, "responses": {"200": {"description": "Credentials stored successfully.", "content": {"application/json": {"example": {"errors": [], "messages": [], "result": null, "success": true}, "schema": {"allOf": [{"$ref": "#/components/schemas/r2-data-catalog_api-response-single"}, {"properties": {"result": {"type": "object", "nullable": true}}, "type": "object"}]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "401": {"description": "Authentication failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "404": {"description": "Catalog not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Credential Management"], "x-api-token-group": ["Workers R2 Data Catalog Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2-data-catalog.credentials", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
