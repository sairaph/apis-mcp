---
title: Put Bucket CORS Policy
page_id: operation-put-accounts-account-id-r2-buckets-bucket-name-cors-03fd60e8
path: operations/r2-bucket
description: Set the CORS policy for a bucket.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/r2/buckets/{bucket_name}/cors
operation_ids:
    - r2-put-bucket-cors-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Put Bucket CORS Policy

`PUT /accounts/{account_id}/r2/buckets/{bucket_name}/cors`

Operation ID: `r2-put-bucket-cors-policy`

Set the CORS policy for a bucket.

## Definition

```yaml
{"operationId": "r2-put-bucket-cors-policy", "summary": "Put Bucket CORS Policy", "description": "Set the CORS policy for a bucket.", "parameters": [{"name": "bucket_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_bucket_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_account_identifier"}}, {"name": "cf-r2-jurisdiction", "in": "header", "schema": {"$ref": "#/components/schemas/r2_jurisdiction"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"rules": {"type": "array", "items": {"$ref": "#/components/schemas/r2_cors-rule"}}}}}}}, "responses": {"200": {"description": "Success Response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/r2_v4_response"}, {"type": "object"}]}}}}, "4XX": {"description": "Error Response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}}, "security": [{"api_token": []}], "tags": ["R2 Bucket"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2.buckets.cors", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
