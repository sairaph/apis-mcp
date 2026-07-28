---
title: List all outputs associated with a specified live input
page_id: operation-get-accounts-account-id-stream-live-inputs-live-input-identifier-outputs-922c0d6f
path: operations/stream-live-inputs
description: Retrieves all outputs associated with a specified live input.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/stream/live_inputs/{live_input_identifier}/outputs
operation_ids:
    - stream-live-inputs-list-all-outputs-associated-with-a-specified-live-input
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List all outputs associated with a specified live input

`GET /accounts/{account_id}/stream/live_inputs/{live_input_identifier}/outputs`

Operation ID: `stream-live-inputs-list-all-outputs-associated-with-a-specified-live-input`

Retrieves all outputs associated with a specified live input.

## Definition

```yaml
{"operationId": "stream-live-inputs-list-all-outputs-associated-with-a-specified-live-input", "summary": "List all outputs associated with a specified live input", "description": "Retrieves all outputs associated with a specified live input.", "parameters": [{"name": "live_input_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_live_input_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier-2"}}], "responses": {"200": {"description": "List all outputs associated with a specified live input response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_output_response_collection"}}}}, "4XX": {"description": "List all outputs associated with a specified live input response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Live Inputs"], "x-api-token-group": ["Stream Write", "Stream Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.live-inputs.outputs", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
