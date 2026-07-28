---
title: Patch Auto-Origin TLS KEX enrollment status for the given zone
page_id: operation-patch-zones-zone-id-settings-auto-origin-tls-kex-27732bf2
path: operations/origin-tls
description: 'Enable or disable Auto-Origin TLS KEX selection for the zone by sending `{"enabled": true}` or `{"enabled": false}`. When enabled, Cloudflare runs a periodic scan of the zone''s origins to determine the preferred key-exchange algorithm and writes that preference to the edge so it is sent first in the TLS ClientHello to the origin.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/settings/auto_origin_tls_kex
operation_ids:
    - ssl-detector-auto-origin-tls-kex-patch-enrollment
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch Auto-Origin TLS KEX enrollment status for the given zone

`PATCH /zones/{zone_id}/settings/auto_origin_tls_kex`

Operation ID: `ssl-detector-auto-origin-tls-kex-patch-enrollment`

Enable or disable Auto-Origin TLS KEX selection for the zone by sending `{"enabled": true}` or `{"enabled": false}`. When enabled, Cloudflare runs a periodic scan of the zone's origins to determine the preferred key-exchange algorithm and writes that preference to the edge so it is sent first in the TLS ClientHello to the origin.

## Definition

```yaml
{"operationId": "ssl-detector-auto-origin-tls-kex-patch-enrollment", "summary": "Patch Auto-Origin TLS KEX enrollment status for the given zone", "description": "Enable or disable Auto-Origin TLS KEX selection for the zone by sending `{\"enabled\": true}` or `{\"enabled\": false}`. When enabled, Cloudflare runs a periodic scan of the zone's origins to determine the preferred key-exchange algorithm and writes that preference to the edge so it is sent first in the TLS ClientHello to the origin.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cache_auto_origin_tls_kex_patch"}}}}, "responses": {"200": {"description": "Patch Auto-Origin TLS KEX enrollment status response.", "content": {"application/json": {"examples": {"Disabled": {"$ref": "#/components/examples/cache_auto_origin_tls_kex_disabled_response"}, "Enabled": {"$ref": "#/components/examples/cache_auto_origin_tls_kex_enabled_response"}}, "schema": {"$ref": "#/components/schemas/cache_auto_origin_tls_kex_response"}}}}, "4XX": {"description": "Patch Auto-Origin TLS KEX enrollment status failure.", "content": {"application/json": {"examples": {"Error": {"$ref": "#/components/examples/cache_auto_origin_tls_kex_error_response"}}, "schema": {"$ref": "#/components/schemas/cache_auto_origin_tls_kex_failure_response"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Origin TLS"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
