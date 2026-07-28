---
title: Get device (deprecated)
page_id: operation-get-accounts-account-id-devices-device-id-f5e17c6f
path: operations/devices
description: |-
    Fetches a single WARP device. Not supported when [multi-user mode](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/warp/deployment/mdm-deployment/windows-multiuser/) is enabled for the account.

    **Deprecated**: please use one of the following endpoints instead:
    - GET /accounts/{account_id}/devices/physical-devices/{device_id}
    - GET /accounts/{account_id}/devices/registrations/{registration_id}
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices/{device_id}
operation_ids:
    - devices-device-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get device (deprecated)

`GET /accounts/{account_id}/devices/{device_id}`

Operation ID: `devices-device-details`

Fetches a single WARP device. Not supported when [multi-user mode](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/warp/deployment/mdm-deployment/windows-multiuser/) is enabled for the account.

**Deprecated**: please use one of the following endpoints instead:
- GET /accounts/{account_id}/devices/physical-devices/{device_id}
- GET /accounts/{account_id}/devices/registrations/{registration_id}

## Definition

```yaml
{"operationId": "devices-device-details", "summary": "Get device (deprecated)", "description": "Fetches a single WARP device. Not supported when [multi-user mode](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/warp/deployment/mdm-deployment/windows-multiuser/) is enabled for the account.\n\n**Deprecated**: please use one of the following endpoints instead:\n- GET /accounts/{account_id}/devices/physical-devices/{device_id}\n- GET /accounts/{account_id}/devices/registrations/{registration_id}\n", "parameters": [{"name": "device_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_registration_id"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "responses": {"200": {"description": "Get device details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_device_response"}}}}, "4XX": {"description": "Get device details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_device_response"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Devices"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
