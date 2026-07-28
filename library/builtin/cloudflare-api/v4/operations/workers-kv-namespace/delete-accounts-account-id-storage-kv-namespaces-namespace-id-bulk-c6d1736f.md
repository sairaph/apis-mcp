---
title: Delete multiple key-value pairs
page_id: operation-delete-accounts-account-id-storage-kv-namespaces-namespace-id-bulk-572fcdcd
path: operations/workers-kv-namespace
description: Remove multiple KV pairs from the namespace. Body should be an array of up to 10,000 keys to be removed.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/storage/kv/namespaces/{namespace_id}/bulk
operation_ids:
    - workers-kv-namespace-delete-multiple-key-value-pairs-deprecated
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete multiple key-value pairs

`DELETE /accounts/{account_id}/storage/kv/namespaces/{namespace_id}/bulk`

Operation ID: `workers-kv-namespace-delete-multiple-key-value-pairs-deprecated`

Remove multiple KV pairs from the namespace. Body should be an array of up to 10,000 keys to be removed.

## Definition

```yaml
{"operationId": "workers-kv-namespace-delete-multiple-key-value-pairs-deprecated", "summary": "Delete multiple key-value pairs", "description": "Remove multiple KV pairs from the namespace. Body should be an array of up to 10,000 keys to be removed.", "parameters": [{"name": "namespace_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_namespace_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers-kv_bulk_delete"}}}}, "responses": {"200": {"description": "Delete multiple key-value pairs response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers-kv_api-response-common-no-result"}, {"properties": {"result": {"$ref": "#/components/schemas/workers-kv_bulk-result"}}, "type": "object"}]}}}}, "4XX": {"description": "Delete multiple key-value pairs response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers-kv_api-response-common-no-result"}, {"properties": {"result": {"$ref": "#/components/schemas/workers-kv_bulk-result"}}, "type": "object"}]}}}}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers KV Namespace"], "x-api-token-group": ["Workers KV Storage Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "bulk-deprecated", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
