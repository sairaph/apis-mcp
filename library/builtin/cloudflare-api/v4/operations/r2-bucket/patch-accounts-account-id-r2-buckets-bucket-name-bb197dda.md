---
title: Patch Bucket
page_id: operation-patch-accounts-account-id-r2-buckets-bucket-name-c592413d
path: operations/r2-bucket
description: Updates properties of an existing R2 bucket.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/r2/buckets/{bucket_name}
operation_ids:
    - r2-patch-bucket
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch Bucket

`PATCH /accounts/{account_id}/r2/buckets/{bucket_name}`

Operation ID: `r2-patch-bucket`

Updates properties of an existing R2 bucket.

## Definition

```yaml
{"operationId": "r2-patch-bucket", "summary": "Patch Bucket", "description": "Updates properties of an existing R2 bucket.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_account_identifier"}}, {"name": "bucket_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_bucket_name"}}, {"name": "cf-r2-jurisdiction", "in": "header", "schema": {"$ref": "#/components/schemas/r2_jurisdiction"}}, {"name": "cf-r2-storage-class", "in": "header", "required": true, "schema": {"allOf": [{"$ref": "#/components/schemas/r2_storage_class"}], "x-stainless-param": "storage_class"}}], "responses": {"200": {"description": "Patch Bucket response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/r2_v4_response"}, {"properties": {"result": {"$ref": "#/components/schemas/r2_bucket"}}, "type": "object"}]}}}}, "4XX": {"description": "Get Bucket response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}}, "security": [{"api_token": []}], "tags": ["R2 Bucket"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.r2.bucket.write"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2.buckets", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
