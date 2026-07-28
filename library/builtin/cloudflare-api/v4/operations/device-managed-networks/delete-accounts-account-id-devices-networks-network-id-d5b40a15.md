---
title: Delete a device managed network
page_id: operation-delete-accounts-account-id-devices-networks-network-id-5d17e427
path: operations/device-managed-networks
description: Deletes a device managed network and fetches a list of the remaining device managed networks for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/devices/networks/{network_id}
operation_ids:
    - device-managed-networks-delete-device-managed-network
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a device managed network

`DELETE /accounts/{account_id}/devices/networks/{network_id}`

Operation ID: `device-managed-networks-delete-device-managed-network`

Deletes a device managed network and fetches a list of the remaining device managed networks for an account.

## Definition

```yaml
{"operationId": "device-managed-networks-delete-device-managed-network", "summary": "Delete a device managed network", "description": "Deletes a device managed network and fetches a list of the remaining device managed networks for an account.", "parameters": [{"name": "network_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete a device managed network response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_components-schemas-response_collection"}}}}, "4XX": {"description": "Delete a device managed network response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_components-schemas-response_collection"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Device Managed Networks"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.networks", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
