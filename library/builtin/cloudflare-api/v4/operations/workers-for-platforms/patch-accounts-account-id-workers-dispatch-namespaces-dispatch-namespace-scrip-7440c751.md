---
title: Patch multiple script secrets
page_id: operation-patch-accounts-account-id-workers-dispatch-namespaces-dispatch-namespace-069da66f
path: operations/workers-for-platforms
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
    - /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/secrets-bulk
operation_ids:
    - namespace-worker-patch-script-secrets-bulk
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch multiple script secrets

`PATCH /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/secrets-bulk`

Operation ID: `namespace-worker-patch-script-secrets-bulk`

Create, update, or delete multiple secrets on a script in a single operation using JSON Merge Patch (RFC 7396).

Usage:

- To create or update a secret, set its value to a secret object.
- To delete a secret, set its value to `null`.
- Secrets not included in the request are left unchanged.

## Definition

```yaml
{"operationId": "namespace-worker-patch-script-secrets-bulk", "summary": "Patch multiple script secrets", "description": "Create, update, or delete multiple secrets on a script in a single operation using JSON Merge Patch (RFC 7396).\n\nUsage:\n\n- To create or update a secret, set its value to a secret object.\n- To delete a secret, set its value to `null`.\n- Secrets not included in the request are left unchanged.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "dispatch_namespace", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_dispatch_namespace_name"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_secret-patch-request"}}, "application/merge-patch+json": {"schema": {"$ref": "#/components/schemas/workers_secret-patch-request"}}}}, "responses": {"200": {"description": "Patch script secrets bulk success (Workers for Platforms).", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_secret-patch-response"}}, "required": ["result"], "type": "object"}]}}}}, "429": {"description": "Too many requests are currently modifying the script.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}, "4XX": {"description": "Patch script secrets bulk failure (Workers for Platforms).", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers for Platforms"]}
```
