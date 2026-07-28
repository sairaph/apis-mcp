---
title: infra_ServiceCommon
page_id: schema-infra-servicecommon-5705e51f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# infra_ServiceCommon

```yaml
{"type": "object", "properties": {"created_at": {"type": "string", "format": "date-time", "example": "2024-01-15T09:30:00Z", "readOnly": true}, "host": {"$ref": "#/components/schemas/infra_ServiceHost"}, "name": {"type": "string", "example": "web-server"}, "service_id": {"type": "string", "format": "uuid", "example": "550e8400-e29b-41d4-a716-446655440000", "readOnly": true}, "tls_settings": {"type": "object", "allOf": [{"$ref": "#/components/schemas/infra_ApiTlsSettings"}], "nullable": true}, "type": {"$ref": "#/components/schemas/infra_ServiceType"}, "updated_at": {"type": "string", "format": "date-time", "example": "2024-01-15T10:45:00Z", "readOnly": true}}, "required": ["name", "type", "host"]}
```
