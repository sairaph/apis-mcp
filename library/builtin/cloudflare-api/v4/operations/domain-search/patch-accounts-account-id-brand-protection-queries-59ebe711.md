---
title: Update saved string queries by ID
page_id: operation-patch-accounts-account-id-brand-protection-queries-d0bc2be3
path: operations/domain-search
description: Update a saved query's tag, scan setting, or string_matches (pattern). When string_matches is provided, the query parameters and hash are updated. At least one of tag, scan, or string_matches is required.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/brand-protection/queries
operation_ids:
    - patchAccountsAccountIdBrandProtectionQueries
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update saved string queries by ID

`PATCH /accounts/{account_id}/brand-protection/queries`

Operation ID: `patchAccountsAccountIdBrandProtectionQueries`

Update a saved query's tag, scan setting, or string_matches (pattern). When string_matches is provided, the query parameters and hash are updated. At least one of tag, scan, or string_matches is required.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "patchAccountsAccountIdBrandProtectionQueries", "summary": "Update saved string queries by ID", "description": "Update a saved query's tag, scan setting, or string_matches (pattern). When string_matches is provided, the query parameters and hash are updated. At least one of tag, scan, or string_matches is required.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"id": {"description": "The query ID to update (required when updating tag or scan)", "type": "integer"}, "scan": {"description": "Whether to scan matches", "type": "boolean"}, "string_matches": {"description": "Updated pattern match constraints. When provided, replaces the existing string_matches.", "type": "array", "items": {"properties": {"max_edit_distance": {"description": "Maximum Levenshtein edit distance for fuzzy matching", "type": "number"}, "pattern": {"description": "The pattern to match against", "type": "string"}}, "required": ["pattern"], "type": "object"}}, "tag": {"description": "Query tag. Required as identifier when updating string_matches.", "type": "string"}}}}}}, "responses": {"default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["domain_search"], "x-api-token-group": ["Intel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "queries", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
