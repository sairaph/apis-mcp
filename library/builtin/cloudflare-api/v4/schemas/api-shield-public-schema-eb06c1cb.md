---
title: api-shield_public_schema
page_id: schema-api-shield-public-schema-eb06c1cb
path: schemas
description: A schema used in schema validation
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_public_schema

A schema used in schema validation

```yaml
{"description": "A schema used in schema validation", "type": "object", "properties": {"created_at": {"$ref": "#/components/schemas/api-shield_timestamp-2"}, "kind": {"description": "The kind of the schema", "type": "string", "example": "openapi_v3", "enum": ["openapi_v3"], "readOnly": true, "x-auditable": true}, "name": {"description": "A human-readable name for the schema", "type": "string", "example": "petstore schema", "readOnly": true, "x-auditable": true}, "schema_id": {"type": "string", "allOf": [{"$ref": "#/components/schemas/api-shield_uuid-2"}, {"description": "A unique identifier of this schema", "format": "uuid", "readOnly": true, "type": "string"}], "x-auditable": true}, "source": {"description": "The raw schema, e.g., the OpenAPI schema, either as JSON or YAML", "type": "string", "example": "<schema file contents>", "readOnly": true, "x-auditable": true}, "validation_enabled": {"description": "An indicator if this schema is enabled", "type": "boolean", "x-auditable": true}}, "required": ["schema_id", "name", "kind", "source", "created_at"]}
```
