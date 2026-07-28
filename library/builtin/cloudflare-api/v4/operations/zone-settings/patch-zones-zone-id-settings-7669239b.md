---
title: Edit multiple zone settings
page_id: operation-patch-zones-zone-id-settings-b705229a
path: operations/zone-settings
description: Edit settings for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/settings
operation_ids:
    - zone-settings-edit-zone-settings-info
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Edit multiple zone settings

`PATCH /zones/{zone_id}/settings`

Operation ID: `zone-settings-edit-zone-settings-info`

Edit settings for a zone.

## Definition

```yaml
{"operationId": "zone-settings-edit-zone-settings-info", "summary": "Edit multiple zone settings", "description": "Edit settings for a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zones_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zones_multiple_settings"}}}}, "responses": {"200": {"description": "Edit zone settings info response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zones_zone_settings_response_collection"}}}}, "4XX": {"description": "Edit zone settings info response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zones_api-response-common-failure-3"}}}}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Settings"], "x-api-token-group": ["Zone Settings Write"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "deprecated", "x-fern-sdk-group-name": "zones.settings", "x-fern-sdk-method-name": "bulk-edit", "x-forge-sunset": {"date": "2030-01-01T00:00:00Z"}, "x-stainless-deprecation-message": "This endpoint is deprecated. Zone settings should instead be managed individually."}
```
