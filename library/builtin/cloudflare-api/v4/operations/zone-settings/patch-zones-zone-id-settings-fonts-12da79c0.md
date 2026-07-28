---
title: Change Cloudflare Fonts setting
page_id: operation-patch-zones-zone-id-settings-fonts-f1eed45d
path: operations/zone-settings
description: |-
    Enhance your website's font delivery with Cloudflare Fonts. Deliver Google Hosted fonts from your own domain,
    boost performance, and enhance user privacy. Refer to the Cloudflare Fonts documentation for more information.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/settings/fonts
operation_ids:
    - zone-settings-change-fonts-setting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Change Cloudflare Fonts setting

`PATCH /zones/{zone_id}/settings/fonts`

Operation ID: `zone-settings-change-fonts-setting`

Enhance your website's font delivery with Cloudflare Fonts. Deliver Google Hosted fonts from your own domain,
boost performance, and enhance user privacy. Refer to the Cloudflare Fonts documentation for more information.

## Definition

```yaml
{"operationId": "zone-settings-change-fonts-setting", "summary": "Change Cloudflare Fonts setting", "description": "Enhance your website's font delivery with Cloudflare Fonts. Deliver Google Hosted fonts from your own domain,\nboost performance, and enhance user privacy. Refer to the Cloudflare Fonts documentation for more information.\n", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/speed_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"value": {"$ref": "#/components/schemas/speed_cloudflare_fonts_value"}}, "required": ["value"]}}}}, "responses": {"200": {"description": "Change Cloudflare Fonts setting response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/speed_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/speed_cloudflare_fonts"}}, "type": "object"}]}}}}, "4XX": {"description": "Change Cloudflare Fonts setting response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/speed_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Settings"], "x-api-token-group": ["Zone Settings Write"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
