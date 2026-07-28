---
title: dlp_DatasetColumn
page_id: schema-dlp-datasetcolumn-0951ec13
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_DatasetColumn

```yaml
{"type": "object", "properties": {"entry_id": {"type": "string", "format": "uuid"}, "header_name": {"type": "string"}, "num_cells": {"type": "integer", "format": "int64"}, "upload_status": {"$ref": "#/components/schemas/dlp_DatasetUploadStatus"}}, "required": ["entry_id", "upload_status", "header_name", "num_cells"]}
```
