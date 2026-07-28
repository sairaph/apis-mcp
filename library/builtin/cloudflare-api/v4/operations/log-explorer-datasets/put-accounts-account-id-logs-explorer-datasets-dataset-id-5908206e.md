---
title: Update an account dataset
page_id: operation-put-accounts-account-id-logs-explorer-datasets-dataset-id-ca589c21
path: operations/log-explorer-datasets
description: Updates the enabled state and/or field configuration of an account dataset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/logs/explorer/datasets/{dataset_id}
operation_ids:
    - accounts-logs-explorer-datasets-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update an account dataset

`PUT /accounts/{account_id}/logs/explorer/datasets/{dataset_id}`

Operation ID: `accounts-logs-explorer-datasets-update`

Updates the enabled state and/or field configuration of an account dataset.

## Definition

```yaml
{"operationId": "accounts-logs-explorer-datasets-update", "summary": "Update an account dataset", "description": "Updates the enabled state and/or field configuration of an account dataset.", "parameters": [{"$ref": "#/components/parameters/lex_AccountID"}, {"$ref": "#/components/parameters/lex_DatasetID"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/lex_UpdateDatasetRequest"}}}}, "responses": {"200": {"description": "The dataset after the update.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/lex_DatasetDetailResponse"}}}}, "400": {"$ref": "#/components/responses/lex_BadRequest-2"}, "403": {"$ref": "#/components/responses/lex_Forbidden-2"}, "404": {"$ref": "#/components/responses/lex_NotFound-2"}}, "security": [{"api_token": []}], "tags": ["Log Explorer Datasets"], "x-api-token-group": ["Logs Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logexplorer.datasets", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
