---
title: Put Local Uploads Configuration
page_id: operation-put-accounts-account-id-r2-buckets-bucket-name-local-uploads-2ae41ce1
path: operations/r2-bucket
description: Set the local uploads configuration for a bucket. When enabled, object's data is written to the nearest region first, then asynchronously replicated to the bucket's primary region.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/r2/buckets/{bucket_name}/local-uploads
operation_ids:
    - r2-put-bucket-local-uploads-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Put Local Uploads Configuration

`PUT /accounts/{account_id}/r2/buckets/{bucket_name}/local-uploads`

Operation ID: `r2-put-bucket-local-uploads-configuration`

Set the local uploads configuration for a bucket. When enabled, object's data is written to the nearest region first, then asynchronously replicated to the bucket's primary region.

## Definition

```yaml
{"operationId": "r2-put-bucket-local-uploads-configuration", "summary": "Put Local Uploads Configuration", "description": "Set the local uploads configuration for a bucket. When enabled, object's data is written to the nearest region first, then asynchronously replicated to the bucket's primary region.", "parameters": [{"name": "bucket_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_bucket_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"enabled": {"description": "Whether to enable local uploads for this bucket.", "type": "boolean", "x-auditable": true}}, "required": ["enabled"]}}}}, "responses": {"200": {"description": "Success Response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/r2_v4_response"}, {"type": "object"}]}}}}, "4XX": {"description": "Error Response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}}, "security": [{"api_token": []}], "tags": ["R2 Bucket"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.r2.bucket.write"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2.buckets.local-uploads", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
