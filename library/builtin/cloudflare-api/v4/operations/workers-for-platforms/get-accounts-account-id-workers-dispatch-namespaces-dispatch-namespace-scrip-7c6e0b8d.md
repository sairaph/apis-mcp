---
title: Get secret binding
page_id: operation-get-accounts-account-id-workers-dispatch-namespaces-dispatch-namespace-s-fa3bcc06
path: operations/workers-for-platforms
description: Get a given secret binding (value omitted) on a script uploaded to a Workers for Platforms namespace.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/secrets/{secret_name}
operation_ids:
    - namespace-worker-get-script-secrets
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get secret binding

`GET /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/secrets/{secret_name}`

Operation ID: `namespace-worker-get-script-secrets`

Get a given secret binding (value omitted) on a script uploaded to a Workers for Platforms namespace.

## Definition

```yaml
{"operationId": "namespace-worker-get-script-secrets", "summary": "Get secret binding", "description": "Get a given secret binding (value omitted) on a script uploaded to a Workers for Platforms namespace.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "dispatch_namespace", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_dispatch_namespace_name"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}, {"name": "secret_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_secret_name"}}, {"name": "url_encoded", "in": "query", "schema": {"$ref": "#/components/schemas/workers_secret_name_url_encoded"}}], "responses": {"200": {"description": "Get script secret (Workers for Platforms).", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_secret"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Get script secret failure (Workers for Platforms).", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers for Platforms"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dispatch-namespaces.scripts.secrets", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
