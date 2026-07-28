---
title: Get device settings for a Zero Trust account
page_id: operation-get-accounts-account-id-devices-settings-e4fc4712
path: operations/zero-trust-accounts
description: Describes the current device settings for a Zero Trust account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices/settings
operation_ids:
    - zero-trust-accounts-get-device-settings-for-zero-trust-account
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get device settings for a Zero Trust account

`GET /accounts/{account_id}/devices/settings`

Operation ID: `zero-trust-accounts-get-device-settings-for-zero-trust-account`

Describes the current device settings for a Zero Trust account.

## Definition

```yaml
{"operationId": "zero-trust-accounts-get-device-settings-for-zero-trust-account", "summary": "Get device settings for a Zero Trust account", "description": "Describes the current device settings for a Zero Trust account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "responses": {"200": {"description": "Get device settings for a Zero Trust account response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_zero-trust-account-device-settings-response"}}}}, "4XX": {"description": "Get device settings for a Zero Trust account response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_zero-trust-account-device-settings-response"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust accounts"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.settings", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
