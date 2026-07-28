---
title: List device posture rules
page_id: operation-get-accounts-account-id-devices-posture-b70bcee5
path: operations/device-posture-rules
description: Fetches device posture rules for a Zero Trust account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices/posture
operation_ids:
    - device-posture-rules-list-device-posture-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List device posture rules

`GET /accounts/{account_id}/devices/posture`

Operation ID: `device-posture-rules-list-device-posture-rules`

Fetches device posture rules for a Zero Trust account.

## Definition

```yaml
{"operationId": "device-posture-rules-list-device-posture-rules", "summary": "List device posture rules", "description": "Fetches device posture rules for a Zero Trust account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "responses": {"200": {"description": "List device posture rules response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_response_collection"}}}}, "4XX": {"description": "List device posture rules response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_response_collection"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Device posture rules"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.posture", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
