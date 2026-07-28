---
title: Lists all target industries from industry map catalog
page_id: operation-get-accounts-account-id-cloudforce-one-events-targetindustries-catalog-59da4581
path: operations/target-industry
description: List all predefined target industries from the industry map catalog.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/targetIndustries/catalog
operation_ids:
    - get_TargetIndustryListComplete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Lists all target industries from industry map catalog

`GET /accounts/{account_id}/cloudforce-one/events/targetIndustries/catalog`

Operation ID: `get_TargetIndustryListComplete`

List all predefined target industries from the industry map catalog.

## Definition

```yaml
{"operationId": "get_TargetIndustryListComplete", "summary": "Lists all target industries from industry map catalog", "description": "List all predefined target industries from the industry map catalog.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}], "responses": {"200": {"description": "Returns all target industries from industry map catalog.", "content": {"application/json": {"schema": {"type": "object", "properties": {"items": {"type": "object", "properties": {"type": {"type": "string", "example": "string"}}, "required": ["type"]}, "type": {"type": "string", "example": "array"}}, "required": ["type", "items"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Target Industry"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
