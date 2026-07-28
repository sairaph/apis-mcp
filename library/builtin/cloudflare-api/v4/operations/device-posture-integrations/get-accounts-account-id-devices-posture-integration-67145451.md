---
title: List your device posture integrations
page_id: operation-get-accounts-account-id-devices-posture-integration-96790543
path: operations/device-posture-integrations
description: Fetches the list of device posture integrations for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices/posture/integration
operation_ids:
    - device-posture-integrations-list-device-posture-integrations
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List your device posture integrations

`GET /accounts/{account_id}/devices/posture/integration`

Operation ID: `device-posture-integrations-list-device-posture-integrations`

Fetches the list of device posture integrations for an account.

## Definition

```yaml
{"operationId": "device-posture-integrations-list-device-posture-integrations", "summary": "List your device posture integrations", "description": "Fetches the list of device posture integrations for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "responses": {"200": {"description": "List your device posture integrations response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_schemas-response_collection"}}}}, "4XX": {"description": "List your device posture integrations response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_schemas-response_collection"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Device Posture Integrations"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.posture.integrations", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
