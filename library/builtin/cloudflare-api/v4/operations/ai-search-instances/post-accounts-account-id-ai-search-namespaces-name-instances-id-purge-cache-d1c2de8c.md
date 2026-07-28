---
title: Purge search cache.
page_id: operation-post-accounts-account-id-ai-search-namespaces-name-instances-id-purge-ca-46be7c4d
path: operations/ai-search-instances
description: Purges all cached search results for an AI Search instance. A new internal cache key is generated, immediately orphaning all prior cached entries.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/ai-search/namespaces/{name}/instances/{id}/purge_cache
operation_ids:
    - ai-search-namespace-purge-instance-cache
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Purge search cache.

`POST /accounts/{account_id}/ai-search/namespaces/{name}/instances/{id}/purge_cache`

Operation ID: `ai-search-namespace-purge-instance-cache`

Purges all cached search results for an AI Search instance. A new internal cache key is generated, immediately orphaning all prior cached entries.

## Definition

```yaml
{"operationId": "ai-search-namespace-purge-instance-cache", "summary": "Purge search cache.", "description": "Purges all cached search results for an AI Search instance. A new internal cache key is generated, immediately orphaning all prior cached entries.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "id", "in": "path", "description": "AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.", "required": true, "schema": {"description": "AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.", "type": "string", "example": "my-ai-search", "maxLength": 64, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$", "x-auditable": true}}, {"name": "name", "in": "path", "description": "Namespace name", "required": true, "schema": {"type": "string"}, "example": "my-namespace"}], "responses": {"200": {"description": "Cache purged successfully.", "content": {"application/json": {"schema": {"type": "object", "properties": {"success": {"type": "boolean"}}, "required": ["success"]}}}}, "404": {"description": "Ai search not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Search Instances"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai-search"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "ai-search", "x-fern-sdk-method-name": "purge-cache", "x-forge-params": {"id": {"description": "AI Search instance ID.", "flagName": "instance-id"}, "name": {"default": "default", "description": "Namespace to use for this operation.", "flagName": "namespace"}}}
```
