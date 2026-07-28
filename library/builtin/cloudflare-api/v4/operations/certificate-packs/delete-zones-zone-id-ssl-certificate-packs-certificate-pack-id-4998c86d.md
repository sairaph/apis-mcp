---
title: Delete Advanced Certificate Manager Certificate Pack
page_id: operation-delete-zones-zone-id-ssl-certificate-packs-certificate-pack-id-4354ca6d
path: operations/certificate-packs
description: For a given zone, delete an advanced certificate pack.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/ssl/certificate_packs/{certificate_pack_id}
operation_ids:
    - certificate-packs-delete-advanced-certificate-manager-certificate-pack
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Advanced Certificate Manager Certificate Pack

`DELETE /zones/{zone_id}/ssl/certificate_packs/{certificate_pack_id}`

Operation ID: `certificate-packs-delete-advanced-certificate-manager-certificate-pack`

For a given zone, delete an advanced certificate pack.

## Definition

```yaml
{"operationId": "certificate-packs-delete-advanced-certificate-manager-certificate-pack", "summary": "Delete Advanced Certificate Manager Certificate Pack", "description": "For a given zone, delete an advanced certificate pack.", "parameters": [{"name": "certificate_pack_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Advanced Certificate Manager Certificate Pack response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_delete_advanced_certificate_pack_response_single"}}}}, "4XX": {"description": "Delete Advanced Certificate Manager Certificate Pack response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_delete_advanced_certificate_pack_response_single"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Certificate Packs"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:read", "#ssl:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ssl.certificate-packs", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
