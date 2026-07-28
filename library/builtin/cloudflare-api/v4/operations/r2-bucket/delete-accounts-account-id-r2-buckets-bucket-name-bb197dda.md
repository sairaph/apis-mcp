---
title: Delete Bucket
page_id: operation-delete-accounts-account-id-r2-buckets-bucket-name-856fc2da
path: operations/r2-bucket
description: Deletes an existing R2 bucket.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/r2/buckets/{bucket_name}
operation_ids:
    - r2-delete-bucket
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Bucket

`DELETE /accounts/{account_id}/r2/buckets/{bucket_name}`

Operation ID: `r2-delete-bucket`

Deletes an existing R2 bucket.

## Definition

```yaml
{"operationId": "r2-delete-bucket", "summary": "Delete Bucket", "description": "Deletes an existing R2 bucket.", "parameters": [{"name": "bucket_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_bucket_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_account_identifier"}}, {"name": "cf-r2-jurisdiction", "in": "header", "schema": {"$ref": "#/components/schemas/r2_jurisdiction"}}], "responses": {"200": {"description": "Delete Bucket response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response"}}}}, "4XX": {"description": "Delete Bucket response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}}, "security": [{"api_token": []}], "tags": ["R2 Bucket"], "x-api-token-group": ["Workers R2 Storage Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2.buckets", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
