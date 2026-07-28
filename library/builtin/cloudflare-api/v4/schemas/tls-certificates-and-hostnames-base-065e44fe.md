---
title: tls-certificates-and-hostnames_base
page_id: schema-tls-certificates-and-hostnames-base-065e44fe
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_base

```yaml
{"type": "object", "properties": {"created_on": {"description": "When the Keyless SSL was created.", "type": "string", "format": "date-time", "example": "2014-01-01T05:20:00Z", "readOnly": true, "x-auditable": true}, "enabled": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_enabled"}, "host": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_host"}, "id": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier-2"}, "modified_on": {"description": "When the Keyless SSL was last modified.", "type": "string", "format": "date-time", "example": "2014-01-01T05:20:00Z", "readOnly": true, "x-auditable": true}, "name": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_name"}, "permissions": {"description": "Available permissions for the Keyless SSL for the current user requesting the item.", "type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["#ssl:read", "#ssl:edit"], "readOnly": true}, "port": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_port"}, "status": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_status-2"}, "tunnel": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_keyless_tunnel"}}, "required": ["id", "name", "host", "port", "status", "enabled", "permissions", "created_on", "modified_on"]}
```
