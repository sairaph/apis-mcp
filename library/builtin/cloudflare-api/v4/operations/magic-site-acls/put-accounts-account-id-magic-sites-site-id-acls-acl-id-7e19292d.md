---
title: Update Site ACL
page_id: operation-put-accounts-account-id-magic-sites-site-id-acls-acl-id-c9cce5d9
path: operations/magic-site-acls
description: Update a specific Site ACL.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/magic/sites/{site_id}/acls/{acl_id}
operation_ids:
    - magic-site-acls-update-acl
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Site ACL

`PUT /accounts/{account_id}/magic/sites/{site_id}/acls/{acl_id}`

Operation ID: `magic-site-acls-update-acl`

Update a specific Site ACL.

## Definition

```yaml
{"operationId": "magic-site-acls-update-acl", "summary": "Update Site ACL", "description": "Update a specific Site ACL.", "parameters": [{"name": "site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "acl_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_acl_update_request"}}}}, "responses": {"200": {"description": "Update Site ACL response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_acl_modified_response"}}}}, "4XX": {"description": "Update Site ACL response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Site ACLs"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPermissionsRequired": {"enum": ["#organization:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.sites.acls", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
