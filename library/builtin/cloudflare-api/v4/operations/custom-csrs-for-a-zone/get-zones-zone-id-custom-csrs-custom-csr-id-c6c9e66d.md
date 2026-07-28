---
title: Custom CSR Details
page_id: operation-get-zones-zone-id-custom-csrs-custom-csr-id-91a60f19
path: operations/custom-csrs-for-a-zone
description: Retrieve details for a specific custom Certificate Signing Request (CSR).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/custom_csrs/{custom_csr_id}
operation_ids:
    - custom-csrs-for-a-zone-custom-csr-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Custom CSR Details

`GET /zones/{zone_id}/custom_csrs/{custom_csr_id}`

Operation ID: `custom-csrs-for-a-zone-custom-csr-details`

Retrieve details for a specific custom Certificate Signing Request (CSR).

## Definition

```yaml
{"operationId": "custom-csrs-for-a-zone-custom-csr-details", "summary": "Custom CSR Details", "description": "Retrieve details for a specific custom Certificate Signing Request (CSR).", "parameters": [{"name": "custom_csr_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_csr_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "responses": {"200": {"description": "Custom CSR Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_csr_response_single"}}}}, "4XX": {"description": "Custom CSR Details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_csr_response_single"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Custom CSRs for a Zone"], "x-api-token-group": ["SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPermissionsRequired": {"enum": ["#ssl:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}}
```
