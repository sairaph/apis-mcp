---
title: Delete Zero Trust certificate
page_id: operation-delete-accounts-account-id-gateway-certificates-certificate-id-7c1e5ab2
path: operations/zero-trust-certificates
description: Delete a gateway-managed Zero Trust certificate. You must deactivate the certificate from the edge (inactive) before deleting it.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/gateway/certificates/{certificate_id}
operation_ids:
    - zero-trust-certificates-delete-zero-trust-certificate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Zero Trust certificate

`DELETE /accounts/{account_id}/gateway/certificates/{certificate_id}`

Operation ID: `zero-trust-certificates-delete-zero-trust-certificate`

Delete a gateway-managed Zero Trust certificate. You must deactivate the certificate from the edge (inactive) before deleting it.

## Definition

```yaml
{"operationId": "zero-trust-certificates-delete-zero-trust-certificate", "summary": "Delete Zero Trust certificate", "description": "Delete a gateway-managed Zero Trust certificate. You must deactivate the certificate from the edge (inactive) before deleting it.", "parameters": [{"name": "certificate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Deletes Zero Trust certificate response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_single_response"}}}}, "4XX": {"description": "Deletes Zero Trust certificate response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_single_response"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust certificates"]}
```
