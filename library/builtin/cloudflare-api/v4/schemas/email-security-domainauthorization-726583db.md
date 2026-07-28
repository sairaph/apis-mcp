---
title: email-security_DomainAuthorization
page_id: schema-email-security-domainauthorization-726583db
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_DomainAuthorization

```yaml
{"type": "object", "properties": {"authorized": {"type": "boolean"}, "status_message": {"type": "string", "nullable": true}, "timestamp": {"type": "string", "format": "date-time", "readOnly": true}}, "required": ["authorized", "timestamp"]}
```
