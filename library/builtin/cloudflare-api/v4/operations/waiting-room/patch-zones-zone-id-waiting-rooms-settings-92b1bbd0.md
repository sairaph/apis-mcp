---
title: Patch zone-level Waiting Room settings
page_id: operation-patch-zones-zone-id-waiting-rooms-settings-e1994456
path: operations/waiting-room
description: Patch zone-level Waiting Room settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/waiting_rooms/settings
operation_ids:
    - waiting-room-patch-zone-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch zone-level Waiting Room settings

`PATCH /zones/{zone_id}/waiting_rooms/settings`

Operation ID: `waiting-room-patch-zone-settings`

Patch zone-level Waiting Room settings.

## Definition

```yaml
{"operationId": "waiting-room-patch-zone-settings", "summary": "Patch zone-level Waiting Room settings", "description": "Patch zone-level Waiting Room settings.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waitingroom_zone_settings"}}}}, "responses": {"200": {"description": "The updated zone-level Waiting Room settings", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waitingroom_zone_settings_response"}}}}, "4XX": {"description": "The zone-level Waiting Room settings response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/waitingroom_zone_settings_response"}, {"$ref": "#/components/schemas/waitingroom_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Waiting Room"], "x-api-token-group": ["Waiting Rooms Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "waiting-rooms.settings", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
