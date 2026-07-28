---
title: Delete Fallback Origin for Custom Hostnames
page_id: operation-delete-zones-zone-id-custom-hostnames-fallback-origin-fbbbc1d8
path: operations/custom-hostname-fallback-origin-for-a-zone
description: Removes the fallback origin configuration for custom hostnames on a zone. Custom hostnames without specific origins will no longer have a fallback.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/custom_hostnames/fallback_origin
operation_ids:
    - custom-hostname-fallback-origin-for-a-zone-delete-fallback-origin-for-custom-hostnames
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Fallback Origin for Custom Hostnames

`DELETE /zones/{zone_id}/custom_hostnames/fallback_origin`

Operation ID: `custom-hostname-fallback-origin-for-a-zone-delete-fallback-origin-for-custom-hostnames`

Removes the fallback origin configuration for custom hostnames on a zone. Custom hostnames without specific origins will no longer have a fallback.

## Definition

```yaml
{"operationId": "custom-hostname-fallback-origin-for-a-zone-delete-fallback-origin-for-custom-hostnames", "summary": "Delete Fallback Origin for Custom Hostnames", "description": "Removes the fallback origin configuration for custom hostnames on a zone. Custom hostnames without specific origins will no longer have a fallback.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Fallback Origin for Custom Hostnames response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_fallback_origin_response"}}}}, "4XX": {"description": "Delete Fallback Origin for Custom Hostnames response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_fallback_origin_response"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Custom Hostname Fallback Origin for a Zone"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "custom-hostnames.fallback-origin", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
