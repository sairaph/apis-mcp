---
title: Create a new TURN key
page_id: operation-post-accounts-account-id-calls-turn-keys-bb67ab20
path: operations/calls-turn-keys
description: Creates a new Cloudflare Calls TURN key.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/calls/turn_keys
operation_ids:
    - calls-turn-key-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a new TURN key

`POST /accounts/{account_id}/calls/turn_keys`

Operation ID: `calls-turn-key-create`

Creates a new Cloudflare Calls TURN key.

## Definition

```yaml
{"operationId": "calls-turn-key-create", "summary": "Create a new TURN key", "description": "Creates a new Cloudflare Calls TURN key.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/calls_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/calls_turn_key_editable_fields"}}}}, "responses": {"201": {"description": "Created a new TURN key", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/calls_turn_key_single_with_secret"}}}}}, "security": [{"api_token": []}], "tags": ["Calls TURN Keys"], "x-api-token-group": ["Calls Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "calls.turn", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
