---
title: Set the Split Tunnel include list
page_id: operation-put-accounts-account-id-devices-policy-include-8ebab722
path: operations/devices
description: Sets the list of routes included in the WARP client's tunnel.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/devices/policy/include
operation_ids:
    - devices-set-split-tunnel-include-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Set the Split Tunnel include list

`PUT /accounts/{account_id}/devices/policy/include`

Operation ID: `devices-set-split-tunnel-include-list`

Sets the list of routes included in the WARP client's tunnel.

## Definition

```yaml
{"operationId": "devices-set-split-tunnel-include-list", "summary": "Set the Split Tunnel include list", "description": "Sets the list of routes included in the WARP client's tunnel.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_split_tunnel_include"}}}}}, "responses": {"200": {"description": "Set the Split Tunnel include list response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_split_tunnel_include_response_collection"}}}}, "4XX": {"description": "Set the Split Tunnel include list response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_split_tunnel_include_response_collection"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Devices"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.policies.default.includes", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
