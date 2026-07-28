---
title: Put Script Tag
page_id: operation-put-accounts-account-id-workers-dispatch-namespaces-dispatch-namespace-s-f044b156
path: operations/workers-for-platforms
description: Put a single tag on a script uploaded to a Workers for Platforms namespace.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/tags/{tag}
operation_ids:
    - namespace-worker-put-script-tag
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Put Script Tag

`PUT /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/tags/{tag}`

Operation ID: `namespace-worker-put-script-tag`

Put a single tag on a script uploaded to a Workers for Platforms namespace.

## Definition

```yaml
{"operationId": "namespace-worker-put-script-tag", "summary": "Put Script Tag", "description": "Put a single tag on a script uploaded to a Workers for Platforms namespace.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "dispatch_namespace", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_dispatch_namespace_name"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}, {"name": "tag", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_tag"}}], "responses": {"200": {"description": "Put script tag.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-null-result"}}}}, "4XX": {"description": "Put script tag failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers for Platforms"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dispatch-namespaces.scripts.tags", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
