---
title: Upload Version
page_id: operation-post-accounts-account-id-workers-scripts-script-name-versions-7545fdc5
path: operations/worker-versions
description: 'Upload a Worker Version without deploying to Cloudflare''s network. You can find more about the multipart metadata on our docs: https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/versions
operation_ids:
    - worker-versions-upload-version
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Upload Version

`POST /accounts/{account_id}/workers/scripts/{script_name}/versions`

Operation ID: `worker-versions-upload-version`

Upload a Worker Version without deploying to Cloudflare's network. You can find more about the multipart metadata on our docs: https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/.

## Definition

```yaml
{"operationId": "worker-versions-upload-version", "summary": "Upload Version", "description": "Upload a Worker Version without deploying to Cloudflare's network. You can find more about the multipart metadata on our docs: https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name-2"}}, {"name": "bindings_inherit", "in": "query", "description": "When set to \"strict\", the upload will fail if any `inherit` type bindings cannot be resolved against the previous version of the Worker. Without this, unresolvable inherit bindings are silently dropped.", "schema": {"type": "string", "enum": ["strict"]}}], "requestBody": {"$ref": "#/components/requestBodies/workers_version-post"}, "responses": {"200": {"description": "Upload Version response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_versions-upload-response"}}}}, "4XX": {"description": "Upload Version response failure. When the upload uses the\ndeclarative `exports` field and one or more entries fail\nreconciliation, the response is the exports reconciliation\nerror envelope (error code 100402) with per-class detail in\n`errors[].meta.details`.\n", "content": {"application/json": {"schema": {"anyOf": [{"allOf": [{"$ref": "#/components/schemas/workers_versions-upload-response"}, {"$ref": "#/components/schemas/workers_api-response-common-failure"}]}, {"$ref": "#/components/schemas/workers_exports_reconciliation_error_response"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Versions"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.versions", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
