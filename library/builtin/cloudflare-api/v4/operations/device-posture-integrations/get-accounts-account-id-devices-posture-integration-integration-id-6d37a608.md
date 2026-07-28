---
title: Get device posture integration details
page_id: operation-get-accounts-account-id-devices-posture-integration-integration-id-c195a311
path: operations/device-posture-integrations
description: Fetches details for a single device posture integration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices/posture/integration/{integration_id}
operation_ids:
    - device-posture-integrations-device-posture-integration-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get device posture integration details

`GET /accounts/{account_id}/devices/posture/integration/{integration_id}`

Operation ID: `device-posture-integrations-device-posture-integration-details`

Fetches details for a single device posture integration.

## Definition

```yaml
{"operationId": "device-posture-integrations-device-posture-integration-details", "summary": "Get device posture integration details", "description": "Fetches details for a single device posture integration.", "parameters": [{"name": "integration_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "responses": {"200": {"description": "Get device posture integration details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_schemas-single_response"}}}}, "4XX": {"description": "Get device posture integration details response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/teams-devices_schemas-single_response"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Device Posture Integrations"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.posture.integrations", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
