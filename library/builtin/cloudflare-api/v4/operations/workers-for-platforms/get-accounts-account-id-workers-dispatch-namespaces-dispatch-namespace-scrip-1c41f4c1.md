---
title: List Scripts in Namespace
page_id: operation-get-accounts-account-id-workers-dispatch-namespaces-dispatch-namespace-s-3fd7b7b7
path: operations/workers-for-platforms
description: Fetch a list of scripts uploaded to a Workers for Platforms namespace.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts
operation_ids:
    - namespace-worker-list-scripts
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Scripts in Namespace

`GET /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts`

Operation ID: `namespace-worker-list-scripts`

Fetch a list of scripts uploaded to a Workers for Platforms namespace.

## Definition

```yaml
{"operationId": "namespace-worker-list-scripts", "summary": "List Scripts in Namespace", "description": "Fetch a list of scripts uploaded to a Workers for Platforms namespace.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "dispatch_namespace", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_dispatch_namespace_name"}}, {"name": "tags", "in": "query", "description": "Filter scripts by tags. Format: comma-separated list of tag:allowed pairs where allowed is 'yes' or 'no'.", "schema": {"type": "string", "example": "production:yes,staging:no"}}], "responses": {"200": {"description": "List scripts in namespace response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/workers_namespace-script-response"}}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "List scripts in namespace response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers for Platforms"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dispatch-namespaces.scripts", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
