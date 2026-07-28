---
title: Replace Custom Certificate and Custom Key In Custom Hostname
page_id: operation-put-zones-zone-id-custom-hostnames-custom-hostname-id-certificate-pack-c-2164f612
path: operations/custom-hostname-for-a-zone
description: Replace a single custom certificate within a certificate pack that contains two bundled certificates. The replacement must adhere to the following constraints. You can only replace an RSA certificate with another RSA certificate or an ECDSA certificate with another ECDSA certificate.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/custom_hostnames/{custom_hostname_id}/certificate_pack/{certificate_pack_id}/certificates/{certificate_id}
operation_ids:
    - custom-hostname-for-a-zone-edit-custom-certificate-custom-hostname
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Replace Custom Certificate and Custom Key In Custom Hostname

`PUT /zones/{zone_id}/custom_hostnames/{custom_hostname_id}/certificate_pack/{certificate_pack_id}/certificates/{certificate_id}`

Operation ID: `custom-hostname-for-a-zone-edit-custom-certificate-custom-hostname`

Replace a single custom certificate within a certificate pack that contains two bundled certificates. The replacement must adhere to the following constraints. You can only replace an RSA certificate with another RSA certificate or an ECDSA certificate with another ECDSA certificate.

## Definition

```yaml
{"operationId": "custom-hostname-for-a-zone-edit-custom-certificate-custom-hostname", "summary": "Replace Custom Certificate and Custom Key In Custom Hostname", "description": "Replace a single custom certificate within a certificate pack that contains two bundled certificates. The replacement must adhere to the following constraints. You can only replace an RSA certificate with another RSA certificate or an ECDSA certificate with another ECDSA certificate.", "parameters": [{"name": "custom_hostname_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "certificate_pack_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "certificate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_cert_and_key"}}}}, "responses": {"202": {"description": "Edit Custom Certificate In a Custom Hostname response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_hostname_response_single"}}}}, "4XX": {"description": "Edit Custom Certificate In a Custom Hostname response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_hostname_response_single"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Custom Hostname for a Zone"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "custom-hostnames.certificate-pack.certificates", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
