---
title: Returns account commands usage, quota, and reset time
page_id: operation-get-accounts-account-id-dex-commands-quota-2cb5e593
path: operations/dex-remote-commands
description: Retrieves the current quota usage and limits for device commands within a specific account, including the time when the quota will reset
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dex/commands/quota
operation_ids:
    - get-commands-quota
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Returns account commands usage, quota, and reset time

`GET /accounts/{account_id}/dex/commands/quota`

Operation ID: `get-commands-quota`

Retrieves the current quota usage and limits for device commands within a specific account, including the time when the quota will reset

## Definition

```yaml
{"operationId": "get-commands-quota", "summary": "Returns account commands usage, quota, and reset time", "description": "Retrieves the current quota usage and limits for device commands within a specific account, including the time when the quota will reset", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}], "responses": {"200": {"description": "Get commands quota response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_api-response-collection"}, {"properties": {"result": {"$ref": "#/components/schemas/digital-experience-monitoring_get_commands_quota_response"}}}]}}}}, "4XX": {"description": "Get commands quota failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["DEX Remote Commands"], "x-api-token-group": ["Cloudflare DEX Write", "Cloudflare DEX Read", "Zero Trust Report", "Zero Trust Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.commands.quota", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
