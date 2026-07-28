---
title: Revoke devices (deprecated)
page_id: operation-post-accounts-account-id-devices-revoke-8cef99ca
path: operations/devices
description: |-
    Revokes a list of devices. Not supported when [multi-user mode](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/warp/deployment/mdm-deployment/windows-multiuser/) is enabled.

    **Deprecated**: please use POST /accounts/{account_id}/devices/registrations/revoke instead.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/devices/revoke
operation_ids:
    - devices-revoke-devices
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Revoke devices (deprecated)

`POST /accounts/{account_id}/devices/revoke`

Operation ID: `devices-revoke-devices`

Revokes a list of devices. Not supported when [multi-user mode](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/warp/deployment/mdm-deployment/windows-multiuser/) is enabled.

**Deprecated**: please use POST /accounts/{account_id}/devices/registrations/revoke instead.

## Definition

```yaml
{"operationId": "devices-revoke-devices", "summary": "Revoke devices (deprecated)", "description": "Revokes a list of devices. Not supported when [multi-user mode](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/warp/deployment/mdm-deployment/windows-multiuser/) is enabled.\n\n**Deprecated**: please use POST /accounts/{account_id}/devices/registrations/revoke instead.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_revoke_devices_request"}}}}, "responses": {"200": {"description": "Revoke devices response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_api-response-single"}}}}, "4XX": {"description": "Revoke devices response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_api-response-single"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Devices"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.revoke", "x-fern-sdk-method-name": "create", "x-forge-hidden": true, "x-forge-require-confirmation": "This operation revokes sessions for a batch of devices destructively."}
```
