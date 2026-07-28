---
title: Get r2.dev Domain of Bucket
page_id: operation-get-accounts-account-id-r2-buckets-bucket-name-domains-managed-099dd275
path: operations/r2-bucket
description: Gets state of public access over the bucket's R2-managed (r2.dev) domain.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/r2/buckets/{bucket_name}/domains/managed
operation_ids:
    - r2-get-bucket-public-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get r2.dev Domain of Bucket

`GET /accounts/{account_id}/r2/buckets/{bucket_name}/domains/managed`

Operation ID: `r2-get-bucket-public-policy`

Gets state of public access over the bucket's R2-managed (r2.dev) domain.

## Definition

```yaml
{"operationId": "r2-get-bucket-public-policy", "summary": "Get r2.dev Domain of Bucket", "description": "Gets state of public access over the bucket's R2-managed (r2.dev) domain.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_account_identifier"}}, {"name": "bucket_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_bucket_name"}}, {"name": "cf-r2-jurisdiction", "in": "header", "schema": {"$ref": "#/components/schemas/r2_jurisdiction"}}], "responses": {"200": {"description": "Get Managed Subdomain response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/r2_v4_response"}, {"properties": {"result": {"$ref": "#/components/schemas/r2_managed_domain_response"}}, "type": "object"}]}}}}, "4XX": {"description": "Get Managed Subdomain response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}}, "security": [{"api_token": []}], "tags": ["R2 Bucket"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.r2.bucket.read"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2.buckets.domains.managed", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
