---
title: Create a device posture rule
page_id: operation-post-accounts-account-id-devices-posture-a1b25e7a
path: operations/device-posture-rules
description: Creates a new device posture rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/devices/posture
operation_ids:
    - device-posture-rules-create-device-posture-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a device posture rule

`POST /accounts/{account_id}/devices/posture`

Operation ID: `device-posture-rules-create-device-posture-rule`

Creates a new device posture rule.

## Definition

```yaml
{"operationId": "device-posture-rules-create-device-posture-rule", "summary": "Create a device posture rule", "description": "Creates a new device posture rule.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"description": {"$ref": "#/components/schemas/teams-devices_description"}, "expiration": {"$ref": "#/components/schemas/teams-devices_expiration"}, "input": {"$ref": "#/components/schemas/teams-devices_input"}, "match": {"$ref": "#/components/schemas/teams-devices_match"}, "name": {"$ref": "#/components/schemas/teams-devices_name"}, "schedule": {"$ref": "#/components/schemas/teams-devices_schedule"}, "type": {"$ref": "#/components/schemas/teams-devices_type"}}, "required": ["name", "type"]}}}}, "responses": {"200": {"description": "Create device posture rule response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_single_response"}}}}, "4XX": {"description": "Create device posture rule response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_single_response"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Device posture rules"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.posture", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
