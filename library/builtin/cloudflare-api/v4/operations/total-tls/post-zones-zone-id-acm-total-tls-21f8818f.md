---
title: Enable or Disable Total TLS
page_id: operation-post-zones-zone-id-acm-total-tls-859eda48
path: operations/total-tls
description: Set Total TLS Settings or disable the feature for a Zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/acm/total_tls
operation_ids:
    - total-tls-enable-or-disable-total-tls
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Enable or Disable Total TLS

`POST /zones/{zone_id}/acm/total_tls`

Operation ID: `total-tls-enable-or-disable-total-tls`

Set Total TLS Settings or disable the feature for a Zone.

## Definition

```yaml
{"operationId": "total-tls-enable-or-disable-total-tls", "summary": "Enable or Disable Total TLS", "description": "Set Total TLS Settings or disable the feature for a Zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"certificate_authority": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_authority-3"}, "enabled": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_enabled-3"}}, "required": ["enabled"]}}}}, "responses": {"200": {"description": "Enable or Disable Total TLS response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_total_tls_settings_response"}}}}, "4XX": {"description": "Enable or Disable Total TLS response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_total_tls_settings_response"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Total TLS"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:read", "#ssl:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "acm.total-tls", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
