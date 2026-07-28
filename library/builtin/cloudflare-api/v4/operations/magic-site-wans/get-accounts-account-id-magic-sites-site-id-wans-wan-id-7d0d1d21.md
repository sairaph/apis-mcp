---
title: Site WAN Details
page_id: operation-get-accounts-account-id-magic-sites-site-id-wans-wan-id-5bbe43db
path: operations/magic-site-wans
description: Get a specific Site WAN.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/sites/{site_id}/wans/{wan_id}
operation_ids:
    - magic-site-wans-wan-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Site WAN Details

`GET /accounts/{account_id}/magic/sites/{site_id}/wans/{wan_id}`

Operation ID: `magic-site-wans-wan-details`

Get a specific Site WAN.

## Definition

```yaml
{"operationId": "magic-site-wans-wan-details", "summary": "Site WAN Details", "description": "Get a specific Site WAN.", "parameters": [{"name": "site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "wan_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "responses": {"200": {"description": "Site WAN Details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_wan_single_response"}}}}, "4XX": {"description": "Site WAN Details response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Site WANs"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read", "Magic Transit Read", "Magic Transit Write"], "x-cfPermissionsRequired": {"enum": ["#organization:read"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.sites.wans", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
