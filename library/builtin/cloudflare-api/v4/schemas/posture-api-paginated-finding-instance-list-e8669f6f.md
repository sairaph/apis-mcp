---
title: posture-api_paginated-finding-instance-list
page_id: schema-posture-api-paginated-finding-instance-list-e8669f6f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_paginated-finding-instance-list

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/posture-api_api-response-collection"}, {"properties": {"result": {"description": "Array of finding instance objects.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_FindingInstance"}}}, "type": "object"}]}
```
