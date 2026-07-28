---
title: Search rules
page_id: operation-get-accounts-account-id-cloudforce-one-rules-search-cc146cdd
path: operations/rules
description: Search rules using hybrid, vector, keyword, or exact retrieval, backed by AI Search with a SQL fallback.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/rules/search
operation_ids:
    - cloudforce-one-search-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Search rules

`GET /accounts/{account_id}/cloudforce-one/rules/search`

Operation ID: `cloudforce-one-search-rules`

Search rules using hybrid, vector, keyword, or exact retrieval, backed by AI Search with a SQL fallback.

## Definition

```yaml
{"operationId": "cloudforce-one-search-rules", "summary": "Search rules", "description": "Search rules using hybrid, vector, keyword, or exact retrieval, backed by AI Search with a SQL fallback.", "parameters": [{"$ref": "#/components/parameters/cloudforce-one_account_id"}, {"name": "namespace", "in": "query", "description": "Filter by namespace. Repeat the parameter to filter by multiple namespaces (e.g. namespace=foo&namespace=bar).", "schema": {"description": "Filter by namespace. Repeat the parameter to filter by multiple namespaces (e.g. namespace=foo&namespace=bar).", "example": ["yara/workers", "yara/dns_record"], "anyOf": [{"type": "string"}, {"items": {"type": "string"}, "maxItems": 50, "type": "array"}]}}, {"name": "path", "in": "query", "description": "Filter by path using exact-match semantics (no descendant matching). Omit the parameter to return rules from all paths. Pass an empty string (path=) to return only rules with an empty/uncategorized path. Pass a non-empty value (e.g. path=yara) for an exact match against that path. Repeat the parameter (e.g. path=yara&path=expr) to OR-match across multiple paths (SQL `IN (...)` semantics). Note: the `recursive` flag does NOT affect path filtering — it only affects namespace filtering for customer accounts.", "schema": {"description": "Filter by path using exact-match semantics (no descendant matching). Omit the parameter to return rules from all paths. Pass an empty string (path=) to return only rules with an empty/uncategorized path. Pass a non-empty value (e.g. path=yara) for an exact match against that path. Repeat the parameter (e.g. path=yara&path=expr) to OR-match across multiple paths (SQL `IN (...)` semantics). Note: the `recursive` flag does NOT affect path filtering — it only affects namespace filtering for customer accounts.", "example": ["yara/workers"], "anyOf": [{"type": "string"}, {"items": {"type": "string"}, "maxItems": 50, "type": "array"}]}}, {"name": "recursive", "in": "query", "description": "Customer accounts only: when true, namespace filtering matches descendants. Does NOT affect path filtering (paths are always exact-match).", "schema": {"description": "Customer accounts only: when true, namespace filtering matches descendants. Does NOT affect path filtering (paths are always exact-match).", "type": "string", "example": "true", "enum": ["true", "false"]}}, {"name": "search", "in": "query", "schema": {"type": "string", "example": "malicious"}}, {"name": "is_public", "in": "query", "description": "Filter by public visibility.", "schema": {"description": "Filter by public visibility.", "type": "string", "example": "true", "enum": ["true", "false"]}}, {"name": "limit", "in": "query", "schema": {"type": "number", "example": 50, "default": 50, "maximum": 100, "minimum": 1}}, {"name": "offset", "in": "query", "schema": {"type": "number", "example": 0, "default": 0, "minimum": 0, "nullable": true}}, {"name": "query", "in": "query", "description": "Natural-language or keyword search query.", "required": true, "schema": {"description": "Natural-language or keyword search query.", "type": "string", "example": "obfuscated proxy worker", "maxLength": 500, "minLength": 1}}, {"name": "mode", "in": "query", "description": "Retrieval strategy used for the query.", "schema": {"description": "Retrieval strategy used for the query.", "type": "string", "example": "hybrid", "default": "hybrid", "enum": ["exact", "hybrid", "vector", "keyword"]}}, {"name": "language", "in": "query", "description": "Filter results by rule language.", "schema": {"description": "Filter results by rule language.", "type": "string", "example": "yara", "enum": ["yara", "js"]}}], "responses": {"200": {"description": "Hybrid/semantic rule search results", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_RulesSearchResponse"}}}}, "401": {"description": "Unauthorized", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}, "503": {"description": "AI Search unavailable", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Rules"]}
```
