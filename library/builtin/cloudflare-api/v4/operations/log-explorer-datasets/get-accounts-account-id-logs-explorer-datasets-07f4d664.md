---
title: List account datasets
page_id: operation-get-accounts-account-id-logs-explorer-datasets-4e6d3681
path: operations/log-explorer-datasets
description: |-
    Returns all Log Explorer datasets configured for the account.

    Pass `include_zones=true` to also include zone-level datasets that
    belong to this account. List responses omit the `fields` property;
    use the single-dataset endpoint to retrieve field configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/logs/explorer/datasets
operation_ids:
    - accounts-logs-explorer-datasets-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List account datasets

`GET /accounts/{account_id}/logs/explorer/datasets`

Operation ID: `accounts-logs-explorer-datasets-list`

Returns all Log Explorer datasets configured for the account.

Pass `include_zones=true` to also include zone-level datasets that
belong to this account. List responses omit the `fields` property;
use the single-dataset endpoint to retrieve field configuration.

## Definition

```yaml
{"operationId": "accounts-logs-explorer-datasets-list", "summary": "List account datasets", "description": "Returns all Log Explorer datasets configured for the account.\n\nPass `include_zones=true` to also include zone-level datasets that\nbelong to this account. List responses omit the `fields` property;\nuse the single-dataset endpoint to retrieve field configuration.\n", "parameters": [{"$ref": "#/components/parameters/lex_AccountID"}, {"name": "include_zones", "in": "query", "description": "Set to true to include zone-scoped datasets belonging to this account.", "schema": {"type": "boolean"}}], "responses": {"200": {"description": "The datasets the account has configured.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/lex_DatasetSummaryListResponse"}}}}, "400": {"$ref": "#/components/responses/lex_BadRequest-2"}, "403": {"$ref": "#/components/responses/lex_Forbidden-2"}, "404": {"$ref": "#/components/responses/lex_NotFound-2"}}, "security": [{"api_token": []}], "tags": ["Log Explorer Datasets"], "x-api-token-group": ["Logs Write", "Logs Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logexplorer.datasets", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
