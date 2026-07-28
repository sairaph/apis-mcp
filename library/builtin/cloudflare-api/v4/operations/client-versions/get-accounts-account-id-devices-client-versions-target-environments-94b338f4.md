---
title: List available target environments
page_id: operation-get-accounts-account-id-devices-client-versions-target-environments-36baf5c5
path: operations/client-versions
description: Retrieves a list of all available target environments with their display names. This endpoint is in Beta.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices/client-versions/target-environments
operation_ids:
    - list-client-target-environments
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List available target environments

`GET /accounts/{account_id}/devices/client-versions/target-environments`

Operation ID: `list-client-target-environments`

Retrieves a list of all available target environments with their display names. This endpoint is in Beta.

## Definition

```yaml
{"operationId": "list-client-target-environments", "summary": "List available target environments", "description": "Retrieves a list of all available target environments with their display names. This endpoint is in Beta.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "List target environments response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "result": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_target_environment_info"}}, "success": {"description": "Indicates whether the API call was successful.", "type": "boolean"}}, "required": ["result", "success", "errors", "messages"]}}}}}, "security": [{"api_token": []}], "tags": ["Client Versions"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.client.target.environments", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
