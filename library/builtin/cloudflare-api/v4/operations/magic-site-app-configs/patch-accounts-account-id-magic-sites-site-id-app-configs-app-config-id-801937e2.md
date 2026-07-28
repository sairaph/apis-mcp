---
title: Update an App Config
page_id: operation-patch-accounts-account-id-magic-sites-site-id-app-configs-app-config-id-b3be060a
path: operations/magic-site-app-configs
description: Updates an App Config for a site
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/magic/sites/{site_id}/app_configs/{app_config_id}
operation_ids:
    - magic-site-app-configs-patch-app-config
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update an App Config

`PATCH /accounts/{account_id}/magic/sites/{site_id}/app_configs/{app_config_id}`

Operation ID: `magic-site-app-configs-patch-app-config`

Updates an App Config for a site

## Definition

```yaml
{"operationId": "magic-site-app-configs-patch-app-config", "summary": "Update an App Config", "description": "Updates an App Config for a site", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "app_config_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_app_config_update_request"}}}}, "responses": {"200": {"description": "Update Site App Config response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_app_config_single_response"}}}}, "4XX": {"description": "Update Site App Config response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Site App Configs"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.sites.app-configuration", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
