---
title: Unrevoke registrations
page_id: operation-post-accounts-account-id-devices-registrations-unrevoke-8c2af4d0
path: operations/registrations
description: Unrevokes a list of WARP registrations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/devices/registrations/unrevoke
operation_ids:
    - unrevoke-registrations
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Unrevoke registrations

`POST /accounts/{account_id}/devices/registrations/unrevoke`

Operation ID: `unrevoke-registrations`

Unrevokes a list of WARP registrations.

## Definition

```yaml
{"operationId": "unrevoke-registrations", "summary": "Unrevoke registrations", "description": "Unrevokes a list of WARP registrations.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "id", "in": "query", "description": "A list of registration IDs to unrevoke.", "required": true, "schema": {"type": "array", "items": {"type": "string"}}}], "responses": {"200": {"description": "Unrevoke registrations response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "result": {"$ref": "#/components/schemas/teams-devices_empty_body"}, "result_info": {"$ref": "#/components/schemas/teams-devices_cursor_result_info"}, "success": {"description": "Whether the API call was successful.", "type": "boolean"}}, "required": ["result", "success", "errors", "messages"]}}}}}, "security": [{"api_token": []}], "tags": ["Registrations"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.registrations", "x-fern-sdk-method-name": "unrevoke", "x-forge-hidden": true}
```
