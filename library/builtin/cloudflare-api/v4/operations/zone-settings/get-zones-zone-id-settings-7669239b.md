---
title: Get all zone settings
page_id: operation-get-zones-zone-id-settings-d18f9a2f
path: operations/zone-settings
description: Available settings for your user in relation to a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/settings
operation_ids:
    - zone-settings-get-all-zone-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get all zone settings

`GET /zones/{zone_id}/settings`

Operation ID: `zone-settings-get-all-zone-settings`

Available settings for your user in relation to a zone.

## Definition

```yaml
{"operationId": "zone-settings-get-all-zone-settings", "summary": "Get all zone settings", "description": "Available settings for your user in relation to a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zones_identifier"}}], "responses": {"200": {"description": "Get all Zone settings response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zones_zone_settings_response_collection"}}}}, "4XX": {"description": "Get all Zone settings response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zones_api-response-common-failure-3"}}}}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Settings"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "deprecated", "x-fern-sdk-group-name": "zones.settings", "x-fern-sdk-method-name": "list", "x-forge-sunset": {"date": "2030-01-01T00:00:00Z"}, "x-stainless-deprecation-message": "This endpoint is deprecated. Zone settings should instead be managed individually."}
```
