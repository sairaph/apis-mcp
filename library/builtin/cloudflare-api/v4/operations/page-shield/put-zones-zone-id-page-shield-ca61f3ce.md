---
title: Update Page Shield settings
page_id: operation-put-zones-zone-id-page-shield-b097b478
path: operations/page-shield
description: Updates Page Shield settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/page_shield
operation_ids:
    - page-shield-update-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Page Shield settings

`PUT /zones/{zone_id}/page_shield`

Operation ID: `page-shield-update-settings`

Updates Page Shield settings.

## Definition

```yaml
{"operationId": "page-shield-update-settings", "summary": "Update Page Shield settings", "description": "Updates Page Shield settings.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/page-shield_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"enabled": {"$ref": "#/components/schemas/page-shield_enabled"}, "use_cloudflare_reporting_endpoint": {"$ref": "#/components/schemas/page-shield_use_cloudflare_reporting_endpoint"}, "use_connection_url_path": {"$ref": "#/components/schemas/page-shield_use_connection_url_path"}}}}}}, "responses": {"200": {"description": "Update Page Shield settings response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/page-shield_zone_settings_response_single"}, {"properties": {"result": {"$ref": "#/components/schemas/page-shield_update-zone-settings-response"}}}]}}}}, "4XX": {"description": "Update Page Shield settings response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/page-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Page Shield"], "x-api-token-group": ["Page Shield", "Domain Page Shield", "Zone Settings Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
