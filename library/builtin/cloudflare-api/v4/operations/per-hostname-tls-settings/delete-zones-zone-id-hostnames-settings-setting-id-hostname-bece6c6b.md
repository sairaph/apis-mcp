---
title: Delete TLS setting for hostname
page_id: operation-delete-zones-zone-id-hostnames-settings-setting-id-hostname-20510a76
path: operations/per-hostname-tls-settings
description: Delete the tls setting value for the hostname.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/hostnames/settings/{setting_id}/{hostname}
operation_ids:
    - per-hostname-tls-settings-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete TLS setting for hostname

`DELETE /zones/{zone_id}/hostnames/settings/{setting_id}/{hostname}`

Operation ID: `per-hostname-tls-settings-delete`

Delete the tls setting value for the hostname.

## Definition

```yaml
{"operationId": "per-hostname-tls-settings-delete", "summary": "Delete TLS setting for hostname", "description": "Delete the tls setting value for the hostname.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "setting_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_setting_id"}}, {"name": "hostname", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_hostname-3"}}], "responses": {"200": {"description": "Delete TLS setting for hostname response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_per_hostname_settings_response_delete"}}}}, "4XX": {"description": "Delete TLS setting for hostname response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_per_hostname_settings_response_delete"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Per-Hostname TLS Settings"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "hostnames.settings.tls", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
