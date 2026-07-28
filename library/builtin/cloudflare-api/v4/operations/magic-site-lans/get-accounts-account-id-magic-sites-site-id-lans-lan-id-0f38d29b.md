---
title: Site LAN Details
page_id: operation-get-accounts-account-id-magic-sites-site-id-lans-lan-id-57e87a60
path: operations/magic-site-lans
description: Get a specific Site LAN.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/sites/{site_id}/lans/{lan_id}
operation_ids:
    - magic-site-lans-lan-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Site LAN Details

`GET /accounts/{account_id}/magic/sites/{site_id}/lans/{lan_id}`

Operation ID: `magic-site-lans-lan-details`

Get a specific Site LAN.

## Definition

```yaml
{"operationId": "magic-site-lans-lan-details", "summary": "Site LAN Details", "description": "Get a specific Site LAN.", "parameters": [{"name": "site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "lan_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "responses": {"200": {"description": "Site LAN Details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_lan_single_response"}}}}, "4XX": {"description": "Site LAN Details response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Site LANs"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read", "Magic Transit Read", "Magic Transit Write"], "x-cfPermissionsRequired": {"enum": ["#organization:read"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.sites.lans", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
