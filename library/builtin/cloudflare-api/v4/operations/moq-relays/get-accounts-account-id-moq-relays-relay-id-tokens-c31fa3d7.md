---
title: List tokens
page_id: operation-get-accounts-account-id-moq-relays-relay-id-tokens-aad897d0
path: operations/moq-relays
description: |-
    Returns metadata for every token in the relay's registry. Secrets are
    never returned. The dashboard derives an `expired` flag by comparing
    each token's `expires` to the current time.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/moq/relays/{relay_id}/tokens
operation_ids:
    - moq-relays-tokens-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List tokens

`GET /accounts/{account_id}/moq/relays/{relay_id}/tokens`

Operation ID: `moq-relays-tokens-list`

Returns metadata for every token in the relay's registry. Secrets are
never returned. The dashboard derives an `expired` flag by comparing
each token's `expires` to the current time.

## Definition

```yaml
{"operationId": "moq-relays-tokens-list", "summary": "List tokens", "description": "Returns metadata for every token in the relay's registry. Secrets are\nnever returned. The dashboard derives an `expired` flag by comparing\neach token's `expires` to the current time.\n", "parameters": [{"$ref": "#/components/parameters/moq_account_id"}, {"$ref": "#/components/parameters/moq_relay_id"}], "responses": {"200": {"description": "Token metadata (no secrets).", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/moq_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/moq_tokens_envelope"}}, "type": "object"}]}}}}, "400": {"description": "Error 21003: Relay ID should be 32 hex characters.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/moq_api-response-error"}}}}, "404": {"description": "Error 21007: Relay not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/moq_api-response-error"}}}}, "500": {"description": "Error 21006: Unexpected server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/moq_api-response-error"}}}}}, "security": [{"api_token": []}], "tags": ["MoQ Relays"], "x-stability": "beta"}
```
