---
title: Get Script Bindings
page_id: operation-get-accounts-account-id-workers-dispatch-namespaces-dispatch-namespace-s-71de8f05
path: operations/workers-for-platforms
description: Fetch script bindings from a script uploaded to a Workers for Platforms namespace.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/bindings
operation_ids:
    - namespace-worker-get-script-bindings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Script Bindings

`GET /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/bindings`

Operation ID: `namespace-worker-get-script-bindings`

Fetch script bindings from a script uploaded to a Workers for Platforms namespace.

## Definition

```yaml
{"operationId": "namespace-worker-get-script-bindings", "summary": "Get Script Bindings", "description": "Fetch script bindings from a script uploaded to a Workers for Platforms namespace.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "dispatch_namespace", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_dispatch_namespace_name"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}], "responses": {"200": {"description": "Fetch script bindings (Workers for Platforms).", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_bindings"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Fetch script bindings failure (Workers for Platforms).", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers for Platforms"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dispatch-namespaces.scripts.bindings", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
