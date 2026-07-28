---
title: Delete Site LAN
page_id: operation-delete-accounts-account-id-magic-sites-site-id-lans-lan-id-7ed56879
path: operations/magic-site-lans
description: Remove a specific Site LAN.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/magic/sites/{site_id}/lans/{lan_id}
operation_ids:
    - magic-site-lans-delete-lan
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Site LAN

`DELETE /accounts/{account_id}/magic/sites/{site_id}/lans/{lan_id}`

Operation ID: `magic-site-lans-delete-lan`

Remove a specific Site LAN.

## Definition

```yaml
{"operationId": "magic-site-lans-delete-lan", "summary": "Delete Site LAN", "description": "Remove a specific Site LAN.", "parameters": [{"name": "site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "lan_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Site LAN response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_lan_deleted_response"}}}}, "4XX": {"description": "Delete Site LAN response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Site LANs"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPermissionsRequired": {"enum": ["#organization:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.sites.lans", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
