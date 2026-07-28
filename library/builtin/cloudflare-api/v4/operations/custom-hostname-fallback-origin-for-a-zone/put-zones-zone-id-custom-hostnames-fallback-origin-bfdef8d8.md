---
title: Update Fallback Origin for Custom Hostnames
page_id: operation-put-zones-zone-id-custom-hostnames-fallback-origin-8c0ded49
path: operations/custom-hostname-fallback-origin-for-a-zone
description: Updates the fallback origin configuration for custom hostnames on a zone. Sets the default origin server for custom hostname traffic.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/custom_hostnames/fallback_origin
operation_ids:
    - custom-hostname-fallback-origin-for-a-zone-update-fallback-origin-for-custom-hostnames
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Fallback Origin for Custom Hostnames

`PUT /zones/{zone_id}/custom_hostnames/fallback_origin`

Operation ID: `custom-hostname-fallback-origin-for-a-zone-update-fallback-origin-for-custom-hostnames`

Updates the fallback origin configuration for custom hostnames on a zone. Sets the default origin server for custom hostname traffic.

## Definition

```yaml
{"operationId": "custom-hostname-fallback-origin-for-a-zone-update-fallback-origin-for-custom-hostnames", "summary": "Update Fallback Origin for Custom Hostnames", "description": "Updates the fallback origin configuration for custom hostnames on a zone. Sets the default origin server for custom hostname traffic.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"origin": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_origin"}}, "required": ["origin"]}}}}, "responses": {"200": {"description": "Update Fallback Origin for Custom Hostnames response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_fallback_origin_response"}}}}, "4XX": {"description": "Update Fallback Origin for Custom Hostnames response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_fallback_origin_response"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Custom Hostname Fallback Origin for a Zone"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "custom-hostnames.fallback-origin", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
