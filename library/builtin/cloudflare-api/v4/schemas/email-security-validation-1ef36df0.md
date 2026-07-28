---
title: email-security_Validation
page_id: schema-email-security-validation-1ef36df0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_Validation

```yaml
{"type": "object", "properties": {"comment": {"type": "string", "nullable": true}, "dkim": {"$ref": "#/components/schemas/email-security_ValidationStatus"}, "dmarc": {"$ref": "#/components/schemas/email-security_ValidationStatus"}, "spf": {"$ref": "#/components/schemas/email-security_ValidationStatus"}}}
```
