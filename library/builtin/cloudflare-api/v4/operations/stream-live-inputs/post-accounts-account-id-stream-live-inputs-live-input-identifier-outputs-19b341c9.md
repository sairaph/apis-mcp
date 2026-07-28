---
title: Create a new output, connected to a live input
page_id: operation-post-accounts-account-id-stream-live-inputs-live-input-identifier-output-0a44474d
path: operations/stream-live-inputs
description: Creates a new output that can be used to simulcast or restream live video to other RTMP or SRT destinations. Outputs are always linked to a specific live input — one live input can have many outputs.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/stream/live_inputs/{live_input_identifier}/outputs
operation_ids:
    - stream-live-inputs-create-a-new-output,-connected-to-a-live-input
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a new output, connected to a live input

`POST /accounts/{account_id}/stream/live_inputs/{live_input_identifier}/outputs`

Operation ID: `stream-live-inputs-create-a-new-output,-connected-to-a-live-input`

Creates a new output that can be used to simulcast or restream live video to other RTMP or SRT destinations. Outputs are always linked to a specific live input — one live input can have many outputs.

## Definition

```yaml
{"operationId": "stream-live-inputs-create-a-new-output,-connected-to-a-live-input", "summary": "Create a new output, connected to a live input", "description": "Creates a new output that can be used to simulcast or restream live video to other RTMP or SRT destinations. Outputs are always linked to a specific live input — one live input can have many outputs.", "parameters": [{"name": "live_input_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_live_input_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_create_output_request"}}}}, "responses": {"200": {"description": "Create a new output, connected to a live input response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_output_response_single"}}}}, "4XX": {"description": "Create a new output, connected to a live input response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Live Inputs"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.live-inputs.outputs", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
