---
title: List your device managed networks
page_id: operation-get-accounts-account-id-devices-networks-03d9a850
path: operations/device-managed-networks
description: Fetches a list of managed networks for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices/networks
operation_ids:
    - device-managed-networks-list-device-managed-networks
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List your device managed networks

`GET /accounts/{account_id}/devices/networks`

Operation ID: `device-managed-networks-list-device-managed-networks`

Fetches a list of managed networks for an account.

## Definition

```yaml
{"operationId": "device-managed-networks-list-device-managed-networks", "summary": "List your device managed networks", "description": "Fetches a list of managed networks for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "responses": {"200": {"description": "List your device managed networks response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_components-schemas-response_collection"}}}}, "4XX": {"description": "List your device managed networks response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_components-schemas-response_collection"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Device Managed Networks"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.networks", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
