---
title: Upload Worker Module
page_id: operation-put-accounts-account-id-workers-scripts-script-name-e1168c2e
path: operations/worker-script
description: 'Upload a worker module. You can find more about the multipart metadata on our docs: https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}
operation_ids:
    - worker-script-upload-worker-module
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Upload Worker Module

`PUT /accounts/{account_id}/workers/scripts/{script_name}`

Operation ID: `worker-script-upload-worker-module`

Upload a worker module. You can find more about the multipart metadata on our docs: https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/.

## Definition

```yaml
{"operationId": "worker-script-upload-worker-module", "summary": "Upload Worker Module", "description": "Upload a worker module. You can find more about the multipart metadata on our docs: https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}, {"name": "bindings_inherit", "in": "query", "description": "When set to \"strict\", the upload will fail if any `inherit` type bindings cannot be resolved against the previous version of the Worker. Without this, unresolvable inherit bindings are silently dropped.", "schema": {"type": "string", "enum": ["strict"]}}], "requestBody": {"$ref": "#/components/requestBodies/workers_script_upload"}, "responses": {"200": {"description": "Upload Worker Module response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_script-response-upload-single"}, {"example": {"errors": [], "messages": [], "result": {"created_on": "2022-05-05T05:15:11.602148Z", "etag": "777f24a43bef5f69174aa69ceaf1dea67968d510a31d1vw3e49d34a0187c06d1", "handlers": ["fetch"], "id": "this-is_my_script-01", "logpush": false, "modified_on": "2022-05-20T19:02:56.446492Z", "placement": {"mode": "smart"}, "startup_time_ms": 10, "tail_consumers": [{"environment": "production", "service": "my-log-consumer"}], "usage_model": "standard"}, "success": true}, "type": "object"}]}}}}, "4XX": {"description": "Upload Worker Module response failure. When the upload uses the\ndeclarative `exports` field and one or more entries fail\nreconciliation, the response is the exports reconciliation\nerror envelope (error code 100402) with per-class detail in\n`errors[].meta.details`.\n", "content": {"application/json": {"schema": {"anyOf": [{"$ref": "#/components/schemas/workers_api-response-common-failure"}, {"$ref": "#/components/schemas/workers_exports_reconciliation_error_response"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Script"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.scripts", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
