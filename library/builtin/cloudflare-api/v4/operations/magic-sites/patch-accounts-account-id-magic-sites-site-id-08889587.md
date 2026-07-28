---
title: Patch Site
page_id: operation-patch-accounts-account-id-magic-sites-site-id-6147eab1
path: operations/magic-sites
description: Patch a specific Site.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/magic/sites/{site_id}
operation_ids:
    - magic-sites-patch-site
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch Site

`PATCH /accounts/{account_id}/magic/sites/{site_id}`

Operation ID: `magic-sites-patch-site`

Patch a specific Site.

## Definition

```yaml
{"operationId": "magic-sites-patch-site", "summary": "Patch Site", "description": "Patch a specific Site.", "parameters": [{"name": "site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_site_update_request"}}}}, "responses": {"200": {"description": "Patch Site response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_site_modified_response"}}}}, "4XX": {"description": "Patch Site response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Sites"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPermissionsRequired": {"enum": ["#organization:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.sites", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
