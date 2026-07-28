---
title: Download Worker
page_id: operation-get-accounts-account-id-workers-scripts-script-name-7de3a221
path: operations/worker-script
description: Fetch raw script content for your worker. Note this is the original script content, not JSON encoded.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}
operation_ids:
    - worker-script-download-worker
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Download Worker

`GET /accounts/{account_id}/workers/scripts/{script_name}`

Operation ID: `worker-script-download-worker`

Fetch raw script content for your worker. Note this is the original script content, not JSON encoded.

## Definition

```yaml
{"operationId": "worker-script-download-worker", "summary": "Download Worker", "description": "Fetch raw script content for your worker. Note this is the original script content, not JSON encoded.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}], "responses": {"200": {"description": "Worker successfully downloaded. Returns script content as a multipart form, with no metadata part and no JSON encoding applied.", "content": {"application/javascript": {"schema": {"type": "string", "example": "export default {\n  async fetch(request, env, ctx) {\n    return new Response(\"Hello, world!\");\n  }\n};\n"}}, "multipart/form-data": {"schema": {"type": "object", "additionalProperties": {"description": "Modules used by the Worker, including JavaScript files and source maps.", "format": "binary", "type": "string"}}}}}, "4XX": {"description": "Download Worker response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Script"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.scripts", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
