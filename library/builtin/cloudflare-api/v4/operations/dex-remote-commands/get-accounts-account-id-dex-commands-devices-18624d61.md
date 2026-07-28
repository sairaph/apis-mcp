---
title: List devices eligible for remote captures
page_id: operation-get-accounts-account-id-dex-commands-devices-14c4390c
path: operations/dex-remote-commands
description: List devices with WARP client support for remote captures which have been connected in the last 1 hour.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dex/commands/devices
operation_ids:
    - get-commands-eligible-devices
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List devices eligible for remote captures

`GET /accounts/{account_id}/dex/commands/devices`

Operation ID: `get-commands-eligible-devices`

List devices with WARP client support for remote captures which have been connected in the last 1 hour.

## Definition

```yaml
{"operationId": "get-commands-eligible-devices", "summary": "List devices eligible for remote captures", "description": "List devices with WARP client support for remote captures which have been connected in the last 1 hour.", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}, {"name": "page", "in": "query", "description": "Page number of paginated results.", "required": true, "schema": {"type": "number", "example": 1, "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "description": "Number of results per page.", "required": true, "schema": {"type": "number", "example": 10, "maximum": 50, "minimum": 1}}, {"name": "search", "in": "query", "description": "Filter devices by name or email.", "schema": {"type": "string"}}], "responses": {"200": {"description": "List of eligible devices.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_api-response-collection"}, {"properties": {"result": {"$ref": "#/components/schemas/digital-experience-monitoring_commands_devices_response"}}}]}}}}, "4XX": {"description": "List eligible devices failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["DEX Remote Commands"], "x-api-token-group": ["Cloudflare DEX Write", "Cloudflare DEX Read", "Zero Trust Report", "Zero Trust Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.commands.eligible-devices", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
