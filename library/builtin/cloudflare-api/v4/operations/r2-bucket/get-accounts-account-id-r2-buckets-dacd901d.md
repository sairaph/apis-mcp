---
title: List Buckets
page_id: operation-get-accounts-account-id-r2-buckets-8ab54323
path: operations/r2-bucket
description: Lists all R2 buckets on your account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/r2/buckets
operation_ids:
    - r2-list-buckets
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Buckets

`GET /accounts/{account_id}/r2/buckets`

Operation ID: `r2-list-buckets`

Lists all R2 buckets on your account.

## Definition

```yaml
{"operationId": "r2-list-buckets", "summary": "List Buckets", "description": "Lists all R2 buckets on your account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_account_identifier"}}, {"name": "name_contains", "in": "query", "schema": {"description": "Bucket names to filter by. Only buckets with this phrase in their name will be returned.", "type": "string", "example": "my-bucket"}}, {"name": "start_after", "in": "query", "schema": {"description": "Bucket name to start searching after. Buckets are ordered lexicographically.", "type": "string", "example": "my-bucket"}}, {"name": "per_page", "in": "query", "schema": {"description": "Maximum number of buckets to return in a single call.", "type": "number", "default": 20, "maximum": 1000, "minimum": 1}}, {"name": "order", "in": "query", "schema": {"description": "Field to order buckets by.", "type": "string", "enum": ["name"]}}, {"name": "direction", "in": "query", "schema": {"description": "Direction to order buckets.", "type": "string", "example": "desc", "enum": ["asc", "desc"]}}, {"name": "cursor", "in": "query", "schema": {"description": "Pagination cursor received during the last List Buckets call. R2 buckets are paginated using cursors instead of page numbers.", "type": "string"}}, {"name": "cf-r2-jurisdiction", "in": "header", "schema": {"$ref": "#/components/schemas/r2_jurisdiction"}}], "responses": {"200": {"description": "List Buckets response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/r2_v4_response_list"}, {"properties": {"result": {"type": "object", "properties": {"buckets": {"type": "array", "items": {"$ref": "#/components/schemas/r2_bucket"}}}}}, "type": "object"}]}}}}, "4XX": {"description": "List Buckets response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}}, "security": [{"api_token": []}], "tags": ["R2 Bucket"], "x-api-token-group": ["Workers R2 Storage Write", "Workers R2 Storage Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2.buckets", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
