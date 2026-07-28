---
title: Configure Custom Domain Settings
page_id: operation-put-accounts-account-id-r2-buckets-bucket-name-domains-custom-domain-19c40b56
path: operations/r2-bucket
description: Edit the configuration for a custom domain on an existing R2 bucket.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/r2/buckets/{bucket_name}/domains/custom/{domain}
operation_ids:
    - r2-edit-custom-domain-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Configure Custom Domain Settings

`PUT /accounts/{account_id}/r2/buckets/{bucket_name}/domains/custom/{domain}`

Operation ID: `r2-edit-custom-domain-settings`

Edit the configuration for a custom domain on an existing R2 bucket.

## Definition

```yaml
{"operationId": "r2-edit-custom-domain-settings", "summary": "Configure Custom Domain Settings", "description": "Edit the configuration for a custom domain on an existing R2 bucket.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_account_identifier"}}, {"name": "bucket_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_bucket_name"}}, {"name": "domain", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_domain_name"}}, {"name": "cf-r2-jurisdiction", "in": "header", "schema": {"$ref": "#/components/schemas/r2_jurisdiction"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_edit_custom_domain_request"}}}}, "responses": {"200": {"description": "Edit Custom Domain Configuration response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/r2_v4_response"}, {"properties": {"result": {"allOf": [{"$ref": "#/components/schemas/r2_edit_custom_domain_response"}]}}, "type": "object"}]}}}}, "4XX": {"description": "Edit Custom Domain Configuration response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}}, "security": [{"api_token": []}], "tags": ["R2 Bucket"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.r2.bucket.write"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2.buckets.domains.custom", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
