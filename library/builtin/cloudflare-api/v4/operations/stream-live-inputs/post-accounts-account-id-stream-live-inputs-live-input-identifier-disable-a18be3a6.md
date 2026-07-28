---
title: Disable a live input
page_id: operation-post-accounts-account-id-stream-live-inputs-live-input-identifier-disabl-cfc1c63b
path: operations/stream-live-inputs
description: Prevents a live input from being streamed to and makes the live input inaccessible to any future API calls until enabled.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/stream/live_inputs/{live_input_identifier}/disable
operation_ids:
    - stream-live-inputs-disable-a-live-input
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Disable a live input

`POST /accounts/{account_id}/stream/live_inputs/{live_input_identifier}/disable`

Operation ID: `stream-live-inputs-disable-a-live-input`

Prevents a live input from being streamed to and makes the live input inaccessible to any future API calls until enabled.

## Definition

```yaml
{"operationId": "stream-live-inputs-disable-a-live-input", "summary": "Disable a live input", "description": "Prevents a live input from being streamed to and makes the live input inaccessible to any future API calls until enabled.", "parameters": [{"name": "live_input_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_live_input_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier-2"}}], "responses": {"200": {"description": "Disable a live input response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_live_input_response_single"}}}}, "4XX": {"description": "Disable a live input response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Live Inputs"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.live-inputs.disable", "x-fern-sdk-method-name": "create", "x-forge-hidden": true, "x-forge-require-confirmation": "This operation disables a live input. Active streams to it are dropped."}
```
