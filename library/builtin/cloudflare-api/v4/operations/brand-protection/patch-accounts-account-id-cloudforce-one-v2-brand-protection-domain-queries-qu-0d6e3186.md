---
title: Update domain query
page_id: operation-patch-accounts-account-id-cloudforce-one-v2-brand-protection-domain-quer-8c4015b9
path: operations/brand-protection
description: Update a saved brand protection domain query with string match patterns
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/domain/queries/{query_id}
operation_ids:
    - patch_UpdateDomainQuery
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update domain query

`PATCH /accounts/{account_id}/cloudforce-one/v2/brand-protection/domain/queries/{query_id}`

Operation ID: `patch_UpdateDomainQuery`

Update a saved brand protection domain query with string match patterns

## Definition

```yaml
{"operationId": "patch_UpdateDomainQuery", "summary": "Update domain query", "description": "Update a saved brand protection domain query with string match patterns", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}, {"name": "query_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"max_time": {"type": "string"}, "min_time": {"type": "string"}, "scan": {"type": "boolean"}, "scan_matches": {"type": "boolean"}, "search_lookback": {"description": "If true, search recent domain lookback entries for matches", "type": "boolean", "default": true}, "string_matches": {"type": "array", "items": {"properties": {"pattern": {"type": "string", "maxLength": 200, "minLength": 1}}, "required": ["pattern"], "type": "object"}, "maxItems": 100, "minItems": 1}, "tag": {"type": "string", "minLength": 1}}, "required": ["tag", "string_matches"]}}}}, "responses": {"200": {"description": "Domain query updated successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"message": {"type": "string"}, "query_id": {"type": "integer"}, "success": {"type": "boolean"}}, "required": ["success", "message"]}}}}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write"]}
```
