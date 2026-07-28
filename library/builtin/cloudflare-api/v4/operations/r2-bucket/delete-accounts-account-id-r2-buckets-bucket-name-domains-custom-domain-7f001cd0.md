---
title: Remove Custom Domain From Bucket
page_id: operation-delete-accounts-account-id-r2-buckets-bucket-name-domains-custom-domain-2b78b26c
path: operations/r2-bucket
description: Remove custom domain registration from an existing R2 bucket.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/r2/buckets/{bucket_name}/domains/custom/{domain}
operation_ids:
    - r2-delete-custom-domain
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Remove Custom Domain From Bucket

`DELETE /accounts/{account_id}/r2/buckets/{bucket_name}/domains/custom/{domain}`

Operation ID: `r2-delete-custom-domain`

Remove custom domain registration from an existing R2 bucket.

## Definition

```yaml
{"operationId": "r2-delete-custom-domain", "summary": "Remove Custom Domain From Bucket", "description": "Remove custom domain registration from an existing R2 bucket.", "parameters": [{"name": "bucket_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_bucket_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_account_identifier"}}, {"name": "domain", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_domain_name"}}, {"name": "cf-r2-jurisdiction", "in": "header", "schema": {"$ref": "#/components/schemas/r2_jurisdiction"}}], "responses": {"200": {"description": "Delete Custom Domain response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/r2_v4_response"}, {"properties": {"result": {"$ref": "#/components/schemas/r2_remove_custom_domain_response"}}, "type": "object"}]}}}}, "4XX": {"description": "Delete Custom Domain response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}}, "security": [{"api_token": []}], "tags": ["R2 Bucket"], "x-api-token-group": ["Workers R2 Storage Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2.buckets.domains.custom", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
