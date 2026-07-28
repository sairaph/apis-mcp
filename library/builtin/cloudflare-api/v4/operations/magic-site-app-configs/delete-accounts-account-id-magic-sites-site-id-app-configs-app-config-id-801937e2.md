---
title: Delete App Config
page_id: operation-delete-accounts-account-id-magic-sites-site-id-app-configs-app-config-id-983f9055
path: operations/magic-site-app-configs
description: Deletes specific App Config associated with a site.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/magic/sites/{site_id}/app_configs/{app_config_id}
operation_ids:
    - magic-site-app-configs-delete-app-config
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete App Config

`DELETE /accounts/{account_id}/magic/sites/{site_id}/app_configs/{app_config_id}`

Operation ID: `magic-site-app-configs-delete-app-config`

Deletes specific App Config associated with a site.

## Definition

```yaml
{"operationId": "magic-site-app-configs-delete-app-config", "summary": "Delete App Config", "description": "Deletes specific App Config associated with a site.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "app_config_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "responses": {"200": {"description": "Delete App Config response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_app_config_single_response"}}}}, "4XX": {"description": "Delete App Config response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Site App Configs"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.sites.app-configuration", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
