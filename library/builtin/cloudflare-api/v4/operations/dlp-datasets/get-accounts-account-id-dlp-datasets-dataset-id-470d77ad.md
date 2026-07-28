---
title: Fetch a specific dataset
page_id: operation-get-accounts-account-id-dlp-datasets-dataset-id-78777e71
path: operations/dlp-datasets
description: Gets a dataset and its latest upload status.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/datasets/{dataset_id}
operation_ids:
    - dlp-datasets-read
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch a specific dataset

`GET /accounts/{account_id}/dlp/datasets/{dataset_id}`

Operation ID: `dlp-datasets-read`

Gets a dataset and its latest upload status.

## Definition

```yaml
{"operationId": "dlp-datasets-read", "summary": "Fetch a specific dataset", "description": "Gets a dataset and its latest upload status.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "dataset_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Dataset read successfully.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_Dataset"}}, "type": "object"}]}}}}, "4XX": {"description": "Dataset read failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Datasets"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
