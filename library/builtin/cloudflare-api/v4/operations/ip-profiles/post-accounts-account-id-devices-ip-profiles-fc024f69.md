---
title: Create IP profile
page_id: operation-post-accounts-account-id-devices-ip-profiles-562e0ae8
path: operations/ip-profiles
description: Creates a WARP Device IP profile. Currently, only IPv4 Device subnets can be associated.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/devices/ip-profiles
operation_ids:
    - create-ip-profile
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create IP profile

`POST /accounts/{account_id}/devices/ip-profiles`

Operation ID: `create-ip-profile`

Creates a WARP Device IP profile. Currently, only IPv4 Device subnets can be associated.

## Definition

```yaml
{"operationId": "create-ip-profile", "summary": "Create IP profile", "description": "Creates a WARP Device IP profile. Currently, only IPv4 Device subnets can be associated.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_ip_profile_create_request"}}}}, "responses": {"200": {"description": "Create Device IP profile response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "result": {"$ref": "#/components/schemas/teams-devices_ip_profile"}, "success": {"description": "Whether the API call was successful.", "type": "boolean"}}, "required": ["result", "success", "errors", "messages"]}}}}}, "security": [{"api_token": []}], "tags": ["IP Profiles"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.ip-profiles", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
