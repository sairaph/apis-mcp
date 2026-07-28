---
title: List devices (deprecated)
page_id: operation-get-accounts-account-id-devices-85278695
path: operations/devices
description: |-
    List WARP devices. Not supported when [multi-user mode](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/warp/deployment/mdm-deployment/windows-multiuser/) is enabled for the account.

    **Deprecated**: please use one of the following endpoints instead:
    - GET /accounts/{account_id}/devices/physical-devices
    - GET /accounts/{account_id}/devices/registrations
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices
operation_ids:
    - devices-list-devices
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List devices (deprecated)

`GET /accounts/{account_id}/devices`

Operation ID: `devices-list-devices`

List WARP devices. Not supported when [multi-user mode](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/warp/deployment/mdm-deployment/windows-multiuser/) is enabled for the account.

**Deprecated**: please use one of the following endpoints instead:
- GET /accounts/{account_id}/devices/physical-devices
- GET /accounts/{account_id}/devices/registrations

## Definition

```yaml
{"operationId": "devices-list-devices", "summary": "List devices (deprecated)", "description": "List WARP devices. Not supported when [multi-user mode](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/warp/deployment/mdm-deployment/windows-multiuser/) is enabled for the account.\n\n**Deprecated**: please use one of the following endpoints instead:\n- GET /accounts/{account_id}/devices/physical-devices\n- GET /accounts/{account_id}/devices/registrations\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "responses": {"200": {"description": "List devices response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_devices_response"}}}}, "4XX": {"description": "List devices response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_devices_response"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Devices"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
