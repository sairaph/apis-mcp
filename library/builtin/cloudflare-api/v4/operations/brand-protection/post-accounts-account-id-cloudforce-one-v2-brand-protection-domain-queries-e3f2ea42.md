---
title: Insert query
page_id: operation-post-accounts-account-id-cloudforce-one-v2-brand-protection-domain-queri-a6c4e976
path: operations/brand-protection
description: Create a new saved brand protection query with string match patterns
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/domain/queries
operation_ids:
    - post_InsertDomainQuery
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Insert query

`POST /accounts/{account_id}/cloudforce-one/v2/brand-protection/domain/queries`

Operation ID: `post_InsertDomainQuery`

Create a new saved brand protection query with string match patterns

## Definition

```yaml
{"operationId": "post_InsertDomainQuery", "summary": "Insert query", "description": "Create a new saved brand protection query with string match patterns", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"max_time": {"type": "string"}, "min_time": {"type": "string"}, "scan": {"type": "boolean"}, "scan_matches": {"type": "boolean"}, "search_lookback": {"description": "If true, search recent domain lookback entries for matches", "type": "boolean", "default": true}, "string_matches": {"type": "array", "items": {"properties": {"pattern": {"type": "string", "maxLength": 200, "minLength": 1}}, "required": ["pattern"], "type": "object"}, "maxItems": 100, "minItems": 1}, "tag": {"type": "string", "minLength": 1}}, "required": ["tag", "string_matches"]}}}}, "responses": {"200": {"description": "Query inserted successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"message": {"type": "string"}, "query_id": {"type": "integer"}, "success": {"type": "boolean"}}, "required": ["success", "message"]}}}}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write"]}
```
