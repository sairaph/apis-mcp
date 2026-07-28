---
title: Lists target industries across multiple datasets
page_id: operation-get-accounts-account-id-cloudforce-one-events-targetindustries-64fb523d
path: operations/target-industry
description: List target industries referenced in events across one or more datasets.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/targetIndustries
operation_ids:
    - get_TargetIndustryList
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Lists target industries across multiple datasets

`GET /accounts/{account_id}/cloudforce-one/events/targetIndustries`

Operation ID: `get_TargetIndustryList`

List target industries referenced in events across one or more datasets.

## Definition

```yaml
{"operationId": "get_TargetIndustryList", "summary": "Lists target industries across multiple datasets", "description": "List target industries referenced in events across one or more datasets.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "datasetIds", "in": "query", "description": "Array of dataset IDs to query target industries from. If not provided, uses the default dataset.", "schema": {"description": "Array of dataset IDs to query target industries from. If not provided, uses the default dataset.", "type": "array", "items": {"type": "string"}}}], "responses": {"200": {"description": "Returns a list of target industries.", "content": {"application/json": {"schema": {"type": "object", "properties": {"items": {"type": "object", "properties": {"type": {"type": "string", "example": "string"}}, "required": ["type"]}, "type": {"type": "string", "example": "array"}}, "required": ["type", "items"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Target Industry"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
