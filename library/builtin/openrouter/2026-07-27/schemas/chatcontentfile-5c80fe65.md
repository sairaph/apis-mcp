---
title: ChatContentFile
page_id: schema-chatcontentfile-5c80fe65
path: schemas
description: File content part for document processing
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatContentFile

File content part for document processing

```yaml
{"description": "File content part for document processing", "example": {"file": {"file_data": "https://example.com/document.pdf", "filename": "document.pdf"}, "type": "file"}, "properties": {"file": {"properties": {"file_data": {"description": "File content as base64 data URL or URL", "type": "string"}, "file_id": {"description": "File ID for previously uploaded files", "type": "string"}, "filename": {"description": "Original filename", "type": "string"}}, "type": "object"}, "type": {"enum": ["file"], "type": "string"}}, "required": ["type", "file"], "type": "object"}
```
