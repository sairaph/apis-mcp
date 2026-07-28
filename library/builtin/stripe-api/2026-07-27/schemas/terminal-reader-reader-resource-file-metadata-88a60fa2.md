---
title: terminal_reader_reader_resource_file_metadata
page_id: schema-terminal-reader-reader-resource-file-metadata-88a60fa2
path: schemas
description: Metadata of an uploaded file
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal_reader_reader_resource_file_metadata

Metadata of an uploaded file

```yaml
{"title": "TerminalReaderReaderResourceFileMetadata", "required": ["created_at", "filename", "size", "type"], "type": "object", "properties": {"created_at": {"type": "integer", "description": "Creation time of the object (in seconds since the Unix epoch).", "format": "unix-time"}, "filename": {"maxLength": 5000, "type": "string", "description": "The original name of the uploaded file (e.g. `receipt.png`)."}, "size": {"type": "integer", "description": "The size (in bytes) of the uploaded file."}, "type": {"maxLength": 5000, "type": "string", "description": "The format of the uploaded file."}}, "description": "Metadata of an uploaded file", "x-expandableFields": []}
```
