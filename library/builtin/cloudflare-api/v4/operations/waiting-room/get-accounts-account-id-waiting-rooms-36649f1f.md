---
title: List waiting rooms for account
page_id: operation-get-accounts-account-id-waiting-rooms-51453666
path: operations/waiting-room
description: Lists waiting rooms for account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/waiting_rooms
operation_ids:
    - waiting-room-list-waiting-rooms-account
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List waiting rooms for account

`GET /accounts/{account_id}/waiting_rooms`

Operation ID: `waiting-room-list-waiting-rooms-account`

Lists waiting rooms for account.

## Definition

```yaml
{"operationId": "waiting-room-list-waiting-rooms-account", "summary": "List waiting rooms for account", "description": "Lists waiting rooms for account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_identifier"}}, {"$ref": "#/components/parameters/waitingroom_page"}, {"$ref": "#/components/parameters/waitingroom_per_page"}], "responses": {"200": {"description": "List waiting rooms for account response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waitingroom_response_collection"}}}}, "4XX": {"description": "List waiting rooms for account response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/waitingroom_response_collection"}, {"$ref": "#/components/schemas/waitingroom_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Waiting Room"], "x-api-token-group": ["Account Waiting Rooms Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "waiting-rooms.account-waiting-rooms", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
