---
title: Create an account dataset
page_id: operation-post-accounts-account-id-logs-explorer-datasets-1968569a
path: operations/log-explorer-datasets
description: |-
    Create a new Log Explorer dataset for the account.

    List available account datasets to see the dataset types and fields you
    can use.

    The `fields` property is optional. If not specified, all available fields
    will be enabled.

    For dataset field definitions, see: https://developers.cloudflare.com/logs/logpush/logpush-job/datasets/
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/logs/explorer/datasets
operation_ids:
    - accounts-logs-explorer-datasets-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create an account dataset

`POST /accounts/{account_id}/logs/explorer/datasets`

Operation ID: `accounts-logs-explorer-datasets-create`

Create a new Log Explorer dataset for the account.

List available account datasets to see the dataset types and fields you
can use.

The `fields` property is optional. If not specified, all available fields
will be enabled.

For dataset field definitions, see: https://developers.cloudflare.com/logs/logpush/logpush-job/datasets/

## Definition

```yaml
{"operationId": "accounts-logs-explorer-datasets-create", "summary": "Create an account dataset", "description": "Create a new Log Explorer dataset for the account.\n\nList available account datasets to see the dataset types and fields you\ncan use.\n\nThe `fields` property is optional. If not specified, all available fields\nwill be enabled.\n\nFor dataset field definitions, see: https://developers.cloudflare.com/logs/logpush/logpush-job/datasets/\n", "parameters": [{"$ref": "#/components/parameters/lex_AccountID"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/lex_CreateDatasetRequest"}}}}, "responses": {"201": {"description": "Dataset created successfully.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/lex_DatasetDetailResponse"}}}}, "400": {"$ref": "#/components/responses/lex_BadRequest-2"}, "403": {"$ref": "#/components/responses/lex_Forbidden-2"}, "404": {"$ref": "#/components/responses/lex_NotFound-2"}, "409": {"$ref": "#/components/responses/lex_Conflict"}}, "security": [{"api_token": []}], "tags": ["Log Explorer Datasets"], "x-api-token-group": ["Logs Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logexplorer.datasets", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
