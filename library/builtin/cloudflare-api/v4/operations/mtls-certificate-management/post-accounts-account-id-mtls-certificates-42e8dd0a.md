---
title: Upload mTLS certificate
page_id: operation-post-accounts-account-id-mtls-certificates-87582412
path: operations/mtls-certificate-management
description: Upload a certificate that you want to use with mTLS-enabled Cloudflare services, such as Bring Your Own CA (BYO-CA) for mTLS. To create certificates issued by the Cloudflare managed CA, use the [Create Client Certificate endpoint](/api/resources/client_certificates/methods/create/).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/mtls_certificates
operation_ids:
    - m-tls-certificate-management-upload-m-tls-certificate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Upload mTLS certificate

`POST /accounts/{account_id}/mtls_certificates`

Operation ID: `m-tls-certificate-management-upload-m-tls-certificate`

Upload a certificate that you want to use with mTLS-enabled Cloudflare services, such as Bring Your Own CA (BYO-CA) for mTLS. To create certificates issued by the Cloudflare managed CA, use the [Create Client Certificate endpoint](/api/resources/client_certificates/methods/create/).

## Definition

```yaml
{"operationId": "m-tls-certificate-management-upload-m-tls-certificate", "summary": "Upload mTLS certificate", "description": "Upload a certificate that you want to use with mTLS-enabled Cloudflare services, such as Bring Your Own CA (BYO-CA) for mTLS. To create certificates issued by the Cloudflare managed CA, use the [Create Client Certificate endpoint](/api/resources/client_certificates/methods/create/).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"ca": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_ca"}, "certificates": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificates-2"}, "name": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_name-2"}, "private_key": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_private_key-4"}}, "required": ["certificates", "ca"]}}}}, "responses": {"200": {"description": "Upload mTLS certificate response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_single_post"}}}}, "4XX": {"description": "Upload mTLS certificate response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_single_post"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["mTLS Certificate Management"], "x-api-token-group": ["Account: SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "mtls-certificates", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
