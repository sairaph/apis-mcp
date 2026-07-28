---
title: Get multiple key-value pairs
page_id: operation-post-accounts-account-id-storage-kv-namespaces-namespace-id-bulk-get-09283bac
path: operations/workers-kv-namespace
description: Retrieve up to 100 KV pairs from the namespace. Keys must contain text-based values. JSON values can optionally be parsed instead of being returned as a string value. Metadata can be included if `withMetadata` is true.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/storage/kv/namespaces/{namespace_id}/bulk/get
operation_ids:
    - workers-kv-namespace-get-multiple-key-value-pairs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get multiple key-value pairs

`POST /accounts/{account_id}/storage/kv/namespaces/{namespace_id}/bulk/get`

Operation ID: `workers-kv-namespace-get-multiple-key-value-pairs`

Retrieve up to 100 KV pairs from the namespace. Keys must contain text-based values. JSON values can optionally be parsed instead of being returned as a string value. Metadata can be included if `withMetadata` is true.

## Definition

```yaml
{"operationId": "workers-kv-namespace-get-multiple-key-value-pairs", "summary": "Get multiple key-value pairs", "description": "Retrieve up to 100 KV pairs from the namespace. Keys must contain text-based values. JSON values can optionally be parsed instead of being returned as a string value. Metadata can be included if `withMetadata` is true.", "parameters": [{"name": "namespace_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_namespace_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"keys": {"description": "Array of keys to retrieve (maximum of 100).", "type": "array", "items": {"$ref": "#/components/schemas/workers-kv_key_name_bulk"}, "maxItems": 100}, "type": {"description": "Whether to parse JSON values in the response.", "type": "string", "default": "text", "enum": ["text", "json"]}, "withMetadata": {"description": "Whether to include metadata in the response.", "type": "boolean", "default": false}}, "required": ["keys"]}}}}, "responses": {"200": {"description": "Get multiple key-value pairs response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers-kv_api-response-common-no-result"}, {"properties": {"result": {"oneOf": [{"$ref": "#/components/schemas/workers-kv_bulk-get-result"}, {"$ref": "#/components/schemas/workers-kv_bulk-get-result-with-metadata"}]}}, "type": "object"}]}}}}, "4XX": {"description": "Get multiple key-value pairs response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers-kv_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers KV Namespace"], "x-api-token-group": ["Workers KV Storage Write", "Workers KV Storage Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "bulk", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
