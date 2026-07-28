---
title: Put Object Lifecycle Rules
page_id: operation-put-accounts-account-id-r2-buckets-bucket-name-lifecycle-26a63ba2
path: operations/r2-bucket
description: Set the object lifecycle rules for a bucket.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/r2/buckets/{bucket_name}/lifecycle
operation_ids:
    - r2-put-bucket-lifecycle-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Put Object Lifecycle Rules

`PUT /accounts/{account_id}/r2/buckets/{bucket_name}/lifecycle`

Operation ID: `r2-put-bucket-lifecycle-configuration`

Set the object lifecycle rules for a bucket.

## Definition

```yaml
{"operationId": "r2-put-bucket-lifecycle-configuration", "summary": "Put Object Lifecycle Rules", "description": "Set the object lifecycle rules for a bucket.", "parameters": [{"name": "bucket_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_bucket_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_account_identifier"}}, {"name": "cf-r2-jurisdiction", "in": "header", "schema": {"$ref": "#/components/schemas/r2_jurisdiction"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"rules": {"type": "array", "items": {"$ref": "#/components/schemas/r2_lifecycle-rule"}}}}}}}, "responses": {"200": {"description": "Success Response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/r2_v4_response"}, {"type": "object"}]}}}}, "4XX": {"description": "Error Response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}}, "security": [{"api_token": []}], "tags": ["R2 Bucket"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2.buckets.lifecycle", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
