---
title: Delete key-value pair
page_id: operation-delete-accounts-account-id-storage-kv-namespaces-namespace-id-values-key-3f88fdc4
path: operations/workers-kv-namespace
description: Remove a KV pair from the namespace. Use URL-encoding to use special characters (for example, `:`, `!`, `%`) in the key name.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/storage/kv/namespaces/{namespace_id}/values/{key_name}
operation_ids:
    - workers-kv-namespace-delete-key-value-pair
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete key-value pair

`DELETE /accounts/{account_id}/storage/kv/namespaces/{namespace_id}/values/{key_name}`

Operation ID: `workers-kv-namespace-delete-key-value-pair`

Remove a KV pair from the namespace. Use URL-encoding to use special characters (for example, `:`, `!`, `%`) in the key name.

## Definition

```yaml
{"operationId": "workers-kv-namespace-delete-key-value-pair", "summary": "Delete key-value pair", "description": "Remove a KV pair from the namespace. Use URL-encoding to use special characters (for example, `:`, `!`, `%`) in the key name.", "parameters": [{"name": "key_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_key_name"}}, {"name": "namespace_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_namespace_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete key-value pair response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers-kv_api-response-common-no-result"}}}}, "4XX": {"description": "Delete key-value pair response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers-kv_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers KV Namespace"], "x-api-token-group": ["Workers KV Storage Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "keys", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
