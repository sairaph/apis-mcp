---
title: Update Google Tag Gateway configuration
page_id: operation-put-zones-zone-id-settings-google-tag-gateway-config-c1ccab3c
path: operations/google-tag-gateway
description: Updates the Google Tag Gateway configuration for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/settings/google-tag-gateway/config
operation_ids:
    - zone-settings-change-google-tag-gateway-config
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Google Tag Gateway configuration

`PUT /zones/{zone_id}/settings/google-tag-gateway/config`

Operation ID: `zone-settings-change-google-tag-gateway-config`

Updates the Google Tag Gateway configuration for a zone.

## Definition

```yaml
{"operationId": "zone-settings-change-google-tag-gateway-config", "summary": "Update Google Tag Gateway configuration", "description": "Updates the Google Tag Gateway configuration for a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/google-tag-gateway_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/google-tag-gateway_google-tag-gateway-config"}}}}, "responses": {"200": {"description": "Update Google Tag Gateway configuration response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/google-tag-gateway_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/google-tag-gateway_google-tag-gateway-config"}}, "type": "object"}]}}}}, "4XX": {"description": "Update Google Tag Gateway configuration response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/google-tag-gateway_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Google Tag Gateway"], "x-api-token-group": ["Zaraz Edit", "Zaraz Admin"], "x-cfPermissionsRequired": {"enum": ["#zaraz:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "google-tag-gateway.config", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
