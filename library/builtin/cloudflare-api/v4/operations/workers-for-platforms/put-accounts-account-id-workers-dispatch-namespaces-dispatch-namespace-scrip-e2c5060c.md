---
title: Upload Worker Module
page_id: operation-put-accounts-account-id-workers-dispatch-namespaces-dispatch-namespace-s-f455d82e
path: operations/workers-for-platforms
description: 'Upload a worker module to a Workers for Platforms namespace. You can find more about the multipart metadata on our docs: https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}
operation_ids:
    - namespace-worker-script-upload-worker-module
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Upload Worker Module

`PUT /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}`

Operation ID: `namespace-worker-script-upload-worker-module`

Upload a worker module to a Workers for Platforms namespace. You can find more about the multipart metadata on our docs: https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/.

## Definition

```yaml
{"operationId": "namespace-worker-script-upload-worker-module", "summary": "Upload Worker Module", "description": "Upload a worker module to a Workers for Platforms namespace. You can find more about the multipart metadata on our docs: https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "dispatch_namespace", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_dispatch_namespace_name"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}, {"name": "bindings_inherit", "in": "query", "description": "When set to \"strict\", the upload will fail if any `inherit` type bindings cannot be resolved against the previous version of the script. Without this, unresolvable inherit bindings are silently dropped.", "schema": {"type": "string", "enum": ["strict"]}}], "requestBody": {"$ref": "#/components/requestBodies/workers_namespace_upload"}, "responses": {"200": {"$ref": "#/components/responses/workers_200"}, "4XX": {"$ref": "#/components/responses/workers_4XX"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers for Platforms"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dispatch-namespaces.scripts", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
