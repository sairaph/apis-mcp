---
title: Get Google Tag Gateway configuration
page_id: operation-get-zones-zone-id-settings-google-tag-gateway-config-d56ad336
path: operations/google-tag-gateway
description: Gets the Google Tag Gateway configuration for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/settings/google-tag-gateway/config
operation_ids:
    - zone-settings-get-google-tag-gateway-config
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Google Tag Gateway configuration

`GET /zones/{zone_id}/settings/google-tag-gateway/config`

Operation ID: `zone-settings-get-google-tag-gateway-config`

Gets the Google Tag Gateway configuration for a zone.

## Definition

```yaml
{"operationId": "zone-settings-get-google-tag-gateway-config", "summary": "Get Google Tag Gateway configuration", "description": "Gets the Google Tag Gateway configuration for a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/google-tag-gateway_identifier"}}], "responses": {"200": {"description": "Get Google Tag Gateway configuration response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/google-tag-gateway_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/google-tag-gateway_google-tag-gateway-config"}}, "type": "object"}]}}}}, "4XX": {"description": "Get Google Tag Gateway configuration response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/google-tag-gateway_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Google Tag Gateway"], "x-api-token-group": ["Zaraz Edit", "Zaraz Read", "Zaraz Admin"], "x-cfPermissionsRequired": {"enum": ["#zaraz:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "google-tag-gateway.config", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
