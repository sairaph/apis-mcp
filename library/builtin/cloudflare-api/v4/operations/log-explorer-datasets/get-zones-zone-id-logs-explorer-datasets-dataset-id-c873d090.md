---
title: Get a zone dataset
page_id: operation-get-zones-zone-id-logs-explorer-datasets-dataset-id-155409a1
path: operations/log-explorer-datasets
description: Retrieve a single Log Explorer dataset by ID for the zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/logs/explorer/datasets/{dataset_id}
operation_ids:
    - zones-logs-explorer-datasets-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a zone dataset

`GET /zones/{zone_id}/logs/explorer/datasets/{dataset_id}`

Operation ID: `zones-logs-explorer-datasets-get`

Retrieve a single Log Explorer dataset by ID for the zone.

## Definition

```yaml
{"operationId": "zones-logs-explorer-datasets-get", "summary": "Get a zone dataset", "description": "Retrieve a single Log Explorer dataset by ID for the zone.", "parameters": [{"$ref": "#/components/parameters/lex_ZoneID"}, {"$ref": "#/components/parameters/lex_DatasetID"}], "responses": {"200": {"description": "Dataset details, including the fields active for ingestion.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/lex_DatasetDetailResponse"}}}}, "400": {"$ref": "#/components/responses/lex_BadRequest-2"}, "403": {"$ref": "#/components/responses/lex_Forbidden-2"}, "404": {"$ref": "#/components/responses/lex_NotFound-2"}}, "security": [{"api_token": []}], "tags": ["Log Explorer Datasets"], "x-api-token-group": ["Logs Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logexplorer.datasets", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
