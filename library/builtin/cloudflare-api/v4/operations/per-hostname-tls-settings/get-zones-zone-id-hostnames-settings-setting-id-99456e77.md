---
title: List TLS setting for hostnames
page_id: operation-get-zones-zone-id-hostnames-settings-setting-id-44b2a462
path: operations/per-hostname-tls-settings
description: List the requested TLS setting for the hostnames under this zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/hostnames/settings/{setting_id}
operation_ids:
    - per-hostname-tls-settings-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List TLS setting for hostnames

`GET /zones/{zone_id}/hostnames/settings/{setting_id}`

Operation ID: `per-hostname-tls-settings-list`

List the requested TLS setting for the hostnames under this zone.

## Definition

```yaml
{"operationId": "per-hostname-tls-settings-list", "summary": "List TLS setting for hostnames", "description": "List the requested TLS setting for the hostnames under this zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "setting_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_setting_id"}}], "responses": {"200": {"description": "List per-hostname TLS settings response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_per_hostname_settings_response_collection"}}}}, "4XX": {"description": "List per-hostname TLS settings response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_per_hostname_settings_response_collection"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Per-Hostname TLS Settings"], "x-api-token-group": ["SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPermissionsRequired": {"enum": ["#ssl:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "hostnames.settings.tls", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
