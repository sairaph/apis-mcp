---
title: Get Enablement Setting for Zone
page_id: operation-get-zones-zone-id-origin-tls-client-auth-settings-327e09a0
path: operations/zone-level-authenticated-origin-pulls
description: Get whether zone-level authenticated origin pulls is enabled or not. It is false by default.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/origin_tls_client_auth/settings
operation_ids:
    - zone-level-authenticated-origin-pulls-get-enablement-setting-for-zone
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Enablement Setting for Zone

`GET /zones/{zone_id}/origin_tls_client_auth/settings`

Operation ID: `zone-level-authenticated-origin-pulls-get-enablement-setting-for-zone`

Get whether zone-level authenticated origin pulls is enabled or not. It is false by default.

## Definition

```yaml
{"operationId": "zone-level-authenticated-origin-pulls-get-enablement-setting-for-zone", "summary": "Get Enablement Setting for Zone", "description": "Get whether zone-level authenticated origin pulls is enabled or not. It is false by default.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "responses": {"200": {"description": "Get Enablement Setting for Zone response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_enabled_response"}}}}, "4XX": {"description": "Get Enablement Setting for Zone response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_enabled_response"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone-Level Authenticated Origin Pulls"], "x-api-token-group": ["SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "origin-tls-client-auth.settings", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
