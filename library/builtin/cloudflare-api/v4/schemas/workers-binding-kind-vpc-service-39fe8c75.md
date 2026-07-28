---
title: workers_binding_kind_vpc_service
page_id: schema-workers-binding-kind-vpc-service-39fe8c75
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_vpc_service

```yaml
{"type": "object", "properties": {"name": {"$ref": "#/components/schemas/workers_binding_name"}, "service_id": {"description": "Identifier of the VPC service to bind to.", "type": "string", "example": "8c8b1387108e49be85669169793e7bd2", "x-auditable": true}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["vpc_service"], "x-auditable": true}}, "required": ["name", "type", "service_id"]}
```
