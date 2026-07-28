---
title: Update a live input
page_id: operation-put-accounts-account-id-stream-live-inputs-live-input-identifier-f6679742
path: operations/stream-live-inputs
description: Updates a specified live input.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/stream/live_inputs/{live_input_identifier}
operation_ids:
    - stream-live-inputs-update-a-live-input
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a live input

`PUT /accounts/{account_id}/stream/live_inputs/{live_input_identifier}`

Operation ID: `stream-live-inputs-update-a-live-input`

Updates a specified live input.

## Definition

```yaml
{"operationId": "stream-live-inputs-update-a-live-input", "summary": "Update a live input", "description": "Updates a specified live input.", "parameters": [{"name": "live_input_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_live_input_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_update_input_request"}}}}, "responses": {"200": {"description": "Update a live input response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_live_input_response_single"}}}}, "4XX": {"description": "Update a live input response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Live Inputs"], "x-api-token-group": ["Stream Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.live-inputs", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
