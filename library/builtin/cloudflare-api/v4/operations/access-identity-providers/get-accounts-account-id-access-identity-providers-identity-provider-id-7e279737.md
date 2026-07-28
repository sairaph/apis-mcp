---
title: Get an Access identity provider
page_id: operation-get-accounts-account-id-access-identity-providers-identity-provider-id-ef0dd456
path: operations/access-identity-providers
description: Fetches a configured identity provider.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/identity_providers/{identity_provider_id}
operation_ids:
    - access-identity-providers-get-an-access-identity-provider
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get an Access identity provider

`GET /accounts/{account_id}/access/identity_providers/{identity_provider_id}`

Operation ID: `access-identity-providers-get-an-access-identity-provider`

Fetches a configured identity provider.

## Definition

```yaml
{"operationId": "access-identity-providers-get-an-access-identity-provider", "summary": "Get an Access identity provider", "description": "Fetches a configured identity provider.", "parameters": [{"name": "identity_provider_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "Get an Access identity provider response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-3"}}}}, "4XX": {"description": "Get an Access identity provider response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access identity providers"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Write", "Access: Organizations, Identity Providers, and Groups Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.identity-providers", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
