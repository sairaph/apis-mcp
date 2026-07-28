---
title: Get mTLS certificate
page_id: operation-get-accounts-account-id-mtls-certificates-mtls-certificate-id-9fa1d69c
path: operations/mtls-certificate-management
description: Fetches a single mTLS certificate uploaded to your account. To get a certificate issued by the Cloudflare managed CA, use the [Client Certificate Details endpoint](/api/resources/client_certificates/methods/get/).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/mtls_certificates/{mtls_certificate_id}
operation_ids:
    - m-tls-certificate-management-get-m-tls-certificate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get mTLS certificate

`GET /accounts/{account_id}/mtls_certificates/{mtls_certificate_id}`

Operation ID: `m-tls-certificate-management-get-m-tls-certificate`

Fetches a single mTLS certificate uploaded to your account. To get a certificate issued by the Cloudflare managed CA, use the [Client Certificate Details endpoint](/api/resources/client_certificates/methods/get/).

## Definition

```yaml
{"operationId": "m-tls-certificate-management-get-m-tls-certificate", "summary": "Get mTLS certificate", "description": "Fetches a single mTLS certificate uploaded to your account. To get a certificate issued by the Cloudflare managed CA, use the [Client Certificate Details endpoint](/api/resources/client_certificates/methods/get/).", "parameters": [{"name": "mtls_certificate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "responses": {"200": {"description": "Get mTLS certificate response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_single-5"}}}}, "4XX": {"description": "Get mTLS certificate response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_single-5"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["mTLS Certificate Management"], "x-api-token-group": ["Account: SSL and Certificates Read", "Account: SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "mtls-certificates", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
