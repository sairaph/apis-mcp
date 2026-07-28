---
title: List dispatch namespaces
page_id: operation-get-accounts-account-id-workers-dispatch-namespaces-2e714c3a
path: operations/workers-for-platforms
description: Fetch a list of Workers for Platforms namespaces.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/dispatch/namespaces
operation_ids:
    - namespace-worker-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List dispatch namespaces

`GET /accounts/{account_id}/workers/dispatch/namespaces`

Operation ID: `namespace-worker-list`

Fetch a list of Workers for Platforms namespaces.

## Definition

```yaml
{"operationId": "namespace-worker-list", "summary": "List dispatch namespaces", "description": "Fetch a list of Workers for Platforms namespaces.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}], "responses": {"200": {"description": "Fetch a list of Workers for Platforms namespaces.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_namespace-list-response"}}}}, "4XX": {"description": "Failure to get list of Workers for Platforms namespaces.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers for Platforms"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dispatch-namespaces", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
