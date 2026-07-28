---
title: Set the Split Tunnel exclude list
page_id: operation-put-accounts-account-id-devices-policy-exclude-3d48131b
path: operations/devices
description: Sets the list of routes excluded from the WARP client's tunnel.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/devices/policy/exclude
operation_ids:
    - devices-set-split-tunnel-exclude-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Set the Split Tunnel exclude list

`PUT /accounts/{account_id}/devices/policy/exclude`

Operation ID: `devices-set-split-tunnel-exclude-list`

Sets the list of routes excluded from the WARP client's tunnel.

## Definition

```yaml
{"operationId": "devices-set-split-tunnel-exclude-list", "summary": "Set the Split Tunnel exclude list", "description": "Sets the list of routes excluded from the WARP client's tunnel.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_split_tunnel"}}}}}, "responses": {"200": {"description": "Set the Split Tunnel exclude list response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_split_tunnel_response_collection"}}}}, "4XX": {"description": "Set the Split Tunnel exclude list response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_split_tunnel_response_collection"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Devices"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.policies.default.excludes", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
