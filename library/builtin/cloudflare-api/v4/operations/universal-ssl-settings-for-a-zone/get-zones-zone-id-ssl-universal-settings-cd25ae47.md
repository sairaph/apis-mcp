---
title: Universal SSL Settings Details
page_id: operation-get-zones-zone-id-ssl-universal-settings-ebb98003
path: operations/universal-ssl-settings-for-a-zone
description: Get Universal SSL Settings for a Zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/ssl/universal/settings
operation_ids:
    - universal-ssl-settings-for-a-zone-universal-ssl-settings-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Universal SSL Settings Details

`GET /zones/{zone_id}/ssl/universal/settings`

Operation ID: `universal-ssl-settings-for-a-zone-universal-ssl-settings-details`

Get Universal SSL Settings for a Zone.

## Definition

```yaml
{"operationId": "universal-ssl-settings-for-a-zone-universal-ssl-settings-details", "summary": "Universal SSL Settings Details", "description": "Get Universal SSL Settings for a Zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "responses": {"200": {"description": "Universal SSL Settings Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_ssl_universal_settings_response"}}}}, "4XX": {"description": "Universal SSL Settings Details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_ssl_universal_settings_response"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Universal SSL Settings for a Zone"], "x-api-token-group": ["SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPermissionsRequired": {"enum": ["#ssl:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ssl.universal.settings", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
