---
title: List saved query matches
page_id: operation-get-accounts-account-id-cloudforce-one-v2-brand-protection-domain-matche-e4389f1d
path: operations/brand-protection
description: Get paginated list of domain matches for one or more brand protection queries. When multiple query_ids are provided (comma-separated), matches are deduplicated across queries and each match includes a match_details array with per-match query metadata and individual dismissed state.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/domain/matches
operation_ids:
    - get_DomainMatchList
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List saved query matches

`GET /accounts/{account_id}/cloudforce-one/v2/brand-protection/domain/matches`

Operation ID: `get_DomainMatchList`

Get paginated list of domain matches for one or more brand protection queries. When multiple query_ids are provided (comma-separated), matches are deduplicated across queries and each match includes a match_details array with per-match query metadata and individual dismissed state.

## Definition

```yaml
{"operationId": "get_DomainMatchList", "summary": "List saved query matches", "description": "Get paginated list of domain matches for one or more brand protection queries. When multiple query_ids are provided (comma-separated), matches are deduplicated across queries and each match includes a match_details array with per-match query metadata and individual dismissed state.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}, {"name": "offset", "in": "query", "schema": {"type": "string", "default": "0"}}, {"name": "limit", "in": "query", "schema": {"type": "string", "default": "50"}}, {"name": "query_id", "in": "query", "required": true, "schema": {"description": "Query ID or comma-separated list of Query IDs. When multiple IDs are provided, matches are deduplicated across queries and each match includes a match_details array with per-match query metadata and dismissed state.", "type": "array", "items": {"type": "string"}}}, {"name": "include_domain_id", "in": "query", "schema": {"type": "string", "default": "false"}}, {"name": "include_dismissed", "in": "query", "schema": {"type": "string"}}, {"name": "domain_search", "in": "query", "schema": {"description": "Filter matches by domain name (substring match)", "type": "string"}}, {"name": "orderBy", "in": "query", "schema": {"description": "Column to sort by. Options: 'domain', 'first_seen', or 'registrar'", "type": "string", "enum": ["domain", "first_seen", "registrar"]}}, {"name": "order", "in": "query", "schema": {"description": "Sort order. Options: 'asc' (ascending) or 'desc' (descending)", "type": "string", "enum": ["asc", "desc"]}}], "responses": {"200": {"description": "Successfully retrieved query matches", "content": {"application/json": {"schema": {"type": "object", "properties": {"matches": {"type": "array", "items": {"properties": {"dismissed": {"description": "Whether the match is dismissed. Only present for single-query requests. For multi-query requests, use the dismissed field in each match_details entry.", "type": "boolean"}, "domain": {"type": "string"}, "first_seen": {"type": "string"}, "match_details": {"description": "Per-match detail objects with query metadata and individual dismissed state. Only present when multiple query_ids are requested.", "type": "array", "items": {"properties": {"dismissed": {"description": "Individual dismissed state for this specific match.", "type": "boolean"}, "match_id": {"type": "integer"}, "query_id": {"type": "integer"}, "query_tag": {"description": "Tag associated with the query, if one exists.", "type": "string", "nullable": true}}, "required": ["query_id", "query_tag", "match_id", "dismissed"], "type": "object"}}, "public_scans": {"type": "object", "nullable": true, "properties": {"submission_id": {"type": "string"}}, "required": ["submission_id"]}, "registrar": {"type": "string", "nullable": true}, "scan_status": {"type": "string"}, "scan_submission_id": {"type": "integer", "nullable": true}, "source": {"type": "string", "nullable": true}}, "required": ["domain", "first_seen", "public_scans", "scan_status", "scan_submission_id", "source", "registrar"], "type": "object"}}, "total": {"type": "integer", "minimum": 0}}, "required": ["matches", "total"]}}}}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
