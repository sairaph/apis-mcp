---
title: email-security_Domain
page_id: schema-email-security-domain-3f60ca6a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_Domain

```yaml
{"type": "object", "properties": {"allowed_delivery_modes": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_DeliveryMode"}}, "authorization": {"$ref": "#/components/schemas/email-security_DomainAuthorization"}, "created_at": {"allOf": [{"$ref": "#/components/schemas/email-security_timestamp"}], "readOnly": true}, "dmarc_status": {"$ref": "#/components/schemas/email-security_DmarcStatus"}, "domain": {"type": "string", "example": "example.com", "x-auditable": true}, "drop_dispositions": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_DispositionLabel"}}, "emails_processed": {"$ref": "#/components/schemas/email-security_EmailsProcessed"}, "folder": {"$ref": "#/components/schemas/email-security_ScannableFolder"}, "id": {"$ref": "#/components/schemas/email-security_DomainId"}, "inbox_provider": {"type": "string", "enum": ["Microsoft", "Google", null], "nullable": true}, "integration_id": {"type": "string", "format": "uuid", "nullable": true}, "ip_restrictions": {"type": "array", "items": {"type": "string"}, "example": ["192.0.2.0/24", "2001:db8::/32"]}, "last_modified": {"description": "Deprecated, use `modified_at` instead. End of life: November 1, 2026.", "allOf": [{"$ref": "#/components/schemas/email-security_timestamp"}], "deprecated": true, "readOnly": true, "x-stainless-deprecation-message": "Use `modified_at` instead."}, "lookback_hops": {"type": "integer", "x-auditable": true}, "modified_at": {"allOf": [{"$ref": "#/components/schemas/email-security_timestamp"}], "readOnly": true}, "o365_tenant_id": {"type": "string", "nullable": true}, "regions": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_Region"}}, "require_tls_inbound": {"type": "boolean", "nullable": true}, "require_tls_outbound": {"type": "boolean", "nullable": true}, "spf_status": {"$ref": "#/components/schemas/email-security_SpfStatus"}, "status": {"$ref": "#/components/schemas/email-security_DomainStatus"}, "transport": {"type": "string", "x-auditable": true}}}
```
