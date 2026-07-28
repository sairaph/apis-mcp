---
title: workers_binding_kind_secrets_store_secret
page_id: schema-workers-binding-kind-secrets-store-secret-f05d4045
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_secrets_store_secret

```yaml
{"type": "object", "properties": {"name": {"$ref": "#/components/schemas/workers_binding_name"}, "secret_name": {"description": "Name of the secret in the store.", "type": "string", "example": "my_secret", "x-auditable": true}, "store_id": {"description": "ID of the store containing the secret.", "type": "string", "example": "8c8b1387108e49be85669169793e7bd2", "x-auditable": true}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["secrets_store_secret"], "x-auditable": true}}, "required": ["name", "type", "store_id", "secret_name"]}
```
