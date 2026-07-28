---
title: List Zero Trust certificates
page_id: operation-get-accounts-account-id-gateway-certificates-bc44085a
path: operations/zero-trust-certificates
description: List all Zero Trust certificates for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/gateway/certificates
operation_ids:
    - zero-trust-certificates-list-zero-trust-certificates
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Zero Trust certificates

`GET /accounts/{account_id}/gateway/certificates`

Operation ID: `zero-trust-certificates-list-zero-trust-certificates`

List all Zero Trust certificates for an account.

## Definition

```yaml
{"operationId": "zero-trust-certificates-list-zero-trust-certificates", "summary": "List Zero Trust certificates", "description": "List all Zero Trust certificates for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "responses": {"200": {"description": "Lists Zero Trust certificates response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_response_collection"}}}}, "4XX": {"description": "Lists Zero Trust certificates response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_response_collection"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust certificates"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.certificates", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
