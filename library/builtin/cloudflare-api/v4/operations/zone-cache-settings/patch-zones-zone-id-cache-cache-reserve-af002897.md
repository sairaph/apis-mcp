---
title: Change Cache Reserve setting
page_id: operation-patch-zones-zone-id-cache-cache-reserve-8511d977
path: operations/zone-cache-settings
description: 'Increase cache lifetimes by automatically storing all cacheable files into Cloudflare''s persistent object storage buckets. Requires Cache Reserve subscription. Note: using Tiered Cache with Cache Reserve is highly recommended to reduce Reserve operations costs. See the [developer docs](https://developers.cloudflare.com/cache/about/cache-reserve) for more information.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/cache/cache_reserve
operation_ids:
    - zone-cache-settings-change-cache-reserve-setting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Change Cache Reserve setting

`PATCH /zones/{zone_id}/cache/cache_reserve`

Operation ID: `zone-cache-settings-change-cache-reserve-setting`

Increase cache lifetimes by automatically storing all cacheable files into Cloudflare's persistent object storage buckets. Requires Cache Reserve subscription. Note: using Tiered Cache with Cache Reserve is highly recommended to reduce Reserve operations costs. See the [developer docs](https://developers.cloudflare.com/cache/about/cache-reserve) for more information.

## Definition

```yaml
{"operationId": "zone-cache-settings-change-cache-reserve-setting", "summary": "Change Cache Reserve setting", "description": "Increase cache lifetimes by automatically storing all cacheable files into Cloudflare's persistent object storage buckets. Requires Cache Reserve subscription. Note: using Tiered Cache with Cache Reserve is highly recommended to reduce Reserve operations costs. See the [developer docs](https://developers.cloudflare.com/cache/about/cache-reserve) for more information.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache-rules_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"value": {"$ref": "#/components/schemas/cache-rules_cache_reserve_value"}}, "required": ["value"]}}}}, "responses": {"200": {"description": "Change Cache Reserve setting response.", "content": {"application/json": {"examples": {"on": {"$ref": "#/components/examples/cache-rules_cache_reserve_on"}}, "schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/cache-rules_zone_cache_settings_response_single"}, {"$ref": "#/components/schemas/cache-rules_cache_reserve_response_value"}]}}}}, "4XX": {"description": "Change Cache Reserve setting response failure.", "content": {"application/json": {"examples": {"Denied": {"$ref": "#/components/examples/cache-rules_cache_reserve_denied_clearing"}}, "schema": {"$ref": "#/components/schemas/cache-rules_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Cache Settings"], "x-api-token-group": ["Zone Settings Write", "Zone Write"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
