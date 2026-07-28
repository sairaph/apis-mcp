---
title: Create dispatch namespace
page_id: operation-post-accounts-account-id-workers-dispatch-namespaces-d9632049
path: operations/workers-for-platforms
description: Create a new Workers for Platforms namespace.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/workers/dispatch/namespaces
operation_ids:
    - namespace-worker-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create dispatch namespace

`POST /accounts/{account_id}/workers/dispatch/namespaces`

Operation ID: `namespace-worker-create`

Create a new Workers for Platforms namespace.

## Definition

```yaml
{"operationId": "namespace-worker-create", "summary": "Create dispatch namespace", "description": "Create a new Workers for Platforms namespace.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"name": {"description": "The name of the dispatch namespace.", "type": "string", "example": "my-dispatch-namespace", "x-auditable": true}}}}}}, "responses": {"200": {"description": "Fetch a list of Workers for Platforms namespaces.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_namespace-single-response"}}}}, "4XX": {"description": "Failure to get list of Workers for Platforms namespaces.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers for Platforms"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dispatch-namespaces", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
