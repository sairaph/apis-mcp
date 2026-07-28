---
title: Upload a Hostname Client Certificate
page_id: operation-post-zones-zone-id-origin-tls-client-auth-hostnames-certificates-4f70168c
path: operations/per-hostname-authenticated-origin-pull
description: Upload a certificate to be used for client authentication on a hostname. 10 hostname certificates per zone are allowed.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/origin_tls_client_auth/hostnames/certificates
operation_ids:
    - per-hostname-authenticated-origin-pull-upload-a-hostname-client-certificate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Upload a Hostname Client Certificate

`POST /zones/{zone_id}/origin_tls_client_auth/hostnames/certificates`

Operation ID: `per-hostname-authenticated-origin-pull-upload-a-hostname-client-certificate`

Upload a certificate to be used for client authentication on a hostname. 10 hostname certificates per zone are allowed.

## Definition

```yaml
{"operationId": "per-hostname-authenticated-origin-pull-upload-a-hostname-client-certificate", "summary": "Upload a Hostname Client Certificate", "description": "Upload a certificate to be used for client authentication on a hostname. 10 hostname certificates per zone are allowed.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"certificate": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate-6"}, "private_key": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_private_key-3"}}, "required": ["certificate", "private_key"]}}}}, "responses": {"200": {"description": "Upload a Hostname Client Certificate response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_single-4"}}}}, "4XX": {"description": "Upload a Hostname Client Certificate response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_single-4"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Per-hostname Authenticated Origin Pull"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "origin-tls-client-auth.hostname-certificates", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
