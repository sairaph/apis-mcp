---
title: workers_versions-list-response
page_id: schema-workers-versions-list-response-99e117a6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_versions-list-response

```yaml
{"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"type": "object", "properties": {"items": {"type": "array", "items": {"$ref": "#/components/schemas/workers_version-item-short"}}}}}, "required": ["result"], "type": "object"}]}
```
