---
title: POST /paas/v4/files
page_id: operation-post-paas-v4-files-62b33aa2
path: operations/untagged
description: This API is designed for uploading auxiliary files (such as glossaries, terminology lists) to support the translation service. It allows users to upload reference materials that can enhance translation accuracy and consistency.
source: https://docs.z.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /paas/v4/files
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# POST /paas/v4/files

`POST /paas/v4/files`

This API is designed for uploading auxiliary files (such as glossaries, terminology lists) to support the translation service. It allows users to upload reference materials that can enhance translation accuracy and consistency.

## Definition

```yaml
{"description": "This API is designed for uploading auxiliary files (such as glossaries, terminology lists) to support the translation service. It allows users to upload reference materials that can enhance translation accuracy and consistency.", "requestBody": {"content": {"multipart/form-data": {"schema": {"type": "object", "properties": {"purpose": {"type": "string", "description": "Upload purpose (agent)", "default": "agent", "enum": ["agent"]}, "file": {"type": "string", "format": "binary", "description": "File to upload. Limit to `100MB`. Allowed formats: `pdf`, `doc`, `xlsx`, `ppt`, `txt`, `jpg`, `png`."}}, "required": ["purpose", "file"]}}}, "required": true}, "responses": {"200": {"description": "Processing successful", "content": {"application/json": {"schema": {"type": "object", "properties": {"id": {"type": "string", "description": "Unique identifier of the uploaded file."}, "object": {"type": "string", "description": "Object type."}, "bytes": {"type": "integer", "description": "File size in bytes."}, "filename": {"type": "string", "description": "Name of the uploaded file."}, "purpose": {"type": "string", "description": "Purpose of the uploaded file."}, "created_at": {"type": "integer", "description": "Timestamp of file creation."}}}}}}, "default": {"description": "The request has failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/Error"}}}}}}
```
