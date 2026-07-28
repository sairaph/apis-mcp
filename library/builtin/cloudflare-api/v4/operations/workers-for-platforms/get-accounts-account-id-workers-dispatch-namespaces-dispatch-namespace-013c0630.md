---
title: Get dispatch namespace
page_id: operation-get-accounts-account-id-workers-dispatch-namespaces-dispatch-namespace-59957006
path: operations/workers-for-platforms
description: Get a Workers for Platforms namespace.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}
operation_ids:
    - namespace-worker-get-namespace
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get dispatch namespace

`GET /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}`

Operation ID: `namespace-worker-get-namespace`

Get a Workers for Platforms namespace.

## Definition

```yaml
{"operationId": "namespace-worker-get-namespace", "summary": "Get dispatch namespace", "description": "Get a Workers for Platforms namespace.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "dispatch_namespace", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_dispatch_namespace_name"}}], "responses": {"200": {"description": "Get a Workers for Platforms namespace.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_namespace-single-response"}}}}, "4XX": {"description": "Failure to get Workers for Platforms namespace.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers for Platforms"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dispatch-namespaces", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
