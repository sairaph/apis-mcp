---
title: Update an Access identity provider
page_id: operation-put-accounts-account-id-access-identity-providers-identity-provider-id-38147863
path: operations/access-identity-providers
description: Updates a configured identity provider.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/access/identity_providers/{identity_provider_id}
operation_ids:
    - access-identity-providers-update-an-access-identity-provider
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update an Access identity provider

`PUT /accounts/{account_id}/access/identity_providers/{identity_provider_id}`

Operation ID: `access-identity-providers-update-an-access-identity-provider`

Updates a configured identity provider.

## Definition

```yaml
{"operationId": "access-identity-providers-update-an-access-identity-provider", "summary": "Update an Access identity provider", "description": "Updates a configured identity provider.", "parameters": [{"name": "identity_provider_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_identity-providers"}}}}, "responses": {"200": {"description": "Update an Access identity provider response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-3"}}}}, "4XX": {"description": "Update an Access identity provider response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access identity providers"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.identity-providers", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
