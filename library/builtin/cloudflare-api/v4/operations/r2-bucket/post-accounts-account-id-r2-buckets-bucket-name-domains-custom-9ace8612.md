---
title: Attach Custom Domain To Bucket
page_id: operation-post-accounts-account-id-r2-buckets-bucket-name-domains-custom-ba90b0d7
path: operations/r2-bucket
description: Register a new custom domain for an existing R2 bucket.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/r2/buckets/{bucket_name}/domains/custom
operation_ids:
    - r2-add-custom-domain
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Attach Custom Domain To Bucket

`POST /accounts/{account_id}/r2/buckets/{bucket_name}/domains/custom`

Operation ID: `r2-add-custom-domain`

Register a new custom domain for an existing R2 bucket.

## Definition

```yaml
{"operationId": "r2-add-custom-domain", "summary": "Attach Custom Domain To Bucket", "description": "Register a new custom domain for an existing R2 bucket.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_account_identifier"}}, {"name": "bucket_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_bucket_name"}}, {"name": "cf-r2-jurisdiction", "in": "header", "schema": {"$ref": "#/components/schemas/r2_jurisdiction"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_add_custom_domain_request"}}}}, "responses": {"200": {"description": "Add Custom Domain response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/r2_v4_response"}, {"properties": {"result": {"$ref": "#/components/schemas/r2_add_custom_domain_response"}}, "type": "object"}]}}}}, "4XX": {"description": "Add Custom Domain response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}}, "security": [{"api_token": []}], "tags": ["R2 Bucket"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.r2.bucket.write"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2.buckets.domains.custom", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
