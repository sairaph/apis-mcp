---
title: Update a zone dataset
page_id: operation-put-zones-zone-id-logs-explorer-datasets-dataset-id-694c7b34
path: operations/log-explorer-datasets
description: Updates the enabled state and/or field configuration of a zone dataset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/logs/explorer/datasets/{dataset_id}
operation_ids:
    - zones-logs-explorer-datasets-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a zone dataset

`PUT /zones/{zone_id}/logs/explorer/datasets/{dataset_id}`

Operation ID: `zones-logs-explorer-datasets-update`

Updates the enabled state and/or field configuration of a zone dataset.

## Definition

```yaml
{"operationId": "zones-logs-explorer-datasets-update", "summary": "Update a zone dataset", "description": "Updates the enabled state and/or field configuration of a zone dataset.", "parameters": [{"$ref": "#/components/parameters/lex_ZoneID"}, {"$ref": "#/components/parameters/lex_DatasetID"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/lex_UpdateDatasetRequest"}}}}, "responses": {"200": {"description": "The dataset after the update.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/lex_DatasetDetailResponse"}}}}, "400": {"$ref": "#/components/responses/lex_BadRequest-2"}, "403": {"$ref": "#/components/responses/lex_Forbidden-2"}, "404": {"$ref": "#/components/responses/lex_NotFound-2"}}, "security": [{"api_token": []}], "tags": ["Log Explorer Datasets"], "x-api-token-group": ["Logs Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logexplorer.datasets", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
