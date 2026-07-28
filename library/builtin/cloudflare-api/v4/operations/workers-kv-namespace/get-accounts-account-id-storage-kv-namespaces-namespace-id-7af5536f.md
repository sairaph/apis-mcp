---
title: Get a Namespace
page_id: operation-get-accounts-account-id-storage-kv-namespaces-namespace-id-6aed5560
path: operations/workers-kv-namespace
description: Get the namespace corresponding to the given ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/storage/kv/namespaces/{namespace_id}
operation_ids:
    - workers-kv-namespace-get-a-namespace
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a Namespace

`GET /accounts/{account_id}/storage/kv/namespaces/{namespace_id}`

Operation ID: `workers-kv-namespace-get-a-namespace`

Get the namespace corresponding to the given ID.

## Definition

```yaml
{"operationId": "workers-kv-namespace-get-a-namespace", "summary": "Get a Namespace", "description": "Get the namespace corresponding to the given ID.", "parameters": [{"name": "namespace_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_namespace_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_identifier"}}], "responses": {"200": {"description": "Get a Namespace response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers-kv_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers-kv_namespace"}}, "type": "object"}]}}}}, "4XX": {"description": "Get a Namespace response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers-kv_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers KV Namespace"], "x-api-token-group": ["Workers KV Storage Write", "Workers KV Storage Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "namespaces", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
