---
title: List Waiting Room Rules
page_id: operation-get-zones-zone-id-waiting-rooms-waiting-room-id-rules-0852a34f
path: operations/waiting-room
description: Lists rules for a waiting room.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/waiting_rooms/{waiting_room_id}/rules
operation_ids:
    - waiting-room-list-waiting-room-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Waiting Room Rules

`GET /zones/{zone_id}/waiting_rooms/{waiting_room_id}/rules`

Operation ID: `waiting-room-list-waiting-room-rules`

Lists rules for a waiting room.

## Definition

```yaml
{"operationId": "waiting-room-list-waiting-room-rules", "summary": "List Waiting Room Rules", "description": "Lists rules for a waiting room.", "parameters": [{"name": "waiting_room_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_waiting_room_id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_identifier"}}], "responses": {"200": {"description": "List Waiting Room Rules response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waitingroom_rules_response_collection"}}}}, "4XX": {"description": "List Waiting Room Rules response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/waitingroom_rules_response_collection"}, {"$ref": "#/components/schemas/waitingroom_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Waiting Room"], "x-api-token-group": ["Waiting Rooms Read", "Waiting Rooms Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "waiting-rooms.rules", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
