---
title: List events
page_id: operation-get-zones-zone-id-waiting-rooms-waiting-room-id-events-f4b51af6
path: operations/waiting-room
description: Lists events for a waiting room.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/waiting_rooms/{waiting_room_id}/events
operation_ids:
    - waiting-room-list-events
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List events

`GET /zones/{zone_id}/waiting_rooms/{waiting_room_id}/events`

Operation ID: `waiting-room-list-events`

Lists events for a waiting room.

## Definition

```yaml
{"operationId": "waiting-room-list-events", "summary": "List events", "description": "Lists events for a waiting room.", "parameters": [{"name": "waiting_room_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_waiting_room_id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_identifier"}}, {"$ref": "#/components/parameters/waitingroom_page"}, {"$ref": "#/components/parameters/waitingroom_per_page"}], "responses": {"200": {"description": "List events response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waitingroom_event_response_collection"}}}}, "4XX": {"description": "List events response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/waitingroom_event_response_collection"}, {"$ref": "#/components/schemas/waitingroom_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Waiting Room"], "x-api-token-group": ["Waiting Rooms Read", "Waiting Rooms Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "waiting-rooms.events", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
