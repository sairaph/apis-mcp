---
title: email_sending_subdomain_config_error
page_id: schema-email-sending-subdomain-config-error-5a46db14
path: schemas
description: A DNS record status detected during preview, status, or fix.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email_sending_subdomain_config_error

A DNS record status detected during preview, status, or fix.

```yaml
{"description": "A DNS record status detected during preview, status, or fix.", "type": "object", "properties": {"code": {"description": "Error code identifying the type of issue. `dkim.conflict` is\nreported whenever 2+ TXT records exist at the DKIM selector,\neven if one matches the canonical Cloudflare content — multi-record\nDKIM can permerror at recipient verification regardless of which\nentry is correct. `domainkey.delegated` indicates that\n`_domainkey.<subdomain>` has NS records in the Cloudflare zone,\ncreating a zone cut that shadows the DKIM TXT record we publish;\nthe customer must remove the NS delegation for DKIM to verify.\n", "type": "string", "example": "dkim.conflict", "enum": ["mx.missing", "mx.foreign", "spf.missing", "spf.foreign", "spf.multiple", "dkim.missing", "dkim.conflict", "dmarc.missing", "dmarc.multiple", "domainkey.delegated"]}, "existing": {"$ref": "#/components/schemas/email_dns_record"}, "missing": {"$ref": "#/components/schemas/email_dns_record"}, "multiple": {"type": "array", "items": {"$ref": "#/components/schemas/email_dns_record"}}}, "required": ["code"]}
```
