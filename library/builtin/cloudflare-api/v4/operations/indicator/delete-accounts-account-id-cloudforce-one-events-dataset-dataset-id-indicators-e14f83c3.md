---
title: Deletes an indicator
page_id: operation-delete-accounts-account-id-cloudforce-one-events-dataset-dataset-id-indi-169921c2
path: operations/indicator
description: Deletes a specific indicator by its UUID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}/indicators/{indicator_id}
operation_ids:
    - delete_IndicatorDelete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Deletes an indicator

`DELETE /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}/indicators/{indicator_id}`

Operation ID: `delete_IndicatorDelete`

Deletes a specific indicator by its UUID.

## Definition

```yaml
{"operationId": "delete_IndicatorDelete", "summary": "Deletes an indicator", "description": "Deletes a specific indicator by its UUID.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "dataset_id", "in": "path", "description": "Dataset ID.", "required": true, "schema": {"description": "Dataset ID.", "type": "string"}}, {"name": "indicator_id", "in": "path", "description": "Indicator UUID.", "required": true, "schema": {"description": "Indicator UUID.", "type": "string"}}], "responses": {"200": {"description": "Indicator deleted successfully.", "content": {"application/json": {"schema": {"type": "object", "properties": {"message": {"type": "string"}, "success": {"type": "boolean"}}}}}}, "404": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Indicator"], "x-api-token-group": ["Cloudforce One Write"]}
```
