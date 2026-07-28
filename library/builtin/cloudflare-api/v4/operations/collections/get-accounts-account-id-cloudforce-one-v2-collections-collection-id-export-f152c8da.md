---
title: Export collection to CSV, JSONL, or Markdown
page_id: operation-get-accounts-account-id-cloudforce-one-v2-collections-collection-id-expo-7b0e1e3c
path: operations/collections
description: Streams collection data in the requested format. Supports CSV (default), JSONL, or Markdown via Accept header. Queries items in batches of 1000 to avoid memory/timeout limits.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/collections/{collection_id}/export
operation_ids:
    - get_CollectionExportEndpoint
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Export collection to CSV, JSONL, or Markdown

`GET /accounts/{account_id}/cloudforce-one/v2/collections/{collection_id}/export`

Operation ID: `get_CollectionExportEndpoint`

Streams collection data in the requested format. Supports CSV (default), JSONL, or Markdown via Accept header. Queries items in batches of 1000 to avoid memory/timeout limits.

## Definition

```yaml
{"operationId": "get_CollectionExportEndpoint", "summary": "Export collection to CSV, JSONL, or Markdown", "description": "Streams collection data in the requested format. Supports CSV (default), JSONL, or Markdown via Accept header. Queries items in batches of 1000 to avoid memory/timeout limits.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID (hex format)", "required": true, "schema": {"description": "Account ID (hex format)", "type": "string"}}, {"name": "collection_id", "in": "path", "description": "Collection UUID", "required": true, "schema": {"description": "Collection UUID", "type": "string"}}, {"name": "include_ids", "in": "query", "description": "Include item IDs in export (default: false)", "schema": {"description": "Include item IDs in export (default: false)", "type": "boolean"}}, {"name": "accept", "in": "header", "description": "Requested format: text/csv (default), application/x-ndjson, text/markdown", "schema": {"description": "Requested format: text/csv (default), application/x-ndjson, text/markdown", "type": "string"}}], "responses": {"200": {"description": "Streaming export response", "content": {"application/x-ndjson": {"schema": {"description": "JSONL export", "type": "string"}}, "text/csv": {"schema": {"description": "CSV export", "type": "string"}}, "text/markdown": {"schema": {"description": "Markdown table export", "type": "string"}}}}, "404": {"description": "Collection not found"}, "406": {"description": "Unsupported format requested"}}, "security": [{"api_token": []}], "tags": ["Collections"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
