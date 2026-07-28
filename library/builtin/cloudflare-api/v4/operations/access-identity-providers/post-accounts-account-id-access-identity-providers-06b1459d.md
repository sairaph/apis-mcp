---
title: Add an Access identity provider
page_id: operation-post-accounts-account-id-access-identity-providers-04d9764c
path: operations/access-identity-providers
description: Adds a new identity provider to Access.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/access/identity_providers
operation_ids:
    - access-identity-providers-add-an-access-identity-provider
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Add an Access identity provider

`POST /accounts/{account_id}/access/identity_providers`

Operation ID: `access-identity-providers-add-an-access-identity-provider`

Adds a new identity provider to Access.

## Definition

```yaml
{"operationId": "access-identity-providers-add-an-access-identity-provider", "summary": "Add an Access identity provider", "description": "Adds a new identity provider to Access.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_identity-providers"}}}}, "responses": {"201": {"description": "Add an Access identity provider response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-3"}}}}, "4XX": {"description": "Add an Access identity provider response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access identity providers"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.identity-providers", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
