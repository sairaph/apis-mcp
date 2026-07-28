---
title: Remove a Namespace
page_id: operation-delete-accounts-account-id-storage-kv-namespaces-namespace-id-1c93aafa
path: operations/workers-kv-namespace
description: Deletes the namespace corresponding to the given ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/storage/kv/namespaces/{namespace_id}
operation_ids:
    - workers-kv-namespace-remove-a-namespace
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Remove a Namespace

`DELETE /accounts/{account_id}/storage/kv/namespaces/{namespace_id}`

Operation ID: `workers-kv-namespace-remove-a-namespace`

Deletes the namespace corresponding to the given ID.

## Definition

```yaml
{"operationId": "workers-kv-namespace-remove-a-namespace", "summary": "Remove a Namespace", "description": "Deletes the namespace corresponding to the given ID.", "parameters": [{"name": "namespace_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_namespace_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Remove a Namespace response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers-kv_api-response-common-no-result"}}}}, "4XX": {"description": "Remove a Namespace response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers-kv_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers KV Namespace"], "x-api-token-group": ["Workers KV Storage Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "namespaces", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
