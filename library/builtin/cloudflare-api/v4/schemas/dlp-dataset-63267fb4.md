---
title: dlp_Dataset
page_id: schema-dlp-dataset-63267fb4
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_Dataset

```yaml
{"type": "object", "properties": {"case_sensitive": {"type": "boolean"}, "columns": {"type": "array", "items": {"$ref": "#/components/schemas/dlp_DatasetColumn"}}, "created_at": {"type": "string", "format": "date-time"}, "description": {"description": "The description of the dataset.", "type": "string", "nullable": true}, "encoding_version": {"type": "integer", "format": "int32", "minimum": 0}, "id": {"type": "string", "format": "uuid"}, "name": {"type": "string"}, "num_cells": {"type": "integer", "format": "int64"}, "secret": {"type": "boolean"}, "status": {"$ref": "#/components/schemas/dlp_DatasetUploadStatus"}, "updated_at": {"description": "Stores when the dataset was last updated.\n\nThis includes name or description changes as well as uploads.", "type": "string", "format": "date-time"}, "uploads": {"type": "array", "items": {"$ref": "#/components/schemas/dlp_DatasetUpload"}}}, "required": ["name", "id", "status", "num_cells", "created_at", "updated_at", "uploads", "secret", "encoding_version", "columns"]}
```
