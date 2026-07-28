---
title: email-security_UpdateDomain
page_id: schema-email-security-updatedomain-1ac73895
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_UpdateDomain

```yaml
{"type": "object", "properties": {"allowed_delivery_modes": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_DeliveryMode"}}, "drop_dispositions": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_DispositionLabel"}}, "folder": {"$ref": "#/components/schemas/email-security_ScannableFolder"}, "integration_id": {"type": "string", "format": "uuid", "nullable": true, "x-auditable": true}, "ip_restrictions": {"type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["192.0.2.0/24", "2001:db8::/32"]}, "lookback_hops": {"type": "integer", "maximum": 20, "minimum": 1, "x-auditable": true}, "regions": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_Region"}}, "require_tls_inbound": {"type": "boolean", "x-auditable": true}, "require_tls_outbound": {"type": "boolean", "x-auditable": true}, "transport": {"type": "string", "x-auditable": true}}}
```
