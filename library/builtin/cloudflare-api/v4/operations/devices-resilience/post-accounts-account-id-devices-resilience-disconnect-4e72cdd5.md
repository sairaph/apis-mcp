---
title: Set Global WARP override state
page_id: operation-post-accounts-account-id-devices-resilience-disconnect-edd27e10
path: operations/devices-resilience
description: Sets the Global WARP override state.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/devices/resilience/disconnect
operation_ids:
    - devices-resilience-set-global-warp-override
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Set Global WARP override state

`POST /accounts/{account_id}/devices/resilience/disconnect`

Operation ID: `devices-resilience-set-global-warp-override`

Sets the Global WARP override state.

## Definition

```yaml
{"operationId": "devices-resilience-set-global-warp-override", "summary": "Set Global WARP override state", "description": "Sets the Global WARP override state.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_global_warp_override_request"}}}}, "responses": {"200": {"description": "Set Global WARP override state response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_global_warp_override_response"}}}}, "4XX": {"description": "Set Global WARP override state response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Devices Resilience"], "x-api-token-group": ["Zero Trust Resilience Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.resilience.global-warp-override", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
