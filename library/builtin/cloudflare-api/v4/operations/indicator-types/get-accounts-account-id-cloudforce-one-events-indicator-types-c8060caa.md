---
title: Lists indicator types across multiple datasets
page_id: operation-get-accounts-account-id-cloudforce-one-events-indicator-types-47fa2459
path: operations/indicator-types
description: List indicator types across one or more datasets for the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/indicator-types
operation_ids:
    - get_IndicatorTypesList
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Lists indicator types across multiple datasets

`GET /accounts/{account_id}/cloudforce-one/events/indicator-types`

Operation ID: `get_IndicatorTypesList`

List indicator types across one or more datasets for the account.

## Definition

```yaml
{"operationId": "get_IndicatorTypesList", "summary": "Lists indicator types across multiple datasets", "description": "List indicator types across one or more datasets for the account.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "datasetIds", "in": "query", "description": "Array of dataset IDs to query indicator types from. If not provided, queries all datasets for the account.", "schema": {"description": "Array of dataset IDs to query indicator types from. If not provided, queries all datasets for the account.", "type": "array", "items": {"type": "string"}}}], "responses": {"200": {"description": "Returns a list of indicator types.", "content": {"application/json": {"schema": {"type": "object", "properties": {"items": {"type": "object", "properties": {"type": {"type": "string", "example": "string"}}, "required": ["type"]}, "type": {"type": "string", "example": "array"}}, "required": ["type", "items"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Indicator Types"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
