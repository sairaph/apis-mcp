---
title: Get the current settings for the active account
page_id: operation-get-accounts-account-id-cni-settings-253e6a1f
path: operations/settings
description: Get the current settings for the active account
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cni/settings
operation_ids:
    - get_settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get the current settings for the active account

`GET /accounts/{account_id}/cni/settings`

Operation ID: `get_settings`

## Definition

```yaml
{"operationId": "get_settings", "summary": "Get the current settings for the active account", "parameters": [{"name": "account_id", "in": "path", "description": "Account tag to retrieve settings for", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "The active account settings values", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/nsc_Settings"}}}}, "400": {"description": "Bad request"}, "404": {"description": "Account not found"}, "500": {"description": "Internal server error"}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Settings"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read", "Magic Transit Read", "Magic Transit Write"]}
```
