---
title: Add script secret
page_id: operation-put-accounts-account-id-workers-scripts-script-name-secrets-5a323800
path: operations/worker-script
description: Add a secret to a script.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/secrets
operation_ids:
    - worker-put-script-secret
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Add script secret

`PUT /accounts/{account_id}/workers/scripts/{script_name}/secrets`

Operation ID: `worker-put-script-secret`

Add a secret to a script.

## Definition

```yaml
{"operationId": "worker-put-script-secret", "summary": "Add script secret", "description": "Add a secret to a script.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_secret"}}}}, "responses": {"200": {"description": "Put script secret binding success.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_secret"}}, "type": "object"}]}}}}, "429": {"description": "Too many requests are currently modifying the script.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}, "4XX": {"description": "Put script secret binding failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Script"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.secrets", "x-fern-sdk-method-name": "update", "x-forge-hidden": true, "x-forge-params": {"name": {"positional": true, "required": true}}}
```
