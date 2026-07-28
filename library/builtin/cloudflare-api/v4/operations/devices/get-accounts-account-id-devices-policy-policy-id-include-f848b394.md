---
title: Get the Split Tunnel include list for a device settings profile
page_id: operation-get-accounts-account-id-devices-policy-policy-id-include-8e4117fc
path: operations/devices
description: Fetches the list of routes included in the WARP client's tunnel for a specific device settings profile.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices/policy/{policy_id}/include
operation_ids:
    - devices-get-split-tunnel-include-list-for-a-device-settings-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get the Split Tunnel include list for a device settings profile

`GET /accounts/{account_id}/devices/policy/{policy_id}/include`

Operation ID: `devices-get-split-tunnel-include-list-for-a-device-settings-policy`

Fetches the list of routes included in the WARP client's tunnel for a specific device settings profile.

## Definition

```yaml
{"operationId": "devices-get-split-tunnel-include-list-for-a-device-settings-policy", "summary": "Get the Split Tunnel include list for a device settings profile", "description": "Fetches the list of routes included in the WARP client's tunnel for a specific device settings profile.", "parameters": [{"name": "policy_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_schemas-uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "responses": {"200": {"description": "Get the Split Tunnel include list for a device settings profile response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_split_tunnel_include_response_collection"}}}}, "4XX": {"description": "Get the Split Tunnel include list for a device settings profile response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_split_tunnel_include_response_collection"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Devices"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.policies.custom.includes", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
