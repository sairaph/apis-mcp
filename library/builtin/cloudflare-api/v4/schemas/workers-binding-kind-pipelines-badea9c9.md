---
title: workers_binding_kind_pipelines
page_id: schema-workers-binding-kind-pipelines-badea9c9
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_pipelines

```yaml
{"type": "object", "properties": {"name": {"$ref": "#/components/schemas/workers_binding_name"}, "pipeline": {"description": "Name of the Pipeline to bind to.", "type": "string", "example": "my-pipeline", "x-auditable": true}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["pipelines"], "x-auditable": true}}, "required": ["name", "type", "pipeline"]}
```
