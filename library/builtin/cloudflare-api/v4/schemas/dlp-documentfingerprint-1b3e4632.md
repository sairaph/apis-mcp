---
title: dlp_DocumentFingerprint
page_id: schema-dlp-documentfingerprint-1b3e4632
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_DocumentFingerprint

```yaml
{"type": "object", "properties": {"created_at": {"type": "string", "format": "date-time"}, "description": {"type": "string", "default": ""}, "entry_id": {"type": "string", "format": "uuid"}, "file_name": {"type": "string", "nullable": true}, "id": {"type": "string", "format": "uuid"}, "match_percent": {"type": "integer", "format": "int32"}, "name": {"type": "string"}, "status": {"$ref": "#/components/schemas/dlp_DatasetUploadStatus"}, "updated_at": {"type": "string", "format": "date-time"}, "version": {"type": "integer", "format": "int64", "nullable": true}}, "required": ["id", "entry_id", "name", "description", "match_percent", "status", "created_at", "updated_at"]}
```
