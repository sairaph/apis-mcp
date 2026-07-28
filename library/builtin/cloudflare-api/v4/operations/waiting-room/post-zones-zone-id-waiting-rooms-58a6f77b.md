---
title: Create waiting room
page_id: operation-post-zones-zone-id-waiting-rooms-ad75f639
path: operations/waiting-room
description: Creates a new waiting room.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/waiting_rooms
operation_ids:
    - waiting-room-create-waiting-room
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create waiting room

`POST /zones/{zone_id}/waiting_rooms`

Operation ID: `waiting-room-create-waiting-room`

Creates a new waiting room.

## Definition

```yaml
{"operationId": "waiting-room-create-waiting-room", "summary": "Create waiting room", "description": "Creates a new waiting room.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waitingroom_query_waitingroom"}}}}, "responses": {"200": {"description": "Create waiting room response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waitingroom_single_response"}}}}, "4XX": {"description": "Create waiting room response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/waitingroom_single_response"}, {"$ref": "#/components/schemas/waitingroom_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Waiting Room"], "x-api-token-group": ["Waiting Rooms Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "waiting-rooms", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
