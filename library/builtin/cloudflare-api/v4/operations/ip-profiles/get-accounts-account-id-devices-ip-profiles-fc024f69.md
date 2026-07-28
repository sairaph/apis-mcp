---
title: List IP profiles
page_id: operation-get-accounts-account-id-devices-ip-profiles-9636fb45
path: operations/ip-profiles
description: Lists WARP Device IP profiles.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices/ip-profiles
operation_ids:
    - list-ip-profiles
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List IP profiles

`GET /accounts/{account_id}/devices/ip-profiles`

Operation ID: `list-ip-profiles`

Lists WARP Device IP profiles.

## Definition

```yaml
{"operationId": "list-ip-profiles", "summary": "List IP profiles", "description": "Lists WARP Device IP profiles.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "page", "in": "query", "description": "The page number to return.", "schema": {"type": "integer", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "description": "The number of IP profiles to return per page.", "schema": {"type": "integer", "default": 50, "maximum": 100, "minimum": 1}}], "responses": {"200": {"description": "List Device IP profiles response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "result": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_ip_profile"}}, "result_info": {"$ref": "#/components/schemas/teams-devices_pagination_info"}, "success": {"description": "Whether the API call was successful.", "type": "boolean"}}, "required": ["result", "success", "errors", "messages"]}}}}}, "security": [{"api_token": []}], "tags": ["IP Profiles"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.ip-profiles", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
