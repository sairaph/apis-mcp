---
title: Get instance statistics.
page_id: operation-get-accounts-account-id-ai-search-namespaces-name-instances-id-stats-ac5e410d
path: operations/ai-search-instances
description: Retrieve usage and indexing statistics for an AI Search instance.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-search/namespaces/{name}/instances/{id}/stats
operation_ids:
    - ai-search-namespace-stats
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get instance statistics.

`GET /accounts/{account_id}/ai-search/namespaces/{name}/instances/{id}/stats`

Operation ID: `ai-search-namespace-stats`

Retrieve usage and indexing statistics for an AI Search instance.

## Definition

```yaml
{"operationId": "ai-search-namespace-stats", "summary": "Get instance statistics.", "description": "Retrieve usage and indexing statistics for an AI Search instance.", "parameters": [{"name": "id", "in": "path", "description": "AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.", "required": true, "schema": {"description": "AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.", "type": "string", "example": "my-ai-search", "maxLength": 64, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$", "x-auditable": true}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "c3dc5f0b34a14ff8e1b3ec04895e1b22"}}, {"name": "name", "in": "path", "description": "Namespace name", "required": true, "schema": {"type": "string"}, "example": "my-namespace"}], "responses": {"200": {"description": "Returns the AI Search stats.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"completed": {"type": "integer"}, "degraded": {"description": "True when status counts are unavailable (e.g. legacy stats query exceeded D1 statement-size limit). Counts are omitted in this case.", "type": "boolean"}, "engine": {"description": "Engine-specific metadata. Present only for managed (v3) instances.", "type": "object", "properties": {"r2": {"description": "R2 bucket storage usage in bytes.", "type": "object", "properties": {"metadataSizeBytes": {"type": "integer"}, "objectCount": {"type": "integer"}, "payloadSizeBytes": {"type": "integer"}}, "required": ["payloadSizeBytes", "metadataSizeBytes", "objectCount"]}, "vectorize": {"description": "Vectorize index metadata (dimensions, vector count).", "type": "object", "properties": {"dimensions": {"type": "integer"}, "vectorsCount": {"type": "integer"}}, "required": ["vectorsCount", "dimensions"]}}}, "error": {"type": "integer"}, "file_embed_errors": {"type": "object", "additionalProperties": true}, "index_source_errors": {"type": "object", "additionalProperties": true}, "last_activity": {"type": "string", "format": "date-time"}, "outdated": {"type": "integer"}, "queued": {"type": "integer"}, "running": {"type": "integer"}, "skipped": {"type": "integer"}}}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "404": {"description": "Not Found", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Search Instances"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai-search"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "ai-search", "x-fern-sdk-method-name": "stats", "x-forge-params": {"id": {"description": "AI Search instance ID.", "flagName": "instance-id"}, "name": {"default": "default", "description": "Namespace to use for this operation.", "flagName": "namespace"}}}
```
