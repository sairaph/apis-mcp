---
title: Get override codes
page_id: operation-get-accounts-account-id-devices-registrations-registration-id-override-c-5a62a8e8
path: operations/warp-teams-device-api-other
description: Fetches one-time use admin override codes for a registration. This relies on the **Admin Override** setting being enabled in your device configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices/registrations/{registration_id}/override_codes
operation_ids:
    - get-registration-override-codes
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get override codes

`GET /accounts/{account_id}/devices/registrations/{registration_id}/override_codes`

Operation ID: `get-registration-override-codes`

Fetches one-time use admin override codes for a registration. This relies on the **Admin Override** setting being enabled in your device configuration.

## Definition

```yaml
{"operationId": "get-registration-override-codes", "summary": "Get override codes", "description": "Fetches one-time use admin override codes for a registration. This relies on the **Admin Override** setting being enabled in your device configuration.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "registration_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Get admin override codes for a registration response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "result": {"$ref": "#/components/schemas/teams-devices_override_codes"}, "success": {"description": "Whether the API call was successful.", "type": "boolean"}}, "required": ["result", "success", "errors", "messages"]}}}}}, "security": [{"api_token": []}], "tags": ["warp-teams-device-api_other"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.override-codes", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
