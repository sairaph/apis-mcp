---
title: Put Script Content
page_id: operation-put-accounts-account-id-workers-dispatch-namespaces-dispatch-namespace-s-6ad78499
path: operations/workers-for-platforms
description: Put script content for a script uploaded to a Workers for Platforms namespace.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/content
operation_ids:
    - namespace-worker-put-script-content
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Put Script Content

`PUT /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/content`

Operation ID: `namespace-worker-put-script-content`

Put script content for a script uploaded to a Workers for Platforms namespace.

## Definition

```yaml
{"operationId": "namespace-worker-put-script-content", "summary": "Put Script Content", "description": "Put script content for a script uploaded to a Workers for Platforms namespace.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "dispatch_namespace", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_dispatch_namespace_name"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}, {"name": "CF-WORKER-BODY-PART", "in": "header", "description": "The multipart name of a script upload part containing script content in service worker format. Alternative to including in a metadata part.", "schema": {"type": "string"}}, {"name": "CF-WORKER-MAIN-MODULE-PART", "in": "header", "description": "The multipart name of a script upload part containing script content in es module format. Alternative to including in a metadata part.", "schema": {"type": "string"}}], "requestBody": {"required": true, "content": {"multipart/form-data": {"encoding": {"files": {"contentType": "application/javascript+module, text/javascript+module, application/javascript, text/javascript, text/x-python, text/x-python-requirement, application/wasm, text/plain, application/octet-stream, application/source-map"}, "metadata": {"contentType": "application/json"}}, "schema": {"type": "object", "properties": {"files": {"description": "An array of modules (often JavaScript files) comprising a Worker script. At least one module must be present and referenced in the metadata as `main_module` or `body_part` by filename.<br/>Possible Content-Type(s) are: `application/javascript+module`, `text/javascript+module`, `application/javascript`, `text/javascript`, `text/x-python`, `text/x-python-requirement`, `application/wasm`, `text/plain`, `application/octet-stream`, `application/source-map`.", "type": "array", "items": {"format": "binary", "type": "string"}, "x-stainless-collection-type": "set"}, "metadata": {"description": "JSON-encoded metadata about the uploaded parts and Worker configuration.", "type": "object", "properties": {"body_part": {"description": "Name of the part in the multipart request that contains the script (e.g. the file adding a listener to the `fetch` event). Indicates a `service worker syntax` Worker.", "type": "string", "example": "worker.js"}, "main_module": {"description": "Name of the part in the multipart request that contains the main module (e.g. the file exporting a `fetch` handler). Indicates a `module syntax` Worker.", "type": "string", "example": "worker.js"}}}}, "required": ["metadata"]}}}}, "responses": {"200": {"description": "Put script content (Workers for Platforms).", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_script-response-single"}}}}, "4XX": {"description": "Put script content failure (Workers for Platforms).", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers for Platforms"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dispatch-namespaces.scripts.content", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
