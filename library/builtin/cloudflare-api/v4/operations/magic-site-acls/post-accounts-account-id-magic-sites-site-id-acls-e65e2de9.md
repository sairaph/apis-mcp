---
title: Create a new Site ACL
page_id: operation-post-accounts-account-id-magic-sites-site-id-acls-8f32ff87
path: operations/magic-site-acls
description: Creates a new Site ACL.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/magic/sites/{site_id}/acls
operation_ids:
    - magic-site-acls-create-acl
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a new Site ACL

`POST /accounts/{account_id}/magic/sites/{site_id}/acls`

Operation ID: `magic-site-acls-create-acl`

Creates a new Site ACL.

## Definition

```yaml
{"operationId": "magic-site-acls-create-acl", "summary": "Create a new Site ACL", "description": "Creates a new Site ACL.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_acls_add_single_request"}}}}, "responses": {"200": {"description": "Create Site ACL response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_acl_single_response"}}}}, "4XX": {"description": "Create Site ACL response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Site ACLs"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.sites.acls", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
