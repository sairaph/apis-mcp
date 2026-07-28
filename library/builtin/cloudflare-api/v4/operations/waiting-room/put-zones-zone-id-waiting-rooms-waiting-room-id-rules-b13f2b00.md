---
title: Replace Waiting Room Rules
page_id: operation-put-zones-zone-id-waiting-rooms-waiting-room-id-rules-a16b1444
path: operations/waiting-room
description: Only available for the Waiting Room Advanced subscription. Replaces all rules for a waiting room.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/waiting_rooms/{waiting_room_id}/rules
operation_ids:
    - waiting-room-replace-waiting-room-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Replace Waiting Room Rules

`PUT /zones/{zone_id}/waiting_rooms/{waiting_room_id}/rules`

Operation ID: `waiting-room-replace-waiting-room-rules`

Only available for the Waiting Room Advanced subscription. Replaces all rules for a waiting room.

## Definition

```yaml
{"operationId": "waiting-room-replace-waiting-room-rules", "summary": "Replace Waiting Room Rules", "description": "Only available for the Waiting Room Advanced subscription. Replaces all rules for a waiting room.", "parameters": [{"name": "waiting_room_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_waiting_room_id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waitingroom_update_rules"}}}}, "responses": {"200": {"description": "Replace Waiting Room Rules response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waitingroom_rules_response_collection"}}}}, "4XX": {"description": "Replace Waiting Room Rules response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/waitingroom_rules_response_collection"}, {"$ref": "#/components/schemas/waitingroom_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Waiting Room"], "x-api-token-group": ["Waiting Rooms Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "waiting-rooms.rules", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
