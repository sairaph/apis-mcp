---
title: Sets the column information for a multi-column upload
page_id: operation-post-accounts-account-id-dlp-datasets-dataset-id-versions-version-990e7d7b
path: operations/dlp-datasets
description: |-
    This is used for multi-column EDMv2 datasets. The EDMv2 format can only be
    created in the Cloudflare dashboard. The columns in the response appear in
    the same order as in the request.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/dlp/datasets/{dataset_id}/versions/{version}
operation_ids:
    - dlp-datasets-define-columns
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Sets the column information for a multi-column upload

`POST /accounts/{account_id}/dlp/datasets/{dataset_id}/versions/{version}`

Operation ID: `dlp-datasets-define-columns`

This is used for multi-column EDMv2 datasets. The EDMv2 format can only be
created in the Cloudflare dashboard. The columns in the response appear in
the same order as in the request.

## Definition

```yaml
{"operationId": "dlp-datasets-define-columns", "summary": "Sets the column information for a multi-column upload", "description": "This is used for multi-column EDMv2 datasets. The EDMv2 format can only be\ncreated in the Cloudflare dashboard. The columns in the response appear in\nthe same order as in the request.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "dataset_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}, {"name": "version", "in": "path", "required": true, "schema": {"type": "integer", "format": "int64"}}], "requestBody": {"description": "array of new columns to create for this dataset version.", "required": true, "content": {"application/json": {"schema": {"type": "array", "items": {"allOf": [{"oneOf": [{"properties": {"entry_id": {"type": "string", "format": "uuid"}}, "required": ["entry_id"], "title": "Existing Column", "type": "object"}, {"properties": {"entry_name": {"type": "string"}}, "required": ["entry_name"], "title": "New Column", "type": "object"}]}, {"properties": {"header_name": {"type": "string"}, "num_cells": {"type": "integer", "format": "int64", "minimum": 0}}, "required": ["header_name", "num_cells"], "type": "object"}]}}}}}, "responses": {"200": {"description": "Dataset columns created successfully.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DatasetColumnArray"}}, "type": "object"}]}}}}, "4XX": {"description": "Failed to create dataset columns.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Datasets"], "x-api-token-group": ["Zero Trust Write"]}
```
