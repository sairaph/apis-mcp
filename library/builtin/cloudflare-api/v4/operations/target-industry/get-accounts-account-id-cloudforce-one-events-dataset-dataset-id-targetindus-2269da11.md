---
title: Lists all target industries for a specific dataset
page_id: operation-get-accounts-account-id-cloudforce-one-events-dataset-dataset-id-targeti-353363f7
path: operations/target-industry
description: List all target industries referenced in events for a specific dataset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}/targetIndustries
operation_ids:
    - get_TargetIndustryListByDataset
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Lists all target industries for a specific dataset

`GET /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}/targetIndustries`

Operation ID: `get_TargetIndustryListByDataset`

List all target industries referenced in events for a specific dataset.

## Definition

```yaml
{"operationId": "get_TargetIndustryListByDataset", "summary": "Lists all target industries for a specific dataset", "description": "List all target industries referenced in events for a specific dataset.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "dataset_id", "in": "path", "description": "Dataset UUID.", "required": true, "schema": {"description": "Dataset UUID.", "type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Returns a list of target industries for the dataset.", "content": {"application/json": {"schema": {"type": "object", "properties": {"items": {"type": "object", "properties": {"type": {"type": "string", "example": "string"}}, "required": ["type"]}, "type": {"type": "string", "example": "array"}}, "required": ["type", "items"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Target Industry"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
