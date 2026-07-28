---
title: Get override codes (deprecated)
page_id: operation-get-accounts-account-id-devices-device-id-override-codes-dafabfed
path: operations/devices
description: |-
    Fetches a one-time use admin override code for a device. This relies on the **Admin Override** setting being enabled in your device configuration. Not supported when [multi-user mode](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/warp/deployment/mdm-deployment/windows-multiuser/) is enabled for the account.
    **Deprecated:** please use GET /accounts/{account_id}/devices/registrations/{registration_id}/override_codes instead.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices/{device_id}/override_codes
operation_ids:
    - devices-list-admin-override-code-for-device
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get override codes (deprecated)

`GET /accounts/{account_id}/devices/{device_id}/override_codes`

Operation ID: `devices-list-admin-override-code-for-device`

Fetches a one-time use admin override code for a device. This relies on the **Admin Override** setting being enabled in your device configuration. Not supported when [multi-user mode](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/warp/deployment/mdm-deployment/windows-multiuser/) is enabled for the account.
**Deprecated:** please use GET /accounts/{account_id}/devices/registrations/{registration_id}/override_codes instead.

## Definition

```yaml
{"operationId": "devices-list-admin-override-code-for-device", "summary": "Get override codes (deprecated)\n", "description": "Fetches a one-time use admin override code for a device. This relies on the **Admin Override** setting being enabled in your device configuration. Not supported when [multi-user mode](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/warp/deployment/mdm-deployment/windows-multiuser/) is enabled for the account.\n**Deprecated:** please use GET /accounts/{account_id}/devices/registrations/{registration_id}/override_codes instead.\n", "parameters": [{"name": "device_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_registration_id"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "responses": {"200": {"description": "Get an admin override code for a device response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_override_codes_response"}}}}, "4XX": {"description": "Get an admin override code for a device response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_override_codes_response"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Devices"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.override-codes", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
