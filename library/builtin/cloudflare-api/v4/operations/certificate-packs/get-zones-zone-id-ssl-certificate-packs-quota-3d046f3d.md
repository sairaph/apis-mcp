---
title: Get Certificate Pack Quotas
page_id: operation-get-zones-zone-id-ssl-certificate-packs-quota-e820cca9
path: operations/certificate-packs
description: For a given zone, list certificate pack quotas.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/ssl/certificate_packs/quota
operation_ids:
    - certificate-packs-get-certificate-pack-quotas
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Certificate Pack Quotas

`GET /zones/{zone_id}/ssl/certificate_packs/quota`

Operation ID: `certificate-packs-get-certificate-pack-quotas`

For a given zone, list certificate pack quotas.

## Definition

```yaml
{"operationId": "certificate-packs-get-certificate-pack-quotas", "summary": "Get Certificate Pack Quotas", "description": "For a given zone, list certificate pack quotas.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "responses": {"200": {"description": "Get Certificate Pack Quotas response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_pack_quota_response"}}}}, "4XX": {"description": "Get Certificate Pack Quotas response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_pack_quota_response"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Certificate Packs"], "x-api-token-group": ["SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPermissionsRequired": {"enum": ["#ssl:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ssl.certificate-packs.quota", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
