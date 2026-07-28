---
title: Create a new dataset
page_id: operation-post-accounts-account-id-dlp-datasets-6a6d8cd9
path: operations/dlp-datasets
description: Creates a new DLP (Data Loss Prevention) dataset for storing custom detection patterns. Datasets can contain exact match data, word lists, or EDM (Exact Data Match) configurations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/dlp/datasets
operation_ids:
    - dlp-datasets-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a new dataset

`POST /accounts/{account_id}/dlp/datasets`

Operation ID: `dlp-datasets-create`

Creates a new DLP (Data Loss Prevention) dataset for storing custom detection patterns. Datasets can contain exact match data, word lists, or EDM (Exact Data Match) configurations.

## Definition

```yaml
{"operationId": "dlp-datasets-create", "summary": "Create a new dataset", "description": "Creates a new DLP (Data Loss Prevention) dataset for storing custom detection patterns. Datasets can contain exact match data, word lists, or EDM (Exact Data Match) configurations.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"description": "Dataset description.", "required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"case_sensitive": {"description": "Only applies to custom word lists.\nDetermines if the words should be matched in a case-sensitive manner\nCannot be set to false if `secret` is true or undefined", "type": "boolean"}, "description": {"description": "The description of the dataset.", "type": "string", "nullable": true}, "encoding_version": {"description": "Dataset encoding version\n\nNon-secret custom word lists with no header are always version 1.\nSecret EDM lists with no header are version 1.\nMulticolumn CSV with headers are version 2.\nOmitting this field provides the default value 0, which is interpreted\nthe same as 1.", "type": "integer", "format": "int32", "minimum": 0}, "name": {"type": "string"}, "secret": {"description": "Generate a secret dataset.\n\nIf true, the response will include a secret to use with the EDM encoder.\nIf false, the response has no secret and the dataset is uploaded in plaintext.", "type": "boolean"}}, "required": ["name"]}}}}, "responses": {"200": {"description": "Dataset created successfully.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DatasetCreation"}}, "type": "object"}]}}}}, "4XX": {"description": "Dataset creation failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Datasets"], "x-api-token-group": ["Zero Trust Write"]}
```
