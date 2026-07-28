---
title: api-shield_schema_hosts
page_id: schema-api-shield-schema-hosts-2e6308ae
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_schema_hosts

```yaml
{"type": "object", "properties": {"created_at": {"$ref": "#/components/schemas/api-shield_timestamp-2"}, "hosts": {"description": "Hosts serving the schema, e.g zone.host.com", "type": "array", "items": {"type": "string"}, "readOnly": true, "x-auditable": true}, "name": {"description": "Name of the schema", "type": "string", "example": "petstore schema", "readOnly": true, "x-auditable": true}, "schema_id": {"type": "string", "allOf": [{"$ref": "#/components/schemas/api-shield_uuid-2"}, {"description": "A unique identifier of this schema", "format": "uuid", "readOnly": true, "type": "string"}], "x-auditable": true}}, "required": ["schema_id", "name", "hosts", "created_at"]}
```
