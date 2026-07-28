---
title: Create a live input
page_id: operation-post-accounts-account-id-stream-live-inputs-eb576b4d
path: operations/stream-live-inputs
description: Creates a live input, and returns credentials that you or your users can use to stream live video to Cloudflare Stream.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/stream/live_inputs
operation_ids:
    - stream-live-inputs-create-a-live-input
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a live input

`POST /accounts/{account_id}/stream/live_inputs`

Operation ID: `stream-live-inputs-create-a-live-input`

Creates a live input, and returns credentials that you or your users can use to stream live video to Cloudflare Stream.

## Definition

```yaml
{"operationId": "stream-live-inputs-create-a-live-input", "summary": "Create a live input", "description": "Creates a live input, and returns credentials that you or your users can use to stream live video to Cloudflare Stream.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_create_input_request"}}}}, "responses": {"200": {"description": "Create a live input response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_live_input_response_single"}}}}, "4XX": {"description": "Create a live input response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Live Inputs"], "x-api-token-group": ["Stream Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.live-inputs", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
