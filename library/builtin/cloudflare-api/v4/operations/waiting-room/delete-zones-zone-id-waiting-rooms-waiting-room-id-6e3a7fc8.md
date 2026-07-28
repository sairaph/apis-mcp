---
title: Delete waiting room
page_id: operation-delete-zones-zone-id-waiting-rooms-waiting-room-id-9e0eb12b
path: operations/waiting-room
description: Deletes a waiting room.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/waiting_rooms/{waiting_room_id}
operation_ids:
    - waiting-room-delete-waiting-room
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete waiting room

`DELETE /zones/{zone_id}/waiting_rooms/{waiting_room_id}`

Operation ID: `waiting-room-delete-waiting-room`

Deletes a waiting room.

## Definition

```yaml
{"operationId": "waiting-room-delete-waiting-room", "summary": "Delete waiting room", "description": "Deletes a waiting room.", "parameters": [{"name": "waiting_room_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_waiting_room_id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete waiting room response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waitingroom_waiting_room_id_response"}}}}, "4XX": {"description": "Delete waiting room response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/waitingroom_waiting_room_id_response"}, {"$ref": "#/components/schemas/waitingroom_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Waiting Room"], "x-api-token-group": ["Waiting Rooms Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "waiting-rooms", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
