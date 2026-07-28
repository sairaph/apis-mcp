---
title: Update BGP Filter Profile
page_id: operation-put-accounts-account-id-magic-bgp-filter-profiles-profile-id-1a484eea
path: operations/magic-bgp-filter-profiles
description: 'Updates a BGP filter profile. Omitted properties are left unchanged. To clear an existing description send `description: ""`.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/magic/bgp/filter_profiles/{profile_id}
operation_ids:
    - magic-bgp-update-filter-profile
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update BGP Filter Profile

`PUT /accounts/{account_id}/magic/bgp/filter_profiles/{profile_id}`

Operation ID: `magic-bgp-update-filter-profile`

Updates a BGP filter profile. Omitted properties are left unchanged. To clear an existing description send `description: ""`.

## Definition

```yaml
{"operationId": "magic-bgp-update-filter-profile", "summary": "Update BGP Filter Profile", "description": "Updates a BGP filter profile. Omitted properties are left unchanged. To clear an existing description send `description: \"\"`.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "profile_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_update_bgp_filter_profile_request"}}}}, "responses": {"200": {"description": "Update BGP Filter Profile response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_magic_bgp_filter_profile_response"}}}}, "4XX": {"description": "Update BGP Filter Profile response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_magic_bgp_filter_profile_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic BGP Filter Profiles"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.bgp-filter-profiles", "x-fern-sdk-method-name": "update", "x-forge-hidden": true, "x-stability": "beta"}
```
