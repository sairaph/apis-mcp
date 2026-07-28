---
title: dlp_Limits
page_id: schema-dlp-limits-868c1981
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_Limits

```yaml
{"type": "object", "properties": {"max_custom_regex_entries": {"description": "Maximum number of custom regex entries allowed for the account.", "type": "integer", "format": "int64", "minimum": 0}, "max_dataset_cells": {"description": "Maximum number of dataset cells allowed for the account, across all EDM and CWL datasets.", "type": "integer", "format": "int64", "minimum": 0}, "max_document_fingerprints": {"description": "Maximum number of document fingerprints allowed for the account.", "type": "integer", "format": "int64", "minimum": 0}, "used_custom_regex_entries": {"description": "Number of custom regex entries currently configured for the account.", "type": "integer", "format": "int64", "minimum": 0}, "used_dataset_cells": {"description": "Number of dataset cells currently configured for the account, across all EDM and CWL datasets. Document fingerprints do not count towards this limit.", "type": "integer", "format": "int64", "minimum": 0}, "used_document_fingerprints": {"description": "Number of document fingerprints currently configured for the account.", "type": "integer", "format": "int64", "minimum": 0}}, "required": ["max_custom_regex_entries", "used_custom_regex_entries", "max_dataset_cells", "used_dataset_cells", "max_document_fingerprints", "used_document_fingerprints"]}
```
