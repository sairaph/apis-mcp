---
title: Delete Site
page_id: operation-delete-accounts-account-id-magic-sites-site-id-3eee928c
path: operations/magic-sites
description: Remove a specific Site.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/magic/sites/{site_id}
operation_ids:
    - magic-sites-delete-site
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Site

`DELETE /accounts/{account_id}/magic/sites/{site_id}`

Operation ID: `magic-sites-delete-site`

Remove a specific Site.

## Definition

```yaml
{"operationId": "magic-sites-delete-site", "summary": "Delete Site", "description": "Remove a specific Site.", "parameters": [{"name": "site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Site response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_site_deleted_response"}}}}, "4XX": {"description": "Delete Site response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Sites"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPermissionsRequired": {"enum": ["#organization:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.sites", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
