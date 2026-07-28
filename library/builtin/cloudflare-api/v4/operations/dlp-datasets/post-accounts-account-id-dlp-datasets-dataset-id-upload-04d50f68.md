---
title: Prepare to upload a new version of a dataset
page_id: operation-post-accounts-account-id-dlp-datasets-dataset-id-upload-b9f1e322
path: operations/dlp-datasets
description: Creates a new version of a DLP dataset, allowing you to stage changes before activation. Used for single-column EDM and custom word lists.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/dlp/datasets/{dataset_id}/upload
operation_ids:
    - dlp-datasets-create-version
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Prepare to upload a new version of a dataset

`POST /accounts/{account_id}/dlp/datasets/{dataset_id}/upload`

Operation ID: `dlp-datasets-create-version`

Creates a new version of a DLP dataset, allowing you to stage changes before activation. Used for single-column EDM and custom word lists.

## Definition

```yaml
{"operationId": "dlp-datasets-create-version", "summary": "Prepare to upload a new version of a dataset", "description": "Creates a new version of a DLP dataset, allowing you to stage changes before activation. Used for single-column EDM and custom word lists.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "dataset_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Dataset version created successfully.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DatasetNewVersion"}}, "type": "object"}]}}}}, "4XX": {"description": "Dataset version creation failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Datasets"], "x-api-token-group": ["Zero Trust Write"]}
```
