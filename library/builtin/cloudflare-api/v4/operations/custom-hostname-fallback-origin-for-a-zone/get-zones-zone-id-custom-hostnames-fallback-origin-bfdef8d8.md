---
title: Get Fallback Origin for Custom Hostnames
page_id: operation-get-zones-zone-id-custom-hostnames-fallback-origin-ce111a4e
path: operations/custom-hostname-fallback-origin-for-a-zone
description: Retrieves the current fallback origin configuration for custom hostnames on a zone. The fallback origin handles traffic when specific custom hostname origins are unavailable.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/custom_hostnames/fallback_origin
operation_ids:
    - custom-hostname-fallback-origin-for-a-zone-get-fallback-origin-for-custom-hostnames
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Fallback Origin for Custom Hostnames

`GET /zones/{zone_id}/custom_hostnames/fallback_origin`

Operation ID: `custom-hostname-fallback-origin-for-a-zone-get-fallback-origin-for-custom-hostnames`

Retrieves the current fallback origin configuration for custom hostnames on a zone. The fallback origin handles traffic when specific custom hostname origins are unavailable.

## Definition

```yaml
{"operationId": "custom-hostname-fallback-origin-for-a-zone-get-fallback-origin-for-custom-hostnames", "summary": "Get Fallback Origin for Custom Hostnames", "description": "Retrieves the current fallback origin configuration for custom hostnames on a zone. The fallback origin handles traffic when specific custom hostname origins are unavailable.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "responses": {"200": {"description": "Get Fallback Origin for Custom Hostnames response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_fallback_origin_response"}}}}, "4XX": {"description": "Get Fallback Origin for Custom Hostnames response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_fallback_origin_response"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Custom Hostname Fallback Origin for a Zone"], "x-api-token-group": ["SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPermissionsRequired": {"enum": ["#ssl:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "custom-hostnames.fallback-origin", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
