---
title: List Jobs
page_id: operation-get-accounts-account-id-autorag-rags-id-jobs-f4e01ecb
path: operations/autorag-jobs
description: Lists jobs for an AutoRAG.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/autorag/rags/{id}/jobs
operation_ids:
    - autorag-config-list-jobs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Jobs

`GET /accounts/{account_id}/autorag/rags/{id}/jobs`

Operation ID: `autorag-config-list-jobs`

Lists jobs for an AutoRAG.

## Definition

```yaml
{"operationId": "autorag-config-list-jobs", "summary": "List Jobs", "description": "Lists jobs for an AutoRAG.", "parameters": [{"name": "id", "in": "path", "description": "rag id", "required": true, "schema": {"description": "rag id", "type": "string", "example": "my-rag", "maxLength": 32, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "c3dc5f0b34a14ff8e1b3ec04895e1b22"}}, {"name": "page", "in": "query", "schema": {"type": "integer", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"type": "integer", "default": 20, "maximum": 50, "minimum": 0}}], "responses": {"200": {"description": "Returns a list of AutoRAG Jobs", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "array", "items": {"properties": {"end_reason": {"type": "string"}, "ended_at": {"type": "string"}, "id": {"type": "string"}, "last_seen_at": {"type": "string"}, "source": {"type": "string", "enum": ["user", "schedule"]}, "started_at": {"type": "string"}}, "required": ["id", "source"], "type": "object"}}, "result_info": {"type": "object", "properties": {"count": {"type": "integer"}, "page": {"type": "integer"}, "per_page": {"type": "integer"}, "total_count": {"type": "integer"}}, "required": ["count", "page", "per_page", "total_count"]}, "success": {"type": "boolean"}}, "required": ["success", "result", "result_info"]}}}}, "404": {"description": "autorag_not_found", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number", "example": 7002}, "message": {"type": "string", "example": "autorag_not_found"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "503": {"description": "unable_to_connect_to_autorag", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number", "example": 7017}, "message": {"type": "string", "example": "unable_to_connect_to_autorag"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AutoRAG Jobs"], "x-api-token-group": ["Auto Rag Write", "Auto Rag Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.rag"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-ignore": true, "x-fern-sdk-group-name": "autorag.jobs", "x-fern-sdk-method-name": "list"}
```
