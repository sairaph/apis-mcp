---
title: List triggers by script
page_id: operation-get-accounts-account-id-builds-workers-external-script-id-triggers-3822fd8b
path: operations/workers
description: Get all triggers for a specific worker script
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/builds/workers/{external_script_id}/triggers
operation_ids:
    - listTriggersByScript
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List triggers by script

`GET /accounts/{account_id}/builds/workers/{external_script_id}/triggers`

Operation ID: `listTriggersByScript`

Get all triggers for a specific worker script

## Definition

```yaml
{"operationId": "listTriggersByScript", "summary": "List triggers by script", "description": "Get all triggers for a specific worker script", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}, {"name": "external_script_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/builds_external_script_id"}}], "responses": {"200": {"description": "Triggers retrieved successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/builds_APIResponse"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/builds_TriggerResponse"}}}, "type": "object"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers"], "x-api-token-group": ["Workers CI Write", "Workers CI Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers-builds.triggers", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
