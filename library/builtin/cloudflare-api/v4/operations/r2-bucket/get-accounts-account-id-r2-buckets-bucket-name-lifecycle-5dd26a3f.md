---
title: Get Object Lifecycle Rules
page_id: operation-get-accounts-account-id-r2-buckets-bucket-name-lifecycle-f7454045
path: operations/r2-bucket
description: Get object lifecycle rules for a bucket.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/r2/buckets/{bucket_name}/lifecycle
operation_ids:
    - r2-get-bucket-lifecycle-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Object Lifecycle Rules

`GET /accounts/{account_id}/r2/buckets/{bucket_name}/lifecycle`

Operation ID: `r2-get-bucket-lifecycle-configuration`

Get object lifecycle rules for a bucket.

## Definition

```yaml
{"operationId": "r2-get-bucket-lifecycle-configuration", "summary": "Get Object Lifecycle Rules", "description": "Get object lifecycle rules for a bucket.", "parameters": [{"name": "bucket_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_bucket_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_account_identifier"}}, {"name": "cf-r2-jurisdiction", "in": "header", "schema": {"$ref": "#/components/schemas/r2_jurisdiction"}}], "responses": {"200": {"description": "Success Response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/r2_v4_response"}, {"properties": {"result": {"$ref": "#/components/schemas/r2_lifecycle-config"}}, "type": "object"}]}}}}, "4XX": {"description": "Error Response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}}, "security": [{"api_token": []}], "tags": ["R2 Bucket"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2.buckets.lifecycle", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
