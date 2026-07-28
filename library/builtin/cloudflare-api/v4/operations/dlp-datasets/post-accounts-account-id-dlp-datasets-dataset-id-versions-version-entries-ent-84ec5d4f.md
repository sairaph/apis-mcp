---
title: Upload a new version of a multi-column dataset
page_id: operation-post-accounts-account-id-dlp-datasets-dataset-id-versions-version-entrie-6c03ee8f
path: operations/dlp-datasets
description: |-
    This is used for multi-column EDMv2 datasets. The EDMv2 format can only be
    created in the Cloudflare dashboard.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/dlp/datasets/{dataset_id}/versions/{version}/entries/{entry_id}
operation_ids:
    - dlp-datasets-upload-dataset-column
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Upload a new version of a multi-column dataset

`POST /accounts/{account_id}/dlp/datasets/{dataset_id}/versions/{version}/entries/{entry_id}`

Operation ID: `dlp-datasets-upload-dataset-column`

This is used for multi-column EDMv2 datasets. The EDMv2 format can only be
created in the Cloudflare dashboard.

## Definition

```yaml
{"operationId": "dlp-datasets-upload-dataset-column", "summary": "Upload a new version of a multi-column dataset", "description": "This is used for multi-column EDMv2 datasets. The EDMv2 format can only be\ncreated in the Cloudflare dashboard.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "dataset_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}, {"name": "version", "in": "path", "required": true, "schema": {"type": "integer", "format": "int64"}}, {"name": "entry_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"description": "Dataset content.", "required": true, "content": {"application/octet-stream": {"schema": {"type": "string", "format": "binary"}}}}, "responses": {"200": {"description": "Dataset column uploaded successfully.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DatasetColumn"}}, "type": "object"}]}}}}, "4XX": {"description": "Failed to upload dataset column.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Datasets"], "x-api-token-group": ["Zero Trust Write"]}
```
