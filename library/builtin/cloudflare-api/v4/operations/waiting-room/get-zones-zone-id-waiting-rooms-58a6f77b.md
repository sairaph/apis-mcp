---
title: List waiting rooms for zone
page_id: operation-get-zones-zone-id-waiting-rooms-12a26603
path: operations/waiting-room
description: Lists waiting rooms for zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/waiting_rooms
operation_ids:
    - waiting-room-list-waiting-rooms
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List waiting rooms for zone

`GET /zones/{zone_id}/waiting_rooms`

Operation ID: `waiting-room-list-waiting-rooms`

Lists waiting rooms for zone.

## Definition

```yaml
{"operationId": "waiting-room-list-waiting-rooms", "summary": "List waiting rooms for zone", "description": "Lists waiting rooms for zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_identifier"}}, {"$ref": "#/components/parameters/waitingroom_page"}, {"$ref": "#/components/parameters/waitingroom_per_page"}], "responses": {"200": {"description": "List waiting rooms for zone response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waitingroom_response_collection"}}}}, "4XX": {"description": "List waiting rooms for zone response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/waitingroom_response_collection"}, {"$ref": "#/components/schemas/waitingroom_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Waiting Room"], "x-api-token-group": ["Waiting Rooms Read", "Waiting Rooms Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "waiting-rooms", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
