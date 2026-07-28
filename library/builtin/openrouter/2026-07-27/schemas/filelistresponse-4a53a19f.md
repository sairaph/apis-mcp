---
title: FileListResponse
page_id: schema-filelistresponse-4a53a19f
path: schemas
description: A page of files belonging to the requesting workspace.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# FileListResponse

A page of files belonging to the requesting workspace.

```yaml
{"description": "A page of files belonging to the requesting workspace.", "example": {"cursor": null, "data": [{"created_at": "2025-01-01T00:00:00Z", "downloadable": false, "filename": "document.pdf", "id": "file_011CNha8iCJcU1wXNR6q4V8w", "mime_type": "application/pdf", "size_bytes": 1024000, "type": "file"}], "first_id": "file_011CNha8iCJcU1wXNR6q4V8w", "has_more": false, "last_id": "file_011CNha8iCJcU1wXNR6q4V8w"}, "properties": {"cursor": {"description": "Opaque cursor for the next page; null when there are no more results.", "type": ["string", "null"]}, "data": {"items": {"$ref": "#/components/schemas/FileMetadata"}, "type": "array"}, "first_id": {"type": ["string", "null"]}, "has_more": {"type": "boolean"}, "last_id": {"type": ["string", "null"]}}, "required": ["data", "has_more", "first_id", "last_id", "cursor"], "type": "object"}
```
