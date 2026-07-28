---
title: Set Enablement for Zone
page_id: operation-put-zones-zone-id-origin-tls-client-auth-settings-6f460939
path: operations/zone-level-authenticated-origin-pulls
description: Enable or disable zone-level authenticated origin pulls. 'enabled' should be set true either before/after the certificate is uploaded to see the certificate in use.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/origin_tls_client_auth/settings
operation_ids:
    - zone-level-authenticated-origin-pulls-set-enablement-for-zone
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Set Enablement for Zone

`PUT /zones/{zone_id}/origin_tls_client_auth/settings`

Operation ID: `zone-level-authenticated-origin-pulls-set-enablement-for-zone`

Enable or disable zone-level authenticated origin pulls. 'enabled' should be set true either before/after the certificate is uploaded to see the certificate in use.

## Definition

```yaml
{"operationId": "zone-level-authenticated-origin-pulls-set-enablement-for-zone", "summary": "Set Enablement for Zone", "description": "Enable or disable zone-level authenticated origin pulls. 'enabled' should be set true either before/after the certificate is uploaded to see the certificate in use.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"enabled": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_enabled-4"}}, "required": ["enabled"]}}}}, "responses": {"200": {"description": "Set Enablement for Zone response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_enabled_response"}}}}, "4XX": {"description": "Set Enablement for Zone response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_enabled_response"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone-Level Authenticated Origin Pulls"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "origin-tls-client-auth.settings", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
