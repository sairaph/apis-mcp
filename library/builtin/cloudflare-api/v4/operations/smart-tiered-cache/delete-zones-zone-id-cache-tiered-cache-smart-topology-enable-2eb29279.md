---
title: Delete Smart Tiered Cache setting
page_id: operation-delete-zones-zone-id-cache-tiered-cache-smart-topology-enable-c45a53ff
path: operations/smart-tiered-cache
description: Smart Tiered Cache dynamically selects the single closest upper tier for each of your website’s origins with no configuration required, using our in-house performance and routing data. Cloudflare collects latency data for each request to an origin, and uses the latency data to determine how well any upper-tier data center is connected with an origin. As a result, Cloudflare can select the data center with the lowest latency to be the upper-tier for an origin.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/cache/tiered_cache_smart_topology_enable
operation_ids:
    - smart-tiered-cache-delete-smart-tiered-cache-setting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Smart Tiered Cache setting

`DELETE /zones/{zone_id}/cache/tiered_cache_smart_topology_enable`

Operation ID: `smart-tiered-cache-delete-smart-tiered-cache-setting`

Smart Tiered Cache dynamically selects the single closest upper tier for each of your website’s origins with no configuration required, using our in-house performance and routing data. Cloudflare collects latency data for each request to an origin, and uses the latency data to determine how well any upper-tier data center is connected with an origin. As a result, Cloudflare can select the data center with the lowest latency to be the upper-tier for an origin.

## Definition

```yaml
{"operationId": "smart-tiered-cache-delete-smart-tiered-cache-setting", "summary": "Delete Smart Tiered Cache setting", "description": "Smart Tiered Cache dynamically selects the single closest upper tier for each of your website’s origins with no configuration required, using our in-house performance and routing data. Cloudflare collects latency data for each request to an origin, and uses the latency data to determine how well any upper-tier data center is connected with an origin. As a result, Cloudflare can select the data center with the lowest latency to be the upper-tier for an origin.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache-rules_identifier"}}], "responses": {"200": {"description": "Delete Smart Tiered Cache setting response.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/cache-rules_zone_cache_settings_delete_response_single"}, {"properties": {"result": {"$ref": "#/components/schemas/cache-rules_smart_tiered_cache"}}, "type": "object"}]}}}}, "4XX": {"description": "Delete Smart Tiered Cache setting response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cache-rules_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Smart Tiered Cache"], "x-api-token-group": ["Zone Settings Write", "Zone Write"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
