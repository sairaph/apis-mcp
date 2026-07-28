---
title: Search scanned images
page_id: operation-post-accounts-account-id-cloudforce-one-v2-brand-protection-logo-search-46fa553c
path: operations/brand-protection
description: Submit an image and find the n closest matches from the scanned images index without creating any match records. Returns similarity scores and metadata for each match.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/logo/search
operation_ids:
    - post_SearchLogoSimilarity
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Search scanned images

`POST /accounts/{account_id}/cloudforce-one/v2/brand-protection/logo/search`

Operation ID: `post_SearchLogoSimilarity`

Submit an image and find the n closest matches from the scanned images index without creating any match records. Returns similarity scores and metadata for each match.

## Definition

```yaml
{"operationId": "post_SearchLogoSimilarity", "summary": "Search scanned images", "description": "Submit an image and find the n closest matches from the scanned images index without creating any match records. Returns similarity scores and metadata for each match.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}, {"name": "showHistoric", "in": "query", "description": "Include scanned images without domain metadata (historic data). Default: false (only show images with domain)", "schema": {"description": "Include scanned images without domain metadata (historic data). Default: false (only show images with domain)", "type": "string", "default": "false"}}, {"name": "download", "in": "query", "description": "If true, include base64-encoded image data in the response", "schema": {"description": "If true, include base64-encoded image data in the response", "type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"image_data": {"description": "Base64 encoded image data. Can include data URI prefix (e.g., 'data:image/png;base64,...') or just the base64 string.", "type": "string", "minLength": 1}, "score_threshold": {"description": "Minimum similarity score threshold for matches (0-1, default: 0)", "type": "number", "default": 0, "maximum": 1, "minimum": 0}, "top_k": {"description": "Number of closest matches to return (1-100, default: 10)", "type": "integer", "default": 10, "maximum": 100, "minimum": 1}}, "required": ["image_data"]}}}}, "responses": {"200": {"description": "Scanned images search completed successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"matches": {"type": "array", "items": {"properties": {"account_id": {"type": "string"}, "content_type": {"description": "MIME type of the image (only present when download=true)", "type": "string"}, "domain": {"type": "string"}, "image_data": {"description": "Base64-encoded image data (only present when download=true)", "type": "string"}, "query_id": {"type": "integer"}, "r2_path": {"type": "string"}, "similarity_score": {"type": "number"}, "similarity_threshold": {"type": "number"}, "tag": {"type": "string"}, "timestamp": {"type": "number"}, "url_scan_id": {"type": "string"}, "vector_id": {"type": "string"}}, "required": ["similarity_score"], "type": "object"}}}, "required": ["matches"]}}}}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write"]}
```
