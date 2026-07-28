---
title: Update r2.dev Domain of Bucket
page_id: operation-put-accounts-account-id-r2-buckets-bucket-name-domains-managed-c12ec930
path: operations/r2-bucket
description: Updates state of public access over the bucket's R2-managed (r2.dev) domain.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/r2/buckets/{bucket_name}/domains/managed
operation_ids:
    - r2-put-bucket-public-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update r2.dev Domain of Bucket

`PUT /accounts/{account_id}/r2/buckets/{bucket_name}/domains/managed`

Operation ID: `r2-put-bucket-public-policy`

Updates state of public access over the bucket's R2-managed (r2.dev) domain.

## Definition

```yaml
{"operationId": "r2-put-bucket-public-policy", "summary": "Update r2.dev Domain of Bucket", "description": "Updates state of public access over the bucket's R2-managed (r2.dev) domain.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_account_identifier"}}, {"name": "bucket_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_bucket_name"}}, {"name": "cf-r2-jurisdiction", "in": "header", "schema": {"$ref": "#/components/schemas/r2_jurisdiction"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_edit_managed_domain_request"}}}}, "responses": {"200": {"description": "Update Managed Subdomain response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/r2_v4_response"}, {"properties": {"result": {"$ref": "#/components/schemas/r2_managed_domain_response"}}, "type": "object"}]}}}}, "4XX": {"description": "Update Managed Subdomain response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}}, "security": [{"api_token": []}], "tags": ["R2 Bucket"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.r2.bucket.write"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2.buckets.domains.managed", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
