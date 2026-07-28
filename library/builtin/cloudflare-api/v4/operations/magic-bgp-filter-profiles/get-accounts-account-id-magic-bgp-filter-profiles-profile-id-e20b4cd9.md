---
title: Get BGP Filter Profile
page_id: operation-get-accounts-account-id-magic-bgp-filter-profiles-profile-id-5279ab87
path: operations/magic-bgp-filter-profiles
description: Gets a specific BGP filter profile for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/bgp/filter_profiles/{profile_id}
operation_ids:
    - magic-bgp-get-filter-profile
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get BGP Filter Profile

`GET /accounts/{account_id}/magic/bgp/filter_profiles/{profile_id}`

Operation ID: `magic-bgp-get-filter-profile`

Gets a specific BGP filter profile for an account.

## Definition

```yaml
{"operationId": "magic-bgp-get-filter-profile", "summary": "Get BGP Filter Profile", "description": "Gets a specific BGP filter profile for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "profile_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "responses": {"200": {"description": "Get BGP Filter Profile response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_magic_bgp_filter_profile_response"}}}}, "4XX": {"description": "Get BGP Filter Profile response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_magic_bgp_filter_profile_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic BGP Filter Profiles"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.bgp-filter-profiles", "x-fern-sdk-method-name": "get", "x-forge-hidden": true, "x-stability": "beta"}
```
