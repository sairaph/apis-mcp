---
title: Get Chrome DevTools Protocol schema.
page_id: operation-get-accounts-account-id-browser-rendering-devtools-browser-session-id-js-af5b3108
path: operations/brapi
description: Returns the complete Chrome DevTools Protocol schema including all domains, commands, events, and types. This schema describes the entire CDP API surface.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/browser-rendering/devtools/browser/{session_id}/json/protocol
operation_ids:
    - brapi-get_DevtoolsJsonProtocol
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Chrome DevTools Protocol schema.

`GET /accounts/{account_id}/browser-rendering/devtools/browser/{session_id}/json/protocol`

Operation ID: `brapi-get_DevtoolsJsonProtocol`

Returns the complete Chrome DevTools Protocol schema including all domains, commands, events, and types. This schema describes the entire CDP API surface.

## Definition

```yaml
{"operationId": "brapi-get_DevtoolsJsonProtocol", "summary": "Get Chrome DevTools Protocol schema.", "description": "Returns the complete Chrome DevTools Protocol schema including all domains, commands, events, and types. This schema describes the entire CDP API surface.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "session_id", "in": "path", "description": "Browser session ID.", "required": true, "schema": {"description": "Browser session ID.", "type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Chrome DevTools Protocol schema.", "content": {"application/json": {"schema": {"type": "object", "properties": {"domains": {"description": "List of protocol domains.", "type": "array", "items": {"properties": {"commands": {"description": "Available commands.", "type": "array", "items": {"additionalProperties": {"nullable": true, "type": "object"}, "type": "object"}}, "dependencies": {"description": "Domain dependencies.", "type": "array", "items": {"type": "string"}}, "domain": {"description": "Domain name.", "type": "string"}, "events": {"description": "Available events.", "type": "array", "items": {"additionalProperties": {"nullable": true, "type": "object"}, "type": "object"}}, "experimental": {"description": "Whether this domain is experimental.", "type": "boolean"}, "types": {"description": "Type definitions.", "type": "array", "items": {"additionalProperties": {"nullable": true, "type": "object"}, "type": "object"}}}, "required": ["domain"], "type": "object"}}, "version": {"description": "Protocol version.", "type": "object", "properties": {"major": {"description": "Major version.", "type": "string"}, "minor": {"description": "Minor version.", "type": "string"}}, "required": ["major", "minor"]}}, "required": ["domains"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"description": "Error code.", "type": "number"}, "message": {"description": "Error message.", "type": "string"}}, "required": ["message", "code"], "type": "object"}}, "success": {"description": "Response status.", "type": "boolean"}}, "required": ["success"]}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"description": "Error code.", "type": "number"}, "message": {"description": "Error message.", "type": "string"}}, "required": ["message", "code"], "type": "object"}}, "success": {"description": "Response status.", "type": "boolean"}}, "required": ["success"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["brapi"], "x-api-token-group": ["Browser Rendering Write", "Browser Rendering Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.brapi.read"], "type": "string"}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "browser-run.devtools.browser", "x-fern-sdk-method-name": "protocol", "x-forge-hidden": true}
```
