---
title: Enable a live input
page_id: operation-post-accounts-account-id-stream-live-inputs-live-input-identifier-enable-e9ac1ae0
path: operations/stream-live-inputs
description: Allows a live input to be streamed to and makes the live input accessible to any future API calls.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/stream/live_inputs/{live_input_identifier}/enable
operation_ids:
    - stream-live-inputs-enable-a-live-input
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Enable a live input

`POST /accounts/{account_id}/stream/live_inputs/{live_input_identifier}/enable`

Operation ID: `stream-live-inputs-enable-a-live-input`

Allows a live input to be streamed to and makes the live input accessible to any future API calls.

## Definition

```yaml
{"operationId": "stream-live-inputs-enable-a-live-input", "summary": "Enable a live input", "description": "Allows a live input to be streamed to and makes the live input accessible to any future API calls.", "parameters": [{"name": "live_input_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_live_input_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier-2"}}], "responses": {"200": {"description": "Enable a live input response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_live_input_response_single"}}}}, "4XX": {"description": "Enable a live input response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Live Inputs"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.live-inputs.enable", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
