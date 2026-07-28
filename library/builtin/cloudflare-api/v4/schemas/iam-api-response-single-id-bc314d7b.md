---
title: iam_api-response-single-id
page_id: schema-iam-api-response-single-id-bc314d7b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_api-response-single-id

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/iam_api-response-common"}, {"properties": {"result": {"type": "object", "nullable": true, "properties": {"id": {"$ref": "#/components/schemas/iam_common_components-schemas-identifier"}}, "required": ["id"]}}}]}
```
