---
title: List live inputs
page_id: operation-get-accounts-account-id-stream-live-inputs-38db2865
path: operations/stream-live-inputs
description: Lists the live inputs created for an account. To get the credentials needed to stream to a specific live input, request a single live input.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/stream/live_inputs
operation_ids:
    - stream-live-inputs-list-live-inputs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List live inputs

`GET /accounts/{account_id}/stream/live_inputs`

Operation ID: `stream-live-inputs-list-live-inputs`

Lists the live inputs created for an account. To get the credentials needed to stream to a specific live input, request a single live input.

## Definition

```yaml
{"operationId": "stream-live-inputs-list-live-inputs", "summary": "List live inputs", "description": "Lists the live inputs created for an account. To get the credentials needed to stream to a specific live input, request a single live input.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier-2"}}, {"name": "include_counts", "in": "query", "schema": {"$ref": "#/components/schemas/stream_include_counts"}}], "responses": {"200": {"description": "List live inputs response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_live_input_response_collection"}}}}, "4XX": {"description": "List live inputs response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Live Inputs"], "x-api-token-group": ["Stream Write", "Stream Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.live-inputs", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
