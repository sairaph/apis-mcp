---
title: email-security_MessageTrace
page_id: schema-email-security-messagetrace-3e377244
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_MessageTrace

```yaml
{"type": "object", "properties": {"inbound": {"type": "object", "properties": {"lines": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_TraceLine"}, "nullable": true}, "pending": {"type": "boolean", "nullable": true}}}, "outbound": {"type": "object", "properties": {"lines": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_TraceLine"}, "nullable": true}, "pending": {"type": "boolean", "nullable": true}}}}, "required": ["inbound", "outbound"]}
```
