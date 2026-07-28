---
title: Get the Split Tunnel exclude list
page_id: operation-get-accounts-account-id-devices-policy-exclude-7044ad96
path: operations/devices
description: Fetches the list of routes excluded from the WARP client's tunnel.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices/policy/exclude
operation_ids:
    - devices-get-split-tunnel-exclude-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get the Split Tunnel exclude list

`GET /accounts/{account_id}/devices/policy/exclude`

Operation ID: `devices-get-split-tunnel-exclude-list`

Fetches the list of routes excluded from the WARP client's tunnel.

## Definition

```yaml
{"operationId": "devices-get-split-tunnel-exclude-list", "summary": "Get the Split Tunnel exclude list", "description": "Fetches the list of routes excluded from the WARP client's tunnel.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "responses": {"200": {"description": "Get the Split Tunnel exclude list response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_split_tunnel_response_collection"}}}}, "4XX": {"description": "Get the Split Tunnel exclude list response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_split_tunnel_response_collection"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Devices"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.policies.default.excludes", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
