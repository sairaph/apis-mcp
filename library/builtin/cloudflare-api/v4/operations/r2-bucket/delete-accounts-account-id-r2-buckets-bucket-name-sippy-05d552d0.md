---
title: Disable Sippy
page_id: operation-delete-accounts-account-id-r2-buckets-bucket-name-sippy-b0c1a599
path: operations/r2-bucket
description: Disables Sippy on this bucket.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/r2/buckets/{bucket_name}/sippy
operation_ids:
    - r2-delete-bucket-sippy-config
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Disable Sippy

`DELETE /accounts/{account_id}/r2/buckets/{bucket_name}/sippy`

Operation ID: `r2-delete-bucket-sippy-config`

Disables Sippy on this bucket.

## Definition

```yaml
{"operationId": "r2-delete-bucket-sippy-config", "summary": "Disable Sippy", "description": "Disables Sippy on this bucket.", "parameters": [{"name": "bucket_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_bucket_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_account_identifier"}}, {"name": "cf-r2-jurisdiction", "in": "header", "schema": {"$ref": "#/components/schemas/r2_jurisdiction"}}], "responses": {"200": {"description": "Delete Sippy Configuration response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/r2_v4_response"}, {"properties": {"result": {"type": "object", "properties": {"enabled": {"type": "boolean", "enum": [false]}}}}, "type": "object"}]}}}}, "4XX": {"description": "Delete Sippy Configuration response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}}, "security": [{"api_token": []}], "tags": ["R2 Bucket"], "x-api-token-group": ["Workers R2 Storage Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2.buckets.sippy", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
