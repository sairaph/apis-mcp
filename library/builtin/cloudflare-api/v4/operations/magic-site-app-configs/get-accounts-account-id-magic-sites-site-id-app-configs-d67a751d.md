---
title: List App Configs
page_id: operation-get-accounts-account-id-magic-sites-site-id-app-configs-e008fb6d
path: operations/magic-site-app-configs
description: Lists App Configs associated with a site.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/sites/{site_id}/app_configs
operation_ids:
    - magic-site-app-configs-list-app-configs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List App Configs

`GET /accounts/{account_id}/magic/sites/{site_id}/app_configs`

Operation ID: `magic-site-app-configs-list-app-configs`

Lists App Configs associated with a site.

## Definition

```yaml
{"operationId": "magic-site-app-configs-list-app-configs", "summary": "List App Configs", "description": "Lists App Configs associated with a site.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "responses": {"200": {"description": "List App Configs response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_app_configs_collection_response"}}}}, "4XX": {"description": "List App Configs response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Site App Configs"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read", "Magic Transit Read", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.sites.app-configuration", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
