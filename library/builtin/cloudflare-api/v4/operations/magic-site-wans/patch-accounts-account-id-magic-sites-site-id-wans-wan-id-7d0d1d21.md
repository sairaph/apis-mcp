---
title: Patch Site WAN
page_id: operation-patch-accounts-account-id-magic-sites-site-id-wans-wan-id-3e3885d7
path: operations/magic-site-wans
description: Patch a specific Site WAN.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/magic/sites/{site_id}/wans/{wan_id}
operation_ids:
    - magic-site-wans-patch-wan
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch Site WAN

`PATCH /accounts/{account_id}/magic/sites/{site_id}/wans/{wan_id}`

Operation ID: `magic-site-wans-patch-wan`

Patch a specific Site WAN.

## Definition

```yaml
{"operationId": "magic-site-wans-patch-wan", "summary": "Patch Site WAN", "description": "Patch a specific Site WAN.", "parameters": [{"name": "site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "wan_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_wan_update_request"}}}}, "responses": {"200": {"description": "Patch Site WAN response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_wan_modified_response"}}}}, "4XX": {"description": "Patch Site WAN response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Site WANs"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPermissionsRequired": {"enum": ["#organization:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.sites.wans", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
