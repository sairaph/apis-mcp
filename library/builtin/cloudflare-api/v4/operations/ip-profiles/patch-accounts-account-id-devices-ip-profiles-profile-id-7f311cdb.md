---
title: Update IP profile
page_id: operation-patch-accounts-account-id-devices-ip-profiles-profile-id-09b71562
path: operations/ip-profiles
description: Updates a WARP Device IP profile. Currently, only IPv4 Device subnets can be associated.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/devices/ip-profiles/{profile_id}
operation_ids:
    - update-ip-profile
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update IP profile

`PATCH /accounts/{account_id}/devices/ip-profiles/{profile_id}`

Operation ID: `update-ip-profile`

Updates a WARP Device IP profile. Currently, only IPv4 Device subnets can be associated.

## Definition

```yaml
{"operationId": "update-ip-profile", "summary": "Update IP profile", "description": "Updates a WARP Device IP profile. Currently, only IPv4 Device subnets can be associated.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "profile_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_ip_profile_update_request"}}}}, "responses": {"200": {"description": "Update Device IP profile response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "result": {"$ref": "#/components/schemas/teams-devices_ip_profile"}, "success": {"description": "Whether the API call was successful.", "type": "boolean"}}, "required": ["result", "success", "errors", "messages"]}}}}}, "security": [{"api_token": []}], "tags": ["IP Profiles"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.ip-profiles", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
