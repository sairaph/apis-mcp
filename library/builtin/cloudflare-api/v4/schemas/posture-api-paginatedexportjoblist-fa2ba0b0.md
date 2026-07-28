---
title: posture-api_PaginatedExportJobList
page_id: schema-posture-api-paginatedexportjoblist-fa2ba0b0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_PaginatedExportJobList

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/posture-api_api-response-collection"}, {"properties": {"result": {"description": "Array of export job objects.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_ExportJob"}}}, "type": "object"}]}
```
