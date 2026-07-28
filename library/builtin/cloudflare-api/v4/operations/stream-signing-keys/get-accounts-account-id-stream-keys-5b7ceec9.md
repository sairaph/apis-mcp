---
title: List signing keys
page_id: operation-get-accounts-account-id-stream-keys-17945ea9
path: operations/stream-signing-keys
description: Lists the video ID and creation date and time when a signing key was created.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/stream/keys
operation_ids:
    - stream-signing-keys-list-signing-keys
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List signing keys

`GET /accounts/{account_id}/stream/keys`

Operation ID: `stream-signing-keys-list-signing-keys`

Lists the video ID and creation date and time when a signing key was created.

## Definition

```yaml
{"operationId": "stream-signing-keys-list-signing-keys", "summary": "List signing keys", "description": "Lists the video ID and creation date and time when a signing key was created.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier-2"}}], "responses": {"200": {"description": "List signing keys response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_key_response_collection"}}}}, "4XX": {"description": "List signing keys response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Signing Keys"], "x-api-token-group": ["Stream Write", "Stream Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.keys", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
