---
title: FileMetadata
page_id: schema-filemetadata-55b3de6c
path: schemas
description: Metadata describing a stored file.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# FileMetadata

Metadata describing a stored file.

```yaml
{"description": "Metadata describing a stored file.", "example": {"created_at": "2025-01-01T00:00:00Z", "downloadable": false, "filename": "document.pdf", "id": "file_011CNha8iCJcU1wXNR6q4V8w", "mime_type": "application/pdf", "size_bytes": 1024000, "type": "file"}, "properties": {"created_at": {"type": "string"}, "downloadable": {"type": "boolean"}, "filename": {"type": "string"}, "id": {"type": "string"}, "mime_type": {"type": "string"}, "size_bytes": {"type": "integer"}, "type": {"enum": ["file"], "type": "string"}}, "required": ["id", "type", "filename", "mime_type", "size_bytes", "created_at", "downloadable"], "type": "object"}
```
