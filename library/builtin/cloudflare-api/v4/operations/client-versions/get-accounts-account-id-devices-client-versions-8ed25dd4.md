---
title: List client versions
page_id: operation-get-accounts-account-id-devices-client-versions-7a8b058c
path: operations/client-versions
description: Lists available WARP client versions for a specific target environment and release track. This endpoint is in Beta.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices/client-versions
operation_ids:
    - list-client-versions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List client versions

`GET /accounts/{account_id}/devices/client-versions`

Operation ID: `list-client-versions`

Lists available WARP client versions for a specific target environment and release track. This endpoint is in Beta.

## Definition

```yaml
{"operationId": "list-client-versions", "summary": "List client versions", "description": "Lists available WARP client versions for a specific target environment and release track. This endpoint is in Beta.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "target_environment", "in": "query", "description": "The target environment for the client version (e.g., windows, macos).", "required": true, "schema": {"type": "string"}}, {"name": "release_track", "in": "query", "description": "The release track (ga for General Availability, beta for Beta releases).", "required": true, "schema": {"type": "string", "enum": ["ga", "beta"]}}, {"name": "page", "in": "query", "description": "The page number to return.", "schema": {"type": "integer", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "description": "The maximum number of versions to return per page.", "schema": {"type": "integer", "default": 20, "maximum": 100, "minimum": 1}}], "responses": {"200": {"description": "List client versions response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "result": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_client_version"}}, "result_info": {"$ref": "#/components/schemas/teams-devices_pagination_info"}, "success": {"description": "Indicates whether the API call was successful.", "type": "boolean"}}, "required": ["result", "result_info", "success", "errors", "messages"]}}}}}, "security": [{"api_token": []}], "tags": ["Client Versions"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.client.versions", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
