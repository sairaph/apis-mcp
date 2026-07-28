---
title: Create account commands
page_id: operation-post-accounts-account-id-dex-commands-82f6e13f
path: operations/dex-remote-commands
description: Initiate commands for up to 10 devices per account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/dex/commands
operation_ids:
    - post-commands
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create account commands

`POST /accounts/{account_id}/dex/commands`

Operation ID: `post-commands`

Initiate commands for up to 10 devices per account.

## Definition

```yaml
{"operationId": "post-commands", "summary": "Create account commands", "description": "Initiate commands for up to 10 devices per account.", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"commands": {"description": "List of device-level commands to execute", "type": "array", "items": {"properties": {"args": {"description": "Command arguments. Allowed fields depend on `type`.", "type": "object", "oneOf": [{"additionalProperties": false, "properties": {"test-all-routes": {"description": "Test an IP address from all included or excluded ranges. Essentially the same as running 'route get <ip>' and collecting the results. This option may increase the time taken to collect the warp-diag.", "type": "boolean", "default": true, "x-auditable": true}}, "title": "warp-diag args", "type": "object"}, {"additionalProperties": false, "properties": {"max-file-size-mb": {"description": "Maximum file size (in MB) for the capture file. If the capture artifact exceeds the specified max file size, it will NOT be uploaded.", "type": "number", "default": 5, "minimum": 1, "x-auditable": true}, "packet-size-bytes": {"description": "Maximum number of bytes to save for each packet", "type": "number", "default": 160, "minimum": 1, "x-auditable": true}, "time-limit-min": {"description": "Limit on capture duration (in minutes)", "type": "number", "default": 5, "minimum": 1, "x-auditable": true}}, "title": "pcap args", "type": "object"}, {"additionalProperties": false, "properties": {"interfaces": {"description": "List of interfaces to run the speed test on", "type": "array", "items": {"enum": ["default", "tunnel"], "type": "string", "x-auditable": true}}}, "title": "speed-test args", "type": "object"}]}, "device_id": {"description": "Unique identifier for the physical device", "type": "string"}, "registration_id": {"description": "Unique identifier for the device registration. Required for multi-user devices to target the correct user session.", "type": "string"}, "type": {"description": "Type of command to execute on the device", "type": "string", "enum": ["pcap", "speed-test", "warp-diag"], "x-auditable": true}, "user_email": {"description": "Email tied to the device", "type": "string"}}, "required": ["device_id", "user_email", "type"], "type": "object"}, "maxItems": 20}}, "required": ["commands"]}}}}, "responses": {"200": {"description": "Create commands response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_api-response-collection"}, {"properties": {"result": {"$ref": "#/components/schemas/digital-experience-monitoring_post_commands_response"}}}]}}}}, "4XX": {"description": "Create commands failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["DEX Remote Commands"], "x-api-token-group": ["Cloudflare DEX Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.commands", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
