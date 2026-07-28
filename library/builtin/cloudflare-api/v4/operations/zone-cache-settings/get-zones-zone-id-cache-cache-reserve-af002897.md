---
title: Get Cache Reserve setting
page_id: operation-get-zones-zone-id-cache-cache-reserve-549274a7
path: operations/zone-cache-settings
description: 'Increase cache lifetimes by automatically storing all cacheable files into Cloudflare''s persistent object storage buckets. Requires Cache Reserve subscription. Note: using Tiered Cache with Cache Reserve is highly recommended to reduce Reserve operations costs. See the [developer docs](https://developers.cloudflare.com/cache/about/cache-reserve) for more information.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/cache/cache_reserve
operation_ids:
    - zone-cache-settings-get-cache-reserve-setting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Cache Reserve setting

`GET /zones/{zone_id}/cache/cache_reserve`

Operation ID: `zone-cache-settings-get-cache-reserve-setting`

Increase cache lifetimes by automatically storing all cacheable files into Cloudflare's persistent object storage buckets. Requires Cache Reserve subscription. Note: using Tiered Cache with Cache Reserve is highly recommended to reduce Reserve operations costs. See the [developer docs](https://developers.cloudflare.com/cache/about/cache-reserve) for more information.

## Definition

```yaml
{"operationId": "zone-cache-settings-get-cache-reserve-setting", "summary": "Get Cache Reserve setting", "description": "Increase cache lifetimes by automatically storing all cacheable files into Cloudflare's persistent object storage buckets. Requires Cache Reserve subscription. Note: using Tiered Cache with Cache Reserve is highly recommended to reduce Reserve operations costs. See the [developer docs](https://developers.cloudflare.com/cache/about/cache-reserve) for more information.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache-rules_identifier"}}], "responses": {"200": {"description": "Get Cache Reserve setting response.", "content": {"application/json": {"examples": {"off": {"$ref": "#/components/examples/cache-rules_cache_reserve_off"}}, "schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/cache-rules_zone_cache_settings_response_single"}, {"$ref": "#/components/schemas/cache-rules_cache_reserve_response_value"}]}}}}, "4XX": {"description": "Get Cache Reserve setting response failure.", "content": {"application/json": {"examples": {"Error": {"$ref": "#/components/examples/cache-rules_dummy_error_response"}}, "schema": {"$ref": "#/components/schemas/cache-rules_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Cache Settings"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read", "Zone Read", "Zone Write"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
