---
title: Delete Site ACL
page_id: operation-delete-accounts-account-id-magic-sites-site-id-acls-acl-id-e7b1e52e
path: operations/magic-site-acls
description: Remove a specific Site ACL.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/magic/sites/{site_id}/acls/{acl_id}
operation_ids:
    - magic-site-acls-delete-acl
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Site ACL

`DELETE /accounts/{account_id}/magic/sites/{site_id}/acls/{acl_id}`

Operation ID: `magic-site-acls-delete-acl`

Remove a specific Site ACL.

## Definition

```yaml
{"operationId": "magic-site-acls-delete-acl", "summary": "Delete Site ACL", "description": "Remove a specific Site ACL.", "parameters": [{"name": "site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "acl_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Site ACL response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_acl_deleted_response"}}}}, "4XX": {"description": "Delete Site ACL response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Site ACLs"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPermissionsRequired": {"enum": ["#organization:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.sites.acls", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
