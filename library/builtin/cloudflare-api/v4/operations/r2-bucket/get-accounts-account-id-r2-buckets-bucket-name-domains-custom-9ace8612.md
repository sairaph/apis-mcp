---
title: List Custom Domains of Bucket
page_id: operation-get-accounts-account-id-r2-buckets-bucket-name-domains-custom-dccb961e
path: operations/r2-bucket
description: Gets a list of all custom domains registered with an existing R2 bucket.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/r2/buckets/{bucket_name}/domains/custom
operation_ids:
    - r2-list-custom-domains
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Custom Domains of Bucket

`GET /accounts/{account_id}/r2/buckets/{bucket_name}/domains/custom`

Operation ID: `r2-list-custom-domains`

Gets a list of all custom domains registered with an existing R2 bucket.

## Definition

```yaml
{"operationId": "r2-list-custom-domains", "summary": "List Custom Domains of Bucket", "description": "Gets a list of all custom domains registered with an existing R2 bucket.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_account_identifier"}}, {"name": "bucket_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_bucket_name"}}, {"name": "cf-r2-jurisdiction", "in": "header", "schema": {"$ref": "#/components/schemas/r2_jurisdiction"}}], "responses": {"200": {"description": "List Custom Domains response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/r2_v4_response"}, {"properties": {"result": {"$ref": "#/components/schemas/r2_list_custom_domains_response"}}, "type": "object"}]}}}}, "4XX": {"description": "List Custom Domains response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}}, "security": [{"api_token": []}], "tags": ["R2 Bucket"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.r2.bucket.read"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2.buckets.domains.custom", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
