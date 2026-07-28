---
title: Create event
page_id: operation-post-zones-zone-id-waiting-rooms-waiting-room-id-events-3c9983a5
path: operations/waiting-room
description: Only available for the Waiting Room Advanced subscription. Creates an event for a waiting room. An event takes place during a specified period of time, temporarily changing the behavior of a waiting room. While the event is active, some of the properties in the event's configuration may either override or inherit from the waiting room's configuration. Note that events cannot overlap with each other, so only one event can be active at a time.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/waiting_rooms/{waiting_room_id}/events
operation_ids:
    - waiting-room-create-event
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create event

`POST /zones/{zone_id}/waiting_rooms/{waiting_room_id}/events`

Operation ID: `waiting-room-create-event`

Only available for the Waiting Room Advanced subscription. Creates an event for a waiting room. An event takes place during a specified period of time, temporarily changing the behavior of a waiting room. While the event is active, some of the properties in the event's configuration may either override or inherit from the waiting room's configuration. Note that events cannot overlap with each other, so only one event can be active at a time.

## Definition

```yaml
{"operationId": "waiting-room-create-event", "summary": "Create event", "description": "Only available for the Waiting Room Advanced subscription. Creates an event for a waiting room. An event takes place during a specified period of time, temporarily changing the behavior of a waiting room. While the event is active, some of the properties in the event's configuration may either override or inherit from the waiting room's configuration. Note that events cannot overlap with each other, so only one event can be active at a time.", "parameters": [{"name": "waiting_room_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_waiting_room_id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waitingroom_query_event"}}}}, "responses": {"200": {"description": "Create event response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waitingroom_event_response"}}}}, "4XX": {"description": "Create event response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/waitingroom_event_response"}, {"$ref": "#/components/schemas/waitingroom_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Waiting Room"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "waiting-rooms.events", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
