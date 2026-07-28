---
title: api-shield_old_response_user_schemas_hosts
page_id: schema-api-shield-old-response-user-schemas-hosts-669b0485
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_old_response_user_schemas_hosts

```yaml
{"type": "object", "properties": {"created_at": {"$ref": "#/components/schemas/api-shield_timestamp-2"}, "hosts": {"description": "Hosts serving the schema, e.g zone.host.com", "type": "array", "items": {"type": "string"}}, "name": {"description": "Name of the schema", "type": "string", "example": "petstore schema", "x-auditable": true}, "schema_id": {"$ref": "#/components/schemas/api-shield_uuid-2"}}, "required": ["schema_id", "name", "hosts", "created_at"]}
```
