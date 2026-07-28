---
title: Show DNS Settings
page_id: operation-get-zones-zone-id-dns-settings-df79a18f
path: operations/dns-settings-for-a-zone
description: Show DNS settings for a zone
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/dns_settings
operation_ids:
    - dns-settings-for-a-zone-list-dns-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Show DNS Settings

`GET /zones/{zone_id}/dns_settings`

Operation ID: `dns-settings-for-a-zone-list-dns-settings`

Show DNS settings for a zone

## Definition

```yaml
{"operationId": "dns-settings-for-a-zone-list-dns-settings", "summary": "Show DNS Settings", "description": "Show DNS settings for a zone", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-settings_identifier"}}], "responses": {"200": {"description": "Show DNS Settings response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-settings_dns_response_single-2"}}}}, "4XX": {"description": "Show DNS Settings response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-settings_dns_response_single-2"}, {"$ref": "#/components/schemas/dns-settings_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNS Settings for a Zone"], "x-api-token-group": ["Zone DNS Settings Write", "Zone DNS Settings Read", "DNS Read", "DNS Write"], "x-cfPermissionsRequired": {"enum": ["#dns_records:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.settings.zone", "x-fern-sdk-method-name": "get"}
```
