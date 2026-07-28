---
title: Change Regional Tiered Cache setting
page_id: operation-patch-zones-zone-id-cache-regional-tiered-cache-d9de4f21
path: operations/zone-cache-settings
description: Instructs Cloudflare to check a regional hub data center on the way to your upper tier. This can help improve performance for smart and custom tiered cache topologies.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/cache/regional_tiered_cache
operation_ids:
    - zone-cache-settings-change-regional-tiered-cache-setting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Change Regional Tiered Cache setting

`PATCH /zones/{zone_id}/cache/regional_tiered_cache`

Operation ID: `zone-cache-settings-change-regional-tiered-cache-setting`

Instructs Cloudflare to check a regional hub data center on the way to your upper tier. This can help improve performance for smart and custom tiered cache topologies.

## Definition

```yaml
{"operationId": "zone-cache-settings-change-regional-tiered-cache-setting", "summary": "Change Regional Tiered Cache setting", "description": "Instructs Cloudflare to check a regional hub data center on the way to your upper tier. This can help improve performance for smart and custom tiered cache topologies.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache-rules_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"value": {"$ref": "#/components/schemas/cache-rules_regional_tiered_cache_value"}}, "required": ["value"]}}}}, "responses": {"200": {"description": "Change Regional Tiered Cache setting response.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/cache-rules_zone_cache_settings_response_single"}, {"$ref": "#/components/schemas/cache-rules_regional_tiered_cache_response_value"}]}}}}, "4XX": {"description": "Change Regional Tiered Cache setting response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cache-rules_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Cache Settings"], "x-api-token-group": ["Zone Settings Write", "Zone Write"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
