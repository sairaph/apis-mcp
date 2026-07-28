---
title: Update a device managed network
page_id: operation-put-accounts-account-id-devices-networks-network-id-05df6fe0
path: operations/device-managed-networks
description: Updates a configured device managed network.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/devices/networks/{network_id}
operation_ids:
    - device-managed-networks-update-device-managed-network
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a device managed network

`PUT /accounts/{account_id}/devices/networks/{network_id}`

Operation ID: `device-managed-networks-update-device-managed-network`

Updates a configured device managed network.

## Definition

```yaml
{"operationId": "device-managed-networks-update-device-managed-network", "summary": "Update a device managed network", "description": "Updates a configured device managed network.", "parameters": [{"name": "network_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"config": {"$ref": "#/components/schemas/teams-devices_schemas-config_request"}, "name": {"$ref": "#/components/schemas/teams-devices_device-managed-networks_components-schemas-name"}, "type": {"$ref": "#/components/schemas/teams-devices_components-schemas-type"}}}}}}, "responses": {"200": {"description": "Update a device managed network response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_components-schemas-single_response"}}}}, "4XX": {"description": "Update a device managed network response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_components-schemas-single_response"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Device Managed Networks"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.networks", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
