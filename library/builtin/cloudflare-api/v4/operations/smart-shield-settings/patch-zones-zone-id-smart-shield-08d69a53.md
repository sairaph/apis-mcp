---
title: Patch Smart Shield Settings
page_id: operation-patch-zones-zone-id-smart-shield-710d66af
path: operations/smart-shield-settings
description: Set Smart Shield Settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/smart_shield
operation_ids:
    - smart-shield-patch-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch Smart Shield Settings

`PATCH /zones/{zone_id}/smart_shield`

Operation ID: `smart-shield-patch-settings`

Set Smart Shield Settings.

## Definition

```yaml
{"operationId": "smart-shield-patch-settings", "summary": "Patch Smart Shield Settings", "description": "Set Smart Shield Settings.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/smartshield_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/smartshield_smart_shield_settings_patch_body"}}}}, "responses": {"200": {"description": "Smart Shield Settings response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/smartshield_api-response-common"}, {"$ref": "#/components/schemas/smartshield_single_smart_shield_patch_response"}]}}}}, "500": {"description": "Get Smart Shield Settings response failure.", "content": {"application/json": {"examples": {"Invalid value": {"$ref": "#/components/examples/smartshield_smart_shield_500_response"}}, "schema": {"$ref": "#/components/schemas/smartshield_api-response-common-failure"}}}}, "502": {"description": "Smart Shield Settings response failure.", "content": {"application/json": {"examples": {"Not found": {"$ref": "#/components/examples/smartshield_smart_shield_502_patch_response"}}, "schema": {"$ref": "#/components/schemas/smartshield_api-response-common-failure"}}}}, "4XX": {"description": "Patch Smart Shield Settings response failure.", "content": {"application/json": {"examples": {"Invalid value": {"$ref": "#/components/examples/smartshield_smart_shield_4xx_get_response"}}, "schema": {"$ref": "#/components/schemas/smartshield_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Smart Shield Settings"], "x-api-token-group": ["Zone Settings Write", "Zone Write"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:read", "#zone_settings:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
