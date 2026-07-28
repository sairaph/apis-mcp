---
title: r2-data-catalog_api-response-errors
page_id: schema-r2-data-catalog-api-response-errors-f8787ce1
path: schemas
description: Contains errors if the API call was unsuccessful.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-data-catalog_api-response-errors

Contains errors if the API call was unsuccessful.

```yaml
{"description": "Contains errors if the API call was unsuccessful.", "type": "array", "items": {"properties": {"code": {"description": "Specifies the error code.", "type": "integer"}, "message": {"description": "Describes the error.", "type": "string"}}, "required": ["code", "message"], "type": "object"}}
```
