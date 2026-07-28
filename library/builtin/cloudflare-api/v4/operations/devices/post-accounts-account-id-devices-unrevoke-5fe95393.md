---
title: Unrevoke devices (deprecated)
page_id: operation-post-accounts-account-id-devices-unrevoke-555decb5
path: operations/devices
description: |-
    Unrevokes a list of devices. Not supported when [multi-user mode](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/warp/deployment/mdm-deployment/windows-multiuser/) is enabled.

    **Deprecated**: please use POST /accounts/{account_id}/devices/registrations/unrevoke instead.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/devices/unrevoke
operation_ids:
    - devices-unrevoke-devices
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Unrevoke devices (deprecated)

`POST /accounts/{account_id}/devices/unrevoke`

Operation ID: `devices-unrevoke-devices`

Unrevokes a list of devices. Not supported when [multi-user mode](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/warp/deployment/mdm-deployment/windows-multiuser/) is enabled.

**Deprecated**: please use POST /accounts/{account_id}/devices/registrations/unrevoke instead.

## Definition

```yaml
{"operationId": "devices-unrevoke-devices", "summary": "Unrevoke devices (deprecated)", "description": "Unrevokes a list of devices. Not supported when [multi-user mode](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/warp/deployment/mdm-deployment/windows-multiuser/) is enabled.\n\n**Deprecated**: please use POST /accounts/{account_id}/devices/registrations/unrevoke instead.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_unrevoke_devices_request"}}}}, "responses": {"200": {"description": "Unrevoke devices response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_api-response-single"}}}}, "4XX": {"description": "Unrevoke devices response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_api-response-single"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Devices"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.unrevoke", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
