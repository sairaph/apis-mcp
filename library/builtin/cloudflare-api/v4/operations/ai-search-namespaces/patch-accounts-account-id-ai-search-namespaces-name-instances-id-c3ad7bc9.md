---
title: Move an instance to a different namespace.
page_id: operation-patch-accounts-account-id-ai-search-namespaces-name-instances-id-f1127014
path: operations/ai-search-namespaces
description: Moves an instance from its current namespace to the specified target namespace. Use 'default' with --destination-namespace to move the instance back to the default namespace. Fails with 400 if the target namespace already has an instance with the same id (ids must be unique within a namespace — the same id can exist in different namespaces).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/ai-search/namespaces/{name}/instances/{id}
operation_ids:
    - ai-search-move-instance
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Move an instance to a different namespace.

`PATCH /accounts/{account_id}/ai-search/namespaces/{name}/instances/{id}`

Operation ID: `ai-search-move-instance`

Moves an instance from its current namespace to the specified target namespace. Use 'default' with --destination-namespace to move the instance back to the default namespace. Fails with 400 if the target namespace already has an instance with the same id (ids must be unique within a namespace — the same id can exist in different namespaces).

## Definition

```yaml
{"operationId": "ai-search-move-instance", "summary": "Move an instance to a different namespace.", "description": "Moves an instance from its current namespace to the specified target namespace. Use 'default' with --destination-namespace to move the instance back to the default namespace. Fails with 400 if the target namespace already has an instance with the same id (ids must be unique within a namespace — the same id can exist in different namespaces).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "c3dc5f0b34a14ff8e1b3ec04895e1b22"}}, {"name": "name", "in": "path", "description": "Current namespace of the instance.", "required": true, "schema": {"description": "Current namespace of the instance.", "type": "string", "example": "production"}}, {"name": "id", "in": "path", "description": "Instance id.", "required": true, "schema": {"description": "Instance id.", "type": "string", "example": "my-blog"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"new_namespace": {"description": "Target namespace to move the instance into.", "type": "string", "example": "staging", "pattern": "^[a-z0-9]([a-z0-9-]{0,26}[a-z0-9])?$"}}, "required": ["new_namespace"]}}}}, "responses": {"200": {"description": "Instance moved.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "additionalProperties": false}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "result"]}}}}, "400": {"description": "Ai search with this name already exist.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "404": {"description": "Ai search not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Search Namespaces"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai-search"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ai-search", "x-fern-sdk-method-name": "move", "x-forge-params": {"id": {"description": "AI Search instance ID.", "flagName": "instance-id"}, "name": {"default": "default", "description": "Namespace currently containing the instance.", "flagName": "source-namespace"}, "new_namespace": {"description": "Namespace to move the instance into.", "flagName": "destination-namespace"}}}
```
