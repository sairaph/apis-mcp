---
title: Get Zero Trust certificate details
page_id: operation-get-accounts-account-id-gateway-certificates-certificate-id-e64422a7
path: operations/zero-trust-certificates
description: Get a single Zero Trust certificate.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/gateway/certificates/{certificate_id}
operation_ids:
    - zero-trust-certificates-zero-trust-certificate-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Zero Trust certificate details

`GET /accounts/{account_id}/gateway/certificates/{certificate_id}`

Operation ID: `zero-trust-certificates-zero-trust-certificate-details`

Get a single Zero Trust certificate.

## Definition

```yaml
{"operationId": "zero-trust-certificates-zero-trust-certificate-details", "summary": "Get Zero Trust certificate details", "description": "Get a single Zero Trust certificate.", "parameters": [{"name": "certificate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "responses": {"200": {"description": "Gets Zero Trust certificate details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_single_response"}}}}, "4XX": {"description": "Gets Zero Trust certificate details response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_single_response"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust certificates"]}
```
