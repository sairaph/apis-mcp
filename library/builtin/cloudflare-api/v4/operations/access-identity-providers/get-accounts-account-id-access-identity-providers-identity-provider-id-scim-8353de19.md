---
title: List SCIM Group resources
page_id: operation-get-accounts-account-id-access-identity-providers-identity-provider-id-s-2240e222
path: operations/access-identity-providers
description: Lists SCIM Group resources synced to Cloudflare via the System for Cross-domain Identity Management (SCIM).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/identity_providers/{identity_provider_id}/scim/groups
operation_ids:
    - access-identity-providers-list-scim-group-resources
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List SCIM Group resources

`GET /accounts/{account_id}/access/identity_providers/{identity_provider_id}/scim/groups`

Operation ID: `access-identity-providers-list-scim-group-resources`

Lists SCIM Group resources synced to Cloudflare via the System for Cross-domain Identity Management (SCIM).

## Definition

```yaml
{"operationId": "access-identity-providers-list-scim-group-resources", "summary": "List SCIM Group resources", "description": "Lists SCIM Group resources synced to Cloudflare via the System for Cross-domain Identity Management (SCIM).", "parameters": [{"name": "identity_provider_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"name": "cf_resource_id", "in": "query", "schema": {"$ref": "#/components/schemas/access_cf_resource_id"}, "explode": true, "style": "form"}, {"name": "idp_resource_id", "in": "query", "schema": {"$ref": "#/components/schemas/access_idp_resource_id"}, "explode": true, "style": "form"}, {"name": "name", "in": "query", "schema": {"$ref": "#/components/schemas/access_name-4"}}, {"$ref": "#/components/parameters/access_page"}, {"name": "per_page", "in": "query", "schema": {"description": "Number of results per page.", "type": "integer", "default": 100, "maximum": 100}}], "responses": {"200": {"description": "List SCIM Group resources response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_scim_groups_response"}}}}, "4XX": {"description": "List SCIM Group resources response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access identity providers"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Write", "Access: Organizations, Identity Providers, and Groups Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.identity-providers.scim.groups", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
