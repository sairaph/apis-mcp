---
title: Edit SSL Configuration
page_id: operation-patch-zones-zone-id-custom-certificates-custom-certificate-id-72b4e50c
path: operations/custom-ssl-for-a-zone
description: 'Upload a new private key and/or PEM/CRT for the SSL certificate. Note: PATCHing a configuration for sni_custom certificates will result in a new resource id being returned, and the previous one being deleted.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/custom_certificates/{custom_certificate_id}
operation_ids:
    - custom-ssl-for-a-zone-edit-ssl-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Edit SSL Configuration

`PATCH /zones/{zone_id}/custom_certificates/{custom_certificate_id}`

Operation ID: `custom-ssl-for-a-zone-edit-ssl-configuration`

Upload a new private key and/or PEM/CRT for the SSL certificate. Note: PATCHing a configuration for sni_custom certificates will result in a new resource id being returned, and the previous one being deleted.

## Definition

```yaml
{"operationId": "custom-ssl-for-a-zone-edit-ssl-configuration", "summary": "Edit SSL Configuration", "description": "Upload a new private key and/or PEM/CRT for the SSL certificate. Note: PATCHing a configuration for sni_custom certificates will result in a new resource id being returned, and the previous one being deleted.", "parameters": [{"name": "custom_certificate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"bundle_method": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_bundle_method"}, "certificate": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate"}, "custom_csr_id": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_csr_id"}, "deploy": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_deploy"}, "geo_restrictions": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_geo_restrictions"}, "policy": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_policy"}, "private_key": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_private_key"}}}}}}, "responses": {"200": {"description": "Edit SSL Configuration response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_single"}}}}, "4XX": {"description": "Edit SSL Configuration response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_single"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Custom SSL for a Zone"], "x-api-token-group": ["Access: Mutual TLS Certificates Write", "SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "custom-certificates", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
