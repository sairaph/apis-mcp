---
title: Create Assets Upload Session
page_id: operation-post-accounts-account-id-workers-scripts-script-name-assets-upload-sessi-8cb709da
path: operations/worker-script
description: Start uploading a collection of assets for use in a Worker version. To learn more about the direct uploads of assets, see https://developers.cloudflare.com/workers/static-assets/direct-upload/.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/assets-upload-session
operation_ids:
    - worker-script-update-create-assets-upload-session
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Assets Upload Session

`POST /accounts/{account_id}/workers/scripts/{script_name}/assets-upload-session`

Operation ID: `worker-script-update-create-assets-upload-session`

Start uploading a collection of assets for use in a Worker version. To learn more about the direct uploads of assets, see https://developers.cloudflare.com/workers/static-assets/direct-upload/.

## Definition

```yaml
{"operationId": "worker-script-update-create-assets-upload-session", "summary": "Create Assets Upload Session", "description": "Start uploading a collection of assets for use in a Worker version. To learn more about the direct uploads of assets, see https://developers.cloudflare.com/workers/static-assets/direct-upload/.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_create-assets-upload-session-object"}}}}, "responses": {"200": {"description": "Create Assets Upload Session response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_create-assets-upload-session-response"}}}}, "4XX": {"description": "Create Assets Upload Session response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_create-assets-upload-session-response"}, {"$ref": "#/components/schemas/workers_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Script"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.scripts.assets.upload", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
