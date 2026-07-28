---
title: workers_binding_kind_r2_bucket
page_id: schema-workers-binding-kind-r2-bucket-d0029d27
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_r2_bucket

```yaml
{"type": "object", "properties": {"bucket_name": {"description": "R2 bucket to bind to.", "type": "string", "example": "my-r2-bucket", "x-auditable": true}, "jurisdiction": {"description": "The [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions) of the R2 bucket.", "type": "string", "example": "eu", "enum": ["eu", "fedramp", "fedramp-high"], "x-auditable": true}, "name": {"$ref": "#/components/schemas/workers_binding_name"}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["r2_bucket"], "x-auditable": true}}, "required": ["name", "type", "bucket_name"]}
```
