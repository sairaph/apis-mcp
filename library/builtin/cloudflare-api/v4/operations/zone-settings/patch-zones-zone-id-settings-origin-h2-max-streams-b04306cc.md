---
title: Change Origin H2 Max Streams Setting
page_id: operation-patch-zones-zone-id-settings-origin-h2-max-streams-5c85710c
path: operations/zone-settings
description: Origin H2 Max Streams configures the max number of concurrent requests that Cloudflare will send within the same connection when communicating with the origin server, if the origin supports it. Note that if your origin does not support H2 multiplexing, 5xx errors may be observed, particularly 520s. Also note that the default value is `100` for all plan types except Enterprise where it is `1`. `1` means that H2 multiplexing is disabled.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/settings/origin_h2_max_streams
operation_ids:
    - zone-cache-settings-change-origin-h2-max-streams-setting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Change Origin H2 Max Streams Setting

`PATCH /zones/{zone_id}/settings/origin_h2_max_streams`

Operation ID: `zone-cache-settings-change-origin-h2-max-streams-setting`

Origin H2 Max Streams configures the max number of concurrent requests that Cloudflare will send within the same connection when communicating with the origin server, if the origin supports it. Note that if your origin does not support H2 multiplexing, 5xx errors may be observed, particularly 520s. Also note that the default value is `100` for all plan types except Enterprise where it is `1`. `1` means that H2 multiplexing is disabled.

## Definition

```yaml
{"operationId": "zone-cache-settings-change-origin-h2-max-streams-setting", "summary": "Change Origin H2 Max Streams Setting", "description": "Origin H2 Max Streams configures the max number of concurrent requests that Cloudflare will send within the same connection when communicating with the origin server, if the origin supports it. Note that if your origin does not support H2 multiplexing, 5xx errors may be observed, particularly 520s. Also note that the default value is `100` for all plan types except Enterprise where it is `1`. `1` means that H2 multiplexing is disabled.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache-rules_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"value": {"$ref": "#/components/schemas/cache-rules_origin_h2_max_streams_value"}}, "required": ["value"]}}}}, "responses": {"200": {"description": "Change Origin H2 Max Streams setting response.", "content": {"application/json": {"examples": {"origin_h2_max_streams_50": {"$ref": "#/components/examples/cache-rules_origin_h2_max_streams_50"}}, "schema": {"$ref": "#/components/schemas/cache-rules_origin_h2_max_streams_response_value"}}}}, "4XX": {"description": "Change Origin H2 Max Streams response failure.", "content": {"application/json": {"examples": {"Error": {"$ref": "#/components/examples/cache-rules_dummy_error_response"}}, "schema": {"$ref": "#/components/schemas/cache-rules_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Settings"], "x-api-token-group": ["Zone Settings Write", "Zone Write"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
