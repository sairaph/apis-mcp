---
title: Create a Namespace
page_id: operation-post-accounts-account-id-storage-kv-namespaces-7bd8d139
path: operations/workers-kv-namespace
description: Creates a namespace under the given title. A `400` is returned if the account already owns a namespace with this title. A namespace must be explicitly deleted to be replaced.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/storage/kv/namespaces
operation_ids:
    - workers-kv-namespace-create-a-namespace
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a Namespace

`POST /accounts/{account_id}/storage/kv/namespaces`

Operation ID: `workers-kv-namespace-create-a-namespace`

Creates a namespace under the given title. A `400` is returned if the account already owns a namespace with this title. A namespace must be explicitly deleted to be replaced.

## Definition

```yaml
{"operationId": "workers-kv-namespace-create-a-namespace", "summary": "Create a Namespace", "description": "Creates a namespace under the given title. A `400` is returned if the account already owns a namespace with this title. A namespace must be explicitly deleted to be replaced.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers-kv_create_rename_namespace_body"}}}}, "responses": {"200": {"description": "Create a Namespace response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers-kv_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers-kv_namespace"}}, "type": "object"}]}}}}, "4XX": {"description": "Create a Namespace response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers-kv_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers KV Namespace"], "x-api-token-group": ["Workers KV Storage Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "namespaces", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
