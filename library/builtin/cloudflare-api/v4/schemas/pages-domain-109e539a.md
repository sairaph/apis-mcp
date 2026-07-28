---
title: pages_domain
page_id: schema-pages-domain-109e539a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# pages_domain

```yaml
{"type": "object", "properties": {"certificate_authority": {"type": "string", "example": "lets_encrypt", "enum": ["google", "lets_encrypt"], "readOnly": true, "x-auditable": true}, "created_on": {"type": "string", "readOnly": true, "x-auditable": true}, "domain_id": {"type": "string", "readOnly": true, "x-auditable": true}, "id": {"type": "string", "readOnly": true, "x-auditable": true}, "name": {"$ref": "#/components/schemas/pages_domain_name"}, "status": {"type": "string", "enum": ["initializing", "pending", "active", "deactivated", "blocked", "error"], "readOnly": true, "x-auditable": true}, "validation_data": {"type": "object", "properties": {"error_message": {"type": "string", "x-auditable": true}, "method": {"type": "string", "enum": ["http", "txt"], "x-auditable": true}, "status": {"type": "string", "enum": ["initializing", "pending", "active", "deactivated", "error"], "x-auditable": true}, "txt_name": {"type": "string", "x-auditable": true}, "txt_value": {"type": "string", "x-auditable": true}}, "readOnly": true, "required": ["status", "method"]}, "verification_data": {"type": "object", "properties": {"error_message": {"type": "string", "x-auditable": true}, "status": {"type": "string", "enum": ["pending", "active", "deactivated", "blocked", "error"]}}, "readOnly": true, "required": ["status"]}, "zone_tag": {"type": "string", "readOnly": true, "x-auditable": true}}, "required": ["id", "domain_id", "name", "status", "created_on", "certificate_authority", "zone_tag", "validation_data", "verification_data"]}
```
