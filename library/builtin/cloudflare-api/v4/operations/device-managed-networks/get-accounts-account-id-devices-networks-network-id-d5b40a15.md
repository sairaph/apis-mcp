---
title: Get device managed network details
page_id: operation-get-accounts-account-id-devices-networks-network-id-805e420e
path: operations/device-managed-networks
description: Fetches details for a single managed network.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices/networks/{network_id}
operation_ids:
    - device-managed-networks-device-managed-network-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get device managed network details

`GET /accounts/{account_id}/devices/networks/{network_id}`

Operation ID: `device-managed-networks-device-managed-network-details`

Fetches details for a single managed network.

## Definition

```yaml
{"operationId": "device-managed-networks-device-managed-network-details", "summary": "Get device managed network details", "description": "Fetches details for a single managed network.", "parameters": [{"name": "network_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "responses": {"200": {"description": "Get device managed network details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_components-schemas-single_response"}}}}, "4XX": {"description": "Get device managed network details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_components-schemas-single_response"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Device Managed Networks"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.networks", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
