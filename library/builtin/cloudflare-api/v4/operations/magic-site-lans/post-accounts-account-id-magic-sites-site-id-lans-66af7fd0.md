---
title: Create a new Site LAN
page_id: operation-post-accounts-account-id-magic-sites-site-id-lans-f1ca4a35
path: operations/magic-site-lans
description: Creates a new Site LAN. If the site is in high availability mode, static_addressing is required along with secondary and virtual address.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/magic/sites/{site_id}/lans
operation_ids:
    - magic-site-lans-create-lan
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a new Site LAN

`POST /accounts/{account_id}/magic/sites/{site_id}/lans`

Operation ID: `magic-site-lans-create-lan`

Creates a new Site LAN. If the site is in high availability mode, static_addressing is required along with secondary and virtual address.

## Definition

```yaml
{"operationId": "magic-site-lans-create-lan", "summary": "Create a new Site LAN", "description": "Creates a new Site LAN. If the site is in high availability mode, static_addressing is required along with secondary and virtual address.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_lans_add_single_request"}}}}, "responses": {"200": {"description": "Create Site LAN response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_lans_collection_response"}}}}, "4XX": {"description": "Create Site LAN response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Site LANs"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.sites.lans", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
