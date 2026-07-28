---
title: Delete a device posture rule
page_id: operation-delete-accounts-account-id-devices-posture-rule-id-978517d9
path: operations/device-posture-rules
description: Deletes a device posture rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/devices/posture/{rule_id}
operation_ids:
    - device-posture-rules-delete-device-posture-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a device posture rule

`DELETE /accounts/{account_id}/devices/posture/{rule_id}`

Operation ID: `device-posture-rules-delete-device-posture-rule`

Deletes a device posture rule.

## Definition

```yaml
{"operationId": "device-posture-rules-delete-device-posture-rule", "summary": "Delete a device posture rule", "description": "Deletes a device posture rule.", "parameters": [{"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete a device posture rule response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_id_response"}}}}, "4XX": {"description": "Delete a device posture rule response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_id_response"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Device posture rules"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.posture", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
