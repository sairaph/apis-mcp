---
title: Create signing keys
page_id: operation-post-accounts-account-id-stream-keys-6c6778d1
path: operations/stream-signing-keys
description: Creates an RSA private key in PEM and JWK formats. Key files are only displayed once after creation. Keys are created, used, and deleted independently of videos, and every key can sign any video.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/stream/keys
operation_ids:
    - stream-signing-keys-create-signing-keys
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create signing keys

`POST /accounts/{account_id}/stream/keys`

Operation ID: `stream-signing-keys-create-signing-keys`

Creates an RSA private key in PEM and JWK formats. Key files are only displayed once after creation. Keys are created, used, and deleted independently of videos, and every key can sign any video.

## Definition

```yaml
{"operationId": "stream-signing-keys-create-signing-keys", "summary": "Create signing keys", "description": "Creates an RSA private key in PEM and JWK formats. Key files are only displayed once after creation. Keys are created, used, and deleted independently of videos, and every key can sign any video.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Create signing keys response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_key_generation_response"}}}}, "4XX": {"description": "Create signing keys response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Signing Keys"], "x-api-token-group": ["Stream Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.keys", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
