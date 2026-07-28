---
title: Delete an Access identity provider
page_id: operation-delete-accounts-account-id-access-identity-providers-identity-provider-i-c3bbfa42
path: operations/access-identity-providers
description: Deletes an identity provider from Access.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/access/identity_providers/{identity_provider_id}
operation_ids:
    - access-identity-providers-delete-an-access-identity-provider
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete an Access identity provider

`DELETE /accounts/{account_id}/access/identity_providers/{identity_provider_id}`

Operation ID: `access-identity-providers-delete-an-access-identity-provider`

Deletes an identity provider from Access.

## Definition

```yaml
{"operationId": "access-identity-providers-delete-an-access-identity-provider", "summary": "Delete an Access identity provider", "description": "Deletes an identity provider from Access.", "parameters": [{"name": "identity_provider_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"202": {"description": "Delete an Access identity provider response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_id_response"}}}}, "4XX": {"description": "Delete an Access identity provider response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access identity providers"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.identity-providers", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
