---
title: Rotate keys for a live input
page_id: operation-post-accounts-account-id-stream-live-inputs-live-input-identifier-rotate-ed8f146d
path: operations/stream-live-inputs
description: Rotates the credentials for a live input without changing its identifier. Old credentials are revoked, broadcasts using stale credentials are automatically disconnected shortly after rotation, and the response returns refreshed credentials.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/stream/live_inputs/{live_input_identifier}/rotate_keys
operation_ids:
    - stream-live-inputs-rotate-keys-for-a-live-input
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Rotate keys for a live input

`POST /accounts/{account_id}/stream/live_inputs/{live_input_identifier}/rotate_keys`

Operation ID: `stream-live-inputs-rotate-keys-for-a-live-input`

Rotates the credentials for a live input without changing its identifier. Old credentials are revoked, broadcasts using stale credentials are automatically disconnected shortly after rotation, and the response returns refreshed credentials.

## Definition

```yaml
{"operationId": "stream-live-inputs-rotate-keys-for-a-live-input", "summary": "Rotate keys for a live input", "description": "Rotates the credentials for a live input without changing its identifier. Old credentials are revoked, broadcasts using stale credentials are automatically disconnected shortly after rotation, and the response returns refreshed credentials.", "parameters": [{"name": "live_input_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_live_input_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier-2"}}], "responses": {"200": {"description": "Rotate keys for a live input response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_live_input_response_single"}}}}, "4XX": {"description": "Rotate keys for a live input response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Live Inputs"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.live-inputs.rotate", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
