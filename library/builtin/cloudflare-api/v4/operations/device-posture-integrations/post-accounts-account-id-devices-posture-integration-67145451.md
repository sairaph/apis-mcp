---
title: Create a device posture integration
page_id: operation-post-accounts-account-id-devices-posture-integration-9bc8f361
path: operations/device-posture-integrations
description: Create a new device posture integration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/devices/posture/integration
operation_ids:
    - device-posture-integrations-create-device-posture-integration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a device posture integration

`POST /accounts/{account_id}/devices/posture/integration`

Operation ID: `device-posture-integrations-create-device-posture-integration`

Create a new device posture integration.

## Definition

```yaml
{"operationId": "device-posture-integrations-create-device-posture-integration", "summary": "Create a device posture integration", "description": "Create a new device posture integration.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"config": {"$ref": "#/components/schemas/teams-devices_config_request"}, "interval": {"$ref": "#/components/schemas/teams-devices_interval"}, "name": {"$ref": "#/components/schemas/teams-devices_components-schemas-name"}, "type": {"$ref": "#/components/schemas/teams-devices_schemas-type"}}, "required": ["name", "type", "interval", "config"]}}}}, "responses": {"200": {"description": "Create a device posture integration response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_schemas-single_response"}}}}, "4XX": {"description": "Create a device posture integration response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_schemas-single_response"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Device Posture Integrations"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.posture.integrations", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
