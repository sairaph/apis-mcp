---
title: Get Auto-Origin TLS KEX enrollment status for the given zone
page_id: operation-get-zones-zone-id-settings-auto-origin-tls-kex-5bfea4de
path: operations/origin-tls
description: When enabled, Cloudflare automatically selects the preferred TLS key-exchange algorithm to use when establishing the TLS connection to the zone's origin, picking from the algorithms permitted by the zone's `origin_tls_compliance_modes` setting. When disabled, the default key-exchange ordering is used.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/settings/auto_origin_tls_kex
operation_ids:
    - ssl-detector-auto-origin-tls-kex-get-enrollment
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Auto-Origin TLS KEX enrollment status for the given zone

`GET /zones/{zone_id}/settings/auto_origin_tls_kex`

Operation ID: `ssl-detector-auto-origin-tls-kex-get-enrollment`

When enabled, Cloudflare automatically selects the preferred TLS key-exchange algorithm to use when establishing the TLS connection to the zone's origin, picking from the algorithms permitted by the zone's `origin_tls_compliance_modes` setting. When disabled, the default key-exchange ordering is used.

## Definition

```yaml
{"operationId": "ssl-detector-auto-origin-tls-kex-get-enrollment", "summary": "Get Auto-Origin TLS KEX enrollment status for the given zone", "description": "When enabled, Cloudflare automatically selects the preferred TLS key-exchange algorithm to use when establishing the TLS connection to the zone's origin, picking from the algorithms permitted by the zone's `origin_tls_compliance_modes` setting. When disabled, the default key-exchange ordering is used.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache_identifier"}}], "responses": {"200": {"description": "Get Auto-Origin TLS KEX enrollment status response.", "content": {"application/json": {"examples": {"Disabled": {"$ref": "#/components/examples/cache_auto_origin_tls_kex_disabled_response"}, "Enabled": {"$ref": "#/components/examples/cache_auto_origin_tls_kex_enabled_response"}}, "schema": {"$ref": "#/components/schemas/cache_auto_origin_tls_kex_response"}}}}, "4XX": {"description": "Get Auto-Origin TLS KEX enrollment status failure.", "content": {"application/json": {"examples": {"Error": {"$ref": "#/components/examples/cache_auto_origin_tls_kex_error_response"}}, "schema": {"$ref": "#/components/schemas/cache_auto_origin_tls_kex_failure_response"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Origin TLS"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
