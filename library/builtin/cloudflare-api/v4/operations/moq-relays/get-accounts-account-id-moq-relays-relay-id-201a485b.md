---
title: Get a relay
page_id: operation-get-accounts-account-id-moq-relays-relay-id-ec3bf83c
path: operations/moq-relays
description: |-
    Retrieves a single MoQ relay including config and status.
    Tokens are NOT included.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/moq/relays/{relay_id}
operation_ids:
    - moq-relays-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a relay

`GET /accounts/{account_id}/moq/relays/{relay_id}`

Operation ID: `moq-relays-get`

Retrieves a single MoQ relay including config and status.
Tokens are NOT included.

## Definition

```yaml
{"operationId": "moq-relays-get", "summary": "Get a relay", "description": "Retrieves a single MoQ relay including config and status.\nTokens are NOT included.\n", "parameters": [{"$ref": "#/components/parameters/moq_account_id"}, {"$ref": "#/components/parameters/moq_relay_id"}], "responses": {"200": {"description": "Relay retrieved successfully.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/moq_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/moq_relay"}}, "type": "object"}]}}}}, "400": {"description": "Error 21003: Relay ID should be 32 hex characters.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/moq_api-response-error"}}}}, "404": {"description": "Error 21007: Relay not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/moq_api-response-error"}}}}, "500": {"description": "Error 21006: Unexpected server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/moq_api-response-error"}}}}}, "security": [{"api_token": []}], "tags": ["MoQ Relays"], "x-stability": "beta"}
```
