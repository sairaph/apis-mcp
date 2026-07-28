---
title: Add a new SSH Certificate Authority (CA)
page_id: operation-post-accounts-account-id-access-gateway-ca-634ebb81
path: operations/gateway-ca
description: Adds a new SSH Certificate Authority (CA).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/access/gateway_ca
operation_ids:
    - access-gateway-ca-add-an-SSH-ca
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Add a new SSH Certificate Authority (CA)

`POST /accounts/{account_id}/access/gateway_ca`

Operation ID: `access-gateway-ca-add-an-SSH-ca`

Adds a new SSH Certificate Authority (CA).

## Definition

```yaml
{"operationId": "access-gateway-ca-add-an-SSH-ca", "summary": "Add a new SSH Certificate Authority (CA)", "description": "Adds a new SSH Certificate Authority (CA).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"201": {"description": "Add a new SSH Certificate Authority (CA) response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-7"}}}}, "4XX": {"description": "Add a new SSH Certificate Authority (CA) response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Gateway CA"], "x-api-token-group": ["Access: SSH Auditing Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.gateway-ca", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
