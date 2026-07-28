---
title: Get device posture rule details
page_id: operation-get-accounts-account-id-devices-posture-rule-id-d2608c90
path: operations/device-posture-rules
description: Fetches a single device posture rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices/posture/{rule_id}
operation_ids:
    - device-posture-rules-device-posture-rules-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get device posture rule details

`GET /accounts/{account_id}/devices/posture/{rule_id}`

Operation ID: `device-posture-rules-device-posture-rules-details`

Fetches a single device posture rule.

## Definition

```yaml
{"operationId": "device-posture-rules-device-posture-rules-details", "summary": "Get device posture rule details", "description": "Fetches a single device posture rule.", "parameters": [{"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "responses": {"200": {"description": "Get device posture rule details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_single_response"}}}}, "4XX": {"description": "Get device posture rule details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_single_response"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Device posture rules"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.posture", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
