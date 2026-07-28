---
title: List logo matches
page_id: operation-get-accounts-account-id-cloudforce-one-v2-brand-protection-logo-matches-ecb26c4b
path: operations/brand-protection
description: Get paginated list of logo matches for a specific brand protection logo query
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/logo/matches
operation_ids:
    - get_LogoMatchList
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List logo matches

`GET /accounts/{account_id}/cloudforce-one/v2/brand-protection/logo/matches`

Operation ID: `get_LogoMatchList`

Get paginated list of logo matches for a specific brand protection logo query

## Definition

```yaml
{"operationId": "get_LogoMatchList", "summary": "List logo matches", "description": "Get paginated list of logo matches for a specific brand protection logo query", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}, {"name": "offset", "in": "query", "schema": {"type": "string", "default": "0"}}, {"name": "limit", "in": "query", "schema": {"type": "string", "default": "50"}}, {"name": "query_id", "in": "query", "required": true, "schema": {"type": "string", "minLength": 1}}, {"name": "download", "in": "query", "schema": {"type": "string"}}, {"name": "orderBy", "in": "query", "schema": {"description": "Column to sort by. Options: 'matchedAt', 'domain', 'similarityScore', or 'registrar'", "type": "string", "enum": ["matchedAt", "domain", "similarityScore", "registrar"]}}, {"name": "order", "in": "query", "schema": {"description": "Sort order. Options: 'asc' (ascending) or 'desc' (descending)", "type": "string", "enum": ["asc", "desc"]}}], "responses": {"200": {"description": "Successfully retrieved logo matches", "content": {"application/json": {"schema": {"type": "object", "properties": {"matches": {"type": "array", "items": {"properties": {"content_type": {"type": "string"}, "domain": {"type": "string", "nullable": true}, "id": {"type": "integer"}, "image_data": {"type": "string"}, "matched_at": {"type": "string", "nullable": true}, "query_id": {"type": "integer"}, "registrar": {"type": "string", "nullable": true}, "similarity_score": {"type": "number"}, "url_scan_id": {"type": "string", "nullable": true}}, "required": ["id", "query_id", "url_scan_id", "similarity_score", "matched_at", "domain", "registrar"], "type": "object"}}, "total": {"type": "integer", "minimum": 0}}, "required": ["matches", "total"]}}}}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
