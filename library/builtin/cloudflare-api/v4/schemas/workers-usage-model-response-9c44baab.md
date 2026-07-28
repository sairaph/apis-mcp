---
title: workers_usage-model-response
page_id: schema-workers-usage-model-response-9c44baab
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_usage-model-response

```yaml
{"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"type": "object", "properties": {"usage_model": {"$ref": "#/components/schemas/workers_usage_model"}, "user_limits": {"$ref": "#/components/schemas/workers_user_limits"}}}}, "required": ["result"], "type": "object"}]}
```
