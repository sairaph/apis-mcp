---
title: Get logo queries
page_id: operation-get-accounts-account-id-cloudforce-one-v2-brand-protection-logo-queries-a50ee95b
path: operations/brand-protection
description: Get all saved brand protection logo queries for an account. Optionally specify id to get a single query. Set download=true to include base64-encoded image data.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/logo/queries
operation_ids:
    - get_GetLogoQueries
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get logo queries

`GET /accounts/{account_id}/cloudforce-one/v2/brand-protection/logo/queries`

Operation ID: `get_GetLogoQueries`

Get all saved brand protection logo queries for an account. Optionally specify id to get a single query. Set download=true to include base64-encoded image data.

## Definition

```yaml
{"operationId": "get_GetLogoQueries", "summary": "Get logo queries", "description": "Get all saved brand protection logo queries for an account. Optionally specify id to get a single query. Set download=true to include base64-encoded image data.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}, {"name": "id", "in": "query", "description": "Optional query ID to retrieve a specific logo query", "schema": {"description": "Optional query ID to retrieve a specific logo query", "type": "string"}}, {"name": "download", "in": "query", "description": "If true, include base64-encoded image data in the response", "schema": {"description": "If true, include base64-encoded image data in the response", "type": "string"}}], "responses": {"200": {"description": "Successfully retrieved logo queries", "content": {"application/json": {"schema": {"type": "array", "items": {"properties": {"content_type": {"description": "MIME type of the image (only present when download=true)", "type": "string"}, "id": {"type": "integer"}, "image_data": {"description": "Base64-encoded image data (only present when download=true)", "type": "string"}, "r2_path": {"type": "string"}, "similarity_threshold": {"type": "number"}, "tag": {"type": "string"}, "uploaded_at": {"type": "string", "nullable": true}}, "required": ["id", "tag", "r2_path", "similarity_threshold", "uploaded_at"], "type": "object"}}}}}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
