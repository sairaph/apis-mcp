---
title: Delete BGP Filter Profile
page_id: operation-delete-accounts-account-id-magic-bgp-filter-profiles-profile-id-cd34a3ab
path: operations/magic-bgp-filter-profiles
description: Deletes a BGP filter profile.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/magic/bgp/filter_profiles/{profile_id}
operation_ids:
    - magic-bgp-delete-filter-profile
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete BGP Filter Profile

`DELETE /accounts/{account_id}/magic/bgp/filter_profiles/{profile_id}`

Operation ID: `magic-bgp-delete-filter-profile`

Deletes a BGP filter profile.

## Definition

```yaml
{"operationId": "magic-bgp-delete-filter-profile", "summary": "Delete BGP Filter Profile", "description": "Deletes a BGP filter profile.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "profile_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "responses": {"200": {"description": "Delete BGP Filter Profile response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_magic_bgp_filter_profile_response"}}}}, "4XX": {"description": "Delete BGP Filter Profile response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_magic_bgp_filter_profile_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic BGP Filter Profiles"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.bgp-filter-profiles", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true, "x-stability": "beta"}
```
