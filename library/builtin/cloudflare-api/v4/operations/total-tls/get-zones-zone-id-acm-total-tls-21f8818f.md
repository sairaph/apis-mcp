---
title: Total TLS Settings Details
page_id: operation-get-zones-zone-id-acm-total-tls-53d3e1a4
path: operations/total-tls
description: Get Total TLS Settings for a Zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/acm/total_tls
operation_ids:
    - total-tls-total-tls-settings-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Total TLS Settings Details

`GET /zones/{zone_id}/acm/total_tls`

Operation ID: `total-tls-total-tls-settings-details`

Get Total TLS Settings for a Zone.

## Definition

```yaml
{"operationId": "total-tls-total-tls-settings-details", "summary": "Total TLS Settings Details", "description": "Get Total TLS Settings for a Zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "responses": {"200": {"description": "Total TLS Settings Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_total_tls_settings_response"}}}}, "4XX": {"description": "Total TLS Settings Details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_total_tls_settings_response"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Total TLS"], "x-api-token-group": ["SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPermissionsRequired": {"enum": ["#ssl:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "acm.total-tls", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
