---
title: Get secret binding
page_id: operation-get-accounts-account-id-workers-scripts-script-name-secrets-secret-name-0801e623
path: operations/worker-script
description: Get a given secret binding (value omitted) on a script.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/secrets/{secret_name}
operation_ids:
    - worker-get-script-secret
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get secret binding

`GET /accounts/{account_id}/workers/scripts/{script_name}/secrets/{secret_name}`

Operation ID: `worker-get-script-secret`

Get a given secret binding (value omitted) on a script.

## Definition

```yaml
{"operationId": "worker-get-script-secret", "summary": "Get secret binding", "description": "Get a given secret binding (value omitted) on a script.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}, {"name": "secret_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_secret_name"}}, {"name": "url_encoded", "in": "query", "schema": {"$ref": "#/components/schemas/workers_secret_name_url_encoded"}}], "responses": {"200": {"description": "Get script secret binding.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_secret"}}, "type": "object"}]}}}}, "4XX": {"description": "Get script secret failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Script"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.secrets", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
