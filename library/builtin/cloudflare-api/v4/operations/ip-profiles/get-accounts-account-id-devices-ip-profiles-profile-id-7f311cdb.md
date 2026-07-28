---
title: Get IP profile
page_id: operation-get-accounts-account-id-devices-ip-profiles-profile-id-66044e49
path: operations/ip-profiles
description: Fetches a single WARP Device IP profile.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices/ip-profiles/{profile_id}
operation_ids:
    - get-ip-profile
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get IP profile

`GET /accounts/{account_id}/devices/ip-profiles/{profile_id}`

Operation ID: `get-ip-profile`

Fetches a single WARP Device IP profile.

## Definition

```yaml
{"operationId": "get-ip-profile", "summary": "Get IP profile", "description": "Fetches a single WARP Device IP profile.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "profile_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Get Device IP profile response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "result": {"$ref": "#/components/schemas/teams-devices_ip_profile"}, "success": {"description": "Whether the API call was successful.", "type": "boolean"}}, "required": ["result", "success", "errors", "messages"]}}}}}, "security": [{"api_token": []}], "tags": ["IP Profiles"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.ip-profiles", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
