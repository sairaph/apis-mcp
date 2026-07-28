---
title: Set the Split Tunnel exclude list for a device settings profile
page_id: operation-put-accounts-account-id-devices-policy-policy-id-exclude-493cecf7
path: operations/devices
description: Sets the list of routes excluded from the WARP client's tunnel for a specific device settings profile.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/devices/policy/{policy_id}/exclude
operation_ids:
    - devices-set-split-tunnel-exclude-list-for-a-device-settings-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Set the Split Tunnel exclude list for a device settings profile

`PUT /accounts/{account_id}/devices/policy/{policy_id}/exclude`

Operation ID: `devices-set-split-tunnel-exclude-list-for-a-device-settings-policy`

Sets the list of routes excluded from the WARP client's tunnel for a specific device settings profile.

## Definition

```yaml
{"operationId": "devices-set-split-tunnel-exclude-list-for-a-device-settings-policy", "summary": "Set the Split Tunnel exclude list for a device settings profile", "description": "Sets the list of routes excluded from the WARP client's tunnel for a specific device settings profile.", "parameters": [{"name": "policy_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_schemas-uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_split_tunnel"}}}}}, "responses": {"200": {"description": "Set the Split Tunnel exclude list for a device settings profile response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_split_tunnel_response_collection"}}}}, "4XX": {"description": "Set the Split Tunnel exclude list for a device settings profile response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_split_tunnel_response_collection"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Devices"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.policies.custom.excludes", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
