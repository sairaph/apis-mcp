---
title: Activate a Zero Trust certificate
page_id: operation-post-accounts-account-id-gateway-certificates-certificate-id-activate-f803df56
path: operations/zero-trust-certificates
description: Bind a single Zero Trust certificate to the edge.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/gateway/certificates/{certificate_id}/activate
operation_ids:
    - zero-trust-certificates-activate-zero-trust-certificate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Activate a Zero Trust certificate

`POST /accounts/{account_id}/gateway/certificates/{certificate_id}/activate`

Operation ID: `zero-trust-certificates-activate-zero-trust-certificate`

Bind a single Zero Trust certificate to the edge.

## Definition

```yaml
{"operationId": "zero-trust-certificates-activate-zero-trust-certificate", "summary": "Activate a Zero Trust certificate", "description": "Bind a single Zero Trust certificate to the edge.", "parameters": [{"name": "certificate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"202": {"description": "Activates Zero Trust certificate details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_single_response"}}}}, "4XX": {"description": "Activates Zero Trust certificate details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_single_response"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust certificates"]}
```
