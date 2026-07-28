---
title: Delete mTLS certificate
page_id: operation-delete-accounts-account-id-mtls-certificates-mtls-certificate-id-d723f0b2
path: operations/mtls-certificate-management
description: Deletes the mTLS certificate unless the certificate is in use by one or more Cloudflare services.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/mtls_certificates/{mtls_certificate_id}
operation_ids:
    - m-tls-certificate-management-delete-m-tls-certificate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete mTLS certificate

`DELETE /accounts/{account_id}/mtls_certificates/{mtls_certificate_id}`

Operation ID: `m-tls-certificate-management-delete-m-tls-certificate`

Deletes the mTLS certificate unless the certificate is in use by one or more Cloudflare services.

## Definition

```yaml
{"operationId": "m-tls-certificate-management-delete-m-tls-certificate", "summary": "Delete mTLS certificate", "description": "Deletes the mTLS certificate unless the certificate is in use by one or more Cloudflare services.", "parameters": [{"name": "mtls_certificate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete mTLS certificate response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_single-5"}}}}, "4XX": {"description": "Delete mTLS certificate response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_single-5"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["mTLS Certificate Management"], "x-api-token-group": ["Account: SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "mtls-certificates", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
