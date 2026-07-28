---
title: Create Zero Trust certificate
page_id: operation-post-accounts-account-id-gateway-certificates-917fa5f4
path: operations/zero-trust-certificates
description: Create a new Zero Trust certificate.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/gateway/certificates
operation_ids:
    - zero-trust-certificates-create-zero-trust-certificate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Zero Trust certificate

`POST /accounts/{account_id}/gateway/certificates`

Operation ID: `zero-trust-certificates-create-zero-trust-certificate`

Create a new Zero Trust certificate.

## Definition

```yaml
{"operationId": "zero-trust-certificates-create-zero-trust-certificate", "summary": "Create Zero Trust certificate", "description": "Create a new Zero Trust certificate.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_generate-cert-request"}}}}, "responses": {"200": {"description": "Creates Zero Trust certificate response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_single_response"}}}}, "4XX": {"description": "Creates Zero Trust certificate response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_single_response"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust certificates"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.certificates", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
