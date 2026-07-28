---
title: Delete a live input
page_id: operation-delete-accounts-account-id-stream-live-inputs-live-input-identifier-e43f3d7b
path: operations/stream-live-inputs
description: Prevents a live input from being streamed to and makes the live input inaccessible to any future API calls.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/stream/live_inputs/{live_input_identifier}
operation_ids:
    - stream-live-inputs-delete-a-live-input
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a live input

`DELETE /accounts/{account_id}/stream/live_inputs/{live_input_identifier}`

Operation ID: `stream-live-inputs-delete-a-live-input`

Prevents a live input from being streamed to and makes the live input inaccessible to any future API calls.

## Definition

```yaml
{"operationId": "stream-live-inputs-delete-a-live-input", "summary": "Delete a live input", "description": "Prevents a live input from being streamed to and makes the live input inaccessible to any future API calls.", "parameters": [{"name": "live_input_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_live_input_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete a live input response.", "content": {"application/json": {}}}, "4XX": {"description": "Delete a live input response failure.", "content": {"application/json": {}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Live Inputs"], "x-api-token-group": ["Stream Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.live-inputs", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
