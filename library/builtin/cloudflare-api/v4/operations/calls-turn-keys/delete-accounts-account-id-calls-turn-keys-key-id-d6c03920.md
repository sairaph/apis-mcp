---
title: Delete TURN key
page_id: operation-delete-accounts-account-id-calls-turn-keys-key-id-2f3a34b3
path: operations/calls-turn-keys
description: Deletes a TURN key from Cloudflare Calls
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/calls/turn_keys/{key_id}
operation_ids:
    - calls-delete-turn-key
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete TURN key

`DELETE /accounts/{account_id}/calls/turn_keys/{key_id}`

Operation ID: `calls-delete-turn-key`

Deletes a TURN key from Cloudflare Calls

## Definition

```yaml
{"operationId": "calls-delete-turn-key", "summary": "Delete TURN key", "description": "Deletes a TURN key from Cloudflare Calls", "parameters": [{"name": "key_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/calls_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/calls_account_identifier"}}], "responses": {"200": {"description": "Delete TURN key response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/calls_turn_key_response_single"}}}}, "4XX": {"description": "Delete TURN key response failure", "content": {"application/json": {}}}}, "security": [{"api_token": []}], "tags": ["Calls TURN Keys"], "x-api-token-group": ["Calls Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "calls.turn", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
