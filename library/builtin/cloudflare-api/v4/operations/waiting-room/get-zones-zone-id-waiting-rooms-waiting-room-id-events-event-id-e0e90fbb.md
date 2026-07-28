---
title: Event details
page_id: operation-get-zones-zone-id-waiting-rooms-waiting-room-id-events-event-id-f0a33146
path: operations/waiting-room
description: Fetches a single configured event for a waiting room.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/waiting_rooms/{waiting_room_id}/events/{event_id}
operation_ids:
    - waiting-room-event-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Event details

`GET /zones/{zone_id}/waiting_rooms/{waiting_room_id}/events/{event_id}`

Operation ID: `waiting-room-event-details`

Fetches a single configured event for a waiting room.

## Definition

```yaml
{"operationId": "waiting-room-event-details", "summary": "Event details", "description": "Fetches a single configured event for a waiting room.", "parameters": [{"name": "event_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_event_id"}}, {"name": "waiting_room_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_waiting_room_id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_identifier"}}], "responses": {"200": {"description": "Event details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waitingroom_event_response"}}}}, "4XX": {"description": "Event details response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/waitingroom_event_response"}, {"$ref": "#/components/schemas/waitingroom_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Waiting Room"], "x-api-token-group": ["Waiting Rooms Read", "Waiting Rooms Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "waiting-rooms.events", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
