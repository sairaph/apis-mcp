---
title: email-security_DomainVerificationInfo
page_id: schema-email-security-domainverificationinfo-fd2840a1
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_DomainVerificationInfo

```yaml
{"type": "object", "properties": {"last_checked_at": {"description": "When the last DNS TXT check was attempted, if any.", "type": "string", "format": "date-time", "nullable": true}, "status": {"$ref": "#/components/schemas/email-security_DomainStatus"}, "txt_record_name": {"description": "Full DNS TXT record name (e.g. `_cf-email-sec-challenge.example.com`).", "type": "string"}, "txt_record_value": {"description": "Token value to publish in the TXT record.", "type": "string"}}, "required": ["status", "txt_record_name", "txt_record_value", "last_checked_at"]}
```
