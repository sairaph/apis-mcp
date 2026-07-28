---
title: Change Origin Max HTTP Version Setting
page_id: operation-patch-zones-zone-id-settings-origin-max-http-version-e6a25d9b
path: operations/zone-settings
description: Origin Max HTTP Setting Version sets the highest HTTP version Cloudflare will attempt to use with your origin. This setting allows Cloudflare to make HTTP/2 requests to your origin. (Refer to [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/), for more information.). The default value is "2" for all plan types except Enterprise where it is "1".
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/settings/origin_max_http_version
operation_ids:
    - zone-cache-settings-change-origin-max-http-version-setting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Change Origin Max HTTP Version Setting

`PATCH /zones/{zone_id}/settings/origin_max_http_version`

Operation ID: `zone-cache-settings-change-origin-max-http-version-setting`

Origin Max HTTP Setting Version sets the highest HTTP version Cloudflare will attempt to use with your origin. This setting allows Cloudflare to make HTTP/2 requests to your origin. (Refer to [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/), for more information.). The default value is "2" for all plan types except Enterprise where it is "1".

## Definition

```yaml
{"operationId": "zone-cache-settings-change-origin-max-http-version-setting", "summary": "Change Origin Max HTTP Version Setting", "description": "Origin Max HTTP Setting Version sets the highest HTTP version Cloudflare will attempt to use with your origin. This setting allows Cloudflare to make HTTP/2 requests to your origin. (Refer to [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/), for more information.). The default value is \"2\" for all plan types except Enterprise where it is \"1\".", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache-rules_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"value": {"$ref": "#/components/schemas/cache-rules_origin_max_http_version_value"}}, "required": ["value"]}}}}, "responses": {"200": {"description": "Change Origin Max HTTP Version setting response.", "content": {"application/json": {"examples": {"max_http_version_2": {"$ref": "#/components/examples/cache-rules_origin_max_http_version_2"}}, "schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/cache-rules_zone_cache_settings_response_single"}, {"$ref": "#/components/schemas/cache-rules_origin_max_http_version_response_value"}]}}}}, "4XX": {"description": "Change Origin Max HTTP Version response failure.", "content": {"application/json": {"examples": {"Error": {"$ref": "#/components/examples/cache-rules_dummy_error_response"}}, "schema": {"$ref": "#/components/schemas/cache-rules_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Settings"], "x-api-token-group": ["Zone Settings Write", "Zone Write"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
