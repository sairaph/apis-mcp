---
title: Get an account dataset
page_id: operation-get-accounts-account-id-logs-explorer-datasets-dataset-id-8dfb52f8
path: operations/log-explorer-datasets
description: Retrieve a single Log Explorer dataset by ID for the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/logs/explorer/datasets/{dataset_id}
operation_ids:
    - accounts-logs-explorer-datasets-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get an account dataset

`GET /accounts/{account_id}/logs/explorer/datasets/{dataset_id}`

Operation ID: `accounts-logs-explorer-datasets-get`

Retrieve a single Log Explorer dataset by ID for the account.

## Definition

```yaml
{"operationId": "accounts-logs-explorer-datasets-get", "summary": "Get an account dataset", "description": "Retrieve a single Log Explorer dataset by ID for the account.", "parameters": [{"$ref": "#/components/parameters/lex_AccountID"}, {"$ref": "#/components/parameters/lex_DatasetID"}], "responses": {"200": {"description": "Dataset details, including the fields active for ingestion.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/lex_DatasetDetailResponse"}}}}, "400": {"$ref": "#/components/responses/lex_BadRequest-2"}, "403": {"$ref": "#/components/responses/lex_Forbidden-2"}, "404": {"$ref": "#/components/responses/lex_NotFound-2"}}, "security": [{"api_token": []}], "tags": ["Log Explorer Datasets"], "x-api-token-group": ["Logs Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logexplorer.datasets", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
