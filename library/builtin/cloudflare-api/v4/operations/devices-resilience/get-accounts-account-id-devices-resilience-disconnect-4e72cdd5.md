---
title: Retrieve Global WARP override state
page_id: operation-get-accounts-account-id-devices-resilience-disconnect-310d846d
path: operations/devices-resilience
description: Fetch the Global WARP override state.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices/resilience/disconnect
operation_ids:
    - devices-resilience-retrieve-global-warp-override
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve Global WARP override state

`GET /accounts/{account_id}/devices/resilience/disconnect`

Operation ID: `devices-resilience-retrieve-global-warp-override`

Fetch the Global WARP override state.

## Definition

```yaml
{"operationId": "devices-resilience-retrieve-global-warp-override", "summary": "Retrieve Global WARP override state", "description": "Fetch the Global WARP override state.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "responses": {"200": {"description": "Fetch Global WARP override state response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_global_warp_override_response"}}}}, "4XX": {"description": "Fetch Global WARP override state failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Devices Resilience"], "x-api-token-group": ["Zero Trust Resilience Read", "Zero Trust Resilience Write", "Zero Trust Read", "Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.resilience.global-warp-override", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
