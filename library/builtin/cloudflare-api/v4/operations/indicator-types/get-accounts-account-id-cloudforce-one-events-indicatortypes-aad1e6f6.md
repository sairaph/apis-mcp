---
title: Lists all indicator types
page_id: operation-get-accounts-account-id-cloudforce-one-events-indicatortypes-4bb052e7
path: operations/indicator-types
description: This Method is deprecated. Please use /events/dataset/:dataset_id/indicatorTypes instead.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/indicatorTypes
operation_ids:
    - get_LegacyIndicatorTypesList
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Lists all indicator types

`GET /accounts/{account_id}/cloudforce-one/events/indicatorTypes`

Operation ID: `get_LegacyIndicatorTypesList`

This Method is deprecated. Please use /events/dataset/:dataset_id/indicatorTypes instead.

## Definition

```yaml
{"operationId": "get_LegacyIndicatorTypesList", "summary": "Lists all indicator types", "description": "This Method is deprecated. Please use /events/dataset/:dataset_id/indicatorTypes instead.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}], "responses": {"200": {"description": "Returns a list of indicator types.", "content": {"application/json": {"schema": {"type": "object", "properties": {"items": {"type": "object", "properties": {"type": {"type": "string", "example": "string"}}, "required": ["type"]}, "type": {"type": "string", "example": "array"}}, "required": ["type", "items"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "deprecated": true, "security": [{"api_token": []}], "tags": ["Indicator Types"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
