---
title: observatory_pages-response-collection
page_id: schema-observatory-pages-response-collection-a3217a60
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# observatory_pages-response-collection

```yaml
{"allOf": [{"$ref": "#/components/schemas/observatory_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"properties": {"region": {"$ref": "#/components/schemas/observatory_labeled_region"}, "scheduleFrequency": {"$ref": "#/components/schemas/observatory_schedule_frequency"}, "tests": {"type": "array", "items": {"$ref": "#/components/schemas/observatory_page_test"}}, "url": {"$ref": "#/components/schemas/observatory_url"}}, "type": "object"}}}, "type": "object"}]}
```
