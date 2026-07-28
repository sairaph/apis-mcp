---
title: Update Site
page_id: operation-put-accounts-account-id-magic-sites-site-id-b1d8e5c4
path: operations/magic-sites
description: Update a specific Site.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/magic/sites/{site_id}
operation_ids:
    - magic-sites-update-site
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Site

`PUT /accounts/{account_id}/magic/sites/{site_id}`

Operation ID: `magic-sites-update-site`

Update a specific Site.

## Definition

```yaml
{"operationId": "magic-sites-update-site", "summary": "Update Site", "description": "Update a specific Site.", "parameters": [{"name": "site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_site_update_request"}}}}, "responses": {"200": {"description": "Update Site response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_site_modified_response"}}}}, "4XX": {"description": "Update Site response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Sites"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPermissionsRequired": {"enum": ["#organization:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.sites", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
