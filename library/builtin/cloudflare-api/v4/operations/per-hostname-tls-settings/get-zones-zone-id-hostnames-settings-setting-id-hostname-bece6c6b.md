---
title: Get TLS setting for hostname
page_id: operation-get-zones-zone-id-hostnames-settings-setting-id-hostname-a8f0f246
path: operations/per-hostname-tls-settings
description: Get the requested TLS setting for the hostname.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/hostnames/settings/{setting_id}/{hostname}
operation_ids:
    - per-hostname-tls-settings-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get TLS setting for hostname

`GET /zones/{zone_id}/hostnames/settings/{setting_id}/{hostname}`

Operation ID: `per-hostname-tls-settings-get`

Get the requested TLS setting for the hostname.

## Definition

```yaml
{"operationId": "per-hostname-tls-settings-get", "summary": "Get TLS setting for hostname", "description": "Get the requested TLS setting for the hostname.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "setting_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_setting_id"}}, {"name": "hostname", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_hostname-3"}}], "responses": {"200": {"description": "Get TLS setting for hostname response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_per_hostname_settings_response"}}}}, "4XX": {"description": "Get TLS setting for hostname response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_per_hostname_settings_response"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Per-Hostname TLS Settings"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "hostnames.settings.tls-single", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
