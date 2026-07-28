---
title: Client Certificate Details
page_id: operation-get-zones-zone-id-client-certificates-client-certificate-id-50bb639f
path: operations/api-shield-client-certificates-for-a-zone
description: Get Details for a single mTLS API Shield Client Certificate.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/client_certificates/{client_certificate_id}
operation_ids:
    - client-certificate-for-a-zone-client-certificate-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Client Certificate Details

`GET /zones/{zone_id}/client_certificates/{client_certificate_id}`

Operation ID: `client-certificate-for-a-zone-client-certificate-details`

Get Details for a single mTLS API Shield Client Certificate.

## Definition

```yaml
{"operationId": "client-certificate-for-a-zone-client-certificate-details", "summary": "Client Certificate Details", "description": "Get Details for a single mTLS API Shield Client Certificate.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "client_certificate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "responses": {"200": {"description": "Client Certificate Details Response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_client_certificate_response_single"}}}}, "4XX": {"description": "Client Certificate Details Response Failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["API Shield Client Certificates for a Zone"], "x-api-token-group": ["SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPermissionsRequired": {"enum": ["#ssl:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "client-certificates", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
