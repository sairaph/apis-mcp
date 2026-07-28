---
title: Create a custom waiting room page preview
page_id: operation-post-zones-zone-id-waiting-rooms-preview-b9778a96
path: operations/waiting-room
description: "Creates a waiting room page preview. Upload a custom waiting room page for preview. You will receive a preview URL in the form `http://waitingrooms.dev/preview/<uuid>`. You can use the following query parameters to change the state of the preview:\n1. `force_queue`: Boolean indicating if all users will be queued in the waiting room and no one will be let into the origin website (also known as queueAll).\n2. `queue_is_full`: Boolean indicating if the waiting room's queue is currently full and not accepting new users at the moment.\n3. `queueing_method`: The queueing method currently used by the waiting room.\n\t- **fifo** indicates a FIFO queue.\n\t- **random** indicates a Random queue.\n\t- **passthrough** indicates a Passthrough queue. Keep in mind that the waiting room page will only be displayed if `force_queue=true` or `event=prequeueing` — for other cases the request will pass through to the origin. For our preview, this will be a fake origin website returning \\\"Welcome\\\". \n\t- **reject** indicates a Reject queue.\n4. `event`: Used to preview a waiting room event.\n\t- **none** indicates no event is occurring.\n\t- **prequeueing** indicates that an event is prequeueing (between `prequeue_start_time` and `event_start_time`).\n\t- **started** indicates that an event has started (between `event_start_time` and `event_end_time`).\n5. `shuffle_at_event_start`: Boolean indicating if the event will shuffle users in the prequeue when it starts. This can only be set to **true** if an event is active (`event` is not **none**).\n\nFor example, you can make a request to `http://waitingrooms.dev/preview/<uuid>?force_queue=false&queue_is_full=false&queueing_method=random&event=started&shuffle_at_event_start=true`\n6. `waitTime`: Non-zero, positive integer indicating the estimated wait time in minutes. The default value is 10 minutes.\n\nFor example, you can make a request to `http://waitingrooms.dev/preview/<uuid>?waitTime=50` to configure the estimated wait time as 50 minutes."
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/waiting_rooms/preview
operation_ids:
    - waiting-room-create-a-custom-waiting-room-page-preview
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a custom waiting room page preview

`POST /zones/{zone_id}/waiting_rooms/preview`

Operation ID: `waiting-room-create-a-custom-waiting-room-page-preview`

Creates a waiting room page preview. Upload a custom waiting room page for preview. You will receive a preview URL in the form `http://waitingrooms.dev/preview/<uuid>`. You can use the following query parameters to change the state of the preview:
1. `force_queue`: Boolean indicating if all users will be queued in the waiting room and no one will be let into the origin website (also known as queueAll).
2. `queue_is_full`: Boolean indicating if the waiting room's queue is currently full and not accepting new users at the moment.
3. `queueing_method`: The queueing method currently used by the waiting room.
	- **fifo** indicates a FIFO queue.
	- **random** indicates a Random queue.
	- **passthrough** indicates a Passthrough queue. Keep in mind that the waiting room page will only be displayed if `force_queue=true` or `event=prequeueing` — for other cases the request will pass through to the origin. For our preview, this will be a fake origin website returning \"Welcome\".
	- **reject** indicates a Reject queue.
4. `event`: Used to preview a waiting room event.
	- **none** indicates no event is occurring.
	- **prequeueing** indicates that an event is prequeueing (between `prequeue_start_time` and `event_start_time`).
	- **started** indicates that an event has started (between `event_start_time` and `event_end_time`).
5. `shuffle_at_event_start`: Boolean indicating if the event will shuffle users in the prequeue when it starts. This can only be set to **true** if an event is active (`event` is not **none**).

For example, you can make a request to `http://waitingrooms.dev/preview/<uuid>?force_queue=false&queue_is_full=false&queueing_method=random&event=started&shuffle_at_event_start=true`
6. `waitTime`: Non-zero, positive integer indicating the estimated wait time in minutes. The default value is 10 minutes.

For example, you can make a request to `http://waitingrooms.dev/preview/<uuid>?waitTime=50` to configure the estimated wait time as 50 minutes.

## Definition

```yaml
{"operationId": "waiting-room-create-a-custom-waiting-room-page-preview", "summary": "Create a custom waiting room page preview", "description": "Creates a waiting room page preview. Upload a custom waiting room page for preview. You will receive a preview URL in the form `http://waitingrooms.dev/preview/<uuid>`. You can use the following query parameters to change the state of the preview:\n1. `force_queue`: Boolean indicating if all users will be queued in the waiting room and no one will be let into the origin website (also known as queueAll).\n2. `queue_is_full`: Boolean indicating if the waiting room's queue is currently full and not accepting new users at the moment.\n3. `queueing_method`: The queueing method currently used by the waiting room.\n\t- **fifo** indicates a FIFO queue.\n\t- **random** indicates a Random queue.\n\t- **passthrough** indicates a Passthrough queue. Keep in mind that the waiting room page will only be displayed if `force_queue=true` or `event=prequeueing` — for other cases the request will pass through to the origin. For our preview, this will be a fake origin website returning \\\"Welcome\\\". \n\t- **reject** indicates a Reject queue.\n4. `event`: Used to preview a waiting room event.\n\t- **none** indicates no event is occurring.\n\t- **prequeueing** indicates that an event is prequeueing (between `prequeue_start_time` and `event_start_time`).\n\t- **started** indicates that an event has started (between `event_start_time` and `event_end_time`).\n5. `shuffle_at_event_start`: Boolean indicating if the event will shuffle users in the prequeue when it starts. This can only be set to **true** if an event is active (`event` is not **none**).\n\nFor example, you can make a request to `http://waitingrooms.dev/preview/<uuid>?force_queue=false&queue_is_full=false&queueing_method=random&event=started&shuffle_at_event_start=true`\n6. `waitTime`: Non-zero, positive integer indicating the estimated wait time in minutes. The default value is 10 minutes.\n\nFor example, you can make a request to `http://waitingrooms.dev/preview/<uuid>?waitTime=50` to configure the estimated wait time as 50 minutes.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waitingroom_query_preview"}}}}, "responses": {"200": {"description": "Create a custom waiting room page preview response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waitingroom_preview_response"}}}}, "4XX": {"description": "Create a custom waiting room page preview response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/waitingroom_preview_response"}, {"$ref": "#/components/schemas/waitingroom_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Waiting Room"], "x-api-token-group": ["Waiting Rooms Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "waiting-rooms.page", "x-fern-sdk-method-name": "preview", "x-forge-hidden": true}
```
