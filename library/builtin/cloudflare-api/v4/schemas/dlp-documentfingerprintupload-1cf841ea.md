---
title: dlp_DocumentFingerprintUpload
page_id: schema-dlp-documentfingerprintupload-1cf841ea
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_DocumentFingerprintUpload

```yaml
{"type": "object", "properties": {"created_at": {"type": "string", "format": "date-time"}, "description": {"type": "string"}, "entry_id": {"type": "string", "format": "uuid"}, "file_name": {"type": "string"}, "id": {"type": "string", "format": "uuid"}, "match_percent": {"type": "integer", "format": "int32"}, "name": {"type": "string"}, "status": {"$ref": "#/components/schemas/dlp_DatasetUploadStatus"}, "updated_at": {"type": "string", "format": "date-time"}, "version": {"type": "integer", "format": "int64"}}, "required": ["id", "entry_id", "name", "description", "match_percent", "version", "file_name", "status", "created_at", "updated_at"]}
```
