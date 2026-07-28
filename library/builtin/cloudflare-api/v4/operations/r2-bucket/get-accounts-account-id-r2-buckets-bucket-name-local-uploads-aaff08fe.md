---
title: Get Local Uploads Configuration
page_id: operation-get-accounts-account-id-r2-buckets-bucket-name-local-uploads-88c17f73
path: operations/r2-bucket
description: Get the local uploads configuration for a bucket. When enabled, object's data is written to the nearest region first, then asynchronously replicated to the bucket's primary region.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/r2/buckets/{bucket_name}/local-uploads
operation_ids:
    - r2-get-bucket-local-uploads-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Local Uploads Configuration

`GET /accounts/{account_id}/r2/buckets/{bucket_name}/local-uploads`

Operation ID: `r2-get-bucket-local-uploads-configuration`

Get the local uploads configuration for a bucket. When enabled, object's data is written to the nearest region first, then asynchronously replicated to the bucket's primary region.

## Definition

```yaml
{"operationId": "r2-get-bucket-local-uploads-configuration", "summary": "Get Local Uploads Configuration", "description": "Get the local uploads configuration for a bucket. When enabled, object's data is written to the nearest region first, then asynchronously replicated to the bucket's primary region.", "parameters": [{"name": "bucket_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_bucket_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_account_identifier"}}], "responses": {"200": {"description": "Success Response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/r2_v4_response"}, {"properties": {"result": {"$ref": "#/components/schemas/r2_local_uploads_configuration"}}, "type": "object"}]}}}}, "4XX": {"description": "Error Response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}}, "security": [{"api_token": []}], "tags": ["R2 Bucket"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.r2.bucket.read"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2.buckets.local-uploads", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
