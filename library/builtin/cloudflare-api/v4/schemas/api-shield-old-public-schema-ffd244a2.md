---
title: api-shield_old_public_schema
page_id: schema-api-shield-old-public-schema-ffd244a2
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_old_public_schema

```yaml
{"type": "object", "properties": {"created_at": {"$ref": "#/components/schemas/api-shield_timestamp-2"}, "kind": {"$ref": "#/components/schemas/api-shield_old_kind"}, "name": {"description": "Name of the schema", "type": "string", "example": "petstore schema", "x-auditable": true}, "schema_id": {"$ref": "#/components/schemas/api-shield_uuid-2"}, "source": {"description": "Source of the schema", "type": "string", "example": "<schema file bytes>", "x-auditable": true}, "validation_enabled": {"$ref": "#/components/schemas/api-shield_old_validation_enabled"}}, "required": ["schema_id", "name", "kind", "created_at"]}
```
