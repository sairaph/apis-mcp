---
title: email-security_SendingDomainRestriction
page_id: schema-email-security-sendingdomainrestriction-0febb166
path: schemas
description: A sending domain restriction that enforces TLS (Transport Layer Security) requirements for emails from specific domains. If TLS is required, the system drops mail without TLS from the specified domain.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_SendingDomainRestriction

A sending domain restriction that enforces TLS (Transport Layer Security) requirements for emails from specific domains. If TLS is required, the system drops mail without TLS from the specified domain.

```yaml
{"description": "A sending domain restriction that enforces TLS (Transport Layer Security) requirements for emails from specific domains. If TLS is required, the system drops mail without TLS from the specified domain.", "type": "object", "properties": {"comments": {"type": "string", "example": "Enforce TLS for all mail from this domain", "maxLength": 1024, "nullable": true}, "created_at": {"allOf": [{"$ref": "#/components/schemas/email-security_timestamp"}], "readOnly": true}, "domain": {"description": "Domain that requires TLS enforcement.", "type": "string", "example": "example.com", "x-auditable": true}, "exclude": {"description": "Subdomains to exempt from TLS requirements.", "type": "array", "items": {"type": "string"}, "example": ["subdomain.example.com"], "x-auditable": true}, "id": {"allOf": [{"$ref": "#/components/schemas/email-security_SendingDomainRestrictionId"}], "readOnly": true}, "last_modified": {"description": "Deprecated, use `modified_at` instead. End of life: November 1, 2026.", "allOf": [{"$ref": "#/components/schemas/email-security_timestamp"}], "deprecated": true, "readOnly": true, "x-stainless-deprecation-message": "Use `modified_at` instead."}, "modified_at": {"allOf": [{"$ref": "#/components/schemas/email-security_timestamp"}], "readOnly": true}}}
```
