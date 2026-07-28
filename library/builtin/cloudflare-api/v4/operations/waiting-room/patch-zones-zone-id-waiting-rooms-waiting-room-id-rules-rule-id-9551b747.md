---
title: Patch Waiting Room Rule
page_id: operation-patch-zones-zone-id-waiting-rooms-waiting-room-id-rules-rule-id-a1ff6509
path: operations/waiting-room
description: Patches a rule for a waiting room.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/waiting_rooms/{waiting_room_id}/rules/{rule_id}
operation_ids:
    - waiting-room-patch-waiting-room-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch Waiting Room Rule

`PATCH /zones/{zone_id}/waiting_rooms/{waiting_room_id}/rules/{rule_id}`

Operation ID: `waiting-room-patch-waiting-room-rule`

Patches a rule for a waiting room.

## Definition

```yaml
{"operationId": "waiting-room-patch-waiting-room-rule", "summary": "Patch Waiting Room Rule", "description": "Patches a rule for a waiting room.", "parameters": [{"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_rule_id"}}, {"name": "waiting_room_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_waiting_room_id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waitingroom_patch_rule"}}}}, "responses": {"200": {"description": "Patch Waiting Room Rule response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waitingroom_rules_response_collection"}}}}, "4XX": {"description": "Patch Waiting Room Rule response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/waitingroom_rules_response_collection"}, {"$ref": "#/components/schemas/waitingroom_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Waiting Room"], "x-api-token-group": ["Waiting Rooms Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "waiting-rooms.rules", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
