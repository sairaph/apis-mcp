---
title: dlp_DatasetUpload
page_id: schema-dlp-datasetupload-2f7d1eba
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_DatasetUpload

```yaml
{"type": "object", "properties": {"num_cells": {"type": "integer", "format": "int64"}, "status": {"$ref": "#/components/schemas/dlp_DatasetUploadStatus"}, "version": {"type": "integer", "format": "int64"}}, "required": ["version", "status", "num_cells"]}
```
