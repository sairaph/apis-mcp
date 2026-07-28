---
title: Create Waiting Room Rule
page_id: operation-post-zones-zone-id-waiting-rooms-waiting-room-id-rules-2c0b1ca2
path: operations/waiting-room
description: Only available for the Waiting Room Advanced subscription. Creates a rule for a waiting room.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/waiting_rooms/{waiting_room_id}/rules
operation_ids:
    - waiting-room-create-waiting-room-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Waiting Room Rule

`POST /zones/{zone_id}/waiting_rooms/{waiting_room_id}/rules`

Operation ID: `waiting-room-create-waiting-room-rule`

Only available for the Waiting Room Advanced subscription. Creates a rule for a waiting room.

## Definition

```yaml
{"operationId": "waiting-room-create-waiting-room-rule", "summary": "Create Waiting Room Rule", "description": "Only available for the Waiting Room Advanced subscription. Creates a rule for a waiting room.", "parameters": [{"name": "waiting_room_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_waiting_room_id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waitingroom_create_rule"}}}}, "responses": {"200": {"description": "Create Waiting Room Rule response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waitingroom_rules_response_collection"}}}}, "4XX": {"description": "Create Waiting Room Rule response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/waitingroom_rules_response_collection"}, {"$ref": "#/components/schemas/waitingroom_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Waiting Room"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "waiting-rooms.rules", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
