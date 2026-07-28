---
title: List mTLS certificates
page_id: operation-get-accounts-account-id-mtls-certificates-bab2a7b6
path: operations/mtls-certificate-management
description: Lists all mTLS certificates uploaded to your account, such as Bring Your Own CA (BYO-CA) for mTLS. To list certificates issued by the Cloudflare managed CA, use the [List Client Certificates endpoint](/api/resources/client_certificates/methods/list/).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/mtls_certificates
operation_ids:
    - m-tls-certificate-management-list-m-tls-certificates
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List mTLS certificates

`GET /accounts/{account_id}/mtls_certificates`

Operation ID: `m-tls-certificate-management-list-m-tls-certificates`

Lists all mTLS certificates uploaded to your account, such as Bring Your Own CA (BYO-CA) for mTLS. To list certificates issued by the Cloudflare managed CA, use the [List Client Certificates endpoint](/api/resources/client_certificates/methods/list/).

## Definition

```yaml
{"operationId": "m-tls-certificate-management-list-m-tls-certificates", "summary": "List mTLS certificates", "description": "Lists all mTLS certificates uploaded to your account, such as Bring Your Own CA (BYO-CA) for mTLS. To list certificates issued by the Cloudflare managed CA, use the [List Client Certificates endpoint](/api/resources/client_certificates/methods/list/).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "type", "in": "query", "description": "Filters results by certificate type. Multiple types can be comma-separated.", "schema": {"type": "array", "items": {"enum": ["custom", "gateway_managed", "access_managed"], "type": "string"}, "example": ["custom"]}, "explode": false, "style": "form"}], "responses": {"200": {"description": "List mTLS certificates response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_collection-5"}}}}, "4XX": {"description": "List mTLS certificates response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_collection-5"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["mTLS Certificate Management"], "x-api-token-group": ["Account: SSL and Certificates Read", "Account: SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "mtls-certificates", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
