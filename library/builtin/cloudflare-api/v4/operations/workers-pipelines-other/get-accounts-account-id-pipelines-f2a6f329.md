---
title: '[DEPRECATED] List Pipelines'
page_id: operation-get-accounts-account-id-pipelines-fe8a58dc
path: operations/workers-pipelines-other
description: '[DEPRECATED] List, filter, and paginate pipelines in an account. Use the new /pipelines/v1/pipelines endpoint instead.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/pipelines
operation_ids:
    - getV4AccountsByAccount_idPipelines_deprecated
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# [DEPRECATED] List Pipelines

`GET /accounts/{account_id}/pipelines`

Operation ID: `getV4AccountsByAccount_idPipelines_deprecated`

[DEPRECATED] List, filter, and paginate pipelines in an account. Use the new /pipelines/v1/pipelines endpoint instead.

## Definition

```yaml
{"operationId": "getV4AccountsByAccount_idPipelines_deprecated", "summary": "[DEPRECATED] List Pipelines", "description": "[DEPRECATED] List, filter, and paginate pipelines in an account. Use the new /pipelines/v1/pipelines endpoint instead.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_workers-pipelines-account-id"}}, {"name": "search", "in": "query", "schema": {"description": "Specifies the prefix of pipeline name to search.", "type": "string"}}, {"name": "page", "in": "query", "schema": {"description": "Specifies which page to retrieve.", "type": "string", "default": "1"}}, {"name": "per_page", "in": "query", "schema": {"description": "Specifies the number of pipelines per page.", "type": "string", "default": "25"}}], "responses": {"200": {"description": "[DEPRECATED] Lists the pipelines. Use /pipelines/v1/pipelines instead.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result_info": {"type": "object", "properties": {"count": {"description": "Indicates the number of items on current page.", "type": "number", "example": 1}, "page": {"description": "Indicates the current page number.", "type": "number", "example": 0}, "per_page": {"description": "Indicates the number of items per page.", "type": "number", "example": 10}, "total_count": {"description": "Indicates the total number of items.", "type": "number", "example": 1}}, "required": ["count", "page", "per_page", "total_count"]}, "results": {"type": "array", "items": {"$ref": "#/components/schemas/cloudflare-pipelines_workers-pipelines-pipeline"}}, "success": {"$ref": "#/components/schemas/cloudflare-pipelines_worker-pipelines-common-success"}}, "required": ["success", "results", "result_info"]}}}}, "4XX": {"description": "Indicates the error trying to list pipelines.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}, "example": []}, "results": {"type": "object", "nullable": true, "x-stainless-empty-object": true}, "success": {"example": false, "allOf": [{"$ref": "#/components/schemas/cloudflare-pipelines_worker-pipelines-common-success"}]}}, "required": ["success", "results", "errors"]}}}}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["workers_pipelines_other"], "x-api-token-group": ["Pipelines Write", "Pipelines Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pipelines", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
