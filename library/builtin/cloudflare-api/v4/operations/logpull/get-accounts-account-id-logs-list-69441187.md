---
title: List log files
page_id: operation-get-accounts-account-id-logs-list-53b03f9b
path: operations/logpull
description: Lists R2 objects containing logs matching the provided query parameters.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/logs/list
operation_ids:
    - logpull-list-log-files
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List log files

`GET /accounts/{account_id}/logs/list`

Operation ID: `logpull-list-log-files`

Lists R2 objects containing logs matching the provided query parameters.

## Definition

```yaml
{"operationId": "logpull-list-log-files", "summary": "List log files", "description": "Lists R2 objects containing logs matching the provided query parameters.", "parameters": [{"$ref": "#/components/parameters/logpull_account_id"}, {"$ref": "#/components/parameters/logpull_accessKeyIdHeader"}, {"$ref": "#/components/parameters/logpull_secretAccessKeyHeader"}, {"$ref": "#/components/parameters/logpull_startParam"}, {"$ref": "#/components/parameters/logpull_endParam"}, {"$ref": "#/components/parameters/logpull_bucketParam"}, {"$ref": "#/components/parameters/logpull_prefixParam"}, {"$ref": "#/components/parameters/logpull_limitParam"}], "responses": {"200": {"$ref": "#/components/responses/logpull_ListKeysOKResponse"}, "400": {"$ref": "#/components/responses/logpull_BadRequestResponse"}, "401": {"$ref": "#/components/responses/logpull_UnauthorizedResponse"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Logpull"], "x-api-token-group": ["Logs Write", "Logs Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logs.list", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
