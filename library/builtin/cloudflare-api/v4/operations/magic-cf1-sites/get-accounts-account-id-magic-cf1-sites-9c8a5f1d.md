---
title: List CF1 Sites
page_id: operation-get-accounts-account-id-magic-cf1-sites-df035d0b
path: operations/magic-cf1-sites
description: Lists CF1 Sites associated with an account. A CF1 Site represents a physical customer network location with optional geographic coordinates.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/cf1_sites
operation_ids:
    - magic-cf1-sites-list-cf1-sites
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List CF1 Sites

`GET /accounts/{account_id}/magic/cf1_sites`

Operation ID: `magic-cf1-sites-list-cf1-sites`

Lists CF1 Sites associated with an account. A CF1 Site represents a physical customer network location with optional geographic coordinates.

## Definition

```yaml
{"operationId": "magic-cf1-sites-list-cf1-sites", "summary": "List CF1 Sites", "description": "Lists CF1 Sites associated with an account. A CF1 Site represents a physical customer network location with optional geographic coordinates.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "responses": {"200": {"description": "List CF1 Sites response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_cf1_sites_collection_response"}}}}, "4XX": {"description": "List CF1 Sites response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic CF1 Sites"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read", "Magic Transit Read", "Magic Transit Write"], "x-cfPermissionsRequired": {"enum": ["#organization:read"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
