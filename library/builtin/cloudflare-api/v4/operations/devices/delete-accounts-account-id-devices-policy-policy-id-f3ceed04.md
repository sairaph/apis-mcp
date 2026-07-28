---
title: Delete a device settings profile
page_id: operation-delete-accounts-account-id-devices-policy-policy-id-1230dd17
path: operations/devices
description: Deletes a device settings profile and fetches a list of the remaining profiles for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/devices/policy/{policy_id}
operation_ids:
    - devices-delete-device-settings-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a device settings profile

`DELETE /accounts/{account_id}/devices/policy/{policy_id}`

Operation ID: `devices-delete-device-settings-policy`

Deletes a device settings profile and fetches a list of the remaining profiles for an account.

## Definition

```yaml
{"operationId": "devices-delete-device-settings-policy", "summary": "Delete a device settings profile", "description": "Deletes a device settings profile and fetches a list of the remaining profiles for an account.", "parameters": [{"name": "policy_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_schemas-uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete a device settings profile response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_device_settings_response_collection"}}}}, "4XX": {"description": "Delete a device settings profile response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_device_settings_response_collection"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Devices"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.policies.custom", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
