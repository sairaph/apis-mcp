---
title: Delete signing keys
page_id: operation-delete-accounts-account-id-stream-keys-identifier-506cd464
path: operations/stream-signing-keys
description: Deletes signing keys and revokes all signed URLs generated with the key.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/stream/keys/{identifier}
operation_ids:
    - stream-signing-keys-delete-signing-keys
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete signing keys

`DELETE /accounts/{account_id}/stream/keys/{identifier}`

Operation ID: `stream-signing-keys-delete-signing-keys`

Deletes signing keys and revokes all signed URLs generated with the key.

## Definition

```yaml
{"operationId": "stream-signing-keys-delete-signing-keys", "summary": "Delete signing keys", "description": "Deletes signing keys and revokes all signed URLs generated with the key.", "parameters": [{"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier-2"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete signing keys response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_deleted_response"}}}}, "4XX": {"description": "Delete signing keys response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Signing Keys"], "x-api-token-group": ["Stream Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.keys", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
