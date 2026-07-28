---
title: workers_binding_kind_durable_object_namespace
page_id: schema-workers-binding-kind-durable-object-namespace-2510990e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_durable_object_namespace

```yaml
{"type": "object", "properties": {"class_name": {"description": "The exported class name of the Durable Object.", "type": "string", "example": "MyDurableObject", "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}, "dispatch_namespace": {"description": "The dispatch namespace the Durable Object script belongs to.", "type": "string", "example": "my-dispatch-namespace", "x-auditable": true}, "environment": {"description": "The environment of the script_name to bind to.", "type": "string", "example": "production", "x-auditable": true}, "name": {"$ref": "#/components/schemas/workers_binding_name"}, "namespace_id": {"allOf": [{"$ref": "#/components/schemas/workers_namespace_identifier"}, {"type": "string", "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}]}, "script_name": {"description": "The script where the Durable Object is defined, if it is external to this Worker.", "type": "string", "example": "my-other-worker", "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["durable_object_namespace"], "x-auditable": true}}, "required": ["name", "type"]}
```
