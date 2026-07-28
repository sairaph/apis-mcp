---
title: Update a device posture rule
page_id: operation-put-accounts-account-id-devices-posture-rule-id-6ba7351c
path: operations/device-posture-rules
description: Updates a device posture rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/devices/posture/{rule_id}
operation_ids:
    - device-posture-rules-update-device-posture-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a device posture rule

`PUT /accounts/{account_id}/devices/posture/{rule_id}`

Operation ID: `device-posture-rules-update-device-posture-rule`

Updates a device posture rule.

## Definition

```yaml
{"operationId": "device-posture-rules-update-device-posture-rule", "summary": "Update a device posture rule", "description": "Updates a device posture rule.", "parameters": [{"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"description": {"$ref": "#/components/schemas/teams-devices_description"}, "expiration": {"$ref": "#/components/schemas/teams-devices_expiration"}, "input": {"$ref": "#/components/schemas/teams-devices_input"}, "match": {"$ref": "#/components/schemas/teams-devices_match"}, "name": {"$ref": "#/components/schemas/teams-devices_name"}, "schedule": {"$ref": "#/components/schemas/teams-devices_schedule"}, "type": {"$ref": "#/components/schemas/teams-devices_type"}}, "required": ["name", "type"]}}}}, "responses": {"200": {"description": "Update a device posture rule response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_single_response"}}}}, "4XX": {"description": "Update a device posture rule response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_single_response"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Device posture rules"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.posture", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
