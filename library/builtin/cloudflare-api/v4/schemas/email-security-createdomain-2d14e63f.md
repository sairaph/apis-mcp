---
title: email-security_CreateDomain
page_id: schema-email-security-createdomain-2d14e63f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_CreateDomain

```yaml
{"type": "object", "properties": {"allowed_delivery_modes": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_DeliveryMode"}, "maxItems": 10, "minItems": 1}, "domain": {"type": "string", "x-auditable": true}, "drop_dispositions": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_DispositionLabel"}, "maxItems": 10}, "folder": {"$ref": "#/components/schemas/email-security_ScannableFolder"}, "integration_id": {"type": "string", "format": "uuid", "nullable": true, "x-auditable": true}, "ip_restrictions": {"type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["192.0.2.0/24", "2001:db8::/32"], "maxItems": 100}, "lookback_hops": {"type": "integer", "maximum": 20, "minimum": 1, "nullable": true, "x-auditable": true}, "regions": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_Region"}, "maxItems": 10, "minItems": 1}, "require_tls_inbound": {"type": "boolean", "nullable": true, "x-auditable": true}, "require_tls_outbound": {"type": "boolean", "nullable": true, "x-auditable": true}, "transport": {"type": "string", "nullable": true, "x-auditable": true}}, "additionalProperties": false, "required": ["domain", "allowed_delivery_modes", "ip_restrictions", "drop_dispositions", "regions"]}
```
