---
title: Patch dispatch namespace
page_id: operation-patch-accounts-account-id-workers-dispatch-namespaces-dispatch-namespace-4eb7f888
path: operations/workers-for-platforms
description: Patch a Workers for Platforms namespace. Omitted fields are left unchanged.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}
operation_ids:
    - namespace-worker-patch-namespace
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch dispatch namespace

`PATCH /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}`

Operation ID: `namespace-worker-patch-namespace`

Patch a Workers for Platforms namespace. Omitted fields are left unchanged.

## Definition

```yaml
{"operationId": "namespace-worker-patch-namespace", "summary": "Patch dispatch namespace", "description": "Patch a Workers for Platforms namespace. Omitted fields are left unchanged.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "dispatch_namespace", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_dispatch_namespace_name"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"name": {"description": "The name of the dispatch namespace.", "type": "string", "example": "my-dispatch-namespace", "x-auditable": true}, "trusted_workers": {"$ref": "#/components/schemas/workers_trusted_workers"}}}}}}, "responses": {"200": {"description": "Patch a Workers for Platforms namespace.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_namespace-single-response"}}}}, "4XX": {"description": "Failure to patch Workers for Platforms namespace.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers for Platforms"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dispatch-namespaces", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
