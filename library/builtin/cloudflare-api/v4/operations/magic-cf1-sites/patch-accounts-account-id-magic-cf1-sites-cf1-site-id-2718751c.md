---
title: Update CF1 Site
page_id: operation-patch-accounts-account-id-magic-cf1-sites-cf1-site-id-2568bc94
path: operations/magic-cf1-sites
description: Partially updates a specific CF1 Site for an account. Only the fields included in the request body are modified; omitted fields retain their existing values.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/magic/cf1_sites/{cf1_site_id}
operation_ids:
    - magic-cf1-sites-update-cf1-site
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update CF1 Site

`PATCH /accounts/{account_id}/magic/cf1_sites/{cf1_site_id}`

Operation ID: `magic-cf1-sites-update-cf1-site`

Partially updates a specific CF1 Site for an account. Only the fields included in the request body are modified; omitted fields retain their existing values.

## Definition

```yaml
{"operationId": "magic-cf1-sites-update-cf1-site", "summary": "Update CF1 Site", "description": "Partially updates a specific CF1 Site for an account. Only the fields included in the request body are modified; omitted fields retain their existing values.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "cf1_site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_cf1_site_update"}}}}, "responses": {"200": {"description": "Update CF1 Site response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_cf1_site_single_response"}}}}, "4XX": {"description": "Update CF1 Site response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic CF1 Sites"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPermissionsRequired": {"enum": ["#organization:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
