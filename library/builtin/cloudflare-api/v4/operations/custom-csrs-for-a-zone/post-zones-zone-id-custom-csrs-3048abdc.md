---
title: Create Custom CSR
page_id: operation-post-zones-zone-id-custom-csrs-62a4cb13
path: operations/custom-csrs-for-a-zone
description: Generate a new custom Certificate Signing Request (CSR) for a zone. Cloudflare generates and securely stores the private key associated with the CSR. The CSR can then be provided to a Certificate Authority for signing. Once signed, the certificate is uploaded via the Custom SSL endpoint using the CSR ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/custom_csrs
operation_ids:
    - custom-csrs-for-a-zone-create-custom-csr
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Custom CSR

`POST /zones/{zone_id}/custom_csrs`

Operation ID: `custom-csrs-for-a-zone-create-custom-csr`

Generate a new custom Certificate Signing Request (CSR) for a zone. Cloudflare generates and securely stores the private key associated with the CSR. The CSR can then be provided to a Certificate Authority for signing. Once signed, the certificate is uploaded via the Custom SSL endpoint using the CSR ID.

## Definition

```yaml
{"operationId": "custom-csrs-for-a-zone-create-custom-csr", "summary": "Create Custom CSR", "description": "Generate a new custom Certificate Signing Request (CSR) for a zone. Cloudflare generates and securely stores the private key associated with the CSR. The CSR can then be provided to a Certificate Authority for signing. Once signed, the certificate is uploaded via the Custom SSL endpoint using the CSR ID.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_csr_create_request"}}}}, "responses": {"201": {"description": "Create Custom CSR response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_csr_response_single"}}}}, "4XX": {"description": "Create Custom CSR response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_csr_response_single"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Custom CSRs for a Zone"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}}
```
