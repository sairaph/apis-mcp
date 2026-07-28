---
title: List Access identity providers
page_id: operation-get-accounts-account-id-access-identity-providers-0085ff57
path: operations/access-identity-providers
description: Lists all configured identity providers.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/identity_providers
operation_ids:
    - access-identity-providers-list-access-identity-providers
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Access identity providers

`GET /accounts/{account_id}/access/identity_providers`

Operation ID: `access-identity-providers-list-access-identity-providers`

Lists all configured identity providers.

## Definition

```yaml
{"operationId": "access-identity-providers-list-access-identity-providers", "summary": "List Access identity providers", "description": "Lists all configured identity providers.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"name": "scim_enabled", "in": "query", "schema": {"description": "Indicates to Access to only retrieve identity providers that have the System for Cross-Domain Identity Management (SCIM) enabled.", "type": "string", "example": true}}, {"$ref": "#/components/parameters/access_page"}, {"name": "per_page", "in": "query", "schema": {"description": "Number of results per page.", "type": "integer", "default": 200, "maximum": 1000}}], "responses": {"200": {"description": "List Access identity providers response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_collection"}}}}, "4XX": {"description": "List Access identity providers response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access identity providers"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Write", "Access: Organizations, Identity Providers, and Groups Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.identity-providers", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
