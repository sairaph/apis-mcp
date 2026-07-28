---
title: Update an output
page_id: operation-put-accounts-account-id-stream-live-inputs-live-input-identifier-outputs-dfe4edb4
path: operations/stream-live-inputs
description: Updates the state of an output.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/stream/live_inputs/{live_input_identifier}/outputs/{output_identifier}
operation_ids:
    - stream-live-inputs-update-an-output
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update an output

`PUT /accounts/{account_id}/stream/live_inputs/{live_input_identifier}/outputs/{output_identifier}`

Operation ID: `stream-live-inputs-update-an-output`

Updates the state of an output.

## Definition

```yaml
{"operationId": "stream-live-inputs-update-an-output", "summary": "Update an output", "description": "Updates the state of an output.", "parameters": [{"name": "output_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_output_identifier"}}, {"name": "live_input_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_live_input_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_update_output_request"}}}}, "responses": {"200": {"description": "Update an output response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_output_response_single"}}}}, "4XX": {"description": "Update an output response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Live Inputs"], "x-api-token-group": ["Stream Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.live-inputs.outputs", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
