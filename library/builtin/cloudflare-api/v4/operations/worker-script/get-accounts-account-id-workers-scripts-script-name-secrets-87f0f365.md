---
title: List script secrets
page_id: operation-get-accounts-account-id-workers-scripts-script-name-secrets-f32f3099
path: operations/worker-script
description: List secrets bound to a script.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/secrets
operation_ids:
    - worker-list-script-secrets
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List script secrets

`GET /accounts/{account_id}/workers/scripts/{script_name}/secrets`

Operation ID: `worker-list-script-secrets`

List secrets bound to a script.

## Definition

```yaml
{"operationId": "worker-list-script-secrets", "summary": "List script secrets", "description": "List secrets bound to a script.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}], "responses": {"200": {"description": "List script secrets.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/workers_secret"}}}, "type": "object"}]}}}}, "4XX": {"description": "List script secrets failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Script"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.secrets", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
