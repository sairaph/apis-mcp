---
title: Edit Universal SSL Settings
page_id: operation-patch-zones-zone-id-ssl-universal-settings-ea992c4b
path: operations/universal-ssl-settings-for-a-zone
description: Patch Universal SSL Settings for a Zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/ssl/universal/settings
operation_ids:
    - universal-ssl-settings-for-a-zone-edit-universal-ssl-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Edit Universal SSL Settings

`PATCH /zones/{zone_id}/ssl/universal/settings`

Operation ID: `universal-ssl-settings-for-a-zone-edit-universal-ssl-settings`

Patch Universal SSL Settings for a Zone.

## Definition

```yaml
{"operationId": "universal-ssl-settings-for-a-zone-edit-universal-ssl-settings", "summary": "Edit Universal SSL Settings", "description": "Patch Universal SSL Settings for a Zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_universal"}}}}, "responses": {"200": {"description": "Edit Universal SSL Settings response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_ssl_universal_settings_response"}}}}, "4XX": {"description": "Edit Universal SSL Settings response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_ssl_universal_settings_response"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Universal SSL Settings for a Zone"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:read", "#ssl:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ssl.universal.settings", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
