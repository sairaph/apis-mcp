---
title: Update dispatch namespace
page_id: operation-put-accounts-account-id-workers-dispatch-namespaces-dispatch-namespace-25b3720d
path: operations/workers-for-platforms
description: Update a Workers for Platforms namespace.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}
operation_ids:
    - namespace-worker-put-namespace
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update dispatch namespace

`PUT /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}`

Operation ID: `namespace-worker-put-namespace`

Update a Workers for Platforms namespace.

## Definition

```yaml
{"operationId": "namespace-worker-put-namespace", "summary": "Update dispatch namespace", "description": "Update a Workers for Platforms namespace.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "dispatch_namespace", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_dispatch_namespace_name"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"name": {"description": "The name of the dispatch namespace.", "type": "string", "example": "my-dispatch-namespace", "x-auditable": true}, "trusted_workers": {"$ref": "#/components/schemas/workers_trusted_workers"}}}}}}, "responses": {"200": {"description": "Update a Workers for Platforms namespace.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_namespace-single-response"}}}}, "4XX": {"description": "Failure to update Workers for Platforms namespace.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers for Platforms"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dispatch-namespaces", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
