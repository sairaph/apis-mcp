---
title: List Site ACLs
page_id: operation-get-accounts-account-id-magic-sites-site-id-acls-00f41ca2
path: operations/magic-site-acls
description: Lists Site ACLs associated with an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/sites/{site_id}/acls
operation_ids:
    - magic-site-acls-list-acls
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Site ACLs

`GET /accounts/{account_id}/magic/sites/{site_id}/acls`

Operation ID: `magic-site-acls-list-acls`

Lists Site ACLs associated with an account.

## Definition

```yaml
{"operationId": "magic-site-acls-list-acls", "summary": "List Site ACLs", "description": "Lists Site ACLs associated with an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "responses": {"200": {"description": "List Site ACLs response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_acls_collection_response"}}}}, "4XX": {"description": "List Site ACLs response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Site ACLs"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read", "Magic Transit Read", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.sites.acls", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
