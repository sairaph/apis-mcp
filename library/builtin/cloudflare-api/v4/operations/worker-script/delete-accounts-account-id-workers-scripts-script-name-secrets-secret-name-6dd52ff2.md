---
title: Delete script secret
page_id: operation-delete-accounts-account-id-workers-scripts-script-name-secrets-secret-na-ddb72fa2
path: operations/worker-script
description: Remove a secret from a script.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/secrets/{secret_name}
operation_ids:
    - worker-delete-script-secret
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete script secret

`DELETE /accounts/{account_id}/workers/scripts/{script_name}/secrets/{secret_name}`

Operation ID: `worker-delete-script-secret`

Remove a secret from a script.

## Definition

```yaml
{"operationId": "worker-delete-script-secret", "summary": "Delete script secret", "description": "Remove a secret from a script.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}, {"name": "secret_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_secret_name"}}, {"name": "url_encoded", "in": "query", "schema": {"$ref": "#/components/schemas/workers_secret_name_url_encoded"}}], "responses": {"200": {"description": "Delete script secret binding.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-null-result"}}}}, "4XX": {"description": "Delete script secret failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Script"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.secrets", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
