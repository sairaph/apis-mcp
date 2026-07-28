---
title: Create BGP Filter Profile
page_id: operation-post-accounts-account-id-magic-bgp-filter-profiles-14083c85
path: operations/magic-bgp-filter-profiles
description: Creates a new BGP filter profile for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/magic/bgp/filter_profiles
operation_ids:
    - magic-bgp-create-filter-profile
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create BGP Filter Profile

`POST /accounts/{account_id}/magic/bgp/filter_profiles`

Operation ID: `magic-bgp-create-filter-profile`

Creates a new BGP filter profile for an account.

## Definition

```yaml
{"operationId": "magic-bgp-create-filter-profile", "summary": "Create BGP Filter Profile", "description": "Creates a new BGP filter profile for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_create_bgp_filter_profile_request"}}}}, "responses": {"200": {"description": "Create BGP Filter Profile response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_magic_bgp_filter_profile_response"}}}}, "4XX": {"description": "Create BGP Filter Profile response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_magic_bgp_filter_profile_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic BGP Filter Profiles"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.bgp-filter-profiles", "x-fern-sdk-method-name": "create", "x-forge-hidden": true, "x-stability": "beta"}
```
