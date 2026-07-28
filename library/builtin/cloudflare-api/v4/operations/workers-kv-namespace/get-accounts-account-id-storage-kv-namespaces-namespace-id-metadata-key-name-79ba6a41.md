---
title: Read the metadata for a key
page_id: operation-get-accounts-account-id-storage-kv-namespaces-namespace-id-metadata-key-d8a97dc0
path: operations/workers-kv-namespace
description: Returns the metadata associated with the given key in the given namespace. Use URL-encoding to use special characters (for example, `:`, `!`, `%`) in the key name.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/storage/kv/namespaces/{namespace_id}/metadata/{key_name}
operation_ids:
    - workers-kv-namespace-read-the-metadata-for-a-key
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Read the metadata for a key

`GET /accounts/{account_id}/storage/kv/namespaces/{namespace_id}/metadata/{key_name}`

Operation ID: `workers-kv-namespace-read-the-metadata-for-a-key`

Returns the metadata associated with the given key in the given namespace. Use URL-encoding to use special characters (for example, `:`, `!`, `%`) in the key name.

## Definition

```yaml
{"operationId": "workers-kv-namespace-read-the-metadata-for-a-key", "summary": "Read the metadata for a key", "description": "Returns the metadata associated with the given key in the given namespace. Use URL-encoding to use special characters (for example, `:`, `!`, `%`) in the key name.", "parameters": [{"name": "key_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_key_name"}}, {"name": "namespace_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_namespace_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_identifier"}}], "responses": {"200": {"description": "Read the metadata for a key response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers-kv_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers-kv_list_metadata"}}, "type": "object"}]}}}}, "4XX": {"description": "Read the metadata for a key response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers-kv_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers KV Namespace"], "x-api-token-group": ["Workers KV Storage Write", "Workers KV Storage Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "metadata", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
