---
title: Create Custom CSR
page_id: operation-post-accounts-account-id-custom-csrs-555b1cf6
path: operations/custom-csrs-for-an-account
description: Generate a new custom Certificate Signing Request (CSR) for an account. Cloudflare generates and securely stores the private key associated with the CSR.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/custom_csrs
operation_ids:
    - custom-csrs-for-an-account-create-custom-csr
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Custom CSR

`POST /accounts/{account_id}/custom_csrs`

Operation ID: `custom-csrs-for-an-account-create-custom-csr`

Generate a new custom Certificate Signing Request (CSR) for an account. Cloudflare generates and securely stores the private key associated with the CSR.

## Definition

```yaml
{"operationId": "custom-csrs-for-an-account-create-custom-csr", "summary": "Create Custom CSR", "description": "Generate a new custom Certificate Signing Request (CSR) for an account. Cloudflare generates and securely stores the private key associated with the CSR.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_csr_create_request-2"}}}}, "responses": {"201": {"description": "Create Custom CSR response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_csr_response_single-2"}}}}, "4XX": {"description": "Create Custom CSR response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_csr_response_single-2"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Custom CSRs for an Account"], "x-api-token-group": ["Account: SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}}
```
