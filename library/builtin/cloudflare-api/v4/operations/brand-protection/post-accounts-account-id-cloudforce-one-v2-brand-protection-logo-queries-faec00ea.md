---
title: Insert logo query
page_id: operation-post-accounts-account-id-cloudforce-one-v2-brand-protection-logo-queries-90cbbbf5
path: operations/brand-protection
description: Create a new saved brand protection logo query for visual similarity matching
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/logo/queries
operation_ids:
    - post_InsertLogoQuery
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Insert logo query

`POST /accounts/{account_id}/cloudforce-one/v2/brand-protection/logo/queries`

Operation ID: `post_InsertLogoQuery`

Create a new saved brand protection logo query for visual similarity matching

## Definition

```yaml
{"operationId": "post_InsertLogoQuery", "summary": "Insert logo query", "description": "Create a new saved brand protection logo query for visual similarity matching", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"image_data": {"description": "Base64 encoded image data. Can include data URI prefix (e.g., 'data:image/png;base64,...') or just the base64 string.", "type": "string", "minLength": 1}, "search_lookback": {"description": "If true, search historic scanned images for matches above the similarity threshold", "type": "boolean", "default": true}, "similarity_threshold": {"description": "Minimum similarity score (0-1) required for visual matches", "type": "number", "maximum": 1, "minimum": 0}, "tag": {"description": "Unique identifier for the logo query", "type": "string", "minLength": 1}}, "required": ["tag", "image_data", "similarity_threshold"]}}}}, "responses": {"200": {"description": "Logo query inserted successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"message": {"type": "string"}, "query_id": {"type": "integer"}, "success": {"type": "boolean"}}, "required": ["success", "message"]}}}}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write"]}
```
