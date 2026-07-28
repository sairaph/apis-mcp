---
title: Update details about a dataset
page_id: operation-put-accounts-account-id-dlp-datasets-dataset-id-7f355393
path: operations/dlp-datasets
description: Updates the configuration of an existing DLP dataset, such as its name, description, or detection settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/dlp/datasets/{dataset_id}
operation_ids:
    - dlp-datasets-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update details about a dataset

`PUT /accounts/{account_id}/dlp/datasets/{dataset_id}`

Operation ID: `dlp-datasets-update`

Updates the configuration of an existing DLP dataset, such as its name, description, or detection settings.

## Definition

```yaml
{"operationId": "dlp-datasets-update", "summary": "Update details about a dataset", "description": "Updates the configuration of an existing DLP dataset, such as its name, description, or detection settings.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "dataset_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"description": "Dataset description.", "required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"case_sensitive": {"description": "Determines if the words should be matched in a case-sensitive manner.\n\nOnly required for custom word lists.", "type": "boolean"}, "description": {"description": "The description of the dataset.", "type": "string", "nullable": true}, "name": {"description": "The name of the dataset, must be unique.", "type": "string", "nullable": true}}}}}}, "responses": {"200": {"description": "Dataset updated successfully.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_Dataset"}}, "type": "object"}]}}}}, "4XX": {"description": "Dataset update failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Datasets"], "x-api-token-group": ["Zero Trust Write"]}
```
