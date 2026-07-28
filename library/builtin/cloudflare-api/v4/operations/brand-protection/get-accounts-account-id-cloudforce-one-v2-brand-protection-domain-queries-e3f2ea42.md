---
title: Get queries
page_id: operation-get-accounts-account-id-cloudforce-one-v2-brand-protection-domain-querie-1e6e36e8
path: operations/brand-protection
description: Get all saved brand protection queries for an account
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/domain/queries
operation_ids:
    - get_GetDomainQueries
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get queries

`GET /accounts/{account_id}/cloudforce-one/v2/brand-protection/domain/queries`

Operation ID: `get_GetDomainQueries`

Get all saved brand protection queries for an account

## Definition

```yaml
{"operationId": "get_GetDomainQueries", "summary": "Get queries", "description": "Get all saved brand protection queries for an account", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}, {"name": "id", "in": "query", "schema": {"type": "string"}}, {"name": "page", "in": "query", "schema": {"description": "Optional page number for paginated list requests. Defaults to 1 when only per_page is supplied. Omit page and per_page to preserve the legacy full-list response.", "type": "integer", "exclusiveMinimum": true, "maximum": 1000000, "minimum": 0}}, {"name": "per_page", "in": "query", "schema": {"description": "Optional number of queries per page for paginated list requests. Defaults to 100 when only page is supplied. Maximum 100. Omit page and per_page to preserve the legacy full-list response.", "type": "integer", "exclusiveMinimum": true, "maximum": 100, "minimum": 0}}], "responses": {"200": {"description": "Successfully retrieved queries", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"anyOf": [{"type": "string"}, {"type": "number"}]}, "message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"code": {"anyOf": [{"type": "string"}, {"type": "number"}]}, "message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"anyOf": [{"items": {"properties": {"created": {"type": "string"}, "parameters": {"required": ["string_matches"], "nullable": true, "properties": {"max_time": {"type": "string"}, "min_time": {"type": "string"}, "string_matches": {"type": "array", "items": {"properties": {"pattern": {"type": "string", "maxLength": 200, "minLength": 1}}, "required": ["pattern"], "type": "object"}}}, "type": "object"}, "query_id": {"type": "integer"}, "query_tag": {"type": "string"}, "scan": {"type": "boolean"}, "updated": {"type": "string"}}, "required": ["query_id", "query_tag", "scan", "parameters", "created", "updated"], "type": "object"}, "type": "array"}, {"properties": {"created": {"type": "string"}, "parameters": {"required": ["string_matches"], "nullable": true, "properties": {"max_time": {"type": "string"}, "min_time": {"type": "string"}, "string_matches": {"type": "array", "items": {"properties": {"pattern": {"type": "string", "maxLength": 200, "minLength": 1}}, "required": ["pattern"], "type": "object"}}}, "type": "object"}, "query_id": {"type": "integer"}, "query_tag": {"type": "string"}, "scan": {"type": "boolean"}, "updated": {"type": "string"}}, "required": ["query_id", "query_tag", "scan", "parameters", "created", "updated"], "type": "object"}]}, "result_info": {"description": "Present on paginated list responses when page or per_page is supplied.", "type": "object", "properties": {"count": {"type": "integer", "minimum": 0}, "page": {"type": "integer", "exclusiveMinimum": true, "minimum": 0}, "per_page": {"type": "integer", "exclusiveMinimum": true, "minimum": 0}, "total_count": {"type": "integer", "minimum": 0}}, "required": ["count", "page", "per_page", "total_count"]}, "success": {"type": "boolean"}}, "required": ["success", "result", "errors", "messages"]}}}}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
