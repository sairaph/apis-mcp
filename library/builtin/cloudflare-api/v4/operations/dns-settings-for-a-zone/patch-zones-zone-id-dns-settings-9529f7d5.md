---
title: Update DNS Settings
page_id: operation-patch-zones-zone-id-dns-settings-ab5b44ae
path: operations/dns-settings-for-a-zone
description: Update DNS settings for a zone
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/dns_settings
operation_ids:
    - dns-settings-for-a-zone-update-dns-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update DNS Settings

`PATCH /zones/{zone_id}/dns_settings`

Operation ID: `dns-settings-for-a-zone-update-dns-settings`

Update DNS settings for a zone

## Definition

```yaml
{"operationId": "dns-settings-for-a-zone-update-dns-settings", "summary": "Update DNS Settings", "description": "Update DNS settings for a zone", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-settings_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-settings_dns-settings-zone-patch"}}}}, "responses": {"200": {"description": "Show DNS Settings response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-settings_dns_response_single-2"}}}}, "4XX": {"description": "Show DNS Settings response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-settings_dns_response_single-2"}, {"$ref": "#/components/schemas/dns-settings_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNS Settings for a Zone"], "x-api-token-group": ["Zone DNS Settings Write", "DNS Write"], "x-cfPermissionsRequired": {"enum": ["#dns_records:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.settings.zone", "x-fern-sdk-method-name": "edit"}
```
