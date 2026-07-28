---
title: Update the current settings for the active account
page_id: operation-put-accounts-account-id-cni-settings-56bdb959
path: operations/settings
description: Update the current settings for the active account
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/cni/settings
operation_ids:
    - update_settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update the current settings for the active account

`PUT /accounts/{account_id}/cni/settings`

Operation ID: `update_settings`

## Definition

```yaml
{"operationId": "update_settings", "summary": "Update the current settings for the active account", "parameters": [{"name": "account_id", "in": "path", "description": "Account tag to update settings for", "required": true, "schema": {"type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/nsc_SettingsRequest"}}}}, "responses": {"200": {"description": "The active account settings values", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/nsc_Settings"}}}}, "400": {"description": "Bad request"}, "404": {"description": "Account not found"}, "500": {"description": "Internal server error"}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Settings"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"]}
```
