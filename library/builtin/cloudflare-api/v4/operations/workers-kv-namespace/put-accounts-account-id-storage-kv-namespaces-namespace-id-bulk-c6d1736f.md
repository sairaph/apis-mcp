---
title: Write multiple key-value pairs
page_id: operation-put-accounts-account-id-storage-kv-namespaces-namespace-id-bulk-7f8a2f50
path: operations/workers-kv-namespace
description: Write multiple keys and values at once. Body should be an array of up to 10,000 key-value pairs to be stored, along with optional expiration information. Existing values and expirations will be overwritten. If neither `expiration` nor `expiration_ttl` is specified, the key-value pair will never expire. If both are set, `expiration_ttl` is used and `expiration` is ignored. The entire request size must be 100 megabytes or less.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/storage/kv/namespaces/{namespace_id}/bulk
operation_ids:
    - workers-kv-namespace-write-multiple-key-value-pairs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Write multiple key-value pairs

`PUT /accounts/{account_id}/storage/kv/namespaces/{namespace_id}/bulk`

Operation ID: `workers-kv-namespace-write-multiple-key-value-pairs`

Write multiple keys and values at once. Body should be an array of up to 10,000 key-value pairs to be stored, along with optional expiration information. Existing values and expirations will be overwritten. If neither `expiration` nor `expiration_ttl` is specified, the key-value pair will never expire. If both are set, `expiration_ttl` is used and `expiration` is ignored. The entire request size must be 100 megabytes or less.

## Definition

```yaml
{"operationId": "workers-kv-namespace-write-multiple-key-value-pairs", "summary": "Write multiple key-value pairs", "description": "Write multiple keys and values at once. Body should be an array of up to 10,000 key-value pairs to be stored, along with optional expiration information. Existing values and expirations will be overwritten. If neither `expiration` nor `expiration_ttl` is specified, the key-value pair will never expire. If both are set, `expiration_ttl` is used and `expiration` is ignored. The entire request size must be 100 megabytes or less.", "parameters": [{"name": "namespace_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_namespace_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers-kv_bulk_write"}}}}, "responses": {"200": {"description": "Write multiple key-value pairs response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers-kv_api-response-common-no-result"}, {"properties": {"result": {"$ref": "#/components/schemas/workers-kv_bulk-result"}}, "type": "object"}]}}}}, "4XX": {"description": "Write multiple key-value pairs response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers-kv_api-response-common-no-result"}, {"properties": {"result": {"$ref": "#/components/schemas/workers-kv_bulk-result"}}, "type": "object"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers KV Namespace"], "x-api-token-group": ["Workers KV Storage Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "bulk", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
