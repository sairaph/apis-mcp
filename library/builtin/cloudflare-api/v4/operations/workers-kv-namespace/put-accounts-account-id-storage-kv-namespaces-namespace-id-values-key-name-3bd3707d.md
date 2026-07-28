---
title: Write key-value pair with optional metadata
page_id: operation-put-accounts-account-id-storage-kv-namespaces-namespace-id-values-key-na-43cba975
path: operations/workers-kv-namespace
description: Write a value identified by a key. Use URL-encoding to use special characters (for example, `:`, `!`, `%`) in the key name. Body should be the value to be stored. If JSON metadata to be associated with the key/value pair is needed, use `multipart/form-data` content type for your PUT request (see dropdown below in `REQUEST BODY SCHEMA`). Existing values, expirations, and metadata will be overwritten. If neither `expiration` nor `expiration_ttl` is specified, the key-value pair will never expire. If both are set, `expiration_ttl` is used and `expiration` is ignored.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/storage/kv/namespaces/{namespace_id}/values/{key_name}
operation_ids:
    - workers-kv-namespace-write-key-value-pair-with-metadata
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Write key-value pair with optional metadata

`PUT /accounts/{account_id}/storage/kv/namespaces/{namespace_id}/values/{key_name}`

Operation ID: `workers-kv-namespace-write-key-value-pair-with-metadata`

Write a value identified by a key. Use URL-encoding to use special characters (for example, `:`, `!`, `%`) in the key name. Body should be the value to be stored. If JSON metadata to be associated with the key/value pair is needed, use `multipart/form-data` content type for your PUT request (see dropdown below in `REQUEST BODY SCHEMA`). Existing values, expirations, and metadata will be overwritten. If neither `expiration` nor `expiration_ttl` is specified, the key-value pair will never expire. If both are set, `expiration_ttl` is used and `expiration` is ignored.

## Definition

```yaml
{"operationId": "workers-kv-namespace-write-key-value-pair-with-metadata", "summary": "Write key-value pair with optional metadata", "description": "Write a value identified by a key. Use URL-encoding to use special characters (for example, `:`, `!`, `%`) in the key name. Body should be the value to be stored. If JSON metadata to be associated with the key/value pair is needed, use `multipart/form-data` content type for your PUT request (see dropdown below in `REQUEST BODY SCHEMA`). Existing values, expirations, and metadata will be overwritten. If neither `expiration` nor `expiration_ttl` is specified, the key-value pair will never expire. If both are set, `expiration_ttl` is used and `expiration` is ignored.", "parameters": [{"name": "key_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_key_name"}}, {"name": "namespace_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_namespace_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_identifier"}}, {"name": "expiration", "in": "query", "schema": {"$ref": "#/components/schemas/workers-kv_expiration"}}, {"name": "expiration_ttl", "in": "query", "schema": {"$ref": "#/components/schemas/workers-kv_expiration_ttl"}}], "requestBody": {"required": true, "content": {"application/octet-stream": {"schema": {"$ref": "#/components/schemas/workers-kv_value"}, "x-stainless-only": ["http"]}, "multipart/form-data": {"encoding": {"metadata": {"contentType": "application/json"}}, "schema": {"type": "object", "properties": {"metadata": {"$ref": "#/components/schemas/workers-kv_metadata"}, "value": {"$ref": "#/components/schemas/workers-kv_value"}}, "required": ["value"]}}}}, "responses": {"200": {"description": "Write key-value pair with metadata response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers-kv_api-response-common-no-result"}}}}, "4XX": {"description": "Write key-value pair with metadata response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers-kv_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers KV Namespace"], "x-api-token-group": ["Workers KV Storage Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "keys", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
