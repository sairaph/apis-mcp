---
title: zones_email_obfuscation
page_id: schema-zones-email-obfuscation-ddc1a324
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_email_obfuscation

```yaml
{"type": "object", "properties": {"id": {"description": "Turn on or off **Email Obfuscation**.", "type": "string", "enum": ["email_obfuscation"], "x-auditable": true}, "value": {"description": "The status of Email Obfuscation.\n", "type": "string", "example": "on", "enum": ["on", "off"], "x-auditable": true}}, "title": "Email Obfuscation", "x-stainless-skip": ["terraform"]}
```
