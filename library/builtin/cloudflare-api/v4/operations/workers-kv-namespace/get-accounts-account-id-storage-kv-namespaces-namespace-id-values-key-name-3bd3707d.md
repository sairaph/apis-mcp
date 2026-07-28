---
title: Read key-value pair
page_id: operation-get-accounts-account-id-storage-kv-namespaces-namespace-id-values-key-na-55d51380
path: operations/workers-kv-namespace
description: Returns the value associated with the given key in the given namespace. Use URL-encoding to use special characters (for example, `:`, `!`, `%`) in the key name. If the KV-pair is set to expire at some point, the expiration time as measured in seconds since the UNIX epoch will be returned in the `expiration` response header.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/storage/kv/namespaces/{namespace_id}/values/{key_name}
operation_ids:
    - workers-kv-namespace-read-key-value-pair
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Read key-value pair

`GET /accounts/{account_id}/storage/kv/namespaces/{namespace_id}/values/{key_name}`

Operation ID: `workers-kv-namespace-read-key-value-pair`

Returns the value associated with the given key in the given namespace. Use URL-encoding to use special characters (for example, `:`, `!`, `%`) in the key name. If the KV-pair is set to expire at some point, the expiration time as measured in seconds since the UNIX epoch will be returned in the `expiration` response header.

## Definition

```yaml
{"operationId": "workers-kv-namespace-read-key-value-pair", "summary": "Read key-value pair", "description": "Returns the value associated with the given key in the given namespace. Use URL-encoding to use special characters (for example, `:`, `!`, `%`) in the key name. If the KV-pair is set to expire at some point, the expiration time as measured in seconds since the UNIX epoch will be returned in the `expiration` response header.", "parameters": [{"name": "key_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_key_name"}}, {"name": "namespace_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_namespace_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_identifier"}}], "responses": {"200": {"description": "Read key-value pair response.", "content": {"application/octet-stream": {"schema": {"$ref": "#/components/schemas/workers-kv_value"}}}}, "4XX": {"description": "Read key-value pair response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers-kv_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers KV Namespace"], "x-api-token-group": ["Workers KV Storage Write", "Workers KV Storage Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "keys", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
