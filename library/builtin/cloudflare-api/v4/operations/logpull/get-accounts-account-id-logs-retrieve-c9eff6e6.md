---
title: Retrieve logs
page_id: operation-get-accounts-account-id-logs-retrieve-579bd714
path: operations/logpull
description: Returns logs stored in R2 that match the provided query parameters.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/logs/retrieve
operation_ids:
    - logpull-retrieve-logs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve logs

`GET /accounts/{account_id}/logs/retrieve`

Operation ID: `logpull-retrieve-logs`

Returns logs stored in R2 that match the provided query parameters.

## Definition

```yaml
{"operationId": "logpull-retrieve-logs", "summary": "Retrieve logs", "description": "Returns logs stored in R2 that match the provided query parameters.", "parameters": [{"$ref": "#/components/parameters/logpull_account_id"}, {"$ref": "#/components/parameters/logpull_accessKeyIdHeader"}, {"$ref": "#/components/parameters/logpull_secretAccessKeyHeader"}, {"$ref": "#/components/parameters/logpull_startParam"}, {"$ref": "#/components/parameters/logpull_endParam"}, {"$ref": "#/components/parameters/logpull_bucketParam"}, {"$ref": "#/components/parameters/logpull_prefixParam"}], "responses": {"200": {"$ref": "#/components/responses/logpull_RetrieveLogsOKResponse"}, "400": {"$ref": "#/components/responses/logpull_BadRequestResponse"}, "401": {"$ref": "#/components/responses/logpull_UnauthorizedResponse"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Logpull"], "x-api-token-group": ["Logs Write", "Logs Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logs.retrieve", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
