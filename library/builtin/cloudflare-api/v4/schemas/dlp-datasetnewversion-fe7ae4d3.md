---
title: dlp_DatasetNewVersion
page_id: schema-dlp-datasetnewversion-fe7ae4d3
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_DatasetNewVersion

```yaml
{"type": "object", "properties": {"case_sensitive": {"type": "boolean"}, "columns": {"type": "array", "items": {"$ref": "#/components/schemas/dlp_DatasetColumn"}}, "encoding_version": {"type": "integer", "format": "int32", "minimum": 0}, "max_cells": {"type": "integer", "format": "int64", "minimum": 0}, "secret": {"type": "string", "format": "password"}, "version": {"type": "integer", "format": "int64"}}, "required": ["version", "max_cells", "encoding_version"]}
```
