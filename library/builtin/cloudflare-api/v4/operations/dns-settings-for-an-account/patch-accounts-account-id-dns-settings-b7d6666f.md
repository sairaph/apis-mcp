---
title: Update DNS Settings
page_id: operation-patch-accounts-account-id-dns-settings-55cc3e2b
path: operations/dns-settings-for-an-account
description: Update DNS settings for an account
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/dns_settings
operation_ids:
    - dns-settings-for-an-account-update-dns-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update DNS Settings

`PATCH /accounts/{account_id}/dns_settings`

Operation ID: `dns-settings-for-an-account-update-dns-settings`

Update DNS settings for an account

## Definition

```yaml
{"operationId": "dns-settings-for-an-account-update-dns-settings", "summary": "Update DNS Settings", "description": "Update DNS settings for an account", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-settings_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-settings_account_settings_patch"}}}}, "responses": {"200": {"description": "Show DNS Settings response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-settings_dns_response_single"}}}}, "4XX": {"description": "Show DNS Settings response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-settings_dns_response_single"}, {"$ref": "#/components/schemas/dns-settings_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNS Settings for an Account"], "x-api-token-group": ["Account DNS Settings Write"], "x-cfPermissionsRequired": {"enum": ["#dns_records:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.settings.account", "x-fern-sdk-method-name": "edit"}
```
