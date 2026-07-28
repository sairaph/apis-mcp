---
title: pages_pages_assets_upload_request
page_id: schema-pages-pages-assets-upload-request-e5f1bbce
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# pages_pages_assets_upload_request

```yaml
{"type": "array", "items": {"properties": {"base64": {"description": "Whether value is base64 encoded.", "type": "boolean", "example": true}, "key": {"description": "File content hash used as the object key in the Pages asset store.", "type": "string", "example": "b026324c6904b2a9cb4b88d6d61c81d1"}, "metadata": {"type": "object", "properties": {"contentType": {"description": "MIME type for the uploaded file.", "type": "string", "example": "text/plain"}}, "required": ["contentType"]}, "value": {"description": "File content. When base64 is true, this value is base64 encoded.", "type": "string", "example": "SGVsbG8sIFdvcmxkIQ=="}}, "required": ["key", "value", "metadata", "base64"], "type": "object"}}
```
