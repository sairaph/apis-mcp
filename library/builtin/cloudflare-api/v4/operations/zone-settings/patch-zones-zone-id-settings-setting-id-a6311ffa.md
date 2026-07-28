---
title: Edit zone setting
page_id: operation-patch-zones-zone-id-settings-setting-id-a24dc374
path: operations/zone-settings
description: Updates a single zone setting by the identifier
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/settings/{setting_id}
operation_ids:
    - zone-settings-edit-single-setting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Edit zone setting

`PATCH /zones/{zone_id}/settings/{setting_id}`

Operation ID: `zone-settings-edit-single-setting`

Updates a single zone setting by the identifier

## Definition

```yaml
{"operationId": "zone-settings-edit-single-setting", "summary": "Edit zone setting", "description": "Updates a single zone setting by the identifier", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zones_identifier"}}, {"name": "setting_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zones_setting_name"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zones_zone_settings_single_request"}}}}, "responses": {"200": {"description": "Edit zone setting response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zones_api-response-common-4"}, {"properties": {"result": {"$ref": "#/components/schemas/zones_setting"}}}]}}}}, "4XX": {"description": "Edit zone settings info response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zones_api-response-common-failure-3"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Settings"], "x-api-token-group": ["Zone Settings Write"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zones.settings", "x-fern-sdk-method-name": "edit"}
```
