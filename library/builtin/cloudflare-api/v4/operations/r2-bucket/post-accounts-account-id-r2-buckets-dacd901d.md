---
title: Create Bucket
page_id: operation-post-accounts-account-id-r2-buckets-d64b0ef5
path: operations/r2-bucket
description: Creates a new R2 bucket.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/r2/buckets
operation_ids:
    - r2-create-bucket
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Bucket

`POST /accounts/{account_id}/r2/buckets`

Operation ID: `r2-create-bucket`

Creates a new R2 bucket.

## Definition

```yaml
{"operationId": "r2-create-bucket", "summary": "Create Bucket", "description": "Creates a new R2 bucket.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_account_identifier"}}, {"name": "cf-r2-jurisdiction", "in": "header", "schema": {"$ref": "#/components/schemas/r2_jurisdiction"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"locationHint": {"$ref": "#/components/schemas/r2_bucket_location"}, "name": {"$ref": "#/components/schemas/r2_bucket_name"}, "storageClass": {"$ref": "#/components/schemas/r2_storage_class"}}, "example": "{\"name\": \"example-bucket\"}", "required": ["name"]}}}}, "responses": {"200": {"description": "Create Bucket response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/r2_v4_response"}, {"properties": {"result": {"$ref": "#/components/schemas/r2_bucket"}}, "type": "object"}]}}}}, "4XX": {"description": "Create Bucket response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}}, "security": [{"api_token": []}], "tags": ["R2 Bucket"], "x-api-token-group": ["Workers R2 Storage Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2.buckets", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
