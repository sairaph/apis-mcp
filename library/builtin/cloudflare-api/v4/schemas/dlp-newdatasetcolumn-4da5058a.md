---
title: dlp_NewDatasetColumn
page_id: schema-dlp-newdatasetcolumn-4da5058a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_NewDatasetColumn

```yaml
{"allOf": [{"oneOf": [{"properties": {"entry_id": {"type": "string", "format": "uuid"}}, "required": ["entry_id"], "title": "Existing Column", "type": "object"}, {"properties": {"entry_name": {"type": "string"}}, "required": ["entry_name"], "title": "New Column", "type": "object"}]}, {"properties": {"header_name": {"type": "string"}, "num_cells": {"type": "integer", "format": "int64", "minimum": 0}}, "required": ["header_name", "num_cells"], "type": "object"}]}
```
