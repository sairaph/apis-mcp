---
title: Delete Site WAN
page_id: operation-delete-accounts-account-id-magic-sites-site-id-wans-wan-id-c84e1882
path: operations/magic-site-wans
description: Remove a specific Site WAN.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/magic/sites/{site_id}/wans/{wan_id}
operation_ids:
    - magic-site-wans-delete-wan
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Site WAN

`DELETE /accounts/{account_id}/magic/sites/{site_id}/wans/{wan_id}`

Operation ID: `magic-site-wans-delete-wan`

Remove a specific Site WAN.

## Definition

```yaml
{"operationId": "magic-site-wans-delete-wan", "summary": "Delete Site WAN", "description": "Remove a specific Site WAN.", "parameters": [{"name": "site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "wan_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Site WAN response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_wan_deleted_response"}}}}, "4XX": {"description": "Delete Site WAN response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Site WANs"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPermissionsRequired": {"enum": ["#organization:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.sites.wans", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
