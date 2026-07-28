---
title: List zone datasets
page_id: operation-get-zones-zone-id-logs-explorer-datasets-be56d412
path: operations/log-explorer-datasets
description: |-
    Returns all Log Explorer datasets configured for the zone.

    List responses omit the `fields` property; use the single-dataset
    endpoint to retrieve field configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/logs/explorer/datasets
operation_ids:
    - zones-logs-explorer-datasets-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List zone datasets

`GET /zones/{zone_id}/logs/explorer/datasets`

Operation ID: `zones-logs-explorer-datasets-list`

Returns all Log Explorer datasets configured for the zone.

List responses omit the `fields` property; use the single-dataset
endpoint to retrieve field configuration.

## Definition

```yaml
{"operationId": "zones-logs-explorer-datasets-list", "summary": "List zone datasets", "description": "Returns all Log Explorer datasets configured for the zone.\n\nList responses omit the `fields` property; use the single-dataset\nendpoint to retrieve field configuration.\n", "parameters": [{"$ref": "#/components/parameters/lex_ZoneID"}], "responses": {"200": {"description": "The datasets the zone has configured.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/lex_DatasetSummaryListResponse"}}}}, "400": {"$ref": "#/components/responses/lex_BadRequest-2"}, "403": {"$ref": "#/components/responses/lex_Forbidden-2"}, "404": {"$ref": "#/components/responses/lex_NotFound-2"}}, "security": [{"api_token": []}], "tags": ["Log Explorer Datasets"], "x-api-token-group": ["Logs Write", "Logs Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logexplorer.datasets", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
