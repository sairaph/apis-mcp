---
title: Rename a Namespace
page_id: operation-put-accounts-account-id-storage-kv-namespaces-namespace-id-3d244e8f
path: operations/workers-kv-namespace
description: Modifies a namespace's title.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/storage/kv/namespaces/{namespace_id}
operation_ids:
    - workers-kv-namespace-rename-a-namespace
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Rename a Namespace

`PUT /accounts/{account_id}/storage/kv/namespaces/{namespace_id}`

Operation ID: `workers-kv-namespace-rename-a-namespace`

Modifies a namespace's title.

## Definition

```yaml
{"operationId": "workers-kv-namespace-rename-a-namespace", "summary": "Rename a Namespace", "description": "Modifies a namespace's title.", "parameters": [{"name": "namespace_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_namespace_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers-kv_create_rename_namespace_body"}}}}, "responses": {"200": {"description": "Rename a Namespace response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers-kv_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers-kv_namespace"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Rename a Namespace response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers-kv_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers KV Namespace"], "x-api-token-group": ["Workers KV Storage Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "namespaces", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
