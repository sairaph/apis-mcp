---
title: Delete dispatch namespace
page_id: operation-delete-accounts-account-id-workers-dispatch-namespaces-dispatch-namespac-43b4fd73
path: operations/workers-for-platforms
description: Delete a Workers for Platforms namespace.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}
operation_ids:
    - namespace-worker-delete-namespace
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete dispatch namespace

`DELETE /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}`

Operation ID: `namespace-worker-delete-namespace`

Delete a Workers for Platforms namespace.

## Definition

```yaml
{"operationId": "namespace-worker-delete-namespace", "summary": "Delete dispatch namespace", "description": "Delete a Workers for Platforms namespace.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "dispatch_namespace", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_dispatch_namespace_name"}}], "responses": {"200": {"description": "Delete a Workers for Platforms namespace.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-null-result"}}}}, "4XX": {"description": "Failure to delete Workers for Platforms namespace.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers for Platforms"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dispatch-namespaces", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
