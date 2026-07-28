---
title: List Site LANs
page_id: operation-get-accounts-account-id-magic-sites-site-id-lans-0503ea14
path: operations/magic-site-lans
description: Lists Site LANs associated with an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/sites/{site_id}/lans
operation_ids:
    - magic-site-lans-list-lans
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Site LANs

`GET /accounts/{account_id}/magic/sites/{site_id}/lans`

Operation ID: `magic-site-lans-list-lans`

Lists Site LANs associated with an account.

## Definition

```yaml
{"operationId": "magic-site-lans-list-lans", "summary": "List Site LANs", "description": "Lists Site LANs associated with an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "responses": {"200": {"description": "List Site LANs response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_lans_collection_response"}}}}, "4XX": {"description": "List Site LANs response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Site LANs"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read", "Magic Transit Read", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.sites.lans", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
