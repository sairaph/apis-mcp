---
title: Patch Argo Smart Routing setting
page_id: operation-patch-zones-zone-id-argo-smart-routing-db05b920
path: operations/argo-smart-routing
description: Configures the value of the Argo Smart Routing enablement setting.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/argo/smart_routing
operation_ids:
    - argo-smart-routing-patch-argo-smart-routing-setting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch Argo Smart Routing setting

`PATCH /zones/{zone_id}/argo/smart_routing`

Operation ID: `argo-smart-routing-patch-argo-smart-routing-setting`

Configures the value of the Argo Smart Routing enablement setting.

## Definition

```yaml
{"operationId": "argo-smart-routing-patch-argo-smart-routing-setting", "summary": "Patch Argo Smart Routing setting", "description": "Configures the value of the Argo Smart Routing enablement setting.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/argo-config_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/argo-config_patch"}}}}, "responses": {"200": {"description": "Patch Argo Smart Routing enablement setting response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/argo-config_api_response_single"}}}}, "4XX": {"description": "Patch Argo Smart Routing enablement setting failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/argo-config_api_response_common_failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Argo Smart Routing"], "x-api-token-group": ["Zone Settings Write"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:read", "#zone_settings:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
