---
title: Preview active event details
page_id: operation-get-zones-zone-id-waiting-rooms-waiting-room-id-events-event-id-details-d971c69e
path: operations/waiting-room
description: Previews an event's configuration as if it was active. Inherited fields from the waiting room will be displayed with their current values.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/waiting_rooms/{waiting_room_id}/events/{event_id}/details
operation_ids:
    - waiting-room-preview-active-event-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Preview active event details

`GET /zones/{zone_id}/waiting_rooms/{waiting_room_id}/events/{event_id}/details`

Operation ID: `waiting-room-preview-active-event-details`

Previews an event's configuration as if it was active. Inherited fields from the waiting room will be displayed with their current values.

## Definition

```yaml
{"operationId": "waiting-room-preview-active-event-details", "summary": "Preview active event details", "description": "Previews an event's configuration as if it was active. Inherited fields from the waiting room will be displayed with their current values.", "parameters": [{"name": "event_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_event_id"}}, {"name": "waiting_room_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_waiting_room_id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_identifier"}}], "responses": {"200": {"description": "Preview active event details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waitingroom_event_details_response"}}}}, "4XX": {"description": "Preview active event details response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/waitingroom_event_details_response"}, {"$ref": "#/components/schemas/waitingroom_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Waiting Room"], "x-api-token-group": ["Waiting Rooms Read", "Waiting Rooms Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "waiting-rooms.events.details", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
