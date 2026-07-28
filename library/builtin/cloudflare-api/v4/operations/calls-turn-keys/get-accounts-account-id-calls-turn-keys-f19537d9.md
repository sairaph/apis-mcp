---
title: List TURN Keys
page_id: operation-get-accounts-account-id-calls-turn-keys-af079148
path: operations/calls-turn-keys
description: Lists all TURN keys in the Cloudflare account
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/calls/turn_keys
operation_ids:
    - calls-turn-key-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List TURN Keys

`GET /accounts/{account_id}/calls/turn_keys`

Operation ID: `calls-turn-key-list`

Lists all TURN keys in the Cloudflare account

## Definition

```yaml
{"operationId": "calls-turn-key-list", "summary": "List TURN Keys", "description": "Lists all TURN keys in the Cloudflare account", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/calls_account_identifier"}}], "responses": {"200": {"description": "List TURN key response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/calls_turn_key_collection"}}}}, "4XX": {"description": "List TURN key response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/calls_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["Calls TURN Keys"], "x-api-token-group": ["Calls Write", "Calls Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "calls.turn", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
