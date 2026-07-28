---
title: posture-api_paginated-finding-type-list
page_id: schema-posture-api-paginated-finding-type-list-bd8f3eab
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_paginated-finding-type-list

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/posture-api_api-response-collection"}, {"properties": {"result": {"description": "Array of finding type objects.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_BaseFindingType"}}}, "type": "object"}]}
```
