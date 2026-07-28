---
title: Get zone setting
page_id: operation-get-zones-zone-id-settings-setting-id-f843a067
path: operations/zone-settings
description: Fetch a single zone setting by name
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/settings/{setting_id}
operation_ids:
    - zone-settings-get-single-setting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get zone setting

`GET /zones/{zone_id}/settings/{setting_id}`

Operation ID: `zone-settings-get-single-setting`

Fetch a single zone setting by name

## Definition

```yaml
{"operationId": "zone-settings-get-single-setting", "summary": "Get zone setting", "description": "Fetch a single zone setting by name", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zones_identifier"}}, {"name": "setting_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zones_setting_name"}}], "responses": {"200": {"description": "Get zone setting response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zones_api-response-common-4"}, {"properties": {"result": {"$ref": "#/components/schemas/zones_setting"}}}]}}}}, "4XX": {"description": "Get zone setting response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zones_api-response-common-failure-3"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Settings"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zones.settings", "x-fern-sdk-method-name": "get"}
```
