---
title: Get zone-level Waiting Room settings
page_id: operation-get-zones-zone-id-waiting-rooms-settings-a97b7e00
path: operations/waiting-room
description: Get zone-level Waiting Room settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/waiting_rooms/settings
operation_ids:
    - waiting-room-get-zone-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get zone-level Waiting Room settings

`GET /zones/{zone_id}/waiting_rooms/settings`

Operation ID: `waiting-room-get-zone-settings`

Get zone-level Waiting Room settings.

## Definition

```yaml
{"operationId": "waiting-room-get-zone-settings", "summary": "Get zone-level Waiting Room settings", "description": "Get zone-level Waiting Room settings.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_identifier"}}], "responses": {"200": {"description": "The current zone-level Waiting Room settings", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waitingroom_zone_settings_response"}}}}, "4XX": {"description": "The current zone-level Waiting Room settings response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/waitingroom_zone_settings_response"}, {"$ref": "#/components/schemas/waitingroom_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Waiting Room"], "x-api-token-group": ["Waiting Rooms Read", "Waiting Rooms Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "waiting-rooms.settings", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
