---
title: Get Certificate Details
page_id: operation-get-zones-zone-id-origin-tls-client-auth-certificate-id-3c910133
path: operations/zone-level-authenticated-origin-pulls
description: Retrieves details for a specific client certificate used in zone-level authenticated origin pulls.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/origin_tls_client_auth/{certificate_id}
operation_ids:
    - zone-level-authenticated-origin-pulls-get-certificate-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Certificate Details

`GET /zones/{zone_id}/origin_tls_client_auth/{certificate_id}`

Operation ID: `zone-level-authenticated-origin-pulls-get-certificate-details`

Retrieves details for a specific client certificate used in zone-level authenticated origin pulls.

## Definition

```yaml
{"operationId": "zone-level-authenticated-origin-pulls-get-certificate-details", "summary": "Get Certificate Details", "description": "Retrieves details for a specific client certificate used in zone-level authenticated origin pulls.", "parameters": [{"name": "certificate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "responses": {"200": {"description": "Get Certificate Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_single-3"}}}}, "4XX": {"description": "Get Certificate Details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_single-3"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone-Level Authenticated Origin Pulls"], "x-api-token-group": ["SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "origin-tls-client-auth.zone-certificates", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
