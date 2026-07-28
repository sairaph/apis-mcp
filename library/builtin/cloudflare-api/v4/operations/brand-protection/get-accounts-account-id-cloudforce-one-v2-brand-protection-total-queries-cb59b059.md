---
title: Get total queries
page_id: operation-get-accounts-account-id-cloudforce-one-v2-brand-protection-total-queries-d0cc3f30
path: operations/brand-protection
description: Get the total number of saved brand protection queries (domain + logo) for an account
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/total-queries
operation_ids:
    - get_TotalQueries
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get total queries

`GET /accounts/{account_id}/cloudforce-one/v2/brand-protection/total-queries`

Operation ID: `get_TotalQueries`

Get the total number of saved brand protection queries (domain + logo) for an account

## Definition

```yaml
{"operationId": "get_TotalQueries", "summary": "Get total queries", "description": "Get the total number of saved brand protection queries (domain + logo) for an account", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}], "responses": {"200": {"description": "Successfully retrieved total query count", "content": {"application/json": {"schema": {"type": "array", "items": {"properties": {"total_queries": {"type": "integer", "minimum": 0}}, "required": ["total_queries"], "type": "object"}}}}}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
