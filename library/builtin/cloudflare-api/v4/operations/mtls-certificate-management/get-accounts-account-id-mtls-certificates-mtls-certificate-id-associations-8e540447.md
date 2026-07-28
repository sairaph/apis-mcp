---
title: List mTLS certificate associations
page_id: operation-get-accounts-account-id-mtls-certificates-mtls-certificate-id-associatio-df453d4d
path: operations/mtls-certificate-management
description: Lists all active associations between the certificate and Cloudflare services.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/mtls_certificates/{mtls_certificate_id}/associations
operation_ids:
    - m-tls-certificate-management-list-m-tls-certificate-associations
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List mTLS certificate associations

`GET /accounts/{account_id}/mtls_certificates/{mtls_certificate_id}/associations`

Operation ID: `m-tls-certificate-management-list-m-tls-certificate-associations`

Lists all active associations between the certificate and Cloudflare services.

## Definition

```yaml
{"operationId": "m-tls-certificate-management-list-m-tls-certificate-associations", "summary": "List mTLS certificate associations", "description": "Lists all active associations between the certificate and Cloudflare services.", "parameters": [{"name": "mtls_certificate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "responses": {"200": {"description": "List mTLS certificate associations response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_association_response_collection"}}}}, "4XX": {"description": "List mTLS certificate associations response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_association_response_collection"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["mTLS Certificate Management"], "x-api-token-group": ["Account: SSL and Certificates Read", "Account: SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "mtls-certificates.associations", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
