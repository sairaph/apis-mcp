---
title: dlp_DatasetCreation
page_id: schema-dlp-datasetcreation-3d0abfc1
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_DatasetCreation

```yaml
{"type": "object", "properties": {"dataset": {"$ref": "#/components/schemas/dlp_Dataset"}, "encoding_version": {"description": "Encoding version to use for dataset.", "type": "integer", "format": "int32", "minimum": 0}, "max_cells": {"type": "integer", "format": "int64", "minimum": 0}, "secret": {"description": "The secret to use for Exact Data Match datasets.\n\nThis is not present in Custom Wordlists.", "type": "string", "format": "password"}, "version": {"description": "The version to use when uploading the dataset.", "type": "integer", "format": "int64"}}, "required": ["version", "max_cells", "dataset", "encoding_version"]}
```
