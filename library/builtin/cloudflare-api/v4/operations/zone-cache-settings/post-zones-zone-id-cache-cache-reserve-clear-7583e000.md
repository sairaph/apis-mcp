---
title: Start Cache Reserve Clear
page_id: operation-post-zones-zone-id-cache-cache-reserve-clear-181d176d
path: operations/zone-cache-settings
description: You can use Cache Reserve Clear to clear your Cache Reserve, but you must first disable Cache Reserve. In most cases, this will be accomplished within 24 hours. You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind that you cannot undo or cancel this operation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/cache/cache_reserve_clear
operation_ids:
    - zone-cache-settings-start-cache-reserve-clear
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Start Cache Reserve Clear

`POST /zones/{zone_id}/cache/cache_reserve_clear`

Operation ID: `zone-cache-settings-start-cache-reserve-clear`

You can use Cache Reserve Clear to clear your Cache Reserve, but you must first disable Cache Reserve. In most cases, this will be accomplished within 24 hours. You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind that you cannot undo or cancel this operation.

## Definition

```yaml
{"operationId": "zone-cache-settings-start-cache-reserve-clear", "summary": "Start Cache Reserve Clear", "description": "You can use Cache Reserve Clear to clear your Cache Reserve, but you must first disable Cache Reserve. In most cases, this will be accomplished within 24 hours. You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind that you cannot undo or cancel this operation.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache-rules_identifier"}}], "requestBody": {"description": "The request body is currently not used.", "required": true, "content": {"application/json": {"example": "{}"}}}, "responses": {"200": {"description": "Start Cache Reserve Clear response.", "content": {"application/json": {"examples": {"In-progress": {"$ref": "#/components/examples/cache-rules_cache_reserve_clear_in_progress"}}, "schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/cache-rules_api-response-common"}, {"$ref": "#/components/schemas/cache-rules_cache_reserve_clear_response_value"}]}}}}, "4XX": {"description": "Start Cache Reserve Clear failure response.", "content": {"application/json": {"examples": {"Rejected": {"$ref": "#/components/examples/cache-rules_cache_reserve_clear_rejected_cr_on"}}, "schema": {"$ref": "#/components/schemas/cache-rules_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Cache Settings"], "x-api-token-group": ["Zone Settings Write", "Zone Write"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
