---
title: Delete an SSH Certificate Authority (CA)
page_id: operation-delete-accounts-account-id-access-gateway-ca-certificate-id-31d29962
path: operations/gateway-ca
description: Deletes an SSH Certificate Authority.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/access/gateway_ca/{certificate_id}
operation_ids:
    - access-gateway-ca-delete-an-SSH-ca
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete an SSH Certificate Authority (CA)

`DELETE /accounts/{account_id}/access/gateway_ca/{certificate_id}`

Operation ID: `access-gateway-ca-delete-an-SSH-ca`

Deletes an SSH Certificate Authority.

## Definition

```yaml
{"operationId": "access-gateway-ca-delete-an-SSH-ca", "summary": "Delete an SSH Certificate Authority (CA)", "description": "Deletes an SSH Certificate Authority.", "parameters": [{"name": "certificate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "Delete an SSH Certificate Authority (CA) response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_id_response"}}}}, "4XX": {"description": "Delete an SSH Certificate Authority (CA) response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Gateway CA"], "x-api-token-group": ["Access: SSH Auditing Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.gateway-ca", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
