---
title: dlp_DatasetUpdate
page_id: schema-dlp-datasetupdate-fb253c2d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_DatasetUpdate

```yaml
{"type": "object", "properties": {"case_sensitive": {"description": "Determines if the words should be matched in a case-sensitive manner.\n\nOnly required for custom word lists.", "type": "boolean"}, "description": {"description": "The description of the dataset.", "type": "string", "nullable": true}, "name": {"description": "The name of the dataset, must be unique.", "type": "string", "nullable": true}}}
```
