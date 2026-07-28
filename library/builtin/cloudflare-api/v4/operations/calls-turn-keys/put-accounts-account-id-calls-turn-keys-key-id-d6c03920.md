---
title: Edit TURN key details
page_id: operation-put-accounts-account-id-calls-turn-keys-key-id-dc28875f
path: operations/calls-turn-keys
description: Edit details for a single TURN key.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/calls/turn_keys/{key_id}
operation_ids:
    - calls-update-turn-key
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Edit TURN key details

`PUT /accounts/{account_id}/calls/turn_keys/{key_id}`

Operation ID: `calls-update-turn-key`

Edit details for a single TURN key.

## Definition

```yaml
{"operationId": "calls-update-turn-key", "summary": "Edit TURN key details", "description": "Edit details for a single TURN key.", "parameters": [{"name": "key_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/calls_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/calls_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/calls_turn_key_editable_fields"}}}}, "responses": {"200": {"description": "Edit TURN key details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/calls_turn_key_response_single"}}}}, "4XX": {"description": "Edit TURN key details response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/calls_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["Calls TURN Keys"], "x-api-token-group": ["Calls Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "calls.turn", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
