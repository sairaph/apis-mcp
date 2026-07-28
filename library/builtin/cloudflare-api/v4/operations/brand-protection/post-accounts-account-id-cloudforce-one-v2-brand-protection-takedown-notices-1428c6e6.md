---
title: Lookup takedown notices by domains
page_id: operation-post-accounts-account-id-cloudforce-one-v2-brand-protection-takedown-not-0c05be26
path: operations/brand-protection
description: Bulk lookup of takedown notices by domain names. Returns an array of { domain, queryId, notice } entries (notice is null when no match exists). Supports optional queryId or queryIds to scope lookups per query. Uses POST to avoid URL length limits when looking up many domains. This endpoint is read-only. Domains are normalized (trimmed, lowercased) and deduplicated.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/takedown-notices/lookup
operation_ids:
    - post_TakedownNoticeLookup
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Lookup takedown notices by domains

`POST /accounts/{account_id}/cloudforce-one/v2/brand-protection/takedown-notices/lookup`

Operation ID: `post_TakedownNoticeLookup`

Bulk lookup of takedown notices by domain names. Returns an array of { domain, queryId, notice } entries (notice is null when no match exists). Supports optional queryId or queryIds to scope lookups per query. Uses POST to avoid URL length limits when looking up many domains. This endpoint is read-only. Domains are normalized (trimmed, lowercased) and deduplicated.

## Definition

```yaml
{"operationId": "post_TakedownNoticeLookup", "summary": "Lookup takedown notices by domains", "description": "Bulk lookup of takedown notices by domain names. Returns an array of { domain, queryId, notice } entries (notice is null when no match exists). Supports optional queryId or queryIds to scope lookups per query. Uses POST to avoid URL length limits when looking up many domains. This endpoint is read-only. Domains are normalized (trimmed, lowercased) and deduplicated.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"domains": {"type": "array", "items": {"maxLength": 253, "minLength": 1, "type": "string"}, "maxItems": 100, "minItems": 1}, "queryId": {"type": "integer", "exclusiveMinimum": true, "minimum": 0}, "queryIds": {"type": "array", "items": {"exclusiveMinimum": true, "minimum": 0, "type": "integer"}, "maxItems": 100}}, "required": ["domains"]}}}}, "responses": {"200": {"description": "Takedown notice lookup results", "content": {"application/json": {"schema": {"type": "object", "properties": {"notices": {"type": "array", "items": {"properties": {"domain": {"type": "string"}, "notice": {"type": "object", "nullable": true, "properties": {"createdAt": {"type": "string", "nullable": true}, "domain": {"type": "string"}, "id": {"type": "number"}, "matchId": {"type": "number", "nullable": true}, "matchType": {"type": "string", "enum": ["logo", "domain"], "nullable": true}, "queryId": {"type": "number", "nullable": true}, "status": {"type": "string", "enum": ["draft", "sent", "resolved", "expired"]}, "updatedAt": {"type": "string", "nullable": true}}, "required": ["id", "domain", "queryId", "status", "matchId", "matchType", "createdAt", "updatedAt"]}, "queryId": {"type": "number", "nullable": true}}, "required": ["domain", "queryId", "notice"], "type": "object"}}}, "required": ["notices"]}}}}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write"]}
```
