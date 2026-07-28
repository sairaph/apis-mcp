---
title: Restart Validation or Update Advanced Certificate Manager Certificate Pack
page_id: operation-patch-zones-zone-id-ssl-certificate-packs-certificate-pack-id-1dc3ddf3
path: operations/certificate-packs
description: For a given zone, restart validation or add cloudflare branding for an advanced certificate pack.  The former is only a validation operation for a Certificate Pack in a validation_timed_out status.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/ssl/certificate_packs/{certificate_pack_id}
operation_ids:
    - certificate-packs-restart-validation-for-advanced-certificate-manager-certificate-pack
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Restart Validation or Update Advanced Certificate Manager Certificate Pack

`PATCH /zones/{zone_id}/ssl/certificate_packs/{certificate_pack_id}`

Operation ID: `certificate-packs-restart-validation-for-advanced-certificate-manager-certificate-pack`

For a given zone, restart validation or add cloudflare branding for an advanced certificate pack.  The former is only a validation operation for a Certificate Pack in a validation_timed_out status.

## Definition

```yaml
{"operationId": "certificate-packs-restart-validation-for-advanced-certificate-manager-certificate-pack", "summary": "Restart Validation or Update Advanced Certificate Manager Certificate Pack", "description": "For a given zone, restart validation or add cloudflare branding for an advanced certificate pack.  The former is only a validation operation for a Certificate Pack in a validation_timed_out status.", "parameters": [{"name": "certificate_pack_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"cloudflare_branding": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_cloudflare_branding"}}}}}}, "responses": {"200": {"description": "Restart Validation for Advanced Certificate Manager Certificate Pack response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_advanced_certificate_pack_response_single"}}}}, "4XX": {"description": "Restart Validation for Advanced Certificate Manager Certificate Pack response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_advanced_certificate_pack_response_single"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Certificate Packs"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:read", "#ssl:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ssl.certificate-packs", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
