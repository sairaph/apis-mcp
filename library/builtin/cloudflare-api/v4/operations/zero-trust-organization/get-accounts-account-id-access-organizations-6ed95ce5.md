---
title: Get your Zero Trust organization
page_id: operation-get-accounts-account-id-access-organizations-1b581312
path: operations/zero-trust-organization
description: Returns the configuration for your Zero Trust organization.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/organizations
operation_ids:
    - zero-trust-organization-get-your-zero-trust-organization
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get your Zero Trust organization

`GET /accounts/{account_id}/access/organizations`

Operation ID: `zero-trust-organization-get-your-zero-trust-organization`

Returns the configuration for your Zero Trust organization.

## Definition

```yaml
{"operationId": "zero-trust-organization-get-your-zero-trust-organization", "summary": "Get your Zero Trust organization", "description": "Returns the configuration for your Zero Trust organization.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "Get your Zero Trust organization response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response"}}}}, "4XX": {"description": "Get your Zero Trust organization response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust organization"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Revoke", "Access: Organizations, Identity Providers, and Groups Write", "Access: Organizations, Identity Providers, and Groups Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.organizations", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
