---
title: Create a device managed network
page_id: operation-post-accounts-account-id-devices-networks-2b444499
path: operations/device-managed-networks
description: Creates a new device managed network.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/devices/networks
operation_ids:
    - device-managed-networks-create-device-managed-network
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a device managed network

`POST /accounts/{account_id}/devices/networks`

Operation ID: `device-managed-networks-create-device-managed-network`

Creates a new device managed network.

## Definition

```yaml
{"operationId": "device-managed-networks-create-device-managed-network", "summary": "Create a device managed network", "description": "Creates a new device managed network.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"config": {"$ref": "#/components/schemas/teams-devices_schemas-config_request"}, "name": {"$ref": "#/components/schemas/teams-devices_device-managed-networks_components-schemas-name"}, "type": {"$ref": "#/components/schemas/teams-devices_components-schemas-type"}}, "required": ["name", "type", "config"]}}}}, "responses": {"200": {"description": "Create a device managed networks response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_components-schemas-single_response"}}}}, "4XX": {"description": "Create a device managed networks response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_components-schemas-single_response"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Device Managed Networks"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.networks", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
