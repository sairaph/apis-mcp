---
title: Search
page_id: operation-post-accounts-account-id-autorag-rags-id-search-2d47c857
path: operations/autorag-rag-search
description: Searches an AutoRAG.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/autorag/rags/{id}/search
operation_ids:
    - autorag-config-search
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Search

`POST /accounts/{account_id}/autorag/rags/{id}/search`

Operation ID: `autorag-config-search`

Searches an AutoRAG.

## Definition

```yaml
{"operationId": "autorag-config-search", "summary": "Search", "description": "Searches an AutoRAG.", "parameters": [{"name": "id", "in": "path", "description": "rag id", "required": true, "schema": {"description": "rag id", "type": "string", "example": "my-rag", "maxLength": 32, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "c3dc5f0b34a14ff8e1b3ec04895e1b22"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"filters": {"anyOf": [{"properties": {"key": {"type": "string"}, "type": {"type": "string", "enum": ["eq", "ne", "gt", "gte", "lt", "lte"]}, "value": {"anyOf": [{"type": "string"}, {"type": "number"}, {"type": "boolean"}]}}, "required": ["key", "type", "value"], "type": "object"}, {"properties": {"filters": {"type": "array", "items": {"properties": {"key": {"type": "string"}, "type": {"type": "string", "enum": ["eq", "ne", "gt", "gte", "lt", "lte"]}, "value": {"anyOf": [{"type": "string"}, {"type": "number"}, {"type": "boolean"}]}}, "required": ["key", "type", "value"], "type": "object"}}, "type": {"type": "string", "enum": ["and", "or"]}}, "required": ["type", "filters"], "type": "object"}]}, "max_num_results": {"type": "integer", "default": 10, "maximum": 50, "minimum": 1}, "query": {"type": "string"}, "ranking_options": {"type": "object", "default": {}, "properties": {"ranker": {"type": "string"}, "score_threshold": {"type": "number", "default": 0.4, "maximum": 1, "minimum": 0}}}, "reranking": {"type": "object", "properties": {"enabled": {"type": "boolean", "default": false}, "model": {"anyOf": [{"enum": ["@cf/baai/bge-reranker-base"], "type": "string"}, {"enum": [""], "type": "string"}]}}}, "rewrite_query": {"type": "boolean", "default": false}}, "required": ["query"]}}}}, "responses": {"200": {"description": "Returns the log details", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"data": {"type": "array", "items": {"properties": {"attributes": {"type": "object"}, "content": {"items": {"properties": {"text": {"type": "string"}, "type": {"type": "string"}}, "type": "object"}, "type": "array"}, "file_id": {"type": "string"}, "filename": {"type": "string"}, "score": {"type": "number"}}, "required": ["score"], "type": "object"}}, "has_more": {"type": "boolean", "default": false}, "next_page": {"type": "string", "nullable": true}, "object": {"type": "string"}, "search_query": {"type": "string"}}, "required": ["search_query"]}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "404": {"description": "Not Found", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number", "example": 7002}, "message": {"type": "string", "example": "Not Found"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AutoRAG RAG Search"], "x-api-token-group": ["Auto Rag Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.rag"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-ignore": true, "x-fern-sdk-group-name": "autorag", "x-fern-sdk-method-name": "search"}
```
