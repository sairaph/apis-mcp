---
title: Change Cloudflare Speed Brain setting
page_id: operation-patch-zones-zone-id-settings-speed-brain-d3ae913d
path: operations/zone-settings
description: |-
    Speed Brain lets compatible browsers speculate on content which can be prefetched or preloaded, making website
    navigation faster. Refer to the Cloudflare Speed Brain documentation for more information.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/settings/speed_brain
operation_ids:
    - zone-settings-change-speed-brain-setting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Change Cloudflare Speed Brain setting

`PATCH /zones/{zone_id}/settings/speed_brain`

Operation ID: `zone-settings-change-speed-brain-setting`

Speed Brain lets compatible browsers speculate on content which can be prefetched or preloaded, making website
navigation faster. Refer to the Cloudflare Speed Brain documentation for more information.

## Definition

```yaml
{"operationId": "zone-settings-change-speed-brain-setting", "summary": "Change Cloudflare Speed Brain setting", "description": "Speed Brain lets compatible browsers speculate on content which can be prefetched or preloaded, making website\nnavigation faster. Refer to the Cloudflare Speed Brain documentation for more information.\n", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/speed_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"value": {"description": "Whether the feature is enabled or disabled.", "type": "string", "enum": ["on", "off"]}}, "required": ["value"]}}}}, "responses": {"200": {"description": "Cloudflare Speed Brain setting response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/speed_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/speed_cloudflare_speed_brain_response"}}, "type": "object"}]}}}}, "4XX": {"description": "Change Cloudflare Speed Brain setting response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/speed_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Settings"], "x-api-token-group": ["Zone Settings Write"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
