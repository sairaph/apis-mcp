---
title: Retrieve a live input
page_id: operation-get-accounts-account-id-stream-live-inputs-live-input-identifier-b2563d66
path: operations/stream-live-inputs
description: Retrieves details of an existing live input.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/stream/live_inputs/{live_input_identifier}
operation_ids:
    - stream-live-inputs-retrieve-a-live-input
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve a live input

`GET /accounts/{account_id}/stream/live_inputs/{live_input_identifier}`

Operation ID: `stream-live-inputs-retrieve-a-live-input`

Retrieves details of an existing live input.

## Definition

```yaml
{"operationId": "stream-live-inputs-retrieve-a-live-input", "summary": "Retrieve a live input", "description": "Retrieves details of an existing live input.", "parameters": [{"name": "live_input_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_live_input_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier-2"}}], "responses": {"200": {"description": "Retrieve a live input response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_live_input_response_single"}}}}, "4XX": {"description": "Retrieve a live input response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Live Inputs"], "x-api-token-group": ["Stream Write", "Stream Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.live-inputs", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
