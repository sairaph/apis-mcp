---
title: Get latest builds by script IDs
page_id: operation-get-accounts-account-id-builds-builds-latest-dac6a588
path: operations/builds
description: Retrieve the most recent builds for multiple worker scripts
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/builds/builds/latest
operation_ids:
    - getLatestBuildsByScripts
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get latest builds by script IDs

`GET /accounts/{account_id}/builds/builds/latest`

Operation ID: `getLatestBuildsByScripts`

Retrieve the most recent builds for multiple worker scripts

## Definition

```yaml
{"operationId": "getLatestBuildsByScripts", "summary": "Get latest builds by script IDs", "description": "Retrieve the most recent builds for multiple worker scripts", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}, {"name": "external_script_ids", "in": "query", "required": true, "schema": {"$ref": "#/components/schemas/builds_external_script_ids"}}], "responses": {"200": {"description": "Latest builds retrieved successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/builds_APIResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/builds_LatestBuildsResponse"}}, "type": "object"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Builds"], "x-api-token-group": ["Workers CI Write", "Workers CI Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers-builds", "x-fern-sdk-method-name": "get-latest-builds", "x-forge-hidden": true}
```
