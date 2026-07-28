---
title: Delete an output
page_id: operation-delete-accounts-account-id-stream-live-inputs-live-input-identifier-outp-da4b9704
path: operations/stream-live-inputs
description: Deletes an output and removes it from the associated live input.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/stream/live_inputs/{live_input_identifier}/outputs/{output_identifier}
operation_ids:
    - stream-live-inputs-delete-an-output
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete an output

`DELETE /accounts/{account_id}/stream/live_inputs/{live_input_identifier}/outputs/{output_identifier}`

Operation ID: `stream-live-inputs-delete-an-output`

Deletes an output and removes it from the associated live input.

## Definition

```yaml
{"operationId": "stream-live-inputs-delete-an-output", "summary": "Delete an output", "description": "Deletes an output and removes it from the associated live input.", "parameters": [{"name": "output_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_output_identifier"}}, {"name": "live_input_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_live_input_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete an output response.", "content": {"application/json": {}}}, "4XX": {"description": "Delete an output response failure.", "content": {"application/json": {}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Live Inputs"], "x-api-token-group": ["Stream Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.live-inputs.outputs", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
