---
title: Get the Split Tunnel include list
page_id: operation-get-accounts-account-id-devices-policy-include-69f4372c
path: operations/devices
description: Fetches the list of routes included in the WARP client's tunnel.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices/policy/include
operation_ids:
    - devices-get-split-tunnel-include-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get the Split Tunnel include list

`GET /accounts/{account_id}/devices/policy/include`

Operation ID: `devices-get-split-tunnel-include-list`

Fetches the list of routes included in the WARP client's tunnel.

## Definition

```yaml
{"operationId": "devices-get-split-tunnel-include-list", "summary": "Get the Split Tunnel include list", "description": "Fetches the list of routes included in the WARP client's tunnel.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "responses": {"200": {"description": "Get the Split Tunnel include list response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_split_tunnel_include_response_collection"}}}}, "4XX": {"description": "Get the Split Tunnel include list response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_split_tunnel_include_response_collection"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Devices"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.policies.default.includes", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
