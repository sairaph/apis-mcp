---
title: Get Worker subdomain
page_id: operation-get-accounts-account-id-workers-scripts-script-name-subdomain-eadbca20
path: operations/worker-script
description: Get if the Worker is available on the workers.dev subdomain.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/subdomain
operation_ids:
    - worker-script-get-subdomain
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Worker subdomain

`GET /accounts/{account_id}/workers/scripts/{script_name}/subdomain`

Operation ID: `worker-script-get-subdomain`

Get if the Worker is available on the workers.dev subdomain.

## Definition

```yaml
{"operationId": "worker-script-get-subdomain", "summary": "Get Worker subdomain", "description": "Get if the Worker is available on the workers.dev subdomain.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}], "responses": {"200": {"description": "Get subdomain response.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_subdomain"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Get subdomain response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Script"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.scripts.subdomain", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
