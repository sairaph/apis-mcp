---
title: Retrieve TURN key details
page_id: operation-get-accounts-account-id-calls-turn-keys-key-id-58ea9eef
path: operations/calls-turn-keys
description: Fetches details for a single TURN key.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/calls/turn_keys/{key_id}
operation_ids:
    - calls-retrieve-turn-key-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve TURN key details

`GET /accounts/{account_id}/calls/turn_keys/{key_id}`

Operation ID: `calls-retrieve-turn-key-details`

Fetches details for a single TURN key.

## Definition

```yaml
{"operationId": "calls-retrieve-turn-key-details", "summary": "Retrieve TURN key details", "description": "Fetches details for a single TURN key.", "parameters": [{"name": "key_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/calls_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/calls_account_identifier"}}], "responses": {"200": {"description": "Retrieve TURN key details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/calls_turn_key_response_single"}}}}, "4XX": {"description": "Retrieve TURN key details failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/calls_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["Calls TURN Keys"], "x-api-token-group": ["Calls Write", "Calls Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "calls.turn", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
