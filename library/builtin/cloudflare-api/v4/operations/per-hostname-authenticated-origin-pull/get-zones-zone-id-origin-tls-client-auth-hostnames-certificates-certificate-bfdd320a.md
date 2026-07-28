---
title: Get the Hostname Client Certificate
page_id: operation-get-zones-zone-id-origin-tls-client-auth-hostnames-certificates-certific-c756ff9a
path: operations/per-hostname-authenticated-origin-pull
description: Get the certificate by ID to be used for client authentication on a hostname.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/origin_tls_client_auth/hostnames/certificates/{certificate_id}
operation_ids:
    - per-hostname-authenticated-origin-pull-get-the-hostname-client-certificate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get the Hostname Client Certificate

`GET /zones/{zone_id}/origin_tls_client_auth/hostnames/certificates/{certificate_id}`

Operation ID: `per-hostname-authenticated-origin-pull-get-the-hostname-client-certificate`

Get the certificate by ID to be used for client authentication on a hostname.

## Definition

```yaml
{"operationId": "per-hostname-authenticated-origin-pull-get-the-hostname-client-certificate", "summary": "Get the Hostname Client Certificate", "description": "Get the certificate by ID to be used for client authentication on a hostname.", "parameters": [{"name": "certificate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "responses": {"200": {"description": "Get the Hostname Client Certificate response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_single-4"}}}}, "4XX": {"description": "Get the Hostname Client Certificate response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_single-4"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Per-hostname Authenticated Origin Pull"], "x-api-token-group": ["SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "origin-tls-client-auth.hostname-certificates", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
