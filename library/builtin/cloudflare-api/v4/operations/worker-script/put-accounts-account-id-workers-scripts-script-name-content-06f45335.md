---
title: Put script content
page_id: operation-put-accounts-account-id-workers-scripts-script-name-content-76850d62
path: operations/worker-script
description: Put script content without touching config or metadata.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/content
operation_ids:
    - worker-script-put-content
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Put script content

`PUT /accounts/{account_id}/workers/scripts/{script_name}/content`

Operation ID: `worker-script-put-content`

Put script content without touching config or metadata.

## Definition

```yaml
{"operationId": "worker-script-put-content", "summary": "Put script content", "description": "Put script content without touching config or metadata.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}, {"name": "CF-WORKER-BODY-PART", "in": "header", "description": "The multipart name of a script upload part containing script content in service worker format. Alternative to including in a metadata part.", "schema": {"type": "string"}}, {"name": "CF-WORKER-MAIN-MODULE-PART", "in": "header", "description": "The multipart name of a script upload part containing script content in es module format. Alternative to including in a metadata part.", "schema": {"type": "string"}}], "requestBody": {"required": true, "content": {"multipart/form-data": {"encoding": {"files": {"contentType": "application/javascript+module, text/javascript+module, application/javascript, text/javascript, text/x-python, text/x-python-requirement, application/wasm, text/plain, application/octet-stream, application/source-map"}, "metadata": {"contentType": "application/json"}}, "schema": {"type": "object", "properties": {"files": {"description": "An array of modules (often JavaScript files) comprising a Worker script. At least one module must be present and referenced in the metadata as `main_module` or `body_part` by filename.<br/>Possible Content-Type(s) are: `application/javascript+module`, `text/javascript+module`, `application/javascript`, `text/javascript`, `text/x-python`, `text/x-python-requirement`, `application/wasm`, `text/plain`, `application/octet-stream`, `application/source-map`.", "type": "array", "items": {"format": "binary", "type": "string"}, "x-stainless-collection-type": "set"}, "metadata": {"description": "JSON-encoded metadata about the uploaded parts and Worker configuration.", "type": "object", "properties": {"body_part": {"description": "Name of the uploaded file that contains the Worker script (e.g. the file adding a listener to the `fetch` event). Indicates a `service worker syntax` Worker.", "type": "string", "example": "worker.js"}, "main_module": {"description": "Name of the uploaded file that contains the main module (e.g. the file exporting a `fetch` handler). Indicates a `module syntax` Worker.", "type": "string", "example": "worker.js"}}}}, "required": ["metadata"]}}}}, "responses": {"200": {"description": "Put script content.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_script-response-single"}}}}, "4XX": {"description": "Put script content failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Script"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.scripts.content", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
