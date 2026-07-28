---
title: Update BGP Settings
page_id: operation-put-accounts-account-id-magic-bgp-settings-be6bc552
path: operations/magic-bgp-settings
description: Modifies the BGP settings for an account, including the default ASN and redistribution configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/magic/bgp/settings
operation_ids:
    - magic-bgp-update-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update BGP Settings

`PUT /accounts/{account_id}/magic/bgp/settings`

Operation ID: `magic-bgp-update-settings`

Modifies the BGP settings for an account, including the default ASN and redistribution configuration.

## Definition

```yaml
{"operationId": "magic-bgp-update-settings", "summary": "Update BGP Settings", "description": "Modifies the BGP settings for an account, including the default ASN and redistribution configuration.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_update_bgp_settings_request"}}}}, "responses": {"200": {"description": "Update BGP Settings response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_magic_bgp_settings_response"}}}}, "4XX": {"description": "Update BGP Settings response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_magic_bgp_settings_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic BGP Settings"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.bgp-settings", "x-fern-sdk-method-name": "update", "x-forge-hidden": true, "x-stability": "beta"}
```
