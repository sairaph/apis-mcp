---
title: Sync
page_id: operation-patch-accounts-account-id-autorag-rags-id-sync-c9c55796
path: operations/autorag-rag
description: Starts synchronization for an AutoRAG.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/autorag/rags/{id}/sync
operation_ids:
    - autorag-config-sync
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Sync

`PATCH /accounts/{account_id}/autorag/rags/{id}/sync`

Operation ID: `autorag-config-sync`

Starts synchronization for an AutoRAG.

## Definition

```yaml
{"operationId": "autorag-config-sync", "summary": "Sync", "description": "Starts synchronization for an AutoRAG.", "parameters": [{"name": "id", "in": "path", "description": "rag id", "required": true, "schema": {"description": "rag id", "type": "string", "example": "my-rag", "maxLength": 32, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "c3dc5f0b34a14ff8e1b3ec04895e1b22"}}], "responses": {"200": {"description": "Returns the autorag sync status", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"job_id": {"type": "string"}}, "required": ["job_id"]}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "400": {"description": "autorag_is_paused", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number", "example": 7032}, "message": {"type": "string", "example": "autorag_is_paused"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "404": {"description": "autorag_not_found", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number", "example": 7002}, "message": {"type": "string", "example": "autorag_not_found"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "429": {"description": "sync_in_cooldown", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number", "example": 7020}, "message": {"type": "string", "example": "sync_in_cooldown"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "503": {"description": "unable_to_connect_to_autorag", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number", "example": 7017}, "message": {"type": "string", "example": "unable_to_connect_to_autorag"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AutoRAG RAG"], "x-api-token-group": ["Auto Rag Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.rag"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-ignore": true, "x-fern-sdk-group-name": "autorag", "x-fern-sdk-method-name": "sync"}
```
