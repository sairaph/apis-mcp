---
title: Create CF1 Sites
page_id: operation-post-accounts-account-id-magic-cf1-sites-5596879c
path: operations/magic-cf1-sites
description: Creates new CF1 Sites for an account. Each site must have a unique name within the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/magic/cf1_sites
operation_ids:
    - magic-cf1-sites-create-cf1-sites
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create CF1 Sites

`POST /accounts/{account_id}/magic/cf1_sites`

Operation ID: `magic-cf1-sites-create-cf1-sites`

Creates new CF1 Sites for an account. Each site must have a unique name within the account.

## Definition

```yaml
{"operationId": "magic-cf1-sites-create-cf1-sites", "summary": "Create CF1 Sites", "description": "Creates new CF1 Sites for an account. Each site must have a unique name within the account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "array", "items": {"$ref": "#/components/schemas/magic_cf1_site"}, "minItems": 1}}}}, "responses": {"200": {"description": "Create CF1 Sites response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_cf1_sites_collection_response"}}}}, "4XX": {"description": "Create CF1 Sites response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic CF1 Sites"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPermissionsRequired": {"enum": ["#organization:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
