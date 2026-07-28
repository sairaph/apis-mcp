---
title: Delete Worker subdomain
page_id: operation-delete-accounts-account-id-workers-scripts-script-name-subdomain-fff064ea
path: operations/worker-script
description: Disable all workers.dev subdomains for a Worker.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/subdomain
operation_ids:
    - worker-script-delete-subdomain
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Worker subdomain

`DELETE /accounts/{account_id}/workers/scripts/{script_name}/subdomain`

Operation ID: `worker-script-delete-subdomain`

Disable all workers.dev subdomains for a Worker.

## Definition

```yaml
{"operationId": "worker-script-delete-subdomain", "summary": "Delete Worker subdomain", "description": "Disable all workers.dev subdomains for a Worker.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}], "responses": {"200": {"description": "Delete subdomain response.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_subdomain"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Delete subdomain response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Script"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.scripts.subdomain", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
