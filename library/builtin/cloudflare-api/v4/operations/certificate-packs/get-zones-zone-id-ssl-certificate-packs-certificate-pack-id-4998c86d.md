---
title: Get Certificate Pack
page_id: operation-get-zones-zone-id-ssl-certificate-packs-certificate-pack-id-2d0c370b
path: operations/certificate-packs
description: For a given zone, get a certificate pack.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/ssl/certificate_packs/{certificate_pack_id}
operation_ids:
    - certificate-packs-get-certificate-pack
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Certificate Pack

`GET /zones/{zone_id}/ssl/certificate_packs/{certificate_pack_id}`

Operation ID: `certificate-packs-get-certificate-pack`

For a given zone, get a certificate pack.

## Definition

```yaml
{"operationId": "certificate-packs-get-certificate-pack", "summary": "Get Certificate Pack", "description": "For a given zone, get a certificate pack.", "parameters": [{"name": "certificate_pack_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "responses": {"200": {"description": "Get Certificate Pack response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_pack_response_single"}}}}, "4XX": {"description": "Get Certificate Pack response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_pack_response_single"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Certificate Packs"], "x-api-token-group": ["SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPermissionsRequired": {"enum": ["#ssl:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ssl.certificate-packs", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
