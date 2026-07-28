---
title: List BGP Filter Profiles
page_id: operation-get-accounts-account-id-magic-bgp-filter-profiles-2cc6abf5
path: operations/magic-bgp-filter-profiles
description: Lists all BGP filter profiles for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/bgp/filter_profiles
operation_ids:
    - magic-bgp-list-filter-profiles
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List BGP Filter Profiles

`GET /accounts/{account_id}/magic/bgp/filter_profiles`

Operation ID: `magic-bgp-list-filter-profiles`

Lists all BGP filter profiles for an account.

## Definition

```yaml
{"operationId": "magic-bgp-list-filter-profiles", "summary": "List BGP Filter Profiles", "description": "Lists all BGP filter profiles for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "responses": {"200": {"description": "List BGP Filter Profiles response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_magic_bgp_filter_profiles_list_response"}}}}, "4XX": {"description": "List BGP Filter Profiles response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_magic_bgp_filter_profiles_list_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic BGP Filter Profiles"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.bgp-filter-profiles", "x-fern-sdk-method-name": "list", "x-forge-hidden": true, "x-stability": "beta"}
```
