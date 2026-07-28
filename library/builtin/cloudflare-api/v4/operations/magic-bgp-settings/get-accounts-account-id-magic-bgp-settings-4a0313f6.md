---
title: Get BGP Settings
page_id: operation-get-accounts-account-id-magic-bgp-settings-1b09a710
path: operations/magic-bgp-settings
description: Gets the BGP settings for an account, including the default ASN and redistribution configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/bgp/settings
operation_ids:
    - magic-bgp-get-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get BGP Settings

`GET /accounts/{account_id}/magic/bgp/settings`

Operation ID: `magic-bgp-get-settings`

Gets the BGP settings for an account, including the default ASN and redistribution configuration.

## Definition

```yaml
{"operationId": "magic-bgp-get-settings", "summary": "Get BGP Settings", "description": "Gets the BGP settings for an account, including the default ASN and redistribution configuration.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "responses": {"200": {"description": "Get BGP Settings response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_magic_bgp_settings_response"}}}}, "4XX": {"description": "Get BGP Settings response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_magic_bgp_settings_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic BGP Settings"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.bgp-settings", "x-fern-sdk-method-name": "get", "x-forge-hidden": true, "x-stability": "beta"}
```
