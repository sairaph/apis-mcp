---
title: Get waiting room status
page_id: operation-get-zones-zone-id-waiting-rooms-waiting-room-id-status-6cd0b724
path: operations/waiting-room
description: |-
    Fetches the status of a configured waiting room. Response fields include:
    1. `status`: String indicating the status of the waiting room. The possible status are:
        - **not_queueing** indicates that the configured thresholds have not been met and all users are going through to the origin.
        - **queueing** indicates that the thresholds have been met and some users are held in the waiting room.
        - **event_prequeueing** indicates that an event is active and is currently prequeueing users before it starts.
        - **suspended** indicates that the room is suspended.
    2. `event_id`: String of the current event's `id` if an event is active, otherwise an empty string.
    3. `estimated_queued_users`: Integer of the estimated number of users currently waiting in the queue.
    4. `estimated_total_active_users`: Integer of the estimated number of users currently active on the origin.
    5. `max_estimated_time_minutes`: Integer of the maximum estimated time currently presented to the users.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/waiting_rooms/{waiting_room_id}/status
operation_ids:
    - waiting-room-get-waiting-room-status
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get waiting room status

`GET /zones/{zone_id}/waiting_rooms/{waiting_room_id}/status`

Operation ID: `waiting-room-get-waiting-room-status`

Fetches the status of a configured waiting room. Response fields include:
1. `status`: String indicating the status of the waiting room. The possible status are:
    - **not_queueing** indicates that the configured thresholds have not been met and all users are going through to the origin.
    - **queueing** indicates that the thresholds have been met and some users are held in the waiting room.
    - **event_prequeueing** indicates that an event is active and is currently prequeueing users before it starts.
    - **suspended** indicates that the room is suspended.
2. `event_id`: String of the current event's `id` if an event is active, otherwise an empty string.
3. `estimated_queued_users`: Integer of the estimated number of users currently waiting in the queue.
4. `estimated_total_active_users`: Integer of the estimated number of users currently active on the origin.
5. `max_estimated_time_minutes`: Integer of the maximum estimated time currently presented to the users.

## Definition

```yaml
{"operationId": "waiting-room-get-waiting-room-status", "summary": "Get waiting room status", "description": "Fetches the status of a configured waiting room. Response fields include:\n1. `status`: String indicating the status of the waiting room. The possible status are:\n\t- **not_queueing** indicates that the configured thresholds have not been met and all users are going through to the origin.\n\t- **queueing** indicates that the thresholds have been met and some users are held in the waiting room.\n\t- **event_prequeueing** indicates that an event is active and is currently prequeueing users before it starts.\n\t- **suspended** indicates that the room is suspended.\n2. `event_id`: String of the current event's `id` if an event is active, otherwise an empty string.\n3. `estimated_queued_users`: Integer of the estimated number of users currently waiting in the queue.\n4. `estimated_total_active_users`: Integer of the estimated number of users currently active on the origin.\n5. `max_estimated_time_minutes`: Integer of the maximum estimated time currently presented to the users.", "parameters": [{"name": "waiting_room_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_waiting_room_id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_identifier"}}], "responses": {"200": {"description": "Get waiting room status response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waitingroom_status_response"}}}}, "4XX": {"description": "Get waiting room status response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/waitingroom_status_response"}, {"$ref": "#/components/schemas/waitingroom_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Waiting Room"], "x-api-token-group": ["Waiting Rooms Read", "Waiting Rooms Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "waiting-rooms.statuses", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
