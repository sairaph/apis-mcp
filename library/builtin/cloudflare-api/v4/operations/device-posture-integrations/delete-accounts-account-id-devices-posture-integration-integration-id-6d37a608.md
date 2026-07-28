---
title: Delete a device posture integration
page_id: operation-delete-accounts-account-id-devices-posture-integration-integration-id-897446eb
path: operations/device-posture-integrations
description: Delete a configured device posture integration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/devices/posture/integration/{integration_id}
operation_ids:
    - device-posture-integrations-delete-device-posture-integration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a device posture integration

`DELETE /accounts/{account_id}/devices/posture/integration/{integration_id}`

Operation ID: `device-posture-integrations-delete-device-posture-integration`

Delete a configured device posture integration.

## Definition

```yaml
{"operationId": "device-posture-integrations-delete-device-posture-integration", "summary": "Delete a device posture integration", "description": "Delete a configured device posture integration.", "parameters": [{"name": "integration_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete a device posture integration response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_schemas-id_response"}}}}, "4XX": {"description": "Delete a device posture integration response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_schemas-id_response"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Device Posture Integrations"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.posture.integrations", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
