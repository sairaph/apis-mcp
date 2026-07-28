---
title: Get Automatic SSL/TLS enrollment status for the given zone
page_id: operation-get-zones-zone-id-settings-ssl-automatic-mode-ae2d6214
path: operations/automatic-ssl-tls
description: If the system is enabled, the response will include next_scheduled_scan, representing the next time this zone will be scanned and the zone's ssl/tls encryption mode is potentially upgraded by the system. If the system is disabled, next_scheduled_scan will not be present in the response body.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/settings/ssl_automatic_mode
operation_ids:
    - ssl-detector-automatic-mode-get-enrollment
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Automatic SSL/TLS enrollment status for the given zone

`GET /zones/{zone_id}/settings/ssl_automatic_mode`

Operation ID: `ssl-detector-automatic-mode-get-enrollment`

If the system is enabled, the response will include next_scheduled_scan, representing the next time this zone will be scanned and the zone's ssl/tls encryption mode is potentially upgraded by the system. If the system is disabled, next_scheduled_scan will not be present in the response body.

## Definition

```yaml
{"operationId": "ssl-detector-automatic-mode-get-enrollment", "summary": "Get Automatic SSL/TLS enrollment status for the given zone", "description": "If the system is enabled, the response will include next_scheduled_scan, representing the next time this zone will be scanned and the zone's ssl/tls encryption mode is potentially upgraded by the system. If the system is disabled, next_scheduled_scan will not be present in the response body.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache_identifier"}}], "responses": {"200": {"description": "Get Automatic SSL/TLS Enrollment status response.", "content": {"application/json": {"examples": {"Disabled": {"$ref": "#/components/examples/cache_dummy_automatic_upgrader_enabled_response"}, "Enabled": {"$ref": "#/components/examples/cache_dummy_automatic_upgrader_disabled_response"}}, "schema": {"$ref": "#/components/schemas/cache_api-response-single-id"}}}}, "4XX": {"description": "Get Automatic SSL/TLS Enrollment status failure.", "content": {"application/json": {"examples": {"Error": {"$ref": "#/components/examples/cache_dummy_error_response"}}, "schema": {"$ref": "#/components/schemas/cache_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Automatic SSL/TLS"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read", "SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
