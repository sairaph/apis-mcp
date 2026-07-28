---
title: Patch multiple script secrets
page_id: operation-patch-accounts-account-id-workers-scripts-script-name-secrets-bulk-c85e18bd
path: operations/worker-script
description: |-
    Create, update, or delete multiple secrets on a script in a single operation using JSON Merge Patch (RFC 7396).

    Usage:

    - To create or update a secret, set its value to a secret object.
    - To delete a secret, set its value to `null`.
    - Secrets not included in the request are left unchanged.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/secrets-bulk
operation_ids:
    - worker-patch-script-secrets-bulk
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch multiple script secrets

`PATCH /accounts/{account_id}/workers/scripts/{script_name}/secrets-bulk`

Operation ID: `worker-patch-script-secrets-bulk`

Create, update, or delete multiple secrets on a script in a single operation using JSON Merge Patch (RFC 7396).

Usage:

- To create or update a secret, set its value to a secret object.
- To delete a secret, set its value to `null`.
- Secrets not included in the request are left unchanged.

## Definition

```yaml
{"operationId": "worker-patch-script-secrets-bulk", "summary": "Patch multiple script secrets", "description": "Create, update, or delete multiple secrets on a script in a single operation using JSON Merge Patch (RFC 7396).\n\nUsage:\n\n- To create or update a secret, set its value to a secret object.\n- To delete a secret, set its value to `null`.\n- Secrets not included in the request are left unchanged.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_secret-patch-request"}}, "application/merge-patch+json": {"schema": {"$ref": "#/components/schemas/workers_secret-patch-request"}}}}, "responses": {"200": {"description": "Patch script secrets bulk success.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_secret-patch-response"}}, "type": "object"}]}}}}, "429": {"description": "Too many requests are currently modifying the script.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}, "4XX": {"description": "Patch script secrets bulk failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Script"]}
```
