---
title: Delete Scripts in Namespace
page_id: operation-delete-accounts-account-id-workers-dispatch-namespaces-dispatch-namespac-1fd6955a
path: operations/workers-for-platforms
description: Delete multiple scripts from a Workers for Platforms namespace based on optional tag filters.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts
operation_ids:
    - namespace-worker-delete-scripts
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Scripts in Namespace

`DELETE /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts`

Operation ID: `namespace-worker-delete-scripts`

Delete multiple scripts from a Workers for Platforms namespace based on optional tag filters.

## Definition

```yaml
{"operationId": "namespace-worker-delete-scripts", "summary": "Delete Scripts in Namespace", "description": "Delete multiple scripts from a Workers for Platforms namespace based on optional tag filters.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "dispatch_namespace", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_dispatch_namespace_name"}}, {"name": "tags", "in": "query", "description": "Filter scripts by tags before deletion. Format: comma-separated list of tag:allowed pairs where allowed is 'yes' or 'no'.", "schema": {"type": "string", "example": "production:yes,staging:no"}}, {"name": "limit", "in": "query", "description": "Limit the number of scripts to delete.", "schema": {"type": "integer", "example": 1000}}], "responses": {"200": {"description": "Delete scripts in namespace response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_namespace-script-delete-bulk-response"}}}}, "4XX": {"description": "Delete scripts in namespace response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers for Platforms"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dispatch-namespaces.scripts", "x-fern-sdk-method-name": "bulk-delete", "x-forge-hidden": true}
```
