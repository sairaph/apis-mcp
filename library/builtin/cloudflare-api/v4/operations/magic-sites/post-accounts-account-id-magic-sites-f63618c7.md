---
title: Create a new Site
page_id: operation-post-accounts-account-id-magic-sites-1e059ad3
path: operations/magic-sites
description: Creates a new Site
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/magic/sites
operation_ids:
    - magic-sites-create-site
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a new Site

`POST /accounts/{account_id}/magic/sites`

Operation ID: `magic-sites-create-site`

Creates a new Site

## Definition

```yaml
{"operationId": "magic-sites-create-site", "summary": "Create a new Site", "description": "Creates a new Site", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_sites_add_single_request"}}}}, "responses": {"200": {"description": "Create Site response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_site_single_response"}}}}, "4XX": {"description": "Create Site response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Sites"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.sites", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
