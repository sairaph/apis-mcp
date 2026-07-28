---
title: Reset device settings for a Zero Trust account with defaults. This turns off all proxying.
page_id: operation-delete-accounts-account-id-devices-settings-c35cead8
path: operations/zero-trust-accounts
description: Resets the current device settings for a Zero Trust account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/devices/settings
operation_ids:
    - zero-trust-accounts-delete-device-settings-for-zero-trust-account
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Reset device settings for a Zero Trust account with defaults. This turns off all proxying.

`DELETE /accounts/{account_id}/devices/settings`

Operation ID: `zero-trust-accounts-delete-device-settings-for-zero-trust-account`

Resets the current device settings for a Zero Trust account.

## Definition

```yaml
{"operationId": "zero-trust-accounts-delete-device-settings-for-zero-trust-account", "summary": "Reset device settings for a Zero Trust account with defaults. This turns off all proxying.", "description": "Resets the current device settings for a Zero Trust account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "responses": {"200": {"description": "Reset response for device settings for a Zero Trust account.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_zero-trust-account-device-settings-response"}}}}, "4XX": {"description": "Reset failure response device settings for a Zero Trust account.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_zero-trust-account-device-settings-response"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust accounts"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.settings", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
