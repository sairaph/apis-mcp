---
title: Create a new App Config
page_id: operation-post-accounts-account-id-magic-sites-site-id-app-configs-ecb6f54f
path: operations/magic-site-app-configs
description: Creates a new App Config for a site
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/magic/sites/{site_id}/app_configs
operation_ids:
    - magic-site-app-configs-add-app-config
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a new App Config

`POST /accounts/{account_id}/magic/sites/{site_id}/app_configs`

Operation ID: `magic-site-app-configs-add-app-config`

Creates a new App Config for a site

## Definition

```yaml
{"operationId": "magic-site-app-configs-add-app-config", "summary": "Create a new App Config", "description": "Creates a new App Config for a site", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_app_config_add_single_request"}}}}, "responses": {"201": {"description": "Create Site App Config response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_app_config_single_response"}}}}, "4XX": {"description": "Create Site App Config response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Site App Configs"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.sites.app-configuration", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
