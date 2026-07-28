---
title: Create a zone dataset
page_id: operation-post-zones-zone-id-logs-explorer-datasets-14eb1820
path: operations/log-explorer-datasets
description: |-
    Create a new Log Explorer dataset for the zone.

    List available zone datasets to see the dataset types and fields you
    can use.

    The `fields` property is optional. If not specified, all available fields
    will be enabled.

    For dataset field definitions, see: https://developers.cloudflare.com/logs/logpush/logpush-job/datasets/
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/logs/explorer/datasets
operation_ids:
    - zones-logs-explorer-datasets-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a zone dataset

`POST /zones/{zone_id}/logs/explorer/datasets`

Operation ID: `zones-logs-explorer-datasets-create`

Create a new Log Explorer dataset for the zone.

List available zone datasets to see the dataset types and fields you
can use.

The `fields` property is optional. If not specified, all available fields
will be enabled.

For dataset field definitions, see: https://developers.cloudflare.com/logs/logpush/logpush-job/datasets/

## Definition

```yaml
{"operationId": "zones-logs-explorer-datasets-create", "summary": "Create a zone dataset", "description": "Create a new Log Explorer dataset for the zone.\n\nList available zone datasets to see the dataset types and fields you\ncan use.\n\nThe `fields` property is optional. If not specified, all available fields\nwill be enabled.\n\nFor dataset field definitions, see: https://developers.cloudflare.com/logs/logpush/logpush-job/datasets/\n", "parameters": [{"$ref": "#/components/parameters/lex_ZoneID"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/lex_CreateDatasetRequest"}}}}, "responses": {"201": {"description": "Dataset created successfully.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/lex_DatasetDetailResponse"}}}}, "400": {"$ref": "#/components/responses/lex_BadRequest-2"}, "403": {"$ref": "#/components/responses/lex_Forbidden-2"}, "404": {"$ref": "#/components/responses/lex_NotFound-2"}, "409": {"$ref": "#/components/responses/lex_Conflict"}}, "security": [{"api_token": []}], "tags": ["Log Explorer Datasets"], "x-api-token-group": ["Logs Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logexplorer.datasets", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
