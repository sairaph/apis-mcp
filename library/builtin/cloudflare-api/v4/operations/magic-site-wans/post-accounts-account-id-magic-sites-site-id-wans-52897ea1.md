---
title: Create a new Site WAN
page_id: operation-post-accounts-account-id-magic-sites-site-id-wans-08f12395
path: operations/magic-site-wans
description: Creates a new Site WAN.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/magic/sites/{site_id}/wans
operation_ids:
    - magic-site-wans-create-wan
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a new Site WAN

`POST /accounts/{account_id}/magic/sites/{site_id}/wans`

Operation ID: `magic-site-wans-create-wan`

Creates a new Site WAN.

## Definition

```yaml
{"operationId": "magic-site-wans-create-wan", "summary": "Create a new Site WAN", "description": "Creates a new Site WAN.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_wans_add_single_request"}}}}, "responses": {"200": {"description": "Create Site WAN response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_wans_collection_response"}}}}, "4XX": {"description": "Create Site WAN response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Site WANs"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.sites.wans", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
