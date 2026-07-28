---
title: Delete Worker
page_id: operation-delete-accounts-account-id-workers-dispatch-namespaces-dispatch-namespac-c3757393
path: operations/workers-for-platforms
description: Delete a worker from a Workers for Platforms namespace. This call has no response body on a successful delete.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}
operation_ids:
    - namespace-worker-script-delete-worker
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Worker

`DELETE /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}`

Operation ID: `namespace-worker-script-delete-worker`

Delete a worker from a Workers for Platforms namespace. This call has no response body on a successful delete.

## Definition

```yaml
{"operationId": "namespace-worker-script-delete-worker", "summary": "Delete Worker", "description": "Delete a worker from a Workers for Platforms namespace. This call has no response body on a successful delete.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "dispatch_namespace", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_dispatch_namespace_name"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}, {"name": "force", "in": "query", "description": "If set to true, delete will not be stopped by associated service binding, durable object, or other binding. Any of these associated bindings/durable objects will be deleted along with the script.", "schema": {"type": "boolean"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Worker response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-null-result"}}}}, "4XX": {"description": "Delete Worker response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers for Platforms"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dispatch-namespaces.scripts", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
