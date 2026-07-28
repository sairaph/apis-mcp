---
title: List account commands
page_id: operation-get-accounts-account-id-dex-commands-18188a39
path: operations/dex-remote-commands
description: Retrieves a paginated list of commands issued to devices under the specified account, optionally filtered by time range, device, or other parameters
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dex/commands
operation_ids:
    - get-commands
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List account commands

`GET /accounts/{account_id}/dex/commands`

Operation ID: `get-commands`

Retrieves a paginated list of commands issued to devices under the specified account, optionally filtered by time range, device, or other parameters

## Definition

```yaml
{"operationId": "get-commands", "summary": "List account commands", "description": "Retrieves a paginated list of commands issued to devices under the specified account, optionally filtered by time range, device, or other parameters", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}, {"name": "page", "in": "query", "description": "Page number of paginated results.", "required": true, "schema": {"type": "number", "example": 1, "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "description": "Number of results per page.", "required": true, "schema": {"type": "number", "example": 10, "maximum": 50, "minimum": 1}}, {"name": "from", "in": "query", "description": "Start time for the query in ISO (RFC3339 - ISO 8601) format.", "schema": {"type": "string", "format": "date-time", "example": "2023-08-20T20:45:00Z"}}, {"name": "to", "in": "query", "description": "End time for the query in ISO (RFC3339 - ISO 8601) format.", "schema": {"type": "string", "format": "date-time", "example": "2023-08-24T20:45:00Z"}}, {"name": "device_id", "in": "query", "description": "Unique identifier for a device.", "schema": {"type": "string"}}, {"name": "user_email", "in": "query", "description": "Email tied to the device.", "schema": {"type": "string"}}, {"name": "command_type", "in": "query", "description": "Optionally filter executed commands by command type.", "schema": {"type": "string", "enum": ["pcap", "speed-test", "warp-diag"]}}, {"name": "status", "in": "query", "description": "Optionally filter executed commands by status.", "schema": {"type": "string", "enum": ["PENDING_EXEC", "PENDING_UPLOAD", "SUCCESS", "FAILED"]}}], "responses": {"200": {"description": "Get commands response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_api-response-collection"}, {"properties": {"result": {"$ref": "#/components/schemas/digital-experience-monitoring_get_commands_response"}}}]}}}}, "4XX": {"description": "Get commands failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["DEX Remote Commands"], "x-api-token-group": ["Cloudflare DEX Write", "Cloudflare DEX Read", "Zero Trust Report", "Zero Trust Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.commands", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
