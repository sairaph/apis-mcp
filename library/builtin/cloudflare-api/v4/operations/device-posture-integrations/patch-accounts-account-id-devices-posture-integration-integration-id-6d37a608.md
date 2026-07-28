---
title: Update a device posture integration
page_id: operation-patch-accounts-account-id-devices-posture-integration-integration-id-90bc15cc
path: operations/device-posture-integrations
description: Updates a configured device posture integration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/devices/posture/integration/{integration_id}
operation_ids:
    - device-posture-integrations-update-device-posture-integration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a device posture integration

`PATCH /accounts/{account_id}/devices/posture/integration/{integration_id}`

Operation ID: `device-posture-integrations-update-device-posture-integration`

Updates a configured device posture integration.

## Definition

```yaml
{"operationId": "device-posture-integrations-update-device-posture-integration", "summary": "Update a device posture integration", "description": "Updates a configured device posture integration.", "parameters": [{"name": "integration_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"config": {"$ref": "#/components/schemas/teams-devices_config_request"}, "interval": {"$ref": "#/components/schemas/teams-devices_interval"}, "name": {"$ref": "#/components/schemas/teams-devices_components-schemas-name"}, "type": {"$ref": "#/components/schemas/teams-devices_schemas-type"}}}}}}, "responses": {"200": {"description": "Update a device posture integration response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_schemas-single_response"}}}}, "4XX": {"description": "Update a device posture integration response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_schemas-single_response"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Device Posture Integrations"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.posture.integrations", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
