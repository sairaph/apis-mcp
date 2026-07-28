---
title: List SSH Certificate Authorities (CA)
page_id: operation-get-accounts-account-id-access-gateway-ca-e24578ff
path: operations/gateway-ca
description: Lists SSH Certificate Authorities (CA).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/gateway_ca
operation_ids:
    - access-gateway-ca-list-SSH-ca
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List SSH Certificate Authorities (CA)

`GET /accounts/{account_id}/access/gateway_ca`

Operation ID: `access-gateway-ca-list-SSH-ca`

Lists SSH Certificate Authorities (CA).

## Definition

```yaml
{"operationId": "access-gateway-ca-list-SSH-ca", "summary": "List SSH Certificate Authorities (CA)", "description": "Lists SSH Certificate Authorities (CA).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "List SSH Certificate Authorities (CA) response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_collection-6"}}}}, "4XX": {"description": "List SSH Certificate Authorities (CA) response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Gateway CA"], "x-api-token-group": ["Access: SSH Auditing Write", "Access: SSH Auditing Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.gateway-ca", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
