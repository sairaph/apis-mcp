---
title: List available zone datasets
page_id: operation-get-zones-zone-id-logs-explorer-datasets-available-fc8db758
path: operations/log-explorer-datasets
description: |-
    Returns all dataset types that this zone can create. Each entry includes
    the dataset schema and timestamp field.

    The schema shows all possible fields for a dataset. However, not all
    fields may be available for your account. When creating or updating a
    dataset, only fields available to your account can be enabled. If you
    request a field that is not available, you will receive an error.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/logs/explorer/datasets/available
operation_ids:
    - zones-logs-explorer-datasets-available-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List available zone datasets

`GET /zones/{zone_id}/logs/explorer/datasets/available`

Operation ID: `zones-logs-explorer-datasets-available-list`

Returns all dataset types that this zone can create. Each entry includes
the dataset schema and timestamp field.

The schema shows all possible fields for a dataset. However, not all
fields may be available for your account. When creating or updating a
dataset, only fields available to your account can be enabled. If you
request a field that is not available, you will receive an error.

## Definition

```yaml
{"operationId": "zones-logs-explorer-datasets-available-list", "summary": "List available zone datasets", "description": "Returns all dataset types that this zone can create. Each entry includes\nthe dataset schema and timestamp field.\n\nThe schema shows all possible fields for a dataset. However, not all\nfields may be available for your account. When creating or updating a\ndataset, only fields available to your account can be enabled. If you\nrequest a field that is not available, you will receive an error.\n", "parameters": [{"$ref": "#/components/parameters/lex_ZoneID"}], "responses": {"200": {"description": "List of dataset types available to create.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/lex_AvailableDestinationListResponse"}}}}, "400": {"$ref": "#/components/responses/lex_BadRequest-2"}, "403": {"$ref": "#/components/responses/lex_Forbidden-2"}, "404": {"$ref": "#/components/responses/lex_NotFound-2"}}, "security": [{"api_token": []}], "tags": ["Log Explorer Datasets"], "x-api-token-group": ["Logs Write", "Logs Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logexplorer.datasets.available", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
