---
title: Get Argo Smart Routing setting
page_id: operation-get-zones-zone-id-argo-smart-routing-c91db48f
path: operations/argo-smart-routing
description: Retrieves the value of Argo Smart Routing enablement setting.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/argo/smart_routing
operation_ids:
    - argo-smart-routing-get-argo-smart-routing-setting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Argo Smart Routing setting

`GET /zones/{zone_id}/argo/smart_routing`

Operation ID: `argo-smart-routing-get-argo-smart-routing-setting`

Retrieves the value of Argo Smart Routing enablement setting.

## Definition

```yaml
{"operationId": "argo-smart-routing-get-argo-smart-routing-setting", "summary": "Get Argo Smart Routing setting", "description": "Retrieves the value of Argo Smart Routing enablement setting.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/argo-config_identifier"}}], "responses": {"200": {"description": "Get Argo Smart Routing enablement setting response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/argo-config_api_response_single"}}}}, "4XX": {"description": "Get Argo Smart Routing enablement setting failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/argo-config_api_response_common_failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Argo Smart Routing"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
