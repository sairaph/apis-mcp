---
title: Deactivate a Zero Trust certificate
page_id: operation-post-accounts-account-id-gateway-certificates-certificate-id-deactivate-d1052649
path: operations/zero-trust-certificates
description: Unbind a single Zero Trust certificate from the edge.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/gateway/certificates/{certificate_id}/deactivate
operation_ids:
    - zero-trust-certificates-deactivate-zero-trust-certificate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Deactivate a Zero Trust certificate

`POST /accounts/{account_id}/gateway/certificates/{certificate_id}/deactivate`

Operation ID: `zero-trust-certificates-deactivate-zero-trust-certificate`

Unbind a single Zero Trust certificate from the edge.

## Definition

```yaml
{"operationId": "zero-trust-certificates-deactivate-zero-trust-certificate", "summary": "Deactivate a Zero Trust certificate", "description": "Unbind a single Zero Trust certificate from the edge.", "parameters": [{"name": "certificate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"201": {"description": "Deactivate Zero Trust certificate details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_single_response"}}}}, "4XX": {"description": "Deactivate Zero Trust certificate details response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_single_response"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust certificates"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.certificates", "x-fern-sdk-method-name": "deactivate", "x-forge-hidden": true}
```
