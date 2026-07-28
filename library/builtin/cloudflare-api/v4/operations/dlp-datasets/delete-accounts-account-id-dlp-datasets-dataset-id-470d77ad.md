---
title: Delete a dataset
page_id: operation-delete-accounts-account-id-dlp-datasets-dataset-id-2c0f0019
path: operations/dlp-datasets
description: This deletes all versions of the dataset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/dlp/datasets/{dataset_id}
operation_ids:
    - dlp-datasets-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a dataset

`DELETE /accounts/{account_id}/dlp/datasets/{dataset_id}`

Operation ID: `dlp-datasets-delete`

This deletes all versions of the dataset.

## Definition

```yaml
{"operationId": "dlp-datasets-delete", "summary": "Delete a dataset", "description": "This deletes all versions of the dataset.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "dataset_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Dataset deleted successfully."}, "4XX": {"description": "Dataset delete failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Datasets"], "x-api-token-group": ["Zero Trust Write"]}
```
