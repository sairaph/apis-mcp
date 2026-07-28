---
title: security-center_securityTxt
page_id: schema-security-center-securitytxt-1df90c06
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# security-center_securityTxt

```yaml
{"type": "object", "properties": {"acknowledgments": {"type": "array", "items": {"format": "uri", "type": "string", "x-auditable": true}, "example": ["https://example.com/hall-of-fame.html"]}, "canonical": {"type": "array", "items": {"format": "uri", "type": "string", "x-auditable": true}, "example": ["https://www.example.com/.well-known/security.txt"]}, "contact": {"type": "array", "items": {"format": "uri", "type": "string", "x-auditable": true}, "example": ["mailto:security@example.com", "tel:+1-201-555-0123", "https://example.com/security-contact.html"]}, "enabled": {"type": "boolean", "example": true, "x-auditable": true}, "encryption": {"type": "array", "items": {"format": "uri", "type": "string"}, "example": ["https://example.com/pgp-key.txt", "dns:5d2d37ab76d47d36._openpgpkey.example.com?type=OPENPGPKEY", "openpgp4fpr:5f2de5521c63a801ab59ccb603d49de44b29100f"]}, "expires": {"type": "string", "format": "date-time", "x-auditable": true}, "hiring": {"type": "array", "items": {"format": "uri", "type": "string", "x-auditable": true}, "example": ["https://example.com/jobs.html"]}, "policy": {"type": "array", "items": {"format": "uri", "type": "string", "x-auditable": true}, "example": ["https://example.com/disclosure-policy.html"]}, "preferred_languages": {"type": "string", "example": "en, es, fr", "x-auditable": true}}}
```
