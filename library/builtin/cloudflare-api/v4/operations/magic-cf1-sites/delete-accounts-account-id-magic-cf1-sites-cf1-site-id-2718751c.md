---
title: Delete CF1 Site
page_id: operation-delete-accounts-account-id-magic-cf1-sites-cf1-site-id-28c22cd8
path: operations/magic-cf1-sites
description: Deletes a specific CF1 Site for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/magic/cf1_sites/{cf1_site_id}
operation_ids:
    - magic-cf1-sites-delete-cf1-site
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete CF1 Site

`DELETE /accounts/{account_id}/magic/cf1_sites/{cf1_site_id}`

Operation ID: `magic-cf1-sites-delete-cf1-site`

Deletes a specific CF1 Site for an account.

## Definition

```yaml
{"operationId": "magic-cf1-sites-delete-cf1-site", "summary": "Delete CF1 Site", "description": "Deletes a specific CF1 Site for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "cf1_site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "responses": {"200": {"description": "Delete CF1 Site response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_cf1_site_single_response"}}}}, "4XX": {"description": "Delete CF1 Site response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic CF1 Sites"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPermissionsRequired": {"enum": ["#organization:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
