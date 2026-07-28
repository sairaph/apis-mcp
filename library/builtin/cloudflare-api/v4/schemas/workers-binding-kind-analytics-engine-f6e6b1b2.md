---
title: workers_binding_kind_analytics_engine
page_id: schema-workers-binding-kind-analytics-engine-f6e6b1b2
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_analytics_engine

```yaml
{"type": "object", "properties": {"dataset": {"description": "The name of the dataset to bind to.", "type": "string", "example": "some_dataset", "x-auditable": true}, "name": {"$ref": "#/components/schemas/workers_binding_name"}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["analytics_engine"], "x-auditable": true}}, "required": ["name", "type", "dataset"]}
```
