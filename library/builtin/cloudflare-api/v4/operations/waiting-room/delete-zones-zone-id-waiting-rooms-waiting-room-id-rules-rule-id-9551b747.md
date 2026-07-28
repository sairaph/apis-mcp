---
title: Delete Waiting Room Rule
page_id: operation-delete-zones-zone-id-waiting-rooms-waiting-room-id-rules-rule-id-6c912c19
path: operations/waiting-room
description: Deletes a rule for a waiting room.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/waiting_rooms/{waiting_room_id}/rules/{rule_id}
operation_ids:
    - waiting-room-delete-waiting-room-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Waiting Room Rule

`DELETE /zones/{zone_id}/waiting_rooms/{waiting_room_id}/rules/{rule_id}`

Operation ID: `waiting-room-delete-waiting-room-rule`

Deletes a rule for a waiting room.

## Definition

```yaml
{"operationId": "waiting-room-delete-waiting-room-rule", "summary": "Delete Waiting Room Rule", "description": "Deletes a rule for a waiting room.", "parameters": [{"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_rule_id"}}, {"name": "waiting_room_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_waiting_room_id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waitingroom_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Waiting Room Rule response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waitingroom_rules_response_collection"}}}}, "4XX": {"description": "Delete Waiting Room Rule response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/waitingroom_rules_response_collection"}, {"$ref": "#/components/schemas/waitingroom_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Waiting Room"], "x-api-token-group": ["Waiting Rooms Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "waiting-rooms.rules", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
