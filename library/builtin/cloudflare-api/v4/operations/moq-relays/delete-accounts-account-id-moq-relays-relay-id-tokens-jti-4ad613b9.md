---
title: Revoke a token
page_id: operation-delete-accounts-account-id-moq-relays-relay-id-tokens-jti-6bd171c5
path: operations/moq-relays
description: |-
    Revokes a token by removing it from the relay's registry. crique rejects
    the token within the cache TTL. Idempotent — revoking an unknown token
    succeeds.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/moq/relays/{relay_id}/tokens/{jti}
operation_ids:
    - moq-relays-tokens-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Revoke a token

`DELETE /accounts/{account_id}/moq/relays/{relay_id}/tokens/{jti}`

Operation ID: `moq-relays-tokens-delete`

Revokes a token by removing it from the relay's registry. crique rejects
the token within the cache TTL. Idempotent — revoking an unknown token
succeeds.

## Definition

```yaml
{"operationId": "moq-relays-tokens-delete", "summary": "Revoke a token", "description": "Revokes a token by removing it from the relay's registry. crique rejects\nthe token within the cache TTL. Idempotent — revoking an unknown token\nsucceeds.\n", "parameters": [{"$ref": "#/components/parameters/moq_account_id"}, {"$ref": "#/components/parameters/moq_relay_id"}, {"$ref": "#/components/parameters/moq_jti"}], "responses": {"200": {"description": "Token revoked.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/moq_api-response-common"}}}}, "400": {"description": "Bad request. Possible errors: 21003 (relay ID should be 32 hex characters), invalid token id.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/moq_api-response-error"}}}}, "404": {"description": "Error 21007: Relay not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/moq_api-response-error"}}}}, "500": {"description": "Error 21006: Unexpected server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/moq_api-response-error"}}}}}, "security": [{"api_token": []}], "tags": ["MoQ Relays"], "x-stability": "beta"}
```
