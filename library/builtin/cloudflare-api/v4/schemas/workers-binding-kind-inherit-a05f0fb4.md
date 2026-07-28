---
title: workers_binding_kind_inherit
page_id: schema-workers-binding-kind-inherit-a05f0fb4
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_inherit

```yaml
{"type": "object", "properties": {"name": {"description": "The name of the inherited binding.", "type": "string", "example": "MY_BINDING", "x-auditable": true}, "old_name": {"description": "The old name of the inherited binding. If set, the binding will be renamed from `old_name` to `name` in the new version. If not set, the binding will keep the same name between versions.", "type": "string", "example": "MY_OLD_BINDING", "x-auditable": true}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["inherit"], "x-auditable": true}, "version_id": {"description": "Identifier for the version to inherit the binding from, which can be the version ID or the literal \"latest\" to inherit from the latest version. Defaults to inheriting the binding from the latest version.", "type": "string", "example": "8969331f-7192-434c-9938-6aea24ed58bf", "default": "latest", "x-auditable": true}}, "required": ["name", "type"]}
```
