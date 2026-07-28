---
title: Custom CSR Details
page_id: operation-get-accounts-account-id-custom-csrs-custom-csr-id-da52df4f
path: operations/custom-csrs-for-an-account
description: Retrieve details for a specific custom Certificate Signing Request (CSR).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/custom_csrs/{custom_csr_id}
operation_ids:
    - custom-csrs-for-an-account-custom-csr-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Custom CSR Details

`GET /accounts/{account_id}/custom_csrs/{custom_csr_id}`

Operation ID: `custom-csrs-for-an-account-custom-csr-details`

Retrieve details for a specific custom Certificate Signing Request (CSR).

## Definition

```yaml
{"operationId": "custom-csrs-for-an-account-custom-csr-details", "summary": "Custom CSR Details", "description": "Retrieve details for a specific custom Certificate Signing Request (CSR).", "parameters": [{"name": "custom_csr_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_csr_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "responses": {"200": {"description": "Custom CSR Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_csr_response_single-2"}}}}, "4XX": {"description": "Custom CSR Details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_csr_response_single-2"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Custom CSRs for an Account"], "x-api-token-group": ["Account: SSL and Certificates Read", "Account: SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}}
```
