---
title: Show DNS Settings
page_id: operation-get-accounts-account-id-dns-settings-f2869e73
path: operations/dns-settings-for-an-account
description: Show DNS settings for an account
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dns_settings
operation_ids:
    - dns-settings-for-an-account-list-dns-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Show DNS Settings

`GET /accounts/{account_id}/dns_settings`

Operation ID: `dns-settings-for-an-account-list-dns-settings`

Show DNS settings for an account

## Definition

```yaml
{"operationId": "dns-settings-for-an-account-list-dns-settings", "summary": "Show DNS Settings", "description": "Show DNS settings for an account", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-settings_identifier"}}], "responses": {"200": {"description": "Show DNS Settings response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-settings_dns_response_single"}}}}, "4XX": {"description": "Show DNS Settings response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-settings_dns_response_single"}, {"$ref": "#/components/schemas/dns-settings_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNS Settings for an Account"], "x-api-token-group": ["Account DNS Settings Write", "Account DNS Settings Read"], "x-cfPermissionsRequired": {"enum": ["#dns_records:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.settings.account", "x-fern-sdk-method-name": "get"}
```
