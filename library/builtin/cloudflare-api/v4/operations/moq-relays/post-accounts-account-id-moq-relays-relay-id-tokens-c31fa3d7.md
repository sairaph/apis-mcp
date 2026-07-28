---
title: Create a token
page_id: operation-post-accounts-account-id-moq-relays-relay-id-tokens-f6727d16
path: operations/moq-relays
description: |-
    Mints a new relay-scoped token and adds it to the relay's accepted-auth
    registry. The token value (secret) is shown once in the response. A relay
    may hold up to 10 tokens; creating an 11th is rejected.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/moq/relays/{relay_id}/tokens
operation_ids:
    - moq-relays-tokens-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a token

`POST /accounts/{account_id}/moq/relays/{relay_id}/tokens`

Operation ID: `moq-relays-tokens-create`

Mints a new relay-scoped token and adds it to the relay's accepted-auth
registry. The token value (secret) is shown once in the response. A relay
may hold up to 10 tokens; creating an 11th is rejected.

## Definition

```yaml
{"operationId": "moq-relays-tokens-create", "summary": "Create a token", "description": "Mints a new relay-scoped token and adds it to the relay's accepted-auth\nregistry. The token value (secret) is shown once in the response. A relay\nmay hold up to 10 tokens; creating an 11th is rejected.\n", "parameters": [{"$ref": "#/components/parameters/moq_account_id"}, {"$ref": "#/components/parameters/moq_relay_id"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"expires": {"description": "Optional expiry (RFC 3339). Defaults to 1 year from creation;\nrejected if more than 1 year in the future.\n", "type": "string", "format": "date-time", "example": "2027-03-27T15:00:00Z"}, "label": {"description": "Optional, customer-set label.", "type": "string", "example": "primary-encoder"}, "operations": {"description": "Non-empty subset of the V1 roles the token is allowed to\nperform. Signed into the token.\n", "type": "array", "items": {"enum": ["publish", "subscribe"], "type": "string"}, "example": ["publish", "subscribe"], "minItems": 1}}, "required": ["operations"]}}}}, "responses": {"201": {"description": "Token created. The secret is shown once.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/moq_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/moq_tokens_envelope"}}, "type": "object"}]}}}}, "400": {"description": "Bad request. Possible errors: 21003 (relay ID should be 32 hex characters), 21004 (failed to decode body, invalid JSON), 21010 (invalid operations), 21012 (expires more than 1 year out).", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/moq_api-response-error"}}}}, "404": {"description": "Error 21007: Relay not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/moq_api-response-error"}}}}, "409": {"description": "Error 21009: Token limit reached (a relay may hold at most 10 tokens).", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/moq_api-response-error"}}}}, "500": {"description": "Error 21006: Unexpected server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/moq_api-response-error"}}}}}, "security": [{"api_token": []}], "tags": ["MoQ Relays"], "x-stability": "beta"}
```
