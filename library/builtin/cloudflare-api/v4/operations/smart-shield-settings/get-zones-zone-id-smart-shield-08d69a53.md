---
title: Get Smart Shield Settings
page_id: operation-get-zones-zone-id-smart-shield-7e60fa19
path: operations/smart-shield-settings
description: Retrieve Smart Shield Settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/smart_shield
operation_ids:
    - smart-shield-get-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Smart Shield Settings

`GET /zones/{zone_id}/smart_shield`

Operation ID: `smart-shield-get-settings`

Retrieve Smart Shield Settings.

## Definition

```yaml
{"operationId": "smart-shield-get-settings", "summary": "Get Smart Shield Settings", "description": "Retrieve Smart Shield Settings.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/smartshield_identifier"}}], "responses": {"200": {"description": "Smart Shield Settings response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/smartshield_api-response-common"}, {"$ref": "#/components/schemas/smartshield_single_smart_shield_get_response"}]}}}}, "500": {"description": "Get Smart Shield Settings response failure.", "content": {"application/json": {"examples": {"Invalid value": {"$ref": "#/components/examples/smartshield_smart_shield_500_response"}}, "schema": {"$ref": "#/components/schemas/smartshield_api-response-common-failure"}}}}, "502": {"description": "Get Smart Shield Settings response failure.", "content": {"application/json": {"examples": {"Not found": {"$ref": "#/components/examples/smartshield_smart_shield_502_get_response"}}, "schema": {"$ref": "#/components/schemas/smartshield_api-response-common-failure"}}}}, "4XX": {"description": "Patch Smart Shield Settings response failure.", "content": {"application/json": {"examples": {"Invalid value": {"$ref": "#/components/examples/smartshield_smart_shield_4xx_get_response"}}, "schema": {"$ref": "#/components/schemas/smartshield_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Smart Shield Settings"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read", "Zone Read", "Zone Write"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
